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

// checks if the KgitproxyGitHubRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KgitproxyGitHubRepository{}

// KgitproxyGitHubRepository struct for KgitproxyGitHubRepository
type KgitproxyGitHubRepository struct {
	GithubId *string `json:"github_id,omitempty"`
}

// NewKgitproxyGitHubRepository instantiates a new KgitproxyGitHubRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKgitproxyGitHubRepository() *KgitproxyGitHubRepository {
	this := KgitproxyGitHubRepository{}
	return &this
}

// NewKgitproxyGitHubRepositoryWithDefaults instantiates a new KgitproxyGitHubRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKgitproxyGitHubRepositoryWithDefaults() *KgitproxyGitHubRepository {
	this := KgitproxyGitHubRepository{}
	return &this
}

// GetGithubId returns the GithubId field value if set, zero value otherwise.
func (o *KgitproxyGitHubRepository) GetGithubId() string {
	if o == nil || IsNil(o.GithubId) {
		var ret string
		return ret
	}
	return *o.GithubId
}

// GetGithubIdOk returns a tuple with the GithubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyGitHubRepository) GetGithubIdOk() (*string, bool) {
	if o == nil || IsNil(o.GithubId) {
		return nil, false
	}
	return o.GithubId, true
}

// HasGithubId returns a boolean if a field has been set.
func (o *KgitproxyGitHubRepository) HasGithubId() bool {
	if o != nil && !IsNil(o.GithubId) {
		return true
	}

	return false
}

// SetGithubId gets a reference to the given string and assigns it to the GithubId field.
func (o *KgitproxyGitHubRepository) SetGithubId(v string) {
	o.GithubId = &v
}

func (o KgitproxyGitHubRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KgitproxyGitHubRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GithubId) {
		toSerialize["github_id"] = o.GithubId
	}
	return toSerialize, nil
}

type NullableKgitproxyGitHubRepository struct {
	value *KgitproxyGitHubRepository
	isSet bool
}

func (v NullableKgitproxyGitHubRepository) Get() *KgitproxyGitHubRepository {
	return v.value
}

func (v *NullableKgitproxyGitHubRepository) Set(val *KgitproxyGitHubRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableKgitproxyGitHubRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableKgitproxyGitHubRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKgitproxyGitHubRepository(val *KgitproxyGitHubRepository) *NullableKgitproxyGitHubRepository {
	return &NullableKgitproxyGitHubRepository{value: val, isSet: true}
}

func (v NullableKgitproxyGitHubRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKgitproxyGitHubRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


