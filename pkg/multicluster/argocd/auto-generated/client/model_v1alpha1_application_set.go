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

// V1alpha1ApplicationSet struct for V1alpha1ApplicationSet
type V1alpha1ApplicationSet struct {
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec *V1alpha1ApplicationSetSpec `json:"spec,omitempty"`
	Status *V1alpha1ApplicationSetStatus `json:"status,omitempty"`
}

// NewV1alpha1ApplicationSet instantiates a new V1alpha1ApplicationSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ApplicationSet() *V1alpha1ApplicationSet {
	this := V1alpha1ApplicationSet{}
	return &this
}

// NewV1alpha1ApplicationSetWithDefaults instantiates a new V1alpha1ApplicationSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ApplicationSetWithDefaults() *V1alpha1ApplicationSet {
	this := V1alpha1ApplicationSet{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSet) GetMetadata() V1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSet) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSet) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1alpha1ApplicationSet) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSet) GetSpec() V1alpha1ApplicationSetSpec {
	if o == nil || o.Spec == nil {
		var ret V1alpha1ApplicationSetSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSet) GetSpecOk() (*V1alpha1ApplicationSetSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSet) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given V1alpha1ApplicationSetSpec and assigns it to the Spec field.
func (o *V1alpha1ApplicationSet) SetSpec(v V1alpha1ApplicationSetSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha1ApplicationSet) GetStatus() V1alpha1ApplicationSetStatus {
	if o == nil || o.Status == nil {
		var ret V1alpha1ApplicationSetStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ApplicationSet) GetStatusOk() (*V1alpha1ApplicationSetStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha1ApplicationSet) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1alpha1ApplicationSetStatus and assigns it to the Status field.
func (o *V1alpha1ApplicationSet) SetStatus(v V1alpha1ApplicationSetStatus) {
	o.Status = &v
}

func (o V1alpha1ApplicationSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ApplicationSet struct {
	value *V1alpha1ApplicationSet
	isSet bool
}

func (v NullableV1alpha1ApplicationSet) Get() *V1alpha1ApplicationSet {
	return v.value
}

func (v *NullableV1alpha1ApplicationSet) Set(val *V1alpha1ApplicationSet) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ApplicationSet) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ApplicationSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ApplicationSet(val *V1alpha1ApplicationSet) *NullableV1alpha1ApplicationSet {
	return &NullableV1alpha1ApplicationSet{value: val, isSet: true}
}

func (v NullableV1alpha1ApplicationSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ApplicationSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


