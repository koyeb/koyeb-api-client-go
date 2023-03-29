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

// checks if the DeploymentProvisioningInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentProvisioningInfo{}

// DeploymentProvisioningInfo struct for DeploymentProvisioningInfo
type DeploymentProvisioningInfo struct {
	// The git sha for this build (we resolve the reference at the start of the build).
	Sha *string `json:"sha,omitempty"`
	// The docker image built as a result of this build.
	Image *string `json:"image,omitempty"`
	// Some info about the build.
	Stages []DeploymentProvisioningInfoStage `json:"stages,omitempty"`
}

// NewDeploymentProvisioningInfo instantiates a new DeploymentProvisioningInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfo() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// NewDeploymentProvisioningInfoWithDefaults instantiates a new DeploymentProvisioningInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoWithDefaults() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *DeploymentProvisioningInfo) SetSha(v string) {
	o.Sha = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DeploymentProvisioningInfo) SetImage(v string) {
	o.Image = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetStages() []DeploymentProvisioningInfoStage {
	if o == nil || IsNil(o.Stages) {
		var ret []DeploymentProvisioningInfoStage
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetStagesOk() ([]DeploymentProvisioningInfoStage, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []DeploymentProvisioningInfoStage and assigns it to the Stages field.
func (o *DeploymentProvisioningInfo) SetStages(v []DeploymentProvisioningInfoStage) {
	o.Stages = v
}

func (o DeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentProvisioningInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	return toSerialize, nil
}

type NullableDeploymentProvisioningInfo struct {
	value *DeploymentProvisioningInfo
	isSet bool
}

func (v NullableDeploymentProvisioningInfo) Get() *DeploymentProvisioningInfo {
	return v.value
}

func (v *NullableDeploymentProvisioningInfo) Set(val *DeploymentProvisioningInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfo(val *DeploymentProvisioningInfo) *NullableDeploymentProvisioningInfo {
	return &NullableDeploymentProvisioningInfo{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


