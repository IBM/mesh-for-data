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

// SessionGetUserInfoResponse struct for SessionGetUserInfoResponse
type SessionGetUserInfoResponse struct {
	Groups *[]string `json:"groups,omitempty"`
	Iss *string `json:"iss,omitempty"`
	LoggedIn *bool `json:"loggedIn,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewSessionGetUserInfoResponse instantiates a new SessionGetUserInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionGetUserInfoResponse() *SessionGetUserInfoResponse {
	this := SessionGetUserInfoResponse{}
	return &this
}

// NewSessionGetUserInfoResponseWithDefaults instantiates a new SessionGetUserInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionGetUserInfoResponseWithDefaults() *SessionGetUserInfoResponse {
	this := SessionGetUserInfoResponse{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *SessionGetUserInfoResponse) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionGetUserInfoResponse) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SessionGetUserInfoResponse) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *SessionGetUserInfoResponse) SetGroups(v []string) {
	o.Groups = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *SessionGetUserInfoResponse) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionGetUserInfoResponse) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *SessionGetUserInfoResponse) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *SessionGetUserInfoResponse) SetIss(v string) {
	o.Iss = &v
}

// GetLoggedIn returns the LoggedIn field value if set, zero value otherwise.
func (o *SessionGetUserInfoResponse) GetLoggedIn() bool {
	if o == nil || o.LoggedIn == nil {
		var ret bool
		return ret
	}
	return *o.LoggedIn
}

// GetLoggedInOk returns a tuple with the LoggedIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionGetUserInfoResponse) GetLoggedInOk() (*bool, bool) {
	if o == nil || o.LoggedIn == nil {
		return nil, false
	}
	return o.LoggedIn, true
}

// HasLoggedIn returns a boolean if a field has been set.
func (o *SessionGetUserInfoResponse) HasLoggedIn() bool {
	if o != nil && o.LoggedIn != nil {
		return true
	}

	return false
}

// SetLoggedIn gets a reference to the given bool and assigns it to the LoggedIn field.
func (o *SessionGetUserInfoResponse) SetLoggedIn(v bool) {
	o.LoggedIn = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SessionGetUserInfoResponse) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionGetUserInfoResponse) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SessionGetUserInfoResponse) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SessionGetUserInfoResponse) SetUsername(v string) {
	o.Username = &v
}

func (o SessionGetUserInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.LoggedIn != nil {
		toSerialize["loggedIn"] = o.LoggedIn
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableSessionGetUserInfoResponse struct {
	value *SessionGetUserInfoResponse
	isSet bool
}

func (v NullableSessionGetUserInfoResponse) Get() *SessionGetUserInfoResponse {
	return v.value
}

func (v *NullableSessionGetUserInfoResponse) Set(val *SessionGetUserInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionGetUserInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionGetUserInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionGetUserInfoResponse(val *SessionGetUserInfoResponse) *NullableSessionGetUserInfoResponse {
	return &NullableSessionGetUserInfoResponse{value: val, isSet: true}
}

func (v NullableSessionGetUserInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionGetUserInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


