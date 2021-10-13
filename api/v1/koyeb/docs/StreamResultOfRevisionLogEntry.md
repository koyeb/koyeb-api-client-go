# StreamResultOfRevisionLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**RevisionLogEntry**](RevisionLogEntry.md) |  | [optional] 
**Error** | Pointer to [**GoogleRpcStatus**](GoogleRpcStatus.md) |  | [optional] 

## Methods

### NewStreamResultOfRevisionLogEntry

`func NewStreamResultOfRevisionLogEntry() *StreamResultOfRevisionLogEntry`

NewStreamResultOfRevisionLogEntry instantiates a new StreamResultOfRevisionLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfRevisionLogEntryWithDefaults

`func NewStreamResultOfRevisionLogEntryWithDefaults() *StreamResultOfRevisionLogEntry`

NewStreamResultOfRevisionLogEntryWithDefaults instantiates a new StreamResultOfRevisionLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfRevisionLogEntry) GetResult() RevisionLogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfRevisionLogEntry) GetResultOk() (*RevisionLogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfRevisionLogEntry) SetResult(v RevisionLogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfRevisionLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfRevisionLogEntry) GetError() GoogleRpcStatus`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfRevisionLogEntry) GetErrorOk() (*GoogleRpcStatus, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfRevisionLogEntry) SetError(v GoogleRpcStatus)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfRevisionLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


