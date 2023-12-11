# ExecCommandRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | Command to exec. Mandatory in the first frame sent | [optional] 
**TtySize** | Pointer to [**ExecCommandRequestTerminalSize**](ExecCommandRequestTerminalSize.md) |  | [optional] 
**Stdin** | Pointer to [**ExecCommandIO**](ExecCommandIO.md) |  | [optional] 
**DisableTty** | Pointer to **bool** | Disable TTY. It&#39;s enough to specify it in the first frame | [optional] 

## Methods

### NewExecCommandRequestBody

`func NewExecCommandRequestBody() *ExecCommandRequestBody`

NewExecCommandRequestBody instantiates a new ExecCommandRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCommandRequestBodyWithDefaults

`func NewExecCommandRequestBodyWithDefaults() *ExecCommandRequestBody`

NewExecCommandRequestBodyWithDefaults instantiates a new ExecCommandRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ExecCommandRequestBody) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ExecCommandRequestBody) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ExecCommandRequestBody) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ExecCommandRequestBody) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetTtySize

`func (o *ExecCommandRequestBody) GetTtySize() ExecCommandRequestTerminalSize`

GetTtySize returns the TtySize field if non-nil, zero value otherwise.

### GetTtySizeOk

`func (o *ExecCommandRequestBody) GetTtySizeOk() (*ExecCommandRequestTerminalSize, bool)`

GetTtySizeOk returns a tuple with the TtySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtySize

`func (o *ExecCommandRequestBody) SetTtySize(v ExecCommandRequestTerminalSize)`

SetTtySize sets TtySize field to given value.

### HasTtySize

`func (o *ExecCommandRequestBody) HasTtySize() bool`

HasTtySize returns a boolean if a field has been set.

### GetStdin

`func (o *ExecCommandRequestBody) GetStdin() ExecCommandIO`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *ExecCommandRequestBody) GetStdinOk() (*ExecCommandIO, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *ExecCommandRequestBody) SetStdin(v ExecCommandIO)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *ExecCommandRequestBody) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetDisableTty

`func (o *ExecCommandRequestBody) GetDisableTty() bool`

GetDisableTty returns the DisableTty field if non-nil, zero value otherwise.

### GetDisableTtyOk

`func (o *ExecCommandRequestBody) GetDisableTtyOk() (*bool, bool)`

GetDisableTtyOk returns a tuple with the DisableTty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTty

`func (o *ExecCommandRequestBody) SetDisableTty(v bool)`

SetDisableTty sets DisableTty field to given value.

### HasDisableTty

`func (o *ExecCommandRequestBody) HasDisableTty() bool`

HasDisableTty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


