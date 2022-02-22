// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package controllers

// This is a collection of constants related to the manager and it's configuration

const ApplicationConcurrentReconcilesConfiguration = "APPLICATION_CONCURRENT_RECONCILES"
const PlotterConcurrentReconcilesConfiguration = "PLOTTER_CONCURRENT_RECONCILES"
const BlueprintConcurrentReconcilesConfiguration = "BLUEPRINT_CONCURRENT_RECONCILES"
const MaximumSecondsUntillReconcile = 60.0

const KubernetesClientQPSConfiguration = "CLIENT_QPS"
const KubernetesClientBurstConfiguration = "CLIENT_BURST"

const DefaultApplicationConcurrentReconciles = 1
const DefaultPlotterConcurrentReconciles = 1
const DefaultBlueprintConcurrentReconciles = 1

const DefaultKubernetesClientQPS = 5.0  // Default from Kubernetes client: 5
const DefaultKubernetesClientBurst = 10 // Default from Kubernetes client: 10

const ManagerPort = 9443
const ListeningPortAddress = 8085
