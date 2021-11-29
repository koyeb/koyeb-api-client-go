# ServiceRevisionStateStageState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ServiceRevisionStateStageStateStatus**](ServiceRevisionStateStageStateStatus.md) |  | [optional] [default to SERVICEREVISIONSTATESTAGESTATESTATUS_UNKNOWN]
**Messages** | Pointer to **[]string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceRevisionStateStageState

`func NewServiceRevisionStateStageState() *ServiceRevisionStateStageState`

NewServiceRevisionStateStageState instantiates a new ServiceRevisionStateStageState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevisionStateStageStateWithDefaults

`func NewServiceRevisionStateStageStateWithDefaults() *ServiceRevisionStateStageState`

NewServiceRevisionStateStageStateWithDefaults instantiates a new ServiceRevisionStateStageState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceRevisionStateStageState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRevisionStateStageState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRevisionStateStageState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceRevisionStateStageState) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ServiceRevisionStateStageState) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ServiceRevisionStateStageState) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ServiceRevisionStateStageState) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ServiceRevisionStateStageState) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceRevisionStateStageState) GetStatus() ServiceRevisionStateStageStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceRevisionStateStageState) GetStatusOk() (*ServiceRevisionStateStageStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceRevisionStateStageState) SetStatus(v ServiceRevisionStateStageStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceRevisionStateStageState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ServiceRevisionStateStageState) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ServiceRevisionStateStageState) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ServiceRevisionStateStageState) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ServiceRevisionStateStageState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetStartedAt

`func (o *ServiceRevisionStateStageState) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ServiceRevisionStateStageState) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ServiceRevisionStateStageState) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ServiceRevisionStateStageState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ServiceRevisionStateStageState) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ServiceRevisionStateStageState) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ServiceRevisionStateStageState) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ServiceRevisionStateStageState) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


