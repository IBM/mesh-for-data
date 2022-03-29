# Configuration Policies

## What are configuration policies?

Configuration policies are the mechanism via which the organization may influence the construction of the data plane, taking into account infrastructure capabilities and costs. Fybrik takes into account the workload context, the data metadata, the data governance policies and the configuration policies when defining the data plane. The configuration policies influence what capabilities should be deployed (e.g. read, copy), in which clusters they should be deployed, and selection of the most appropriate module that implements the capability.

## Input to policies

The `input` object includes general application data such as workload cluster and application properties, as well as dataset details (user requirements, metadata).

Available properties:
- `cluster.name`: name of the workload cluster
- `cluster.metadata.region`: region of the workload cluster
- `properties`: application/workload properties defined in FybrikApplication, e.g. `properties.intent`
- `request.metadata`: asset metadata as defined in catalog taxonomy, e.g `request.metadata.geography`
- `usage`: a set of boolean properties associated with data use: `usage.read`, `usage.write`, `usage.copy`

## Syntax 

Policies are written in rego files. Each file declares a package `adminconfig`.

Rules are written in the following syntax: `config[{"capability": capability, "decision": decision}]` where

`capability` represents a required module capability, such as "read", "write", "transform" and "copy".

`decision` is a JSON structure that matches `Decision` defined above. 

```
{ 
	"policy": {"ID": <id>, "policySetID": <setId>, "description": <description>, "version": <version>}, 
	"deploy": <"True", "False">,
	"restrictions": {
		"modules": <list of restrictions>,
		"clusters": <list of restrictions>,
        "storageaccounts": <list of restrictions>,
	},
}
```
`restriction` restricts a `property` to either a set of `values` or a value in a given `range`.

For example, the policy above restricts the choice of clusters and modules for a read capability by narrowing the choice of deployment clusters to the workload cluster, and restricting the module type to service.

```
config[{"capability": "read", "decision": decision}] {
    input.request.usage == "read"
    policy := {"ID": "read-location", "description":"Deploy read in the workload cluster", "version": "0.1"}
    cluster_restrict := {"property": "name", "values": [ input.workload.cluster.name ] }
    module_restrict := {"property": "type", "values": ["service"]}
    decision := {"policy": policy, "restrictions": {"clusters": [cluster_restrict], "modules": [module_restrict]}}
}
```


