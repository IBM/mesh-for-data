// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package local

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	app "fybrik.io/fybrik/manager/apis/app/v1beta1"
	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/multicluster"
)

// localClusterManager for local cluster configuration
type localClusterManager struct {
	Client client.Client
}

// GetClusters returns a list of registered clusters
func (cm *localClusterManager) GetClusters() ([]multicluster.Cluster, error) {
	clusters := []multicluster.Cluster{{
		Name: environment.GetLocalClusterName(),
		Metadata: multicluster.ClusterMetadata{
			Region:        environment.GetLocalRegion(),
			Zone:          environment.GetLocalZone(),
			VaultAuthPath: environment.GetLocalVaultAuthPath(),
		},
	}}
	return clusters, nil
}

func (cm *localClusterManager) IsMultiClusterSetup() bool {
	return false
}

// GetBlueprint returns a blueprint matching the given name, namespace and cluster details
func (cm *localClusterManager) GetBlueprint(cluster, namespace, name string) (*app.Blueprint, error) {
	if cluster != environment.GetLocalClusterName() {
		return nil, fmt.Errorf("unregistered cluster: %s", cluster)
	}
	blueprint := &app.Blueprint{}
	namespacedName := client.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}

	err := cm.Client.Get(context.Background(), namespacedName, blueprint)
	return blueprint, err
}

// CreateBlueprint creates a blueprint resource or updates an existing one
func (cm *localClusterManager) CreateBlueprint(cluster string, blueprint *app.Blueprint) error {
	return cm.UpdateBlueprint(cluster, blueprint)
}

// UpdateBlueprint updates the given blueprint or creates a new one if it does not exist
func (cm *localClusterManager) UpdateBlueprint(cluster string, blueprint *app.Blueprint) error {
	if cluster != environment.GetLocalClusterName() {
		return fmt.Errorf("unregistered cluster: %s", cluster) //nolint:revive // Ignore repetitive error msg
	}
	resource := &app.Blueprint{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blueprint.Name,
			Namespace: blueprint.Namespace,
		},
	}
	if _, err := ctrl.CreateOrUpdate(context.Background(), cm.Client, resource, func() error {
		resource.Spec = blueprint.Spec
		resource.ObjectMeta.Finalizers = blueprint.ObjectMeta.Finalizers
		resource.ObjectMeta.Labels = blueprint.ObjectMeta.Labels
		resource.ObjectMeta.Annotations = blueprint.ObjectMeta.Annotations
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteBlueprint deletes the blueprint resource
func (cm *localClusterManager) DeleteBlueprint(cluster, namespace, name string) error {
	blueprint, err := cm.GetBlueprint(cluster, namespace, name)
	if err != nil {
		return err
	}
	return cm.Client.Delete(context.Background(), blueprint)
}

// NewClusterManager creates an instance of ClusterManager for a local cluster configuration
func NewClusterManager(cl client.Client) (multicluster.ClusterManager, error) {
	return &localClusterManager{
		Client: cl,
	}, nil
}
