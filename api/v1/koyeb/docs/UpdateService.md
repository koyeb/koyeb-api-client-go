# UpdateService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to [**DeploymentDefinition**](DeploymentDefinition.md) |  | [optional] 
**Metadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 
**SkipBuild** | Pointer to **bool** | If set to true, the build stage will be skipped and the image coming from the last successful build step will be used instead. The call fails if no previous successful builds happened. | [optional] 
**SaveOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateService

`func NewUpdateService() *UpdateService`

NewUpdateService instantiates a new UpdateService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceWithDefaults

`func NewUpdateServiceWithDefaults() *UpdateService`

NewUpdateServiceWithDefaults instantiates a new UpdateService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *UpdateService) GetDefinition() DeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateService) GetDefinitionOk() (*DeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateService) SetDefinition(v DeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *UpdateService) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateService) GetMetadata() DeploymentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateService) GetMetadataOk() (*DeploymentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateService) SetMetadata(v DeploymentMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateService) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSkipBuild

`func (o *UpdateService) GetSkipBuild() bool`

GetSkipBuild returns the SkipBuild field if non-nil, zero value otherwise.

### GetSkipBuildOk

`func (o *UpdateService) GetSkipBuildOk() (*bool, bool)`

GetSkipBuildOk returns a tuple with the SkipBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBuild

`func (o *UpdateService) SetSkipBuild(v bool)`

SetSkipBuild sets SkipBuild field to given value.

### HasSkipBuild

`func (o *UpdateService) HasSkipBuild() bool`

HasSkipBuild returns a boolean if a field has been set.

### GetSaveOnly

`func (o *UpdateService) GetSaveOnly() bool`

GetSaveOnly returns the SaveOnly field if non-nil, zero value otherwise.

### GetSaveOnlyOk

`func (o *UpdateService) GetSaveOnlyOk() (*bool, bool)`

GetSaveOnlyOk returns a tuple with the SaveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveOnly

`func (o *UpdateService) SetSaveOnly(v bool)`

SetSaveOnly sets SaveOnly field to given value.

### HasSaveOnly

`func (o *UpdateService) HasSaveOnly() bool`

HasSaveOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


