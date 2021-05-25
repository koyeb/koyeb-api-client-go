# GitSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A url to a git repository (contains the provider as well) .e.g: github.com/koyeb/test. | [optional] 
**Reference** | Pointer to **string** | A git reference of the code that needs to be build (.e.g: refs/heads/my-branch refs/tags/my-tag deadbeef). | [optional] 
**DeployOnPush** | Pointer to **bool** | If true a new revision will be triggered when the hash of the reference is pushed. | [optional] 

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

### GetReference

`func (o *GitSource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GitSource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GitSource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GitSource) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDeployOnPush

`func (o *GitSource) GetDeployOnPush() bool`

GetDeployOnPush returns the DeployOnPush field if non-nil, zero value otherwise.

### GetDeployOnPushOk

`func (o *GitSource) GetDeployOnPushOk() (*bool, bool)`

GetDeployOnPushOk returns a tuple with the DeployOnPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployOnPush

`func (o *GitSource) SetDeployOnPush(v bool)`

SetDeployOnPush sets DeployOnPush field to given value.

### HasDeployOnPush

`func (o *GitSource) HasDeployOnPush() bool`

HasDeployOnPush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


