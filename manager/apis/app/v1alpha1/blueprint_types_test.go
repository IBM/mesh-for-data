// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"testing"

	"context"

	"fybrik.io/fybrik/pkg/model/taxonomy"
	"fybrik.io/fybrik/pkg/serde"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestBlueprint(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &Blueprint{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: BlueprintSpec{
			Application: ApplicationDetails{
				UUID:    "123",
				Context: taxonomy.AppInfo{Properties: serde.Properties{Items: make(map[string]interface{})}},
			},
			Cluster: "cluster1",
			Modules: map[string]BlueprintModule{"start-instance1": {
				Name: "start",
				Arguments: ModuleArguments{
					Verbosity:  2,
					Capability: "start",
					Assets: []AssetContext{
						{
							AssetID: "test-asset",
						},
					},
				},
				Chart: ChartSpec{Name: "start-image"}}},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &Blueprint{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}
