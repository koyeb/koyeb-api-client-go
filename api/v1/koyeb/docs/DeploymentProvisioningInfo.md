# DeploymentProvisioningInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | Pointer to **string** | The git sha for this build (we resolve the reference at the start of the build). | [optional] 
**Image** | Pointer to **string** | The docker image built as a result of this build. | [optional] 
**Stages** | Pointer to [**[]DeploymentProvisioningInfoStage**](DeploymentProvisioningInfoStage.md) | Some info about the build. | [optional] 

## Methods

### NewDeploymentProvisioningInfo

`func NewDeploymentProvisioningInfo() *DeploymentProvisioningInfo`

NewDeploymentProvisioningInfo instantiates a new DeploymentProvisioningInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProvisioningInfoWithDefaults

`func NewDeploymentProvisioningInfoWithDefaults() *DeploymentProvisioningInfo`

NewDeploymentProvisioningInfoWithDefaults instantiates a new DeploymentProvisioningInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *DeploymentProvisioningInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DeploymentProvisioningInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DeploymentProvisioningInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *DeploymentProvisioningInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetImage

`func (o *DeploymentProvisioningInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DeploymentProvisioningInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DeploymentProvisioningInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DeploymentProvisioningInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetStages

`func (o *DeploymentProvisioningInfo) GetStages() []DeploymentProvisioningInfoStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *DeploymentProvisioningInfo) GetStagesOk() (*[]DeploymentProvisioningInfoStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *DeploymentProvisioningInfo) SetStages(v []DeploymentProvisioningInfoStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *DeploymentProvisioningInfo) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


