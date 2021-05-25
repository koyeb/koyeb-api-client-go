# FunctionExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**FunctionRunInfoState**](FunctionRunInfoState.md) |  | [optional] [default to FUNCTIONRUNINFOSTATE_UNKNOWN]
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewFunctionExecution

`func NewFunctionExecution() *FunctionExecution`

NewFunctionExecution instantiates a new FunctionExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionExecutionWithDefaults

`func NewFunctionExecutionWithDefaults() *FunctionExecution`

NewFunctionExecutionWithDefaults instantiates a new FunctionExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *FunctionExecution) GetState() FunctionRunInfoState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FunctionExecution) GetStateOk() (*FunctionRunInfoState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FunctionExecution) SetState(v FunctionRunInfoState)`

SetState sets State field to given value.

### HasState

`func (o *FunctionExecution) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStart

`func (o *FunctionExecution) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *FunctionExecution) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *FunctionExecution) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *FunctionExecution) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *FunctionExecution) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *FunctionExecution) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *FunctionExecution) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *FunctionExecution) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExitCode

`func (o *FunctionExecution) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *FunctionExecution) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *FunctionExecution) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *FunctionExecution) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


