# ErrorWithFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]ErrorField**](ErrorField.md) |  | [optional] 

## Methods

### NewErrorWithFields

`func NewErrorWithFields() *ErrorWithFields`

NewErrorWithFields instantiates a new ErrorWithFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithFieldsWithDefaults

`func NewErrorWithFieldsWithDefaults() *ErrorWithFields`

NewErrorWithFieldsWithDefaults instantiates a new ErrorWithFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorWithFields) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorWithFields) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorWithFields) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorWithFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ErrorWithFields) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorWithFields) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorWithFields) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorWithFields) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorWithFields) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorWithFields) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorWithFields) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorWithFields) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *ErrorWithFields) GetFields() []ErrorField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorWithFields) GetFieldsOk() (*[]ErrorField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorWithFields) SetFields(v []ErrorField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorWithFields) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


