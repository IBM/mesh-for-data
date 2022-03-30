// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"emperror.dev/errors"
	distributionref "github.com/distribution/distribution/reference"
	"github.com/rs/zerolog"
	credentialprovider "github.com/vdemeester/k8s-pkg-credentialprovider"
	credentialprovidersecrets "github.com/vdemeester/k8s-pkg-credentialprovider/secrets"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/release"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	fapp "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers"
	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/helm"
	"fybrik.io/fybrik/pkg/logging"
)

const (
	BlueprintFinalizerName string = "Blueprint.finalizer"
)

// BlueprintReconciler reconciles a Blueprint object
type BlueprintReconciler struct {
	client.Client
	Name   string
	Log    zerolog.Logger
	Scheme *runtime.Scheme
	Helmer helm.Interface
}

type ExtendedArguments struct {
	fapp.ModuleArguments     `json:",inline"`
	*fapp.ApplicationDetails `json:",inline"`
	Labels                   map[string]string `json:"labels"`
	UUID                     string            `json:"uuid"`
}

// Reconcile receives a Blueprint CRD
func (r *BlueprintReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	blueprint := fapp.Blueprint{}
	if err := r.Get(ctx, req.NamespacedName, &blueprint); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	uuid := utils.GetFybrikApplicationUUIDfromAnnotations(blueprint.GetAnnotations())
	log := r.Log.With().Str(utils.FybrikAppUUID, uuid).
		Str("blueprint", req.NamespacedName.String()).Logger()

	if res, err := r.reconcileFinalizers(ctx, &blueprint); err != nil {
		log.Error().Err(err).Msg("Could not reconcile blueprint " + blueprint.GetName() + " finalizers")
		return res, err
	}

	// If the object has a scheduled deletion time, update status and return
	if !blueprint.DeletionTimestamp.IsZero() {
		// The object is being deleted
		log.Trace().Str(logging.ACTION, logging.DELETE).Msg("Deleting blueprint " + blueprint.GetName())
		return ctrl.Result{}, nil
	}

	observedStatus := blueprint.Status.DeepCopy()
	log.Trace().Str(logging.ACTION, logging.CREATE).Msg("Installing/Updating blueprint " + blueprint.GetName())

	result, err := r.reconcile(ctx, &log, &blueprint)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to reconcile blueprint")
	}

	if !equality.Semantic.DeepEqual(&blueprint.Status, observedStatus) {
		if err := utils.UpdateStatus(ctx, r.Client, &blueprint, observedStatus); err != nil {
			return ctrl.Result{}, errors.WrapWithDetails(err, "failed to update blueprint status", "status", blueprint.Status)
		}
	}

	log.Debug().Msg("blueprint reconcile cycle completed.") // TODO - Add result to log?
	return result, nil
}

