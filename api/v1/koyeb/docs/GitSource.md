# GitSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A url to a git repository (contains the provider as well) .e.g: github.com/koyeb/test. | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**BuildCommand** | Pointer to **string** |  | [optional] 
**RunCommand** | Pointer to **string** |  | [optional] 
**NoDeployOnPush** | Pointer to **bool** |  | [optional] 

## Methods

### NewGitSource

`func NewGitSource() *GitSource`

NewGitSource instantiates a new GitSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitSourceWithDefaults

`func NewGitSourceWithDefaults() *GitSource`

NewGitSourceWithDefaults instantiates a new GitSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *GitSource) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitSource) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitSource) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GitSource) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBranch

`func (o *GitSource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitSource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitSource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitSource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTag

`func (o *GitSource) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GitSource) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GitSource) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GitSource) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSha

`func (o *GitSource) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitSource) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitSource) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitSource) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetBuildCommand

`func (o *GitSource) GetBuildCommand() string`

GetBuildCommand returns the BuildCommand field if non-nil, zero value otherwise.

### GetBuildCommandOk

`func (o *GitSource) GetBuildCommandOk() (*string, bool)`

GetBuildCommandOk returns a tuple with the BuildCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommand

`func (o *GitSource) SetBuildCommand(v string)`

SetBuildCommand sets BuildCommand field to given value.

### HasBuildCommand

`func (o *GitSource) HasBuildCommand() bool`

HasBuildCommand returns a boolean if a field has been set.

### GetRunCommand

`func (o *GitSource) GetRunCommand() string`

GetRunCommand returns the RunCommand field if non-nil, zero value otherwise.

### GetRunCommandOk

`func (o *GitSource) GetRunCommandOk() (*string, bool)`

GetRunCommandOk returns a tuple with the RunCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCommand

`func (o *GitSource) SetRunCommand(v string)`

SetRunCommand sets RunCommand field to given value.

### HasRunCommand

`func (o *GitSource) HasRunCommand() bool`

HasRunCommand returns a boolean if a field has been set.

### GetNoDeployOnPush

`func (o *GitSource) GetNoDeployOnPush() bool`

GetNoDeployOnPush returns the NoDeployOnPush field if non-nil, zero value otherwise.

### GetNoDeployOnPushOk

`func (o *GitSource) GetNoDeployOnPushOk() (*bool, bool)`

GetNoDeployOnPushOk returns a tuple with the NoDeployOnPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeployOnPush

`func (o *GitSource) SetNoDeployOnPush(v bool)`

SetNoDeployOnPush sets NoDeployOnPush field to given value.

### HasNoDeployOnPush

`func (o *GitSource) HasNoDeployOnPush() bool`

HasNoDeployOnPush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


