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

// V1EventSource EventSource contains information for an event.
type V1EventSource struct {
	Component *string `json:"component,omitempty"`
	Host *string `json:"host,omitempty"`
}

// NewV1EventSource instantiates a new V1EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EventSource() *V1EventSource {
	this := V1EventSource{}
	return &this
}

// NewV1EventSourceWithDefaults instantiates a new V1EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EventSourceWithDefaults() *V1EventSource {
	this := V1EventSource{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *V1EventSource) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EventSource) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *V1EventSource) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *V1EventSource) SetComponent(v string) {
	o.Component = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V1EventSource) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EventSource) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V1EventSource) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V1EventSource) SetHost(v string) {
	o.Host = &v
}

func (o V1EventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableV1EventSource struct {
	value *V1EventSource
	isSet bool
}

func (v NullableV1EventSource) Get() *V1EventSource {
	return v.value
}

func (v *NullableV1EventSource) Set(val *V1EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EventSource(val *V1EventSource) *NullableV1EventSource {
	return &NullableV1EventSource{value: val, isSet: true}
}

func (v NullableV1EventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


