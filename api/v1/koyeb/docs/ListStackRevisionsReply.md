# ListStackRevisionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revisions** | Pointer to [**[]StackRevisionListItem**](StackRevisionListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListStackRevisionsReply

`func NewListStackRevisionsReply() *ListStackRevisionsReply`

NewListStackRevisionsReply instantiates a new ListStackRevisionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStackRevisionsReplyWithDefaults

`func NewListStackRevisionsReplyWithDefaults() *ListStackRevisionsReply`

NewListStackRevisionsReplyWithDefaults instantiates a new ListStackRevisionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisions

`func (o *ListStackRevisionsReply) GetRevisions() []StackRevisionListItem`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ListStackRevisionsReply) GetRevisionsOk() (*[]StackRevisionListItem, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ListStackRevisionsReply) SetRevisions(v []StackRevisionListItem)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *ListStackRevisionsReply) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetLimit

`func (o *ListStackRevisionsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListStackRevisionsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListStackRevisionsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListStackRevisionsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListStackRevisionsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListStackRevisionsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListStackRevisionsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListStackRevisionsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListStackRevisionsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListStackRevisionsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListStackRevisionsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListStackRevisionsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


