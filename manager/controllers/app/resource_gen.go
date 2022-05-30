// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"

	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/utils"
)

// ContextInterface is an interface for communication with a generated resource
type ContextInterface interface {
	ResourceExists(ref *app.ResourceReference) bool
	CreateOrUpdateResource(owner *app.ResourceReference, ref *app.ResourceReference, plotterSpec *app.PlotterSpec,
		labels map[string]string, uuid string) error
	DeleteResource(ref *app.ResourceReference) error
	GetResourceStatus(ref *app.ResourceReference) (app.ObservedState, error)
	CreateResourceReference(owner *app.ResourceReference) *app.ResourceReference
	GetManagedObject() runtime.Object
}

// Interface for managing Plotter resources

// PlotterInterface context implementation for communication with a single Plotter resource
type PlotterInterface struct {
	Client client.Client
}

// GetManagedObject returns the type of the managed runtime object
func (c *PlotterInterface) GetManagedObject() runtime.Object {
	return &app.Plotter{}
}

// CreateResourceReference returns an identifier (name and namespace) of the generated resource.
func (c *PlotterInterface) CreateResourceReference(owner *app.ResourceReference) *app.ResourceReference {
	// Plotter runs in the control plane namespace. Plotter name identifies fybrikapplication (name and namespace)
	return &app.ResourceReference{
		Name:       owner.Name + "-" + owner.Namespace,
		Namespace:  utils.GetSystemNamespace(),
		Kind:       "Plotter",
		AppVersion: owner.AppVersion,
	}
}

// ResourceExists checks whether the Plotter resource generated by FybrikApplication controller is active
func (c *PlotterInterface) ResourceExists(ref *app.ResourceReference) bool {
	if ref == nil || ref.Namespace == "" {
		return false
	}
	resource := c.GetResourceSignature(ref)
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, resource); err != nil {
		return false
	}
	return true
}

// GetResourceSignature returns the namespaced information of the generated Plotter resource
func (c *PlotterInterface) GetResourceSignature(ref *app.ResourceReference) *app.Plotter {
	return &app.Plotter{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ref.Name,
			Namespace: ref.Namespace,
		},
	}
}

// CreateOrUpdateResource creates a new Plotter resource or updates an existing one
func (c *PlotterInterface) CreateOrUpdateResource(owner, ref *app.ResourceReference, plotterSpec *app.PlotterSpec,
	labels map[string]string, uuid string) error {
	plotter := c.GetResourceSignature(ref)
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, plotter); err == nil {
		if equality.Semantic.DeepEqual(&plotter.Spec, plotterSpec) {
			// nothing needs to be done
			return nil
		}
	}
	if _, err := ctrl.CreateOrUpdate(context.Background(), c.Client, plotter, func() error {
		plotter.Spec = *plotterSpec
		ctrlutil.AddFinalizer(plotter, PlotterFinalizerName)
		plotter.Labels = labels
		if plotter.Labels == nil {
			plotter.Labels = make(map[string]string)
		}
		debugLabels := ownerLabels(types.NamespacedName{Namespace: owner.Namespace, Name: owner.Name})
		for key, val := range debugLabels {
			plotter.Labels[key] = val
		}
		if plotter.Annotations == nil {
			plotter.Annotations = make(map[string]string)
			plotter.Annotations[utils.FybrikAppUUID] = uuid // For logging
		}
		modulesNamespace := utils.GetFybrikModuleNamespaceFromLabels(labels)
		if plotter.Labels == nil && modulesNamespace != "" {
			plotter.Labels = make(map[string]string)
			plotter.Labels[utils.FybrikModuleNamespace] = modulesNamespace
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteResource deletes the generated Plotter resource
func (c *PlotterInterface) DeleteResource(ref *app.ResourceReference) error {
	resource := c.GetResourceSignature(ref)
	err := c.Client.Delete(context.Background(), resource)
	return err
}

// GetResourceStatus returns the generated Plotter status
func (c *PlotterInterface) GetResourceStatus(ref *app.ResourceReference) (app.ObservedState, error) {
	if ref == nil || ref.Namespace == "" {
		return app.ObservedState{}, nil
	}
	resource := c.GetResourceSignature(ref)
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, resource); err != nil {
		return app.ObservedState{}, err
	}
	return resource.Status.ObservedState, nil
}

// NewPlotterInterface creates a new plotter interface for FybrikApplication controller
func NewPlotterInterface(cl client.Client) *PlotterInterface {
	return &PlotterInterface{
		Client: cl,
	}
}
