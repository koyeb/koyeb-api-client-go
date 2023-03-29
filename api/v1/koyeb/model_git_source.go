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

// checks if the GitSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitSource{}

// GitSource struct for GitSource
type GitSource struct {
	// A url to a git repository (contains the provider as well) .e.g: github.com/koyeb/test.
	Repository *string `json:"repository,omitempty"`
	Branch *string `json:"branch,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Sha *string `json:"sha,omitempty"`
	BuildCommand *string `json:"build_command,omitempty"`
	RunCommand *string `json:"run_command,omitempty"`
	NoDeployOnPush *bool `json:"no_deploy_on_push,omitempty"`
	Workdir *string `json:"workdir,omitempty"`
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
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GitSource) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
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
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GitSource) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
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
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GitSource) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
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
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitSource) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitSource) SetSha(v string) {
	o.Sha = &v
}

// GetBuildCommand returns the BuildCommand field value if set, zero value otherwise.
func (o *GitSource) GetBuildCommand() string {
	if o == nil || IsNil(o.BuildCommand) {
		var ret string
		return ret
	}
	return *o.BuildCommand
}

// GetBuildCommandOk returns a tuple with the BuildCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBuildCommandOk() (*string, bool) {
	if o == nil || IsNil(o.BuildCommand) {
		return nil, false
	}
	return o.BuildCommand, true
}

// HasBuildCommand returns a boolean if a field has been set.
func (o *GitSource) HasBuildCommand() bool {
	if o != nil && !IsNil(o.BuildCommand) {
		return true
	}

	return false
}

// SetBuildCommand gets a reference to the given string and assigns it to the BuildCommand field.
func (o *GitSource) SetBuildCommand(v string) {
	o.BuildCommand = &v
}

// GetRunCommand returns the RunCommand field value if set, zero value otherwise.
func (o *GitSource) GetRunCommand() string {
	if o == nil || IsNil(o.RunCommand) {
		var ret string
		return ret
	}
	return *o.RunCommand
}

// GetRunCommandOk returns a tuple with the RunCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetRunCommandOk() (*string, bool) {
	if o == nil || IsNil(o.RunCommand) {
		return nil, false
	}
	return o.RunCommand, true
}

// HasRunCommand returns a boolean if a field has been set.
func (o *GitSource) HasRunCommand() bool {
	if o != nil && !IsNil(o.RunCommand) {
		return true
	}

	return false
}

// SetRunCommand gets a reference to the given string and assigns it to the RunCommand field.
func (o *GitSource) SetRunCommand(v string) {
	o.RunCommand = &v
}

// GetNoDeployOnPush returns the NoDeployOnPush field value if set, zero value otherwise.
func (o *GitSource) GetNoDeployOnPush() bool {
	if o == nil || IsNil(o.NoDeployOnPush) {
		var ret bool
		return ret
	}
	return *o.NoDeployOnPush
}

// GetNoDeployOnPushOk returns a tuple with the NoDeployOnPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetNoDeployOnPushOk() (*bool, bool) {
	if o == nil || IsNil(o.NoDeployOnPush) {
		return nil, false
	}
	return o.NoDeployOnPush, true
}

// HasNoDeployOnPush returns a boolean if a field has been set.
func (o *GitSource) HasNoDeployOnPush() bool {
	if o != nil && !IsNil(o.NoDeployOnPush) {
		return true
	}

	return false
}

// SetNoDeployOnPush gets a reference to the given bool and assigns it to the NoDeployOnPush field.
func (o *GitSource) SetNoDeployOnPush(v bool) {
	o.NoDeployOnPush = &v
}

// GetWorkdir returns the Workdir field value if set, zero value otherwise.
func (o *GitSource) GetWorkdir() string {
	if o == nil || IsNil(o.Workdir) {
		var ret string
		return ret
	}
	return *o.Workdir
}

// GetWorkdirOk returns a tuple with the Workdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetWorkdirOk() (*string, bool) {
	if o == nil || IsNil(o.Workdir) {
		return nil, false
	}
	return o.Workdir, true
}

// HasWorkdir returns a boolean if a field has been set.
func (o *GitSource) HasWorkdir() bool {
	if o != nil && !IsNil(o.Workdir) {
		return true
	}

	return false
}

// SetWorkdir gets a reference to the given string and assigns it to the Workdir field.
func (o *GitSource) SetWorkdir(v string) {
	o.Workdir = &v
}

func (o GitSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.BuildCommand) {
		toSerialize["build_command"] = o.BuildCommand
	}
	if !IsNil(o.RunCommand) {
		toSerialize["run_command"] = o.RunCommand
	}
	if !IsNil(o.NoDeployOnPush) {
		toSerialize["no_deploy_on_push"] = o.NoDeployOnPush
	}
	if !IsNil(o.Workdir) {
		toSerialize["workdir"] = o.Workdir
	}
	return toSerialize, nil
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


