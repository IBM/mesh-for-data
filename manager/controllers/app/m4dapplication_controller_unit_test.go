// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"io/ioutil"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ibm/the-mesh-for-data/manager/controllers/app/modules"

	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"

	"github.com/ibm/the-mesh-for-data/manager/controllers/mockup"
	"github.com/ibm/the-mesh-for-data/pkg/storage"

	"github.com/ibm/the-mesh-for-data/pkg/vault"

	app "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/yaml"
)

// TestM4DApplicationController runs M4DApplicationReconciler.Reconcile() against a
// fake client that tracks a M4dApplication object.
// This test does not require a Kubernetes environment to run.
// This mechanism of testing can be used to test corner cases of the reconcile function.
func TestM4DApplicationControllerCSVCopyAndRead(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	// Set the logger to development mode for verbose logs.
	logf.SetLogger(zap.New(zap.UseDevMode(true)))

	var (
		name      = "notebook"
		namespace = "default"
	)

	application, err := readApplication("../../testdata/e2e/m4dapplication.yaml")
	g.Expect(err).To(gomega.BeNil(), "Cannot read m4dapplication file for test")
	application.Spec.Data[0].Requirements.Copy.Required = true // Require copy to enforce creation of copy module

	// Objects to track in the fake client.
	objs := []runtime.Object{
		application,
	}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	err = app.AddToScheme(scheme.Scheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create a fake client to mock API calls.
	cl := fake.NewFakeClientWithScheme(s, objs...)

	copyModule, err := readModule("../../../modules/implicit-copy-batch-module.yaml")
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create module in fake K8s agent
	err = cl.Create(context.Background(), copyModule)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	readModule, err := readModule("../../testdata/e2e/module-read.yaml")
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(readModule).NotTo(gomega.BeNil())

	readModule.Spec.Capabilities.SupportedInterfaces = readModule.Spec.Capabilities.SupportedInterfaces[1:] // remove parquet interface
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create module in fake K8s agent
	err = cl.Create(context.Background(), readModule)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create storage account
	accountShire := &app.M4DStorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "account1",
			Namespace: utils.GetSystemNamespace(),
		},
		Spec: app.M4DStorageAccountSpec{
			Endpoint:  "http://endpoint1",
			SecretRef: "dummy-secret",
			Regions:   []string{"theshire"},
		},
	}
	err = cl.Create(context.Background(), accountShire)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	dummySecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dummy-secret",
			Namespace: utils.GetSystemNamespace(),
		},
		Data: map[string][]byte{"accessKeyID": []byte("value1"), "secretAccessKey": []byte("value2")},
		Type: "Opaque",
	}
	err = cl.Create(context.Background(), dummySecret)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	dummyCatalog := DataCatalogDummy{
		credentials: make(map[string]pb.Credentials),
		dataDetails: make(map[string]modules.DataDetails),
	}

	dummyCatalog.credentials[application.Spec.Data[0].DataSetID] = pb.Credentials{ApiKey: "1234"}
	dummyCatalog.dataDetails[application.Spec.Data[0].DataSetID] = modules.DataDetails{
		Name: application.Spec.Data[0].DataSetID,
		Interface: app.InterfaceDetails{
			Protocol:   app.S3,
			DataFormat: app.CSV,
		},
		Geography:  "theshire",
		Connection: runtime.RawExtension{},
		Metadata: &pb.DatasetMetadata{
			DatasetNamedMetadata: nil,
			DatasetTags:          nil,
			ComponentsMetadata:   nil,
		},
	}

	// Create a M4DApplicationReconciler object with the scheme and fake client.
	r := &M4DApplicationReconciler{
		Client:          cl,
		Name:            "TestReconciler",
		Log:             ctrl.Log.WithName("test-controller"),
		Scheme:          s,
		VaultConnection: vault.NewDummyConnection(),
		PolicyCompiler:  &mockup.MockPolicyCompiler{},
		ResourceInterface: &PlotterInterface{
			Client: cl,
		},
		ClusterManager: &mockup.ClusterLister{},
		Provision:      &storage.ProvisionTest{},
		DataCatalog:    &dummyCatalog,
	}

	// Mock request to simulate Reconcile() being called on an event for a
	// watched resource .
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}
	res, err := r.Reconcile(req)
	g.Expect(err).To(gomega.BeNil())

	// Check the result of reconciliation to make sure it has the desired state.
	g.Expect(res.Requeue).To(gomega.BeFalse(), "reconcile did not requeue request as expected")

	// Check if Application generated a plotter
	err = cl.Get(context.TODO(), req.NamespacedName, application)
	g.Expect(err).To(gomega.BeNil(), "Can fetch plotter")
	g.Expect(application.Status.Generated.Kind).To(gomega.Equal("Plotter"))

	plotterObjectKey := types.NamespacedName{
		Namespace: "m4d-system",
		Name:      "notebook-default",
	}
	plotter := &app.Plotter{}
	err = cl.Get(context.Background(), plotterObjectKey, plotter)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	bpSpec := plotter.Spec.Blueprints["thegreendragon"]
	g.Expect(bpSpec.Flow.Steps[0].Template).To(gomega.Equal("implicit-copy-batch"))
	g.Expect(bpSpec.Flow.Steps[0].Arguments.Copy.Source.Format).To(gomega.Equal("csv"))
	g.Expect(bpSpec.Flow.Steps[0].Arguments.Copy.Destination.Format).To(gomega.Equal("csv"))
	g.Expect(bpSpec.Flow.Steps[0].Arguments.Copy.Destination.Format).To(gomega.Equal(bpSpec.Flow.Steps[1].Arguments.Read[0].Source.Format))
}

func readModule(f string) (*app.M4DModule, error) {
	moduleYAML, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	module := &app.M4DModule{}
	err = yaml.Unmarshal(moduleYAML, module)
	if err != nil {
		return nil, err
	}
	return module, nil
}

func readApplication(f string) (*app.M4DApplication, error) {
	applicationYAML, err := ioutil.ReadFile("../../testdata/e2e/m4dapplication.yaml")
	if err != nil {
		return nil, err
	}
	application := &app.M4DApplication{}
	err = yaml.Unmarshal(applicationYAML, application)
	if err != nil {
		return nil, err
	}
	return application, nil
}
