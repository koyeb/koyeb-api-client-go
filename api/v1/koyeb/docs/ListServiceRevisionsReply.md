# ListServiceRevisionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revisions** | Pointer to [**[]ServiceRevisionListItem**](ServiceRevisionListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListServiceRevisionsReply

`func NewListServiceRevisionsReply() *ListServiceRevisionsReply`

NewListServiceRevisionsReply instantiates a new ListServiceRevisionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceRevisionsReplyWithDefaults

`func NewListServiceRevisionsReplyWithDefaults() *ListServiceRevisionsReply`

NewListServiceRevisionsReplyWithDefaults instantiates a new ListServiceRevisionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisions

`func (o *ListServiceRevisionsReply) GetRevisions() []ServiceRevisionListItem`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ListServiceRevisionsReply) GetRevisionsOk() (*[]ServiceRevisionListItem, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ListServiceRevisionsReply) SetRevisions(v []ServiceRevisionListItem)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *ListServiceRevisionsReply) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetLimit

`func (o *ListServiceRevisionsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListServiceRevisionsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListServiceRevisionsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListServiceRevisionsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListServiceRevisionsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListServiceRevisionsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListServiceRevisionsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListServiceRevisionsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListServiceRevisionsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListServiceRevisionsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListServiceRevisionsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListServiceRevisionsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