`policy` provides policy metadata: unique ID, human-readable description, version and `policySetID` (see ### Policy Set ID)

`restrictions` provides restrictions for `modules`, `clusters` and `storageaccounts`.
Each restriction provides a list or a range of allowed values for a property of module/cluster/storageaccount object. For example, to restrict a module type to either "service" or "plugin", we'll use "type" as a property, and [ "service","plugin ] as a list of allowed values.
Properties of a module can be found inside [`FybrikModule`](../reference/crds.md#fybrikmodule) Spec.
Properties of a storage account are listed inside [`FybrikStorageAccount`](../reference/crds.md#fybrikstorageaccount).
Cluster is not a custom resource. It has the following properties:
- name: cluster name
- metadata.region: cluster region
- metadata.zone: cluster zone

`deploy` receives "True"/"False" values. These values indicate whether the capability should or should not be deployed. If not specified in the policy, it's up to Fybrik to decide on the capability deployment.

### Policy Set ID

Fybrik supports evaluating different sets of policies for different FybrikApplications. It is possible to define a policy for a specific `policySetID` which will be trigered only if it matches the `policySetID` defined in FybrikApplication. 
If a policy does not specify a policy set id, it will be considered as relevant for all FybrikApplications.
In a similar way, all policies are relevant for a FybrikApplication that does not specify a policy set id, to support a use-case of a single policy set for all.

### Out of the box policies

Out of the box policies come with the fybrik deployment. They define the deployment of basic capabilities, such as read, write, copy and delete. 

```
package adminconfig

# read capability deployment
config[{"capability": "read", "decision": decision}] {
    input.request.usage == "read"
    policy := {"ID": "read-default-enabled", "description":"Read capability is requested for read workloads", "version": "0.1"}
    decision := {"policy": policy, "deploy": "True"}
}

# write capability deployment
config[{"capability": "write", "decision": decision}] {
    input.request.usage == "write"
    policy := {"ID": "write-default-enabled", "description":"Write capability is requested for workloads that write data", "version": "0.1"}
    decision := {"policy": policy, "deploy": "True"}
}

# copy requested by the user
config[{"capability": "copy", "decision": decision}] {
    input.request.usage == "copy"
    policy := {"ID": "copy-request", "description":"Copy (ingest) capability is requested by the user", "version": "0.1"}
    decision := {"policy": policy, "deploy": "True"}
}

# delete capability deployment
config[{"capability": "delete", "decision": decision}] {
    input.request.usage == "delete"
    policy := {"ID": "delete-request", "description":"Delete capability is requested by the user", "version": "0.1"}
    decision := {"policy": policy, "deploy": "True"}
}

# do not deploy copy in scenarios different from read or copy
config[{"capability": "copy", "decision": decision}] {
    input.request.usage != "read"
    input.request.usage != "copy"
    policy := {"ID": "copy-disabled", "description":"Copy capability is not requested", "version": "0.1"}
    decision := {"policy": policy, "deploy": "False"}
}

# do not deploy read in other scenarios
config[{"capability": "read", "decision": decision}] {
    input.request.usage != "read"
    policy := {"ID": "read-disabled", "description":"Read capability is not requested", "version": "0.1"}
    decision := {"policy": policy, "deploy": "False"}
}

# do not deploy write in other scenarios
config[{"capability": "write", "decision": decision}] {
    input.request.usage != "write"
    policy := {"ID": "write-disabled", "description":"Write capability is not requested", "version": "0.1"}
    decision := {"policy": policy, "deploy": "False"}
}

# do not deploy delete in other scenarios
config[{"capability": "delete", "decision": decision}] {
    input.request.usage != "delete"
    policy := {"ID": "delete-disabled", "description":"Delete capability is not requested", "version": "0.1"}
    decision := {"policy": policy, "deploy": "False"}
}

```

### Extended policies

The extended policies define advanced deployment requirements, such as where read or transform modules should run, what should be the scope of module deployments, and more. 

The policies below are provided as a sample and can be updated for the production deployment.

```
package adminconfig

# configure where transformations take place
config[{"capability": "transform", "decision": decision}] {
    policy := {"ID": "transform-geo", "description":"Governance based transformations must take place in the geography where the data is stored", "version": "0.1"}
    cluster_restrict := {"property": "metadata.region", "values": [input.request.dataset.geography]}
    decision := {"policy": policy, "restrictions": {"clusters": [cluster_restrict]}}
}

# configure the scope of the read capability
config[{"capability": "read", "decision": decision}] {
    input.request.usage == "read"
    policy := {"ID": "read-scope", "description":"Deploy read at the workload scope", "version": "0.1"}
    decision := {"policy": policy, "restrictions": {"modules": [{"property": "capabilities.scope", "values" : ["workload"]}]}}
}

# configure where the read capability will be deployed
config[{"capability": "read", "decision": decision}] {
    input.request.usage == "read"
    policy := {"ID": "read-location", "description":"Deploy read in the workload cluster", "version": "0.1"}
    cluster_restrict := {"property": "name", "values": [ input.workload.cluster.name ] }
    decision := {"policy": policy, "restrictions": {"clusters": [cluster_restrict]}}
}

# allow implicit copies by default
config[{"capability": "copy", "decision": decision}] {
    input.request.usage == "read"
    policy := {"ID": "copy-default", "description":"Implicit copies are allowed in read scenarios", "version": "0.1"}
    decision := {"policy": policy}
}

```

### How to provide custom policies

In order to deploy Fybrik with customized policies, perform the following steps 

1. Clone the github repository of Fybrik for the required release: `git clone -b releases/<version> https://github.com/fybrik/fybrik.git`
2. Copy the rego files containing customized policies to fybrik/charts/fybrik/files/adminconfig/ folder 
3. Install Fybrik:
```
cd fybrik
helm install fybrik-crd charts/fybrik-crd -n fybrik-system --wait
helm install fybrik charts/fybrik --set global.tag=master --set global.imagePullPolicy=Always -n fybrik-system --wait
```
