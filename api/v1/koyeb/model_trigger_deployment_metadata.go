/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// TriggerDeploymentMetadata struct for TriggerDeploymentMetadata
type TriggerDeploymentMetadata struct {
	Type *TriggerDeploymentMetadataTriggerType `json:"type,omitempty"`
	Actor *TriggerDeploymentMetadataActorType `json:"actor,omitempty"`
	Git *TriggerGitDeploymentMetadata `json:"git,omitempty"`
}

// NewTriggerDeploymentMetadata instantiates a new TriggerDeploymentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerDeploymentMetadata() *TriggerDeploymentMetadata {
	this := TriggerDeploymentMetadata{}
	var type_ TriggerDeploymentMetadataTriggerType = TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN_TYPE
	this.Type = &type_
	var actor TriggerDeploymentMetadataActorType = TRIGGERDEPLOYMENTMETADATAACTORTYPE_UNKNOWN_ACTOR
	this.Actor = &actor
	return &this
}

// NewTriggerDeploymentMetadataWithDefaults instantiates a new TriggerDeploymentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerDeploymentMetadataWithDefaults() *TriggerDeploymentMetadata {
	this := TriggerDeploymentMetadata{}
	var type_ TriggerDeploymentMetadataTriggerType = TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN_TYPE
	this.Type = &type_
	var actor TriggerDeploymentMetadataActorType = TRIGGERDEPLOYMENTMETADATAACTORTYPE_UNKNOWN_ACTOR
	this.Actor = &actor
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetType() TriggerDeploymentMetadataTriggerType {
	if o == nil || isNil(o.Type) {
		var ret TriggerDeploymentMetadataTriggerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetTypeOk() (*TriggerDeploymentMetadataTriggerType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TriggerDeploymentMetadataTriggerType and assigns it to the Type field.
func (o *TriggerDeploymentMetadata) SetType(v TriggerDeploymentMetadataTriggerType) {
	o.Type = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetActor() TriggerDeploymentMetadataActorType {
	if o == nil || isNil(o.Actor) {
		var ret TriggerDeploymentMetadataActorType
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetActorOk() (*TriggerDeploymentMetadataActorType, bool) {
	if o == nil || isNil(o.Actor) {
    return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasActor() bool {
	if o != nil && !isNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given TriggerDeploymentMetadataActorType and assigns it to the Actor field.
func (o *TriggerDeploymentMetadata) SetActor(v TriggerDeploymentMetadataActorType) {
	o.Actor = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetGit() TriggerGitDeploymentMetadata {
	if o == nil || isNil(o.Git) {
		var ret TriggerGitDeploymentMetadata
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetGitOk() (*TriggerGitDeploymentMetadata, bool) {
	if o == nil || isNil(o.Git) {
    return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given TriggerGitDeploymentMetadata and assigns it to the Git field.
func (o *TriggerDeploymentMetadata) SetGit(v TriggerGitDeploymentMetadata) {
	o.Git = &v
}

func (o TriggerDeploymentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerDeploymentMetadata struct {
	value *TriggerDeploymentMetadata
	isSet bool
}

func (v NullableTriggerDeploymentMetadata) Get() *TriggerDeploymentMetadata {
	return v.value
}

func (v *NullableTriggerDeploymentMetadata) Set(val *TriggerDeploymentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerDeploymentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerDeploymentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerDeploymentMetadata(val *TriggerDeploymentMetadata) *NullableTriggerDeploymentMetadata {
	return &NullableTriggerDeploymentMetadata{value: val, isSet: true}
}

func (v NullableTriggerDeploymentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerDeploymentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


