# KgitproxyListRepositoriesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]KgitproxyRepository**](KgitproxyRepository.md) | The collection of repositories. | [optional] 
**Limit** | Pointer to **int64** | The limit in the request. | [optional] 
**Offset** | Pointer to **int64** | The offset in the request. | [optional] 
**Count** | Pointer to **int64** | The total number of items. | [optional] 

## Methods

### NewKgitproxyListRepositoriesReply

`func NewKgitproxyListRepositoriesReply() *KgitproxyListRepositoriesReply`

NewKgitproxyListRepositoriesReply instantiates a new KgitproxyListRepositoriesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKgitproxyListRepositoriesReplyWithDefaults

`func NewKgitproxyListRepositoriesReplyWithDefaults() *KgitproxyListRepositoriesReply`

NewKgitproxyListRepositoriesReplyWithDefaults instantiates a new KgitproxyListRepositoriesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *KgitproxyListRepositoriesReply) GetRepositories() []KgitproxyRepository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *KgitproxyListRepositoriesReply) GetRepositoriesOk() (*[]KgitproxyRepository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *KgitproxyListRepositoriesReply) SetRepositories(v []KgitproxyRepository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *KgitproxyListRepositoriesReply) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetLimit

`func (o *KgitproxyListRepositoriesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KgitproxyListRepositoriesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KgitproxyListRepositoriesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KgitproxyListRepositoriesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *KgitproxyListRepositoriesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KgitproxyListRepositoriesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KgitproxyListRepositoriesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KgitproxyListRepositoriesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *KgitproxyListRepositoriesReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KgitproxyListRepositoriesReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KgitproxyListRepositoriesReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *KgitproxyListRepositoriesReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


