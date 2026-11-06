# ListArchivesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archives** | Pointer to [**[]Archive**](Archive.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListArchivesReply

`func NewListArchivesReply() *ListArchivesReply`

NewListArchivesReply instantiates a new ListArchivesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchivesReplyWithDefaults

`func NewListArchivesReplyWithDefaults() *ListArchivesReply`

NewListArchivesReplyWithDefaults instantiates a new ListArchivesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchives

`func (o *ListArchivesReply) GetArchives() []Archive`

GetArchives returns the Archives field if non-nil, zero value otherwise.

### GetArchivesOk

`func (o *ListArchivesReply) GetArchivesOk() (*[]Archive, bool)`

GetArchivesOk returns a tuple with the Archives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchives

`func (o *ListArchivesReply) SetArchives(v []Archive)`

SetArchives sets Archives field to given value.

### HasArchives

`func (o *ListArchivesReply) HasArchives() bool`

HasArchives returns a boolean if a field has been set.

### GetLimit

`func (o *ListArchivesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListArchivesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListArchivesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListArchivesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListArchivesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListArchivesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListArchivesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListArchivesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListArchivesReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListArchivesReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListArchivesReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListArchivesReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *ListArchivesReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListArchivesReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListArchivesReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListArchivesReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


