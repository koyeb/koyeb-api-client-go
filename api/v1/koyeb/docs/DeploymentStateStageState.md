# DeploymentStateStageState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentStateStageStateStatus**](DeploymentStateStageStateStatus.md) |  | [optional] [default to DEPLOYMENTSTATESTAGESTATESTATUS_UNKNOWN]
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeploymentStateStageState

`func NewDeploymentStateStageState() *DeploymentStateStageState`

NewDeploymentStateStageState instantiates a new DeploymentStateStageState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStateStageStateWithDefaults

`func NewDeploymentStateStageStateWithDefaults() *DeploymentStateStageState`

NewDeploymentStateStageStateWithDefaults instantiates a new DeploymentStateStageState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentStateStageState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStateStageState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStateStageState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentStateStageState) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DeploymentStateStageState) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DeploymentStateStageState) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DeploymentStateStageState) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DeploymentStateStageState) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentStateStageState) GetStatus() DeploymentStateStageStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentStateStageState) GetStatusOk() (*DeploymentStateStageStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentStateStageState) SetStatus(v DeploymentStateStageStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentStateStageState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *DeploymentStateStageState) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeploymentStateStageState) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeploymentStateStageState) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeploymentStateStageState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


