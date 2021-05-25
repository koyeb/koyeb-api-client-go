# ListCredentialsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]Credential**](Credential.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCredentialsReply

`func NewListCredentialsReply() *ListCredentialsReply`

NewListCredentialsReply instantiates a new ListCredentialsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialsReplyWithDefaults

`func NewListCredentialsReplyWithDefaults() *ListCredentialsReply`

NewListCredentialsReplyWithDefaults instantiates a new ListCredentialsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ListCredentialsReply) GetCredentials() []Credential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ListCredentialsReply) GetCredentialsOk() (*[]Credential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ListCredentialsReply) SetCredentials(v []Credential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ListCredentialsReply) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLimit

`func (o *ListCredentialsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCredentialsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCredentialsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCredentialsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCredentialsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCredentialsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCredentialsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCredentialsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListCredentialsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCredentialsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCredentialsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCredentialsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


