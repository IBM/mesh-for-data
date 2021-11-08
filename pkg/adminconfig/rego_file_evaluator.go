// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package adminconfig

import (
	"context"
	"fmt"
	"io/ioutil"

	"fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/utils"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

// Type definitions for parsing OPA response
// A list of decisions per capability, e.g. {"read": {"deploy": true}, "write": {"deploy": false}}
type RuleDecisionList []DecisionPerCapabilityMap

// A list of {capability : decision} per rule ("config")
type RuleExpressions map[string]RuleDecisionList

// Rule decisions per package
type PackageExpressions map[string]RuleExpressions

// A directory containing rego files that define admin config policies
const RegoPolicyDirectory = "/tmp/adminconfig/"

// RegoPolicyEvaluator implements EvaluatorInterface
// It evaluates policies from pkg/adminconfig/policies directory
type RegoPolicyEvaluator struct {
	Manager      *InfrastructureManager
	Data         *Infrastructure
	Query        rego.PreparedEvalQuery
	ReadyForEval bool
}

// NewRegoPolicyEvaluator constructs a new RegoPolicyEvaluator object
func NewRegoPolicyEvaluator() *RegoPolicyEvaluator {
	return &RegoPolicyEvaluator{
		Manager:      nil,
		Data:         nil,
		Query:        rego.PreparedEvalQuery{},
		ReadyForEval: false,
	}
}

// SetupWithInfrastructureManager connects the evaluator to the infrastructure manager for obtaining infrastructure details
func (r *RegoPolicyEvaluator) SetupWithInfrastructureManager(mgr *InfrastructureManager) {
	r.Manager = mgr
	r.Data = nil
	// get infrastructure details using a new manager
	if data, err := mgr.SetInfrastructure(); err != nil {
		r.Data = data
	}
}

// prepareQuery prepares a query for OPA evaluation - data object and compiled modules
// This function is called upon the change in either the infrastructure data or rego files
func (r *RegoPolicyEvaluator) prepareQuery() (rego.PreparedEvalQuery, error) {
	var err error
	if r.Data == nil {
		if r.Data, err = r.Manager.SetInfrastructure(); err != nil {
			return rego.PreparedEvalQuery{}, errors.Wrap(err, "could not set infrastructure")
		}
	}
	var json map[string]interface{}
	var bytes []byte
	if bytes, err = yaml.Marshal(r.Data); err != nil {
		return rego.PreparedEvalQuery{}, errors.Wrap(err, "couldn't marshall Data structure")
	}
	fmt.Println("Data\n" + string(bytes))
	if err = yaml.Unmarshal([]byte(bytes), &json); err != nil {
		return rego.PreparedEvalQuery{}, errors.Wrap(err, "couldn't unmarshall Data structure")
	}
	// Manually create the storage layer. inmem.NewFromObject returns an
	// in-memory store containing the supplied data.
	store := inmem.NewFromObject(json)

	// read and compile rego files
	files, err := ioutil.ReadDir(RegoPolicyDirectory)
	if err != nil {
		return rego.PreparedEvalQuery{}, err
	}
	modules := map[string]string{}
	for _, info := range files {
		name := info.Name()
		module, err := ioutil.ReadFile(RegoPolicyDirectory + name)
		if err != nil {
			return rego.PreparedEvalQuery{}, err
		}
		if err != nil {
			return rego.PreparedEvalQuery{}, err
		}
		modules[name] = string(module)
	}
	compiler, err := ast.CompileModules(modules)

	if err != nil {
		return rego.PreparedEvalQuery{}, errors.Wrap(err, "couldn't compile modules")
	}
	rego := rego.New(
		rego.Query("data.adminconfig"),
		rego.Store(store),
		rego.Compiler(compiler),
	)
	return rego.PrepareForEval(context.Background())
}

// Evaluate method evaluates the rego files based on the dynamic input object
func (r *RegoPolicyEvaluator) Evaluate(in *EvaluatorInput) (EvaluatorOutput, error) {
	if !r.ReadyForEval {
		var err error
		if r.Query, err = r.prepareQuery(); err != nil {
			return EvaluatorOutput{Valid: false}, errors.Wrap(err, "failed to prepare a query")
		}
		r.ReadyForEval = true
	}
	input, err := r.prepareInputForOPA(in)

	if err != nil {
		return EvaluatorOutput{Valid: false}, errors.Wrap(err, "failed to prepare an input for OPA")
	}
	// Run the evaluation with the new input
	rs, err := r.Query.Eval(context.Background(), rego.EvalInput(input))
	if err != nil {
		return EvaluatorOutput{Valid: false}, errors.Wrap(err, "failed to evaluate a query")
	}
	bytes, _ := yaml.Marshal(&rs)
	fmt.Println("Response: " + string(bytes))
	valid := true
	// merge decisions and build an output object for the manager
	decisions, err := r.getOPADecisions(rs)
	if err != nil {
		return EvaluatorOutput{Valid: valid, DatasetID: in.AssetRequirements.DatasetID, ConfigDecisions: decisions}, err
	}
	return EvaluatorOutput{Valid: valid, DatasetID: in.AssetRequirements.DatasetID, ConfigDecisions: decisions}, nil
}

// prepares an input in OPA format
func (r *RegoPolicyEvaluator) prepareInputForOPA(in *EvaluatorInput) (map[string]interface{}, error) {
	var input map[string]interface{}
	bytes, err := yaml.Marshal(in)
	if err != nil {
		return input, errors.Wrap(err, "failed to marshal the input structure")
	}
	fmt.Println("Input:\n" + string(bytes))
	err = yaml.Unmarshal(bytes, &input)
	return input, errors.Wrap(err, "failed  to unmarshal the input structure")
}

func (r *RegoPolicyEvaluator) initDecisions() map[v1alpha1.CapabilityType]Decision {
	return map[v1alpha1.CapabilityType]Decision{
		v1alpha1.Read:      DefaultDecision(r.Data),
		v1alpha1.Copy:      DefaultDecision(r.Data),
		v1alpha1.Transform: DefaultDecision(r.Data),
		v1alpha1.Write:     DefaultDecision(r.Data),
	}
}

// getOPADecisions parses the OPA decisions and merges decisions for the same capability
func (r *RegoPolicyEvaluator) getOPADecisions(rs rego.ResultSet) (DecisionPerCapabilityMap, error) {
	decisions := r.initDecisions()
	if len(rs) == 0 {
		return decisions, errors.New("invalid opa evaluation - an empty result set has been received")
	}
	defaultDecision := DefaultDecision(r.Data)
	for _, result := range rs {
		for _, expr := range result.Expressions {
			bytes, err := yaml.Marshal(expr.Value)
			if err != nil {
				return nil, err
			}
			exprStruct := PackageExpressions{}
			if err = yaml.Unmarshal(bytes, &exprStruct); err != nil {
				return nil, errors.Wrap(err, "Unexpected OPA response structure")
			}
			for _, packageRules := range exprStruct {
				for _, rule := range packageRules["config"] {
					for capability, newDecision := range rule {
						// apply defaults for undefined fields
						// string -> ConditionStatus conversion
						if newDecision.Deploy == "" {
							newDecision.Deploy = defaultDecision.Deploy
						} else if newDecision.Deploy == "true" {
							newDecision.Deploy = corev1.ConditionTrue
						} else if newDecision.Deploy == "false" {
							newDecision.Deploy = corev1.ConditionFalse
						}
						if len(newDecision.DeploymentRestrictions.Clusters) == 0 {
							newDecision.DeploymentRestrictions.Clusters = defaultDecision.DeploymentRestrictions.Clusters
						}
						// a single decision should be made for a capability
						valid, mergedDecision := r.merge(newDecision, decisions[capability])
						if !valid {
							return decisions, errors.New("Conflict while merging OPA decision " + newDecision.Policy.Description)
						}
						decisions[capability] = mergedDecision
					}
				}
			}
		}
	}
	return decisions, nil
}

// This function merges two decisions for the same capability using the following logic:
// deploy: true/false take precedence over undefined, true and false result in a conflict.
// cluster restrictions: the result of merge is an intersection of cluster sets from both decisions.
// module restrictions: new pairs <key, value> are added, if both exist - compatibility is checked.
// policy: concatenation of IDs and descriptions.
func (r *RegoPolicyEvaluator) merge(newDecision Decision, oldDecision Decision) (bool, Decision) {
	mergedDecision := Decision{}
	// merge deployment decisions
	deploy := oldDecision.Deploy
	if deploy == corev1.ConditionUnknown {
		deploy = newDecision.Deploy
	} else if newDecision.Deploy != corev1.ConditionUnknown {
		if newDecision.Deploy != deploy {
			return false, mergedDecision
		}
	}
	mergedDecision.Deploy = deploy
	// merge cluster restricitions
	mergedDecision.DeploymentRestrictions.Clusters = utils.Intersection(
		newDecision.DeploymentRestrictions.Clusters,
		oldDecision.DeploymentRestrictions.Clusters,
	)
	if len(mergedDecision.DeploymentRestrictions.Clusters) == 0 {
		return false, mergedDecision
	}
	// merge module restrictions
	mergedDecision.DeploymentRestrictions.ModuleRestrictions = oldDecision.DeploymentRestrictions.ModuleRestrictions
	for key, val := range newDecision.DeploymentRestrictions.ModuleRestrictions {
		if mergedDecision.DeploymentRestrictions.ModuleRestrictions[key] == "" {
			mergedDecision.DeploymentRestrictions.ModuleRestrictions[key] = val
		} else if mergedDecision.DeploymentRestrictions.ModuleRestrictions[key] != val {
			return false, mergedDecision
		}
	}
	// merge policies descriptions/ids
	mergedDecision.Policy = oldDecision.Policy
	if mergedDecision.Policy.ID != "" {
		mergedDecision.Policy.ID += ";"
	}
	mergedDecision.Policy.ID += newDecision.Policy.ID
	if mergedDecision.Policy.Description != "" {
		mergedDecision.Policy.Description += ";"
	}
	mergedDecision.Policy.Description += newDecision.Policy.Description
	return true, mergedDecision
}
