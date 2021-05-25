# FunctionRunInfoListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**FnName** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**FunctionRunInfoState**](FunctionRunInfoState.md) |  | [optional] [default to FUNCTIONRUNINFOSTATE_UNKNOWN]
**Executions** | Pointer to [**[]FunctionExecution**](FunctionExecution.md) | A list of executions, with the first one being the most recent one. | [optional] 

## Methods

### NewFunctionRunInfoListItem

`func NewFunctionRunInfoListItem() *FunctionRunInfoListItem`

NewFunctionRunInfoListItem instantiates a new FunctionRunInfoListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionRunInfoListItemWithDefaults

`func NewFunctionRunInfoListItemWithDefaults() *FunctionRunInfoListItem`

NewFunctionRunInfoListItemWithDefaults instantiates a new FunctionRunInfoListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *FunctionRunInfoListItem) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *FunctionRunInfoListItem) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *FunctionRunInfoListItem) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *FunctionRunInfoListItem) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetEventId

`func (o *FunctionRunInfoListItem) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *FunctionRunInfoListItem) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *FunctionRunInfoListItem) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *FunctionRunInfoListItem) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetFnName

`func (o *FunctionRunInfoListItem) GetFnName() string`

GetFnName returns the FnName field if non-nil, zero value otherwise.

### GetFnNameOk

`func (o *FunctionRunInfoListItem) GetFnNameOk() (*string, bool)`

GetFnNameOk returns a tuple with the FnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnName

`func (o *FunctionRunInfoListItem) SetFnName(v string)`

SetFnName sets FnName field to given value.

### HasFnName

`func (o *FunctionRunInfoListItem) HasFnName() bool`

HasFnName returns a boolean if a field has been set.

### GetState

`func (o *FunctionRunInfoListItem) GetState() FunctionRunInfoState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FunctionRunInfoListItem) GetStateOk() (*FunctionRunInfoState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FunctionRunInfoListItem) SetState(v FunctionRunInfoState)`

SetState sets State field to given value.

### HasState

`func (o *FunctionRunInfoListItem) HasState() bool`

HasState returns a boolean if a field has been set.

### GetExecutions

`func (o *FunctionRunInfoListItem) GetExecutions() []FunctionExecution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *FunctionRunInfoListItem) GetExecutionsOk() (*[]FunctionExecution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *FunctionRunInfoListItem) SetExecutions(v []FunctionExecution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *FunctionRunInfoListItem) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


