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

// ApplicationLinkInfo struct for ApplicationLinkInfo
type ApplicationLinkInfo struct {
	Description *string `json:"description,omitempty"`
	IconClass *string `json:"iconClass,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewApplicationLinkInfo instantiates a new ApplicationLinkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLinkInfo() *ApplicationLinkInfo {
	this := ApplicationLinkInfo{}
	return &this
}

// NewApplicationLinkInfoWithDefaults instantiates a new ApplicationLinkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLinkInfoWithDefaults() *ApplicationLinkInfo {
	this := ApplicationLinkInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationLinkInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationLinkInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationLinkInfo) SetDescription(v string) {
	o.Description = &v
}

// GetIconClass returns the IconClass field value if set, zero value otherwise.
func (o *ApplicationLinkInfo) GetIconClass() string {
	if o == nil || o.IconClass == nil {
		var ret string
		return ret
	}
	return *o.IconClass
}

// GetIconClassOk returns a tuple with the IconClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkInfo) GetIconClassOk() (*string, bool) {
	if o == nil || o.IconClass == nil {
		return nil, false
	}
	return o.IconClass, true
}

// HasIconClass returns a boolean if a field has been set.
func (o *ApplicationLinkInfo) HasIconClass() bool {
	if o != nil && o.IconClass != nil {
		return true
	}

	return false
}

// SetIconClass gets a reference to the given string and assigns it to the IconClass field.
func (o *ApplicationLinkInfo) SetIconClass(v string) {
	o.IconClass = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApplicationLinkInfo) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkInfo) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApplicationLinkInfo) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApplicationLinkInfo) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApplicationLinkInfo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkInfo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApplicationLinkInfo) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApplicationLinkInfo) SetUrl(v string) {
	o.Url = &v
}

func (o ApplicationLinkInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IconClass != nil {
		toSerialize["iconClass"] = o.IconClass
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationLinkInfo struct {
	value *ApplicationLinkInfo
	isSet bool
}

func (v NullableApplicationLinkInfo) Get() *ApplicationLinkInfo {
	return v.value
}

func (v *NullableApplicationLinkInfo) Set(val *ApplicationLinkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLinkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLinkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLinkInfo(val *ApplicationLinkInfo) *NullableApplicationLinkInfo {
	return &NullableApplicationLinkInfo{value: val, isSet: true}
}

func (v NullableApplicationLinkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLinkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


