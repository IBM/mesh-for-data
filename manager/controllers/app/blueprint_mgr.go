// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/app/modules"
	"fybrik.io/fybrik/manager/controllers/utils"
	// Temporary - shouldn't have something specific to implicit copies
)

// RefineInstances collects all instances of the same read/write module and creates a new instance instead, with accumulated arguments.
// Copy modules are left unchanged.
func (r *FybrikApplicationReconciler) RefineInstances(instances []modules.ModuleInstanceSpec) []modules.ModuleInstanceSpec {
	newInstances := make([]modules.ModuleInstanceSpec, 0)
	// map instances to be unified, according to the cluster and module
	instanceMap := make(map[string]modules.ModuleInstanceSpec)
	for _, moduleInstance := range instances {
		if moduleInstance.Args.Copy != nil {
			newInstances = append(newInstances, moduleInstance)
			continue
		}
		key := moduleInstance.Module.GetName() + "," + moduleInstance.ClusterName
		if instance, ok := instanceMap[key]; !ok {
			instanceMap[key] = moduleInstance
		} else {
			instance.Args.Read = append(instance.Args.Read, moduleInstance.Args.Read...)
			instance.Args.Write = append(instance.Args.Write, moduleInstance.Args.Write...)
			// AssetID is used for step name generation
			instance.AssetID += "," + moduleInstance.AssetID
			instanceMap[key] = instance
		}
	}
	for _, moduleInstance := range instanceMap {
		newInstances = append(newInstances, moduleInstance)
	}
	return newInstances
}

// GenerateBlueprints creates Blueprint specs (one per cluster)
func (r *FybrikApplicationReconciler) GenerateBlueprints(instances []modules.ModuleInstanceSpec, appContext *app.FybrikApplication) map[string]app.BlueprintSpec {
	blueprintMap := make(map[string]app.BlueprintSpec)
	instanceMap := make(map[string][]modules.ModuleInstanceSpec)
	for _, moduleInstance := range instances {
		instanceMap[moduleInstance.ClusterName] = append(instanceMap[moduleInstance.ClusterName], moduleInstance)
	}
	for key, instanceList := range instanceMap {
		// unite several instances of a read/write module
		instances := r.RefineInstances(instanceList)
		blueprintMap[key] = r.GenerateBlueprint(instances, appContext, key)
	}
	utils.PrintStructure(blueprintMap, r.Log, "BlueprintMap")
	return blueprintMap
}

// GenerateBlueprint creates the Blueprint spec based on the datasets and the governance actions required, which dictate the modules that must run in the fybrik
// Credentials for accessing data set are stored in a credential management system (such as vault) and the paths for accessing them are included in the blueprint.
// The credentials themselves are not included in the blueprint.
func (r *FybrikApplicationReconciler) GenerateBlueprint(instances []modules.ModuleInstanceSpec, appContext *app.FybrikApplication, clusterName string) app.BlueprintSpec {
	var spec app.BlueprintSpec

	spec.Cluster = clusterName

	// Create the list of BlueprintModules
	var blueprintModules []app.BlueprintModule
	for _, moduleInstance := range instances {
		modulename := moduleInstance.Module.GetName()

		var blueprintModule app.BlueprintModule
		blueprintModule.Name = modulename
		blueprintModule.InstanceName = utils.CreateStepName(modulename, moduleInstance.AssetID) // Need unique name for each module so include ids for dataset
		blueprintModule.Arguments = *moduleInstance.Args
		blueprintModule.Chart = moduleInstance.Module.Spec.Chart
		blueprintModules = append(blueprintModules, blueprintModule)
	}

	spec.Modules = blueprintModules
	return spec
}
