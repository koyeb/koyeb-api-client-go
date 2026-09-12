# ListServicePoolsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePools** | Pointer to [**[]ServicePool**](ServicePool.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListServicePoolsReply

`func NewListServicePoolsReply() *ListServicePoolsReply`

NewListServicePoolsReply instantiates a new ListServicePoolsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePoolsReplyWithDefaults

`func NewListServicePoolsReplyWithDefaults() *ListServicePoolsReply`

NewListServicePoolsReplyWithDefaults instantiates a new ListServicePoolsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePools

`func (o *ListServicePoolsReply) GetServicePools() []ServicePool`

GetServicePools returns the ServicePools field if non-nil, zero value otherwise.

### GetServicePoolsOk

`func (o *ListServicePoolsReply) GetServicePoolsOk() (*[]ServicePool, bool)`

GetServicePoolsOk returns a tuple with the ServicePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePools

`func (o *ListServicePoolsReply) SetServicePools(v []ServicePool)`

SetServicePools sets ServicePools field to given value.

### HasServicePools

`func (o *ListServicePoolsReply) HasServicePools() bool`

HasServicePools returns a boolean if a field has been set.

### GetLimit

`func (o *ListServicePoolsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListServicePoolsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListServicePoolsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListServicePoolsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListServicePoolsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListServicePoolsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListServicePoolsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListServicePoolsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListServicePoolsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListServicePoolsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListServicePoolsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListServicePoolsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *ListServicePoolsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListServicePoolsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListServicePoolsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListServicePoolsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


