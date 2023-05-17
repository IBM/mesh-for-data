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

// V1alpha1RepositoryCertificate struct for V1alpha1RepositoryCertificate
type V1alpha1RepositoryCertificate struct {
	CertData *string `json:"certData,omitempty"`
	CertInfo *string `json:"certInfo,omitempty"`
	CertSubType *string `json:"certSubType,omitempty"`
	CertType *string `json:"certType,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
}

// NewV1alpha1RepositoryCertificate instantiates a new V1alpha1RepositoryCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1RepositoryCertificate() *V1alpha1RepositoryCertificate {
	this := V1alpha1RepositoryCertificate{}
	return &this
}

// NewV1alpha1RepositoryCertificateWithDefaults instantiates a new V1alpha1RepositoryCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1RepositoryCertificateWithDefaults() *V1alpha1RepositoryCertificate {
	this := V1alpha1RepositoryCertificate{}
	return &this
}

// GetCertData returns the CertData field value if set, zero value otherwise.
func (o *V1alpha1RepositoryCertificate) GetCertData() string {
	if o == nil || o.CertData == nil {
		var ret string
		return ret
	}
	return *o.CertData
}

// GetCertDataOk returns a tuple with the CertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryCertificate) GetCertDataOk() (*string, bool) {
	if o == nil || o.CertData == nil {
		return nil, false
	}
	return o.CertData, true
}

// HasCertData returns a boolean if a field has been set.
func (o *V1alpha1RepositoryCertificate) HasCertData() bool {
	if o != nil && o.CertData != nil {
		return true
	}

	return false
}

// SetCertData gets a reference to the given string and assigns it to the CertData field.
func (o *V1alpha1RepositoryCertificate) SetCertData(v string) {
	o.CertData = &v
}

// GetCertInfo returns the CertInfo field value if set, zero value otherwise.
func (o *V1alpha1RepositoryCertificate) GetCertInfo() string {
	if o == nil || o.CertInfo == nil {
		var ret string
		return ret
	}
	return *o.CertInfo
}

// GetCertInfoOk returns a tuple with the CertInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryCertificate) GetCertInfoOk() (*string, bool) {
	if o == nil || o.CertInfo == nil {
		return nil, false
	}
	return o.CertInfo, true
}

// HasCertInfo returns a boolean if a field has been set.
func (o *V1alpha1RepositoryCertificate) HasCertInfo() bool {
	if o != nil && o.CertInfo != nil {
		return true
	}

	return false
}

// SetCertInfo gets a reference to the given string and assigns it to the CertInfo field.
func (o *V1alpha1RepositoryCertificate) SetCertInfo(v string) {
	o.CertInfo = &v
}

// GetCertSubType returns the CertSubType field value if set, zero value otherwise.
func (o *V1alpha1RepositoryCertificate) GetCertSubType() string {
	if o == nil || o.CertSubType == nil {
		var ret string
		return ret
	}
	return *o.CertSubType
}

// GetCertSubTypeOk returns a tuple with the CertSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryCertificate) GetCertSubTypeOk() (*string, bool) {
	if o == nil || o.CertSubType == nil {
		return nil, false
	}
	return o.CertSubType, true
}

// HasCertSubType returns a boolean if a field has been set.
func (o *V1alpha1RepositoryCertificate) HasCertSubType() bool {
	if o != nil && o.CertSubType != nil {
		return true
	}

	return false
}

// SetCertSubType gets a reference to the given string and assigns it to the CertSubType field.
func (o *V1alpha1RepositoryCertificate) SetCertSubType(v string) {
	o.CertSubType = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *V1alpha1RepositoryCertificate) GetCertType() string {
	if o == nil || o.CertType == nil {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryCertificate) GetCertTypeOk() (*string, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *V1alpha1RepositoryCertificate) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *V1alpha1RepositoryCertificate) SetCertType(v string) {
	o.CertType = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *V1alpha1RepositoryCertificate) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1RepositoryCertificate) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *V1alpha1RepositoryCertificate) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *V1alpha1RepositoryCertificate) SetServerName(v string) {
	o.ServerName = &v
}

func (o V1alpha1RepositoryCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertData != nil {
		toSerialize["certData"] = o.CertData
	}
	if o.CertInfo != nil {
		toSerialize["certInfo"] = o.CertInfo
	}
	if o.CertSubType != nil {
		toSerialize["certSubType"] = o.CertSubType
	}
	if o.CertType != nil {
		toSerialize["certType"] = o.CertType
	}
	if o.ServerName != nil {
		toSerialize["serverName"] = o.ServerName
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1RepositoryCertificate struct {
	value *V1alpha1RepositoryCertificate
	isSet bool
}

func (v NullableV1alpha1RepositoryCertificate) Get() *V1alpha1RepositoryCertificate {
	return v.value
}

func (v *NullableV1alpha1RepositoryCertificate) Set(val *V1alpha1RepositoryCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1RepositoryCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1RepositoryCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1RepositoryCertificate(val *V1alpha1RepositoryCertificate) *NullableV1alpha1RepositoryCertificate {
	return &NullableV1alpha1RepositoryCertificate{value: val, isSet: true}
}

func (v NullableV1alpha1RepositoryCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1RepositoryCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


