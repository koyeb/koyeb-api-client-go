# ExecCommandReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stdout** | Pointer to [**ExecCommandIO**](ExecCommandIO.md) |  | [optional] 
**Stderr** | Pointer to [**ExecCommandIO**](ExecCommandIO.md) |  | [optional] 
**Exited** | Pointer to **bool** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewExecCommandReply

`func NewExecCommandReply() *ExecCommandReply`

NewExecCommandReply instantiates a new ExecCommandReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCommandReplyWithDefaults

`func NewExecCommandReplyWithDefaults() *ExecCommandReply`

NewExecCommandReplyWithDefaults instantiates a new ExecCommandReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStdout

`func (o *ExecCommandReply) GetStdout() ExecCommandIO`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *ExecCommandReply) GetStdoutOk() (*ExecCommandIO, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *ExecCommandReply) SetStdout(v ExecCommandIO)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *ExecCommandReply) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *ExecCommandReply) GetStderr() ExecCommandIO`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *ExecCommandReply) GetStderrOk() (*ExecCommandIO, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *ExecCommandReply) SetStderr(v ExecCommandIO)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *ExecCommandReply) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetExited

`func (o *ExecCommandReply) GetExited() bool`

GetExited returns the Exited field if non-nil, zero value otherwise.

### GetExitedOk

`func (o *ExecCommandReply) GetExitedOk() (*bool, bool)`

GetExitedOk returns a tuple with the Exited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExited

`func (o *ExecCommandReply) SetExited(v bool)`

SetExited sets Exited field to given value.

### HasExited

`func (o *ExecCommandReply) HasExited() bool`

HasExited returns a boolean if a field has been set.

### GetExitCode

`func (o *ExecCommandReply) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ExecCommandReply) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ExecCommandReply) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ExecCommandReply) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


