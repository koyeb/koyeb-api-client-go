# DeploymentStateBuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | Pointer to **string** | The git sha for this build (we resolve the reference at the start of the build). | [optional] 
**Image** | Pointer to **string** | The docker image built as a result of this build. | [optional] 
**BuildJobId** | Pointer to **string** | The id of the job that ran the build. | [optional] 
**StageState** | Pointer to [**[]DeploymentStateStageState**](DeploymentStateStageState.md) | Some info about the build. | [optional] 

## Methods

### NewDeploymentStateBuildInfo

`func NewDeploymentStateBuildInfo() *DeploymentStateBuildInfo`

NewDeploymentStateBuildInfo instantiates a new DeploymentStateBuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateBuildInfoWithDefaults

`func NewDeploymentStateBuildInfoWithDefaults() *DeploymentStateBuildInfo`

NewDeploymentStateBuildInfoWithDefaults instantiates a new DeploymentStateBuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *DeploymentStateBuildInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DeploymentStateBuildInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DeploymentStateBuildInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *DeploymentStateBuildInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetImage

`func (o *DeploymentStateBuildInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DeploymentStateBuildInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DeploymentStateBuildInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DeploymentStateBuildInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetBuildJobId

`func (o *DeploymentStateBuildInfo) GetBuildJobId() string`

GetBuildJobId returns the BuildJobId field if non-nil, zero value otherwise.

### GetBuildJobIdOk

`func (o *DeploymentStateBuildInfo) GetBuildJobIdOk() (*string, bool)`

GetBuildJobIdOk returns a tuple with the BuildJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildJobId

`func (o *DeploymentStateBuildInfo) SetBuildJobId(v string)`

SetBuildJobId sets BuildJobId field to given value.

### HasBuildJobId

`func (o *DeploymentStateBuildInfo) HasBuildJobId() bool`

HasBuildJobId returns a boolean if a field has been set.

### GetStageState

`func (o *DeploymentStateBuildInfo) GetStageState() []DeploymentStateStageState`

GetStageState returns the StageState field if non-nil, zero value otherwise.

### GetStageStateOk

`func (o *DeploymentStateBuildInfo) GetStageStateOk() (*[]DeploymentStateStageState, bool)`

GetStageStateOk returns a tuple with the StageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageState

`func (o *DeploymentStateBuildInfo) SetStageState(v []DeploymentStateStageState)`

SetStageState sets StageState field to given value.

### HasStageState

`func (o *DeploymentStateBuildInfo) HasStageState() bool`

HasStageState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


