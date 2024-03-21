# RedeployRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**UseCache** | Pointer to **bool** |  | [optional] 
**SkipBuild** | Pointer to **bool** | If set to true, the build stage will be skipped and the image coming from the last successful build step will be used instead. The call fails if no previous successful builds happened. | [optional] 

## Methods

### NewRedeployRequestInfo

`func NewRedeployRequestInfo() *RedeployRequestInfo`

NewRedeployRequestInfo instantiates a new RedeployRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeployRequestInfoWithDefaults

`func NewRedeployRequestInfoWithDefaults() *RedeployRequestInfo`

NewRedeployRequestInfoWithDefaults instantiates a new RedeployRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentGroup

`func (o *RedeployRequestInfo) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *RedeployRequestInfo) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *RedeployRequestInfo) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *RedeployRequestInfo) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetSha

`func (o *RedeployRequestInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *RedeployRequestInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *RedeployRequestInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *RedeployRequestInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetUseCache

`func (o *RedeployRequestInfo) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *RedeployRequestInfo) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *RedeployRequestInfo) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *RedeployRequestInfo) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### GetSkipBuild

`func (o *RedeployRequestInfo) GetSkipBuild() bool`

GetSkipBuild returns the SkipBuild field if non-nil, zero value otherwise.

### GetSkipBuildOk

`func (o *RedeployRequestInfo) GetSkipBuildOk() (*bool, bool)`

GetSkipBuildOk returns a tuple with the SkipBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBuild

`func (o *RedeployRequestInfo) SetSkipBuild(v bool)`

SetSkipBuild sets SkipBuild field to given value.

### HasSkipBuild

`func (o *RedeployRequestInfo) HasSkipBuild() bool`

HasSkipBuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


