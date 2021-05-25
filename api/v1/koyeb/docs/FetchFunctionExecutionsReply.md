# FetchFunctionExecutionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | Pointer to [**[]FunctionRunInfoListItem**](FunctionRunInfoListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewFetchFunctionExecutionsReply

`func NewFetchFunctionExecutionsReply() *FetchFunctionExecutionsReply`

NewFetchFunctionExecutionsReply instantiates a new FetchFunctionExecutionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchFunctionExecutionsReplyWithDefaults

`func NewFetchFunctionExecutionsReplyWithDefaults() *FetchFunctionExecutionsReply`

NewFetchFunctionExecutionsReplyWithDefaults instantiates a new FetchFunctionExecutionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *FetchFunctionExecutionsReply) GetExecutions() []FunctionRunInfoListItem`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *FetchFunctionExecutionsReply) GetExecutionsOk() (*[]FunctionRunInfoListItem, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *FetchFunctionExecutionsReply) SetExecutions(v []FunctionRunInfoListItem)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *FetchFunctionExecutionsReply) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetLimit

`func (o *FetchFunctionExecutionsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FetchFunctionExecutionsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FetchFunctionExecutionsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FetchFunctionExecutionsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *FetchFunctionExecutionsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FetchFunctionExecutionsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FetchFunctionExecutionsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *FetchFunctionExecutionsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *FetchFunctionExecutionsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FetchFunctionExecutionsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FetchFunctionExecutionsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FetchFunctionExecutionsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


