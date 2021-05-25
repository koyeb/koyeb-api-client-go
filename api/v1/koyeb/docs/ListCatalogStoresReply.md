# ListCatalogStoresReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogStores** | Pointer to [**[]CatalogStore**](CatalogStore.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCatalogStoresReply

`func NewListCatalogStoresReply() *ListCatalogStoresReply`

NewListCatalogStoresReply instantiates a new ListCatalogStoresReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogStoresReplyWithDefaults

`func NewListCatalogStoresReplyWithDefaults() *ListCatalogStoresReply`

NewListCatalogStoresReplyWithDefaults instantiates a new ListCatalogStoresReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogStores

`func (o *ListCatalogStoresReply) GetCatalogStores() []CatalogStore`

GetCatalogStores returns the CatalogStores field if non-nil, zero value otherwise.

### GetCatalogStoresOk

`func (o *ListCatalogStoresReply) GetCatalogStoresOk() (*[]CatalogStore, bool)`

GetCatalogStoresOk returns a tuple with the CatalogStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogStores

`func (o *ListCatalogStoresReply) SetCatalogStores(v []CatalogStore)`

SetCatalogStores sets CatalogStores field to given value.

### HasCatalogStores

`func (o *ListCatalogStoresReply) HasCatalogStores() bool`

HasCatalogStores returns a boolean if a field has been set.

### GetLimit

`func (o *ListCatalogStoresReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCatalogStoresReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCatalogStoresReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCatalogStoresReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCatalogStoresReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCatalogStoresReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCatalogStoresReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCatalogStoresReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListCatalogStoresReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCatalogStoresReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCatalogStoresReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCatalogStoresReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


