// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package datapath

import (
	fappv1 "fybrik.io/fybrik/manager/apis/app/v1beta1"
	fappv2 "fybrik.io/fybrik/manager/apis/app/v1beta2"
	"fybrik.io/fybrik/pkg/adminconfig"
	"fybrik.io/fybrik/pkg/infrastructure"
	"fybrik.io/fybrik/pkg/model/datacatalog"
	"fybrik.io/fybrik/pkg/model/taxonomy"
	"fybrik.io/fybrik/pkg/multicluster"
)

// DataInfo defines all the information about the given data set that comes from the fybrikapplication spec and from the connectors.
type DataInfo struct {
	// Source connection details
	DataDetails *datacatalog.GetAssetResponse
	// Pointer to the relevant data context in the Fybrik application spec
	Context *fappv1.DataContext
	// Evaluated config policies
	Configuration adminconfig.EvaluatorOutput
	// Workload cluster
	WorkloadCluster multicluster.Cluster
	// Required governance actions to perform on this asset
	Actions []taxonomy.Action
	// Potential actions to be taken on storing this asset in a specific location
	StorageRequirements map[taxonomy.ProcessingLocation][]taxonomy.Action
}

// Environment defines the available resources (clusters, modules, storageAccounts)
// It also contains the results of queries to policy manager regarding writing data to storage accounts
type Environment struct {
	Modules          map[string]*fappv1.FybrikModule
	Clusters         []multicluster.Cluster
	StorageAccounts  []*fappv2.FybrikStorageAccount
	AttributeManager *infrastructure.AttributeManager
}