// reconcileFinalizers reconciles finalizers for Blueprint
func (r *BlueprintReconciler) reconcileFinalizers(ctx context.Context, blueprint *fapp.Blueprint) (ctrl.Result, error) {
	// finalizer
	hasFinalizer := ctrlutil.ContainsFinalizer(blueprint, BlueprintFinalizerName)

	// If the object has a scheduled deletion time, delete it and its associated resources
	if !blueprint.DeletionTimestamp.IsZero() {
		// The object is being deleted
		if hasFinalizer { // Finalizer was created when the object was created
			// the finalizer is present - delete the allocated resources
			if err := r.deleteExternalResources(blueprint); err != nil {
				r.Log.Error().Err(err).Msg("Error while deleting owned resources")
			}
			// remove the finalizer from the list and update it, because it needs to be deleted together with the object
			ctrlutil.RemoveFinalizer(blueprint, BlueprintFinalizerName)

			if err := utils.UpdateFinalizers(ctx, r.Client, blueprint); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}
	// Make sure this CRD instance has a finalizer
	if !hasFinalizer {
		ctrlutil.AddFinalizer(blueprint, BlueprintFinalizerName)
		if err := utils.UpdateFinalizers(ctx, r.Client, blueprint); err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *BlueprintReconciler) deleteExternalResources(blueprint *fapp.Blueprint) error {
	errs := make([]string, 0)
	for release := range blueprint.Status.Releases {
		if _, err := r.Helmer.Uninstall(blueprint.Spec.ModulesNamespace, release); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "; "))
}

func getDomainFromImageName(image string) (string, error) {
	named, err := distributionref.ParseNormalizedNamed(image)
	if err != nil {
		return "", errors.WithMessage(err, "couldn't parse image name: "+image)
	}

	return distributionref.Domain(named), nil
}

func (r *BlueprintReconciler) obtainSecrets(ctx context.Context, log *zerolog.Logger, chartSpec fapp.ChartSpec) (string, error) {
	var registrySuccessfulLogin string
	if chartSpec.ChartPullSecret != "" {
		// obtain ChartPullSecret
		pullSecret := corev1.Secret{}
		pullSecrets := []corev1.Secret{}

		if err := r.Get(ctx, types.NamespacedName{Namespace: utils.GetSystemNamespace(), Name: chartSpec.ChartPullSecret},
			&pullSecret); err == nil {
			// if this is not a dockerconfigjson, ignore
			if pullSecret.Type == "kubernetes.io/dockerconfigjson" {
				pullSecrets = append(pullSecrets, pullSecret)
			}
		} else {
			return "", errors.WithMessage(err, "could not find ChartPullSecret: "+chartSpec.ChartPullSecret)
		}

		if len(pullSecrets) != 0 {
			// create a keyring of all dockerconfigjson secrets, to be used for lookup
			keyring := credentialprovider.NewDockerKeyring()
			keyring, _ = credentialprovidersecrets.MakeDockerKeyring(pullSecrets, keyring)
			repoToPull, err := getDomainFromImageName(chartSpec.Name)
			if err != nil {
				return "", errors.WithMessage(err, chartSpec.Name+": failed to parse image name")
			}

			creds, withCredentials := keyring.Lookup(chartSpec.Name)
			if withCredentials {
				for _, cred := range creds {
					err := r.Helmer.RegistryLogin(repoToPull, cred.Username, cred.Password, false)
					if err == nil {
						registrySuccessfulLogin = repoToPull
						break
					}
					log.Error().Msg("Failed to login to helm registry: " + repoToPull)
				}
			} else {
				log.Error().Msg("there is a mismatch between helm chart: " + chartSpec.Name +
					" and the registries associated with secret: " + chartSpec.ChartPullSecret)
			}
		}
	}
	return registrySuccessfulLogin, nil
}

func (r *BlueprintReconciler) applyChartResource(ctx context.Context, log *zerolog.Logger, chartSpec fapp.ChartSpec,
	args map[string]interface{}, blueprint *fapp.Blueprint, releaseName string) (ctrl.Result, error) {
	log.Trace().Str(logging.ACTION, logging.CREATE).Msg("--- Chart Ref ---\n\n" + chartSpec.Name + "\n\n")
	kubeNamespace := blueprint.Spec.ModulesNamespace

	args = CopyMap(args)
	for k, v := range chartSpec.Values {
		SetMapField(args, k, v)
	}
	nbytes, _ := yaml.Marshal(args)
	log.Trace().Str(logging.ACTION, logging.CREATE).Msg("--- Values.yaml ---\n\n" + string(nbytes) + "\n\n")

	var registrySuccessfulLogin string
	var err error
	registrySuccessfulLogin, err = r.obtainSecrets(ctx, log, chartSpec)
	if err != nil {
		return ctrl.Result{}, err
	}

	tmpDir, err := ioutil.TempDir("", "fybrik-helm-")
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed to create temporary directory for chart pull")
	}
	defer func(log *zerolog.Logger) {
		if err = os.RemoveAll(tmpDir); err != nil {
			log.Error().Msgf("Error while calling RemoveAll on directory %s created for pulling helm chart", tmpDir)
		}
	}(log)
	err = r.Helmer.Pull(chartSpec.Name, tmpDir)
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed chart pull")
	}
	// if we logged into a registry, let us try to logout
	if registrySuccessfulLogin != "" {
		logoutErr := r.Helmer.RegistryLogout(registrySuccessfulLogin)
		if logoutErr != nil {
			return ctrl.Result{}, errors.WithMessage(err, "failed to logout from helm registry: "+registrySuccessfulLogin)
		}
	}
	chart, err := r.Helmer.Load(chartSpec.Name, tmpDir)
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed chart load")
	}

	rel, err := r.Helmer.Status(kubeNamespace, releaseName)
	if err == nil && rel != nil {
		rel, err = r.Helmer.Upgrade(chart, kubeNamespace, releaseName, args)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed upgrade")
		}
	} else {
		rel, err = r.Helmer.Install(chart, kubeNamespace, releaseName, args)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed install")
		}
	}
	log.Trace().Str(logging.ACTION, logging.CREATE).Msg("--- Release Status ---\n\n" + string(rel.Info.Status) + "\n\n")
	return ctrl.Result{}, nil
}

