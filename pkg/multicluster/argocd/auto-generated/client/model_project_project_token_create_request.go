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

// ProjectProjectTokenCreateRequest ProjectTokenCreateRequest defines project token creation parameters.
type ProjectProjectTokenCreateRequest struct {
	Description *string `json:"description,omitempty"`
	ExpiresIn *string `json:"expiresIn,omitempty"`
	Id *string `json:"id,omitempty"`
	Project *string `json:"project,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NewProjectProjectTokenCreateRequest instantiates a new ProjectProjectTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectTokenCreateRequest() *ProjectProjectTokenCreateRequest {
	this := ProjectProjectTokenCreateRequest{}
	return &this
}

// NewProjectProjectTokenCreateRequestWithDefaults instantiates a new ProjectProjectTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectTokenCreateRequestWithDefaults() *ProjectProjectTokenCreateRequest {
	this := ProjectProjectTokenCreateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectProjectTokenCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectProjectTokenCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectProjectTokenCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *ProjectProjectTokenCreateRequest) GetExpiresIn() string {
	if o == nil || o.ExpiresIn == nil {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenCreateRequest) GetExpiresInOk() (*string, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *ProjectProjectTokenCreateRequest) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *ProjectProjectTokenCreateRequest) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectProjectTokenCreateRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenCreateRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectProjectTokenCreateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectProjectTokenCreateRequest) SetId(v string) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectProjectTokenCreateRequest) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenCreateRequest) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectProjectTokenCreateRequest) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ProjectProjectTokenCreateRequest) SetProject(v string) {
	o.Project = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ProjectProjectTokenCreateRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectTokenCreateRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ProjectProjectTokenCreateRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ProjectProjectTokenCreateRequest) SetRole(v string) {
	o.Role = &v
}

func (o ProjectProjectTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpiresIn != nil {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableProjectProjectTokenCreateRequest struct {
	value *ProjectProjectTokenCreateRequest
	isSet bool
}

func (v NullableProjectProjectTokenCreateRequest) Get() *ProjectProjectTokenCreateRequest {
	return v.value
}

func (v *NullableProjectProjectTokenCreateRequest) Set(val *ProjectProjectTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectTokenCreateRequest(val *ProjectProjectTokenCreateRequest) *NullableProjectProjectTokenCreateRequest {
	return &NullableProjectProjectTokenCreateRequest{value: val, isSet: true}
}

func (v NullableProjectProjectTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


