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

// ArchiveSource struct for ArchiveSource
type ArchiveSource struct {
	Id *string `json:"id,omitempty"`
	Buildpack *BuildpackBuilder `json:"buildpack,omitempty"`
	Docker *DockerBuilder `json:"docker,omitempty"`
}

// NewArchiveSource instantiates a new ArchiveSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveSource() *ArchiveSource {
	this := ArchiveSource{}
	return &this
}

// NewArchiveSourceWithDefaults instantiates a new ArchiveSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveSourceWithDefaults() *ArchiveSource {
	this := ArchiveSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArchiveSource) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSource) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArchiveSource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArchiveSource) SetId(v string) {
	o.Id = &v
}

// GetBuildpack returns the Buildpack field value if set, zero value otherwise.
func (o *ArchiveSource) GetBuildpack() BuildpackBuilder {
	if o == nil || isNil(o.Buildpack) {
		var ret BuildpackBuilder
		return ret
	}
	return *o.Buildpack
}

// GetBuildpackOk returns a tuple with the Buildpack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSource) GetBuildpackOk() (*BuildpackBuilder, bool) {
	if o == nil || isNil(o.Buildpack) {
    return nil, false
	}
	return o.Buildpack, true
}

// HasBuildpack returns a boolean if a field has been set.
func (o *ArchiveSource) HasBuildpack() bool {
	if o != nil && !isNil(o.Buildpack) {
		return true
	}

	return false
}

// SetBuildpack gets a reference to the given BuildpackBuilder and assigns it to the Buildpack field.
func (o *ArchiveSource) SetBuildpack(v BuildpackBuilder) {
	o.Buildpack = &v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *ArchiveSource) GetDocker() DockerBuilder {
	if o == nil || isNil(o.Docker) {
		var ret DockerBuilder
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveSource) GetDockerOk() (*DockerBuilder, bool) {
	if o == nil || isNil(o.Docker) {
    return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *ArchiveSource) HasDocker() bool {
	if o != nil && !isNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerBuilder and assigns it to the Docker field.
func (o *ArchiveSource) SetDocker(v DockerBuilder) {
	o.Docker = &v
}

func (o ArchiveSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Buildpack) {
		toSerialize["buildpack"] = o.Buildpack
	}
	if !isNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveSource struct {
	value *ArchiveSource
	isSet bool
}

func (v NullableArchiveSource) Get() *ArchiveSource {
	return v.value
}

func (v *NullableArchiveSource) Set(val *ArchiveSource) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveSource) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveSource(val *ArchiveSource) *NullableArchiveSource {
	return &NullableArchiveSource{value: val, isSet: true}
}

func (v NullableArchiveSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

