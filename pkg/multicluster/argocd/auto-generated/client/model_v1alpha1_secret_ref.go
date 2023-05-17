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

// V1alpha1SecretRef Utility struct for a reference to a secret key.
type V1alpha1SecretRef struct {
	Key *string `json:"key,omitempty"`
	SecretName *string `json:"secretName,omitempty"`
}

// NewV1alpha1SecretRef instantiates a new V1alpha1SecretRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1SecretRef() *V1alpha1SecretRef {
	this := V1alpha1SecretRef{}
	return &this
}

// NewV1alpha1SecretRefWithDefaults instantiates a new V1alpha1SecretRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1SecretRefWithDefaults() *V1alpha1SecretRef {
	this := V1alpha1SecretRef{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *V1alpha1SecretRef) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SecretRef) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *V1alpha1SecretRef) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *V1alpha1SecretRef) SetKey(v string) {
	o.Key = &v
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *V1alpha1SecretRef) GetSecretName() string {
	if o == nil || o.SecretName == nil {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1SecretRef) GetSecretNameOk() (*string, bool) {
	if o == nil || o.SecretName == nil {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *V1alpha1SecretRef) HasSecretName() bool {
	if o != nil && o.SecretName != nil {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the SecretName field.
func (o *V1alpha1SecretRef) SetSecretName(v string) {
	o.SecretName = &v
}

func (o V1alpha1SecretRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.SecretName != nil {
		toSerialize["secretName"] = o.SecretName
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1SecretRef struct {
	value *V1alpha1SecretRef
	isSet bool
}

func (v NullableV1alpha1SecretRef) Get() *V1alpha1SecretRef {
	return v.value
}

func (v *NullableV1alpha1SecretRef) Set(val *V1alpha1SecretRef) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1SecretRef) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1SecretRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1SecretRef(val *V1alpha1SecretRef) *NullableV1alpha1SecretRef {
	return &NullableV1alpha1SecretRef{value: val, isSet: true}
}

func (v NullableV1alpha1SecretRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1SecretRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


