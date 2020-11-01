// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	appapi "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/ibm/the-mesh-for-data/manager/controllers/mockup"
	"github.com/ibm/the-mesh-for-data/pkg/helm"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var mgr ctrl.Manager
var testEnv *envtest.Environment

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "config", "crd", "bases"),
		},
		AttachControlPlaneOutput: true,
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	// Mockup connectors
	go mockup.MockCatalogConnector()

	err = appapi.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	if os.Getenv("USE_EXISTING_CONTROLLER") == "true" {
		logf.Log.Info("Using existing controller in existing cluster...")
		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	} else {
		mgr, err = ctrl.NewManager(cfg, ctrl.Options{
			Scheme:             scheme.Scheme,
			MetricsBindAddress: "localhost:8086",
		})
		Expect(err).ToNot(HaveOccurred())

		policyCompiler := &mockup.MockPolicyCompiler{}
		err = NewM4DApplicationReconciler(mgr, "M4DApplication", nil, policyCompiler).SetupWithManager(mgr)
		Expect(err).ToNot(HaveOccurred())
		err = NewBlueprintReconciler(mgr, "Blueprint", new(helm.Fake)).SetupWithManager(mgr)
		Expect(err).ToNot(HaveOccurred())

		go func() {
			err = mgr.Start(ctrl.SetupSignalHandler())
			Expect(err).ToNot(HaveOccurred())
		}()

		k8sClient = mgr.GetClient()
	}
	Expect(k8sClient).ToNot(BeNil())

	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	gexec.KillAndWait(5 * time.Second)
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
