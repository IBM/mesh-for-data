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

// ClusterGoogleAnalyticsConfig struct for ClusterGoogleAnalyticsConfig
type ClusterGoogleAnalyticsConfig struct {
	AnonymizeUsers *bool `json:"anonymizeUsers,omitempty"`
	TrackingID *string `json:"trackingID,omitempty"`
}

// NewClusterGoogleAnalyticsConfig instantiates a new ClusterGoogleAnalyticsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterGoogleAnalyticsConfig() *ClusterGoogleAnalyticsConfig {
	this := ClusterGoogleAnalyticsConfig{}
	return &this
}

// NewClusterGoogleAnalyticsConfigWithDefaults instantiates a new ClusterGoogleAnalyticsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterGoogleAnalyticsConfigWithDefaults() *ClusterGoogleAnalyticsConfig {
	this := ClusterGoogleAnalyticsConfig{}
	return &this
}

// GetAnonymizeUsers returns the AnonymizeUsers field value if set, zero value otherwise.
func (o *ClusterGoogleAnalyticsConfig) GetAnonymizeUsers() bool {
	if o == nil || o.AnonymizeUsers == nil {
		var ret bool
		return ret
	}
	return *o.AnonymizeUsers
}

// GetAnonymizeUsersOk returns a tuple with the AnonymizeUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterGoogleAnalyticsConfig) GetAnonymizeUsersOk() (*bool, bool) {
	if o == nil || o.AnonymizeUsers == nil {
		return nil, false
	}
	return o.AnonymizeUsers, true
}

// HasAnonymizeUsers returns a boolean if a field has been set.
func (o *ClusterGoogleAnalyticsConfig) HasAnonymizeUsers() bool {
	if o != nil && o.AnonymizeUsers != nil {
		return true
	}

	return false
}

// SetAnonymizeUsers gets a reference to the given bool and assigns it to the AnonymizeUsers field.
func (o *ClusterGoogleAnalyticsConfig) SetAnonymizeUsers(v bool) {
	o.AnonymizeUsers = &v
}

// GetTrackingID returns the TrackingID field value if set, zero value otherwise.
func (o *ClusterGoogleAnalyticsConfig) GetTrackingID() string {
	if o == nil || o.TrackingID == nil {
		var ret string
		return ret
	}
	return *o.TrackingID
}

// GetTrackingIDOk returns a tuple with the TrackingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterGoogleAnalyticsConfig) GetTrackingIDOk() (*string, bool) {
	if o == nil || o.TrackingID == nil {
		return nil, false
	}
	return o.TrackingID, true
}

// HasTrackingID returns a boolean if a field has been set.
func (o *ClusterGoogleAnalyticsConfig) HasTrackingID() bool {
	if o != nil && o.TrackingID != nil {
		return true
	}

	return false
}

// SetTrackingID gets a reference to the given string and assigns it to the TrackingID field.
func (o *ClusterGoogleAnalyticsConfig) SetTrackingID(v string) {
	o.TrackingID = &v
}

func (o ClusterGoogleAnalyticsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnonymizeUsers != nil {
		toSerialize["anonymizeUsers"] = o.AnonymizeUsers
	}
	if o.TrackingID != nil {
		toSerialize["trackingID"] = o.TrackingID
	}
	return json.Marshal(toSerialize)
}

type NullableClusterGoogleAnalyticsConfig struct {
	value *ClusterGoogleAnalyticsConfig
	isSet bool
}

func (v NullableClusterGoogleAnalyticsConfig) Get() *ClusterGoogleAnalyticsConfig {
	return v.value
}

func (v *NullableClusterGoogleAnalyticsConfig) Set(val *ClusterGoogleAnalyticsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterGoogleAnalyticsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterGoogleAnalyticsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterGoogleAnalyticsConfig(val *ClusterGoogleAnalyticsConfig) *NullableClusterGoogleAnalyticsConfig {
	return &NullableClusterGoogleAnalyticsConfig{value: val, isSet: true}
}

func (v NullableClusterGoogleAnalyticsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterGoogleAnalyticsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