// CopyMap copies a map
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

// SetMapField updates a map
func SetMapField(obj map[string]interface{}, k string, v interface{}) bool {
	components := strings.Split(k, ".")
	for n, component := range components {
		if n == len(components)-1 {
			obj[component] = v
		} else {
			m, ok := obj[component]
			if !ok {
				m := make(map[string]interface{})
				obj[component] = m
				obj = m
			} else if obj, ok = m.(map[string]interface{}); !ok {
				return false
			}
		}
	}
	return true
}

// updateModuleState updates the module state
func (r *BlueprintReconciler) updateModuleState(blueprint *fapp.Blueprint, instanceName string, isReady bool, err string) {
	state := fapp.ObservedState{
		Ready: isReady,
		Error: err,
	}
	blueprint.Status.ModulesState[instanceName] = state
}

//nolint:gocyclo
func (r *BlueprintReconciler) reconcile(ctx context.Context, log *zerolog.Logger, blueprint *fapp.Blueprint) (ctrl.Result, error) {
	uuid := utils.GetFybrikApplicationUUIDfromAnnotations(blueprint.GetAnnotations())
	modulesNamespace := blueprint.Spec.ModulesNamespace
	// Gather all templates and process them into a list of resources to apply
	// force-update if the blueprint spec is different
	updateRequired := blueprint.Status.ObservedGeneration != blueprint.GetGeneration()
	blueprint.Status.ObservedGeneration = blueprint.GetGeneration()
	// reset blueprint state
	blueprint.Status.ObservedState.Ready = false
	blueprint.Status.ObservedState.Error = ""
	if blueprint.Status.Releases == nil {
		blueprint.Status.Releases = map[string]int64{}
	}
	if blueprint.Status.ModulesState == nil {
		blueprint.Status.ModulesState = make(map[string]fapp.ObservedState)
	}
	// count the overall number of Helm releases and how many of them are ready
	numReleases, numReady := 0, 0
	// Add debug information to module labels
	if blueprint.Labels == nil {
		blueprint.Labels = map[string]string{}
	}
	blueprint.Labels[fapp.BlueprintNameLabel] = blueprint.Name
	blueprint.Labels[fapp.BlueprintNamespaceLabel] = blueprint.Namespace

	for instanceName, module := range blueprint.Spec.Modules {
		// Get arguments by type
		extendedArguments := ExtendedArguments{
			ModuleArguments:    module.Arguments,
			ApplicationDetails: blueprint.Spec.Application,
			Labels:             blueprint.Labels,
			UUID:               uuid,
		}
		var args map[string]interface{}
		args, err := utils.StructToMap(&extendedArguments)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, "Blueprint step arguments are invalid")
		}

		logging.LogStructure("Arguments", args, log, false, false)
		releaseName := utils.GetReleaseName(blueprint.Labels[fapp.ApplicationNameLabel],
			blueprint.Labels[fapp.ApplicationNamespaceLabel], instanceName)
		log.Trace().Msg("Release name: " + releaseName)
		numReleases++
		// check the release status
		rel, err := r.Helmer.Status(modulesNamespace, releaseName)
		// unexisting release or a failed release - re-apply the chart
		if updateRequired || err != nil || rel == nil || rel.Info.Status == release.StatusFailed {
			// Process templates with arguments
			chart := module.Chart
			if _, err := r.applyChartResource(ctx, log, chart, args, blueprint, releaseName); err != nil {
				blueprint.Status.ObservedState.Error += errors.Wrap(err, "ChartDeploymentFailure: ").Error() + "\n"
				r.updateModuleState(blueprint, instanceName, false, err.Error())
			} else {
				r.updateModuleState(blueprint, instanceName, false, "")
			}
		} else if rel.Info.Status == release.StatusDeployed {
			status, errMsg := r.checkReleaseStatus(releaseName, modulesNamespace, uuid)
			if status == corev1.ConditionFalse {
				blueprint.Status.ObservedState.Error += "ResourceAllocationFailure: " + errMsg + "\n"
				r.updateModuleState(blueprint, instanceName, false, errMsg)
			} else if status == corev1.ConditionTrue {
				r.updateModuleState(blueprint, instanceName, true, "")
				numReady++
			}
		}
		blueprint.Status.Releases[releaseName] = blueprint.Status.ObservedGeneration
	}
	// clean-up
	for release, version := range blueprint.Status.Releases {
		if version != blueprint.Status.ObservedGeneration {
			_, err := r.Helmer.Uninstall(modulesNamespace, release)
			if err != nil {
				log.Error().Err(err).Str(logging.ACTION, logging.DELETE).Msg("Error uninstalling release " + release)
			} else {
				delete(blueprint.Status.Releases, release)
			}
		}
	}
	// check if all releases reached the ready state
	if numReady == numReleases {
		// all modules have been orhestrated successfully - the data is ready for use
		blueprint.Status.ObservedState.Ready = true
		return ctrl.Result{}, nil
	}

	// the status is unknown yet - continue polling
	if blueprint.Status.ObservedState.Error == "" {
		return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
	}
	return ctrl.Result{}, nil
}

