# KgitproxyListBranchesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]KgitproxyBranch**](KgitproxyBranch.md) | The collection of branches. | [optional] 
**Limit** | Pointer to **int64** | The limit in the request. | [optional] 
**Offset** | Pointer to **int64** | The offset in the request. | [optional] 
**Count** | Pointer to **int64** | The total number of items. | [optional] 

## Methods

### NewKgitproxyListBranchesReply

`func NewKgitproxyListBranchesReply() *KgitproxyListBranchesReply`

NewKgitproxyListBranchesReply instantiates a new KgitproxyListBranchesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKgitproxyListBranchesReplyWithDefaults

`func NewKgitproxyListBranchesReplyWithDefaults() *KgitproxyListBranchesReply`

NewKgitproxyListBranchesReplyWithDefaults instantiates a new KgitproxyListBranchesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *KgitproxyListBranchesReply) GetBranches() []KgitproxyBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *KgitproxyListBranchesReply) GetBranchesOk() (*[]KgitproxyBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *KgitproxyListBranchesReply) SetBranches(v []KgitproxyBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *KgitproxyListBranchesReply) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetLimit

`func (o *KgitproxyListBranchesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KgitproxyListBranchesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KgitproxyListBranchesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KgitproxyListBranchesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *KgitproxyListBranchesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KgitproxyListBranchesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KgitproxyListBranchesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KgitproxyListBranchesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *KgitproxyListBranchesReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KgitproxyListBranchesReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KgitproxyListBranchesReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *KgitproxyListBranchesReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


