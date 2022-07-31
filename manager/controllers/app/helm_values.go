// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import api "fybrik.io/fybrik/manager/apis/app/v1alpha1"

// HelmValues are the values passed to modules during orchestration of the data plane
type HelmValues struct {
	// Asset specific arguments such as data stores and transformations
	api.ModuleArguments `json:",inline"`
	// Application details such as workload selector and user information
	*api.ApplicationDetails `json:",inline"`
	// Application and debug labels
	Labels map[string]string `json:"labels"`
	// Application unique identifier
	UUID string `json:"uuid"`
}
