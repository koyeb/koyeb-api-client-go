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

// checks if the DockerHubRegistryConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DockerHubRegistryConfiguration{}

// DockerHubRegistryConfiguration struct for DockerHubRegistryConfiguration
type DockerHubRegistryConfiguration struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewDockerHubRegistryConfiguration instantiates a new DockerHubRegistryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerHubRegistryConfiguration() *DockerHubRegistryConfiguration {
	this := DockerHubRegistryConfiguration{}
	return &this
}

// NewDockerHubRegistryConfigurationWithDefaults instantiates a new DockerHubRegistryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerHubRegistryConfigurationWithDefaults() *DockerHubRegistryConfiguration {
	this := DockerHubRegistryConfiguration{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DockerHubRegistryConfiguration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerHubRegistryConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DockerHubRegistryConfiguration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DockerHubRegistryConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DockerHubRegistryConfiguration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerHubRegistryConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DockerHubRegistryConfiguration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DockerHubRegistryConfiguration) SetPassword(v string) {
	o.Password = &v
}

func (o DockerHubRegistryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DockerHubRegistryConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableDockerHubRegistryConfiguration struct {
	value *DockerHubRegistryConfiguration
	isSet bool
}

func (v NullableDockerHubRegistryConfiguration) Get() *DockerHubRegistryConfiguration {
	return v.value
}

func (v *NullableDockerHubRegistryConfiguration) Set(val *DockerHubRegistryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerHubRegistryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerHubRegistryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerHubRegistryConfiguration(val *DockerHubRegistryConfiguration) *NullableDockerHubRegistryConfiguration {
	return &NullableDockerHubRegistryConfiguration{value: val, isSet: true}
}

func (v NullableDockerHubRegistryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerHubRegistryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


