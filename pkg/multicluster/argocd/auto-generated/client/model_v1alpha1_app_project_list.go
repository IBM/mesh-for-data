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

// V1alpha1AppProjectList struct for V1alpha1AppProjectList
type V1alpha1AppProjectList struct {
	Items *[]V1alpha1AppProject `json:"items,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewV1alpha1AppProjectList instantiates a new V1alpha1AppProjectList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1AppProjectList() *V1alpha1AppProjectList {
	this := V1alpha1AppProjectList{}
	return &this
}

// NewV1alpha1AppProjectListWithDefaults instantiates a new V1alpha1AppProjectList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1AppProjectListWithDefaults() *V1alpha1AppProjectList {
	this := V1alpha1AppProjectList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *V1alpha1AppProjectList) GetItems() []V1alpha1AppProject {
	if o == nil || o.Items == nil {
		var ret []V1alpha1AppProject
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectList) GetItemsOk() (*[]V1alpha1AppProject, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *V1alpha1AppProjectList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []V1alpha1AppProject and assigns it to the Items field.
func (o *V1alpha1AppProjectList) SetItems(v []V1alpha1AppProject) {
	o.Items = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1AppProjectList) GetMetadata() V1ListMeta {
	if o == nil || o.Metadata == nil {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1AppProjectList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1AppProjectList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1alpha1AppProjectList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1alpha1AppProjectList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1AppProjectList struct {
	value *V1alpha1AppProjectList
	isSet bool
}

func (v NullableV1alpha1AppProjectList) Get() *V1alpha1AppProjectList {
	return v.value
}

func (v *NullableV1alpha1AppProjectList) Set(val *V1alpha1AppProjectList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1AppProjectList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1AppProjectList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1AppProjectList(val *V1alpha1AppProjectList) *NullableV1alpha1AppProjectList {
	return &NullableV1alpha1AppProjectList{value: val, isSet: true}
}

func (v NullableV1alpha1AppProjectList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1AppProjectList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


