# ListCatalogInstancesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]CatalogInstanceListItem**](CatalogInstanceListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCatalogInstancesReply

`func NewListCatalogInstancesReply() *ListCatalogInstancesReply`

NewListCatalogInstancesReply instantiates a new ListCatalogInstancesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogInstancesReplyWithDefaults

`func NewListCatalogInstancesReplyWithDefaults() *ListCatalogInstancesReply`

NewListCatalogInstancesReplyWithDefaults instantiates a new ListCatalogInstancesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListCatalogInstancesReply) GetInstances() []CatalogInstanceListItem`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListCatalogInstancesReply) GetInstancesOk() (*[]CatalogInstanceListItem, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListCatalogInstancesReply) SetInstances(v []CatalogInstanceListItem)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListCatalogInstancesReply) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLimit

`func (o *ListCatalogInstancesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCatalogInstancesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCatalogInstancesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCatalogInstancesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCatalogInstancesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCatalogInstancesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCatalogInstancesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCatalogInstancesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListCatalogInstancesReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCatalogInstancesReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCatalogInstancesReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCatalogInstancesReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


