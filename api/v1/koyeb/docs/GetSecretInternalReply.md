# GetSecretInternalReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to [**InternalSecret**](InternalSecret.md) |  | [optional] 

## Methods

### NewGetSecretInternalReply

`func NewGetSecretInternalReply() *GetSecretInternalReply`

NewGetSecretInternalReply instantiates a new GetSecretInternalReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecretInternalReplyWithDefaults

`func NewGetSecretInternalReplyWithDefaults() *GetSecretInternalReply`

NewGetSecretInternalReplyWithDefaults instantiates a new GetSecretInternalReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *GetSecretInternalReply) GetSecret() InternalSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetSecretInternalReply) GetSecretOk() (*InternalSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetSecretInternalReply) SetSecret(v InternalSecret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GetSecretInternalReply) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


