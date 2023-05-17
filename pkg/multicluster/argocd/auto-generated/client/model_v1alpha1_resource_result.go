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

// V1alpha1ResourceResult struct for V1alpha1ResourceResult
type V1alpha1ResourceResult struct {
	Group *string `json:"group,omitempty"`
	// HookPhase contains the state of any operation associated with this resource OR hook This can also contain values for non-hook resources.
	HookPhase *string `json:"hookPhase,omitempty"`
	HookType *string `json:"hookType,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Message *string `json:"message,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Status *string `json:"status,omitempty"`
	SyncPhase *string `json:"syncPhase,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewV1alpha1ResourceResult instantiates a new V1alpha1ResourceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1ResourceResult() *V1alpha1ResourceResult {
	this := V1alpha1ResourceResult{}
	return &this
}

// NewV1alpha1ResourceResultWithDefaults instantiates a new V1alpha1ResourceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1ResourceResultWithDefaults() *V1alpha1ResourceResult {
	this := V1alpha1ResourceResult{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1alpha1ResourceResult) SetGroup(v string) {
	o.Group = &v
}

// GetHookPhase returns the HookPhase field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetHookPhase() string {
	if o == nil || o.HookPhase == nil {
		var ret string
		return ret
	}
	return *o.HookPhase
}

// GetHookPhaseOk returns a tuple with the HookPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetHookPhaseOk() (*string, bool) {
	if o == nil || o.HookPhase == nil {
		return nil, false
	}
	return o.HookPhase, true
}

// HasHookPhase returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasHookPhase() bool {
	if o != nil && o.HookPhase != nil {
		return true
	}

	return false
}

// SetHookPhase gets a reference to the given string and assigns it to the HookPhase field.
func (o *V1alpha1ResourceResult) SetHookPhase(v string) {
	o.HookPhase = &v
}

// GetHookType returns the HookType field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetHookType() string {
	if o == nil || o.HookType == nil {
		var ret string
		return ret
	}
	return *o.HookType
}

// GetHookTypeOk returns a tuple with the HookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetHookTypeOk() (*string, bool) {
	if o == nil || o.HookType == nil {
		return nil, false
	}
	return o.HookType, true
}

// HasHookType returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasHookType() bool {
	if o != nil && o.HookType != nil {
		return true
	}

	return false
}

// SetHookType gets a reference to the given string and assigns it to the HookType field.
func (o *V1alpha1ResourceResult) SetHookType(v string) {
	o.HookType = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1ResourceResult) SetKind(v string) {
	o.Kind = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1alpha1ResourceResult) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha1ResourceResult) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1alpha1ResourceResult) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V1alpha1ResourceResult) SetStatus(v string) {
	o.Status = &v
}

// GetSyncPhase returns the SyncPhase field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetSyncPhase() string {
	if o == nil || o.SyncPhase == nil {
		var ret string
		return ret
	}
	return *o.SyncPhase
}

// GetSyncPhaseOk returns a tuple with the SyncPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetSyncPhaseOk() (*string, bool) {
	if o == nil || o.SyncPhase == nil {
		return nil, false
	}
	return o.SyncPhase, true
}

// HasSyncPhase returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasSyncPhase() bool {
	if o != nil && o.SyncPhase != nil {
		return true
	}

	return false
}

// SetSyncPhase gets a reference to the given string and assigns it to the SyncPhase field.
func (o *V1alpha1ResourceResult) SetSyncPhase(v string) {
	o.SyncPhase = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1alpha1ResourceResult) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1ResourceResult) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1alpha1ResourceResult) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1alpha1ResourceResult) SetVersion(v string) {
	o.Version = &v
}

func (o V1alpha1ResourceResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.HookPhase != nil {
		toSerialize["hookPhase"] = o.HookPhase
	}
	if o.HookType != nil {
		toSerialize["hookType"] = o.HookType
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SyncPhase != nil {
		toSerialize["syncPhase"] = o.SyncPhase
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableV1alpha1ResourceResult struct {
	value *V1alpha1ResourceResult
	isSet bool
}

func (v NullableV1alpha1ResourceResult) Get() *V1alpha1ResourceResult {
	return v.value
}

func (v *NullableV1alpha1ResourceResult) Set(val *V1alpha1ResourceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1ResourceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1ResourceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1ResourceResult(val *V1alpha1ResourceResult) *NullableV1alpha1ResourceResult {
	return &NullableV1alpha1ResourceResult{value: val, isSet: true}
}

func (v NullableV1alpha1ResourceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1ResourceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


