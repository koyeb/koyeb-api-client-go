# LoginMethodReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to [**LoginMethodReplyMethod**](LoginMethodReplyMethod.md) |  | [optional] [default to LOGINMETHODREPLYMETHOD_KOYEB]

## Methods

### NewLoginMethodReply

`func NewLoginMethodReply() *LoginMethodReply`

NewLoginMethodReply instantiates a new LoginMethodReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginMethodReplyWithDefaults

`func NewLoginMethodReplyWithDefaults() *LoginMethodReply`

NewLoginMethodReplyWithDefaults instantiates a new LoginMethodReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LoginMethodReply) GetMethod() LoginMethodReplyMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoginMethodReply) GetMethodOk() (*LoginMethodReplyMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoginMethodReply) SetMethod(v LoginMethodReplyMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoginMethodReply) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


