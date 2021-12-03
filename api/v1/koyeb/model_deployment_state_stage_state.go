/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// DeploymentStateStageState struct for DeploymentStateStageState
type DeploymentStateStageState struct {
	Name *string `json:"name,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	Status *DeploymentStateStageStateStatus `json:"status,omitempty"`
	Messages *[]string `json:"messages,omitempty"`
}

// NewDeploymentStateStageState instantiates a new DeploymentStateStageState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateStageState() *DeploymentStateStageState {
	this := DeploymentStateStageState{}
	var status DeploymentStateStageStateStatus = DEPLOYMENTSTATESTAGESTATESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewDeploymentStateStageStateWithDefaults instantiates a new DeploymentStateStageState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateStageStateWithDefaults() *DeploymentStateStageState {
	this := DeploymentStateStageState{}
	var status DeploymentStateStageStateStatus = DEPLOYMENTSTATESTAGESTATESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateStageState) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateStageState) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateStageState) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateStageState) SetName(v string) {
	o.Name = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *DeploymentStateStageState) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateStageState) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *DeploymentStateStageState) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *DeploymentStateStageState) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentStateStageState) GetStatus() DeploymentStateStageStateStatus {
	if o == nil || o.Status == nil {
		var ret DeploymentStateStageStateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateStageState) GetStatusOk() (*DeploymentStateStageStateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentStateStageState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentStateStageStateStatus and assigns it to the Status field.
func (o *DeploymentStateStageState) SetStatus(v DeploymentStateStageStateStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DeploymentStateStageState) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateStageState) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DeploymentStateStageState) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *DeploymentStateStageState) SetMessages(v []string) {
	o.Messages = &v
}

func (o DeploymentStateStageState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateStageState struct {
	value *DeploymentStateStageState
	isSet bool
}

func (v NullableDeploymentStateStageState) Get() *DeploymentStateStageState {
	return v.value
}

func (v *NullableDeploymentStateStageState) Set(val *DeploymentStateStageState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateStageState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateStageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateStageState(val *DeploymentStateStageState) *NullableDeploymentStateStageState {
	return &NullableDeploymentStateStageState{value: val, isSet: true}
}

func (v NullableDeploymentStateStageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateStageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


