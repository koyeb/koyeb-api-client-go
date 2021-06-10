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

// GitSource struct for GitSource
type GitSource struct {
	// A url to a git repository (contains the provider as well) .e.g: github.com/koyeb/test.
	Repository *string `json:"repository,omitempty"`
	Branch *string `json:"branch,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Sha *string `json:"sha,omitempty"`
}

// NewGitSource instantiates a new GitSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitSource() *GitSource {
	this := GitSource{}
	return &this
}

// NewGitSourceWithDefaults instantiates a new GitSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitSourceWithDefaults() *GitSource {
	this := GitSource{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GitSource) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GitSource) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *GitSource) SetRepository(v string) {
	o.Repository = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *GitSource) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GitSource) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *GitSource) SetBranch(v string) {
	o.Branch = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GitSource) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GitSource) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GitSource) SetTag(v string) {
	o.Tag = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitSource) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitSource) HasSha() bool {
	if o != nil && o.Sha != nil {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitSource) SetSha(v string) {
	o.Sha = &v
}

func (o GitSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}
	return json.Marshal(toSerialize)
}

type NullableGitSource struct {
	value *GitSource
	isSet bool
}

func (v NullableGitSource) Get() *GitSource {
	return v.value
}

func (v *NullableGitSource) Set(val *GitSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGitSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGitSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitSource(val *GitSource) *NullableGitSource {
	return &NullableGitSource{value: val, isSet: true}
}

func (v NullableGitSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


