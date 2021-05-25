# GoogleRpcStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]GoogleProtobufAny**](GoogleProtobufAny.md) |  | [optional] 

## Methods

### NewGoogleRpcStatus

`func NewGoogleRpcStatus() *GoogleRpcStatus`

NewGoogleRpcStatus instantiates a new GoogleRpcStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleRpcStatusWithDefaults

`func NewGoogleRpcStatusWithDefaults() *GoogleRpcStatus`

NewGoogleRpcStatusWithDefaults instantiates a new GoogleRpcStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GoogleRpcStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GoogleRpcStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GoogleRpcStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GoogleRpcStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GoogleRpcStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GoogleRpcStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GoogleRpcStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GoogleRpcStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *GoogleRpcStatus) GetDetails() []GoogleProtobufAny`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GoogleRpcStatus) GetDetailsOk() (*[]GoogleProtobufAny, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GoogleRpcStatus) SetDetails(v []GoogleProtobufAny)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GoogleRpcStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


