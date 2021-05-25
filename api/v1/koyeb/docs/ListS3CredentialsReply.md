# ListS3CredentialsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Credentials** | Pointer to [**[]S3Credential**](S3Credential.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListS3CredentialsReply

`func NewListS3CredentialsReply() *ListS3CredentialsReply`

NewListS3CredentialsReply instantiates a new ListS3CredentialsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListS3CredentialsReplyWithDefaults

`func NewListS3CredentialsReplyWithDefaults() *ListS3CredentialsReply`

NewListS3CredentialsReplyWithDefaults instantiates a new ListS3CredentialsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Credentials

`func (o *ListS3CredentialsReply) GetS3Credentials() []S3Credential`

GetS3Credentials returns the S3Credentials field if non-nil, zero value otherwise.

### GetS3CredentialsOk

`func (o *ListS3CredentialsReply) GetS3CredentialsOk() (*[]S3Credential, bool)`

GetS3CredentialsOk returns a tuple with the S3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Credentials

`func (o *ListS3CredentialsReply) SetS3Credentials(v []S3Credential)`

SetS3Credentials sets S3Credentials field to given value.

### HasS3Credentials

`func (o *ListS3CredentialsReply) HasS3Credentials() bool`

HasS3Credentials returns a boolean if a field has been set.

### GetLimit

`func (o *ListS3CredentialsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListS3CredentialsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListS3CredentialsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListS3CredentialsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListS3CredentialsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListS3CredentialsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListS3CredentialsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListS3CredentialsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListS3CredentialsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListS3CredentialsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListS3CredentialsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListS3CredentialsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


