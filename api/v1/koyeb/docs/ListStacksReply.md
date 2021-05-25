# ListStacksReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stacks** | Pointer to [**[]StackListItem**](StackListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListStacksReply

`func NewListStacksReply() *ListStacksReply`

NewListStacksReply instantiates a new ListStacksReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStacksReplyWithDefaults

`func NewListStacksReplyWithDefaults() *ListStacksReply`

NewListStacksReplyWithDefaults instantiates a new ListStacksReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStacks

`func (o *ListStacksReply) GetStacks() []StackListItem`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *ListStacksReply) GetStacksOk() (*[]StackListItem, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *ListStacksReply) SetStacks(v []StackListItem)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *ListStacksReply) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetLimit

`func (o *ListStacksReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListStacksReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListStacksReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListStacksReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListStacksReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListStacksReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListStacksReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListStacksReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListStacksReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListStacksReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListStacksReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListStacksReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


