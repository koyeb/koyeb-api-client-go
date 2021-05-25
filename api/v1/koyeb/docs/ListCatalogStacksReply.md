# ListCatalogStacksReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogStacks** | Pointer to [**[]CatalogStackListItem**](CatalogStackListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCatalogStacksReply

`func NewListCatalogStacksReply() *ListCatalogStacksReply`

NewListCatalogStacksReply instantiates a new ListCatalogStacksReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogStacksReplyWithDefaults

`func NewListCatalogStacksReplyWithDefaults() *ListCatalogStacksReply`

NewListCatalogStacksReplyWithDefaults instantiates a new ListCatalogStacksReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogStacks

`func (o *ListCatalogStacksReply) GetCatalogStacks() []CatalogStackListItem`

GetCatalogStacks returns the CatalogStacks field if non-nil, zero value otherwise.

### GetCatalogStacksOk

`func (o *ListCatalogStacksReply) GetCatalogStacksOk() (*[]CatalogStackListItem, bool)`

GetCatalogStacksOk returns a tuple with the CatalogStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogStacks

`func (o *ListCatalogStacksReply) SetCatalogStacks(v []CatalogStackListItem)`

SetCatalogStacks sets CatalogStacks field to given value.

### HasCatalogStacks

`func (o *ListCatalogStacksReply) HasCatalogStacks() bool`

HasCatalogStacks returns a boolean if a field has been set.

### GetLimit

`func (o *ListCatalogStacksReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCatalogStacksReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCatalogStacksReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCatalogStacksReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCatalogStacksReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCatalogStacksReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCatalogStacksReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCatalogStacksReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListCatalogStacksReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCatalogStacksReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCatalogStacksReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCatalogStacksReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


