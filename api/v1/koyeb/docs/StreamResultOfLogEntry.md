# StreamResultOfLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**LogEntry**](LogEntry.md) |  | [optional] 
**Error** | Pointer to [**GoogleRpcStatus**](GoogleRpcStatus.md) |  | [optional] 

## Methods

### NewStreamResultOfLogEntry

`func NewStreamResultOfLogEntry() *StreamResultOfLogEntry`

NewStreamResultOfLogEntry instantiates a new StreamResultOfLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfLogEntryWithDefaults

`func NewStreamResultOfLogEntryWithDefaults() *StreamResultOfLogEntry`

NewStreamResultOfLogEntryWithDefaults instantiates a new StreamResultOfLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfLogEntry) GetResult() LogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfLogEntry) GetResultOk() (*LogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfLogEntry) SetResult(v LogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfLogEntry) GetError() GoogleRpcStatus`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfLogEntry) GetErrorOk() (*GoogleRpcStatus, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfLogEntry) SetError(v GoogleRpcStatus)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


