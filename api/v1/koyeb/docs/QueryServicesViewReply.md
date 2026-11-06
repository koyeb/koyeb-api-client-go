# QueryServicesViewReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]ServiceViewRow**](ServiceViewRow.md) |  | [optional] 
**Groups** | Pointer to [**[]ServiceViewGroup**](ServiceViewGroup.md) |  | [optional] 
**Facets** | Pointer to [**[]ServiceViewFacet**](ServiceViewFacet.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueryServicesViewReply

`func NewQueryServicesViewReply() *QueryServicesViewReply`

NewQueryServicesViewReply instantiates a new QueryServicesViewReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesViewReplyWithDefaults

`func NewQueryServicesViewReplyWithDefaults() *QueryServicesViewReply`

NewQueryServicesViewReplyWithDefaults instantiates a new QueryServicesViewReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *QueryServicesViewReply) GetServices() []ServiceViewRow`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *QueryServicesViewReply) GetServicesOk() (*[]ServiceViewRow, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *QueryServicesViewReply) SetServices(v []ServiceViewRow)`

SetServices sets Services field to given value.

### HasServices

`func (o *QueryServicesViewReply) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetGroups

`func (o *QueryServicesViewReply) GetGroups() []ServiceViewGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *QueryServicesViewReply) GetGroupsOk() (*[]ServiceViewGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *QueryServicesViewReply) SetGroups(v []ServiceViewGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *QueryServicesViewReply) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetFacets

`func (o *QueryServicesViewReply) GetFacets() []ServiceViewFacet`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *QueryServicesViewReply) GetFacetsOk() (*[]ServiceViewFacet, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *QueryServicesViewReply) SetFacets(v []ServiceViewFacet)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *QueryServicesViewReply) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetLimit

`func (o *QueryServicesViewReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryServicesViewReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryServicesViewReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryServicesViewReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QueryServicesViewReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryServicesViewReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryServicesViewReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryServicesViewReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *QueryServicesViewReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryServicesViewReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryServicesViewReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryServicesViewReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *QueryServicesViewReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *QueryServicesViewReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *QueryServicesViewReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *QueryServicesViewReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