// NewBlueprintReconciler creates a new reconciler for Blueprint resources
func NewBlueprintReconciler(mgr ctrl.Manager, name string, helmer helm.Interface) *BlueprintReconciler {
	return &BlueprintReconciler{
		Client: mgr.GetClient(),
		Name:   name,
		Log:    logging.LogInit(logging.CONTROLLER, name),
		Scheme: mgr.GetScheme(),
		Helmer: helmer,
	}
}

// SetupWithManager registers Blueprint controller
func (r *BlueprintReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// 'UpdateFunc' and 'CreateFunc' used to judge if the event came from within the blueprint's namespace.
	// If that is true, the event will be processed by the reconciler.
	// If it's not then it is a rogue event created by someone outside of the control plane.

	blueprintNamespace := utils.GetSystemNamespace()
	p := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return e.Object.GetNamespace() == blueprintNamespace
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return e.ObjectOld.GetNamespace() == blueprintNamespace
		},
	}
	numReconciles := environment.GetEnvAsInt(controllers.BlueprintConcurrentReconcilesConfiguration,
		controllers.DefaultBlueprintConcurrentReconciles)
	r.Log.Trace().Msg("Concurrent blueprint reconciles: " + fmt.Sprint(numReconciles))

	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{MaxConcurrentReconciles: numReconciles}).
		For(&fapp.Blueprint{}).
		WithEventFilter(p).
		Complete(r)
}

func (r *BlueprintReconciler) getExpectedResults(kind string) (*fapp.ResourceStatusIndicator, error) {
	// Assumption: specification for each resource kind is done in one place.
	ctx := context.Background()

	var moduleList fapp.FybrikModuleList
	if err := r.List(ctx, &moduleList, client.InNamespace(utils.GetSystemNamespace())); err != nil {
		return nil, err
	}
	for ind := range moduleList.Items {
		for _, res := range moduleList.Items[ind].Spec.StatusIndicators {
			if res.Kind == kind {
				return res.DeepCopy(), nil
			}
		}
	}
	return nil, nil
}

