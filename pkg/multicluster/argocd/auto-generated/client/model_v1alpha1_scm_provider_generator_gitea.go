/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// V1alpha1SCMProviderGeneratorGitea SCMProviderGeneratorGitea defines a connection info specific to Gitea.
type V1alpha1SCMProviderGeneratorGitea struct {
	// Scan all branches instead of just the default branch.
	AllBranches *bool `json:"allBranches,omitempty"`
	// The Gitea URL to talk to. For example https://gitea.mydomain.com/.
	Api *string `json:"api,omitempty"`
	Insecure *bool `json:"insecure,omitempty"`
	// Gitea organization or user to scan. Required.
	Owner *string `json:"owner,omitempty"`
	TokenRef *V1alpha1SecretRef `json:"tokenRef,omitempty"`
}

// NewV1alpha1SCMProviderGeneratorGitea instantiates a new V1alpha1SCMProviderGeneratorGitea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SCMProviderGeneratorGitea() *V1alpha1SCMProviderGeneratorGitea {
	this := V1alpha1SCMProviderGeneratorGitea{}
	return &this
}

// NewV1alpha1SCMProviderGeneratorGiteaWithDefaults instantiates a new V1alpha1SCMProviderGeneratorGitea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SCMProviderGeneratorGiteaWithDefaults() *V1alpha1SCMProviderGeneratorGitea {
	this := V1alpha1SCMProviderGeneratorGitea{}
	return &this
}

// GetAllBranches returns the AllBranches field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGitea) GetAllBranches() bool {
	if o == nil || o.AllBranches == nil {
		var ret bool
		return ret
	}
	return *o.AllBranches
}

// GetAllBranchesOk returns a tuple with the AllBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) GetAllBranchesOk() (*bool, bool) {
	if o == nil || o.AllBranches == nil {
		return nil, false
	}
	return o.AllBranches, true
}

// HasAllBranches returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) HasAllBranches() bool {
	if o != nil && o.AllBranches != nil {
		return true
	}

	return false
}

// SetAllBranches gets a reference to the given bool and assigns it to the AllBranches field.
func (o *V1alpha1SCMProviderGeneratorGitea) SetAllBranches(v bool) {
	o.AllBranches = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGitea) GetApi() string {
	if o == nil || o.Api == nil {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) GetApiOk() (*string, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *V1alpha1SCMProviderGeneratorGitea) SetApi(v string) {
	o.Api = &v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGitea) GetInsecure() bool {
	if o == nil || o.Insecure == nil {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) GetInsecureOk() (*bool, bool) {
	if o == nil || o.Insecure == nil {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) HasInsecure() bool {
	if o != nil && o.Insecure != nil {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *V1alpha1SCMProviderGeneratorGitea) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGitea) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *V1alpha1SCMProviderGeneratorGitea) SetOwner(v string) {
	o.Owner = &v
}

// GetTokenRef returns the TokenRef field value if set, zero value otherwise.
func (o *V1alpha1SCMProviderGeneratorGitea) GetTokenRef() V1alpha1SecretRef {
	if o == nil || o.TokenRef == nil {
		var ret V1alpha1SecretRef
		return ret
	}
	return *o.TokenRef
}

// GetTokenRefOk returns a tuple with the TokenRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) GetTokenRefOk() (*V1alpha1SecretRef, bool) {
	if o == nil || o.TokenRef == nil {
		return nil, false
	}
	return o.TokenRef, true
}

// HasTokenRef returns a boolean if a field has been set.
func (o *V1alpha1SCMProviderGeneratorGitea) HasTokenRef() bool {
	if o != nil && o.TokenRef != nil {
		return true
	}

	return false
}

// SetTokenRef gets a reference to the given V1alpha1SecretRef and assigns it to the TokenRef field.
func (o *V1alpha1SCMProviderGeneratorGitea) SetTokenRef(v V1alpha1SecretRef) {
	o.TokenRef = &v
}

func (o V1alpha1SCMProviderGeneratorGitea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllBranches != nil {
		toSerialize["allBranches"] = o.AllBranches
	}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	if o.Insecure != nil {
		toSerialize["insecure"] = o.Insecure
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.TokenRef != nil {
		toSerialize["tokenRef"] = o.TokenRef
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1SCMProviderGeneratorGitea struct {
	value *V1alpha1SCMProviderGeneratorGitea
	isSet bool
}

func (v NullableV1alpha1SCMProviderGeneratorGitea) Get() *V1alpha1SCMProviderGeneratorGitea {
	return v.value
}

func (v *NullableV1alpha1SCMProviderGeneratorGitea) Set(val *V1alpha1SCMProviderGeneratorGitea) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SCMProviderGeneratorGitea) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SCMProviderGeneratorGitea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SCMProviderGeneratorGitea(val *V1alpha1SCMProviderGeneratorGitea) *NullableV1alpha1SCMProviderGeneratorGitea {
	return &NullableV1alpha1SCMProviderGeneratorGitea{value: val, isSet: true}
}

func (v NullableV1alpha1SCMProviderGeneratorGitea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SCMProviderGeneratorGitea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


