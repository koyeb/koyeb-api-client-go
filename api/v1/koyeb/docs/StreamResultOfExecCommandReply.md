# StreamResultOfExecCommandReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ExecCommandReply**](ExecCommandReply.md) |  | [optional] 
**Error** | Pointer to [**GoogleRpcStatus**](GoogleRpcStatus.md) |  | [optional] 

## Methods

### NewStreamResultOfExecCommandReply

`func NewStreamResultOfExecCommandReply() *StreamResultOfExecCommandReply`

NewStreamResultOfExecCommandReply instantiates a new StreamResultOfExecCommandReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfExecCommandReplyWithDefaults

`func NewStreamResultOfExecCommandReplyWithDefaults() *StreamResultOfExecCommandReply`

NewStreamResultOfExecCommandReplyWithDefaults instantiates a new StreamResultOfExecCommandReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfExecCommandReply) GetResult() ExecCommandReply`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfExecCommandReply) GetResultOk() (*ExecCommandReply, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfExecCommandReply) SetResult(v ExecCommandReply)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfExecCommandReply) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfExecCommandReply) GetError() GoogleRpcStatus`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfExecCommandReply) GetErrorOk() (*GoogleRpcStatus, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfExecCommandReply) SetError(v GoogleRpcStatus)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfExecCommandReply) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