// checkResourceStatus returns the computed state and an error message if exists
func (r *BlueprintReconciler) checkResourceStatus(res *unstructured.Unstructured) (corev1.ConditionStatus, string) {
	// get indications how to compute the resource status based on a module spec
	expected, err := r.getExpectedResults(res.GetKind())
	if err != nil {
		// Could not retrieve the list of modules, will retry later
		return corev1.ConditionUnknown, ""
	}
	uuid := utils.GetFybrikApplicationUUIDfromAnnotations(res.GetAnnotations())

	if expected == nil {
		// use kstatus to compute the status of the resources for them the expected results have not been specified
		// Current status of a deployed release indicates that the resource has been successfully reconciled
		// Failed status indicates a failure
		computedResult, err := kstatus.Compute(res)
		if err != nil {
			r.Log.Error().Err(err).Msg("Error computing the status of " + res.GetKind())
			return corev1.ConditionUnknown, ""
		}
		switch computedResult.Status {
		case kstatus.FailedStatus:
			return corev1.ConditionFalse, computedResult.Message
		case kstatus.CurrentStatus:
			return corev1.ConditionTrue, ""
		default:
			return corev1.ConditionUnknown, ""
		}
	}
	// use expected values to compute the status
	if r.matchesCondition(res, expected.SuccessCondition, uuid) {
		return corev1.ConditionTrue, ""
	}
	if r.matchesCondition(res, expected.FailureCondition, uuid) {
		return corev1.ConditionFalse, getErrorMessage(res, expected.ErrorMessage)
	}
	return corev1.ConditionUnknown, ""
}

func getErrorMessage(res *unstructured.Unstructured, fieldPath string) string {
	// convert the unstructured data to Labels interface to use Has and Get methods for retrieving data
	labelsImpl := utils.UnstructuredAsLabels{Data: res}
	if !labelsImpl.Has(fieldPath) {
		return ""
	}
	return labelsImpl.Get(fieldPath)
}

func (r *BlueprintReconciler) matchesCondition(res *unstructured.Unstructured, condition, uuid string) bool {
	selector, err := labels.Parse(condition)
	if err != nil {
		r.Log.Error().Err(err).Str(utils.FybrikAppUUID, uuid).
			Msg("condition " + condition + "failed to parse")
		return false
	}
	// get selector requirements, 'selectable' property is ignored
	requirements, _ := selector.Requirements()
	// convert the unstructured data to Labels interface to leverage the package capability of parsing and evaluating conditions
	labelsImpl := utils.UnstructuredAsLabels{Data: res}
	for _, req := range requirements {
		if !req.Matches(labelsImpl) {
			return false
		}
	}
	return true
}

func (r *BlueprintReconciler) checkReleaseStatus(releaseName, namespace, uuid string) (corev1.ConditionStatus, string) {
	log := r.Log.With().Str(utils.FybrikAppUUID, uuid).Logger()

	// get all resources for the given helm release in their current state
	resources, err := r.Helmer.GetResources(namespace, releaseName)
	if err != nil {
		log.Error().Err(err).Msg("Error getting resources")
		return corev1.ConditionUnknown, ""
	}
	// return True if all resources are ready, False - if any resource failed, Unknown - otherwise
	numReady := 0
	for _, res := range resources {
		state, errMsg := r.checkResourceStatus(res)
		log.Debug().Msg("Status of " + res.GetKind() + " " + res.GetName() + " is " + string(state))
		if state == corev1.ConditionFalse {
			return state, errMsg
		}
		if state == corev1.ConditionTrue {
			numReady++
		}
	}
	if numReady == len(resources) {
		return corev1.ConditionTrue, ""
	}
	return corev1.ConditionUnknown, ""
}
