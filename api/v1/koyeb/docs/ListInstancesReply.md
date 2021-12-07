# ListInstancesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]InstanceListItem**](InstanceListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewListInstancesReply

`func NewListInstancesReply() *ListInstancesReply`

NewListInstancesReply instantiates a new ListInstancesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstancesReplyWithDefaults

`func NewListInstancesReplyWithDefaults() *ListInstancesReply`

NewListInstancesReplyWithDefaults instantiates a new ListInstancesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListInstancesReply) GetInstances() []InstanceListItem`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListInstancesReply) GetInstancesOk() (*[]InstanceListItem, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListInstancesReply) SetInstances(v []InstanceListItem)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListInstancesReply) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLimit

`func (o *ListInstancesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInstancesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInstancesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListInstancesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListInstancesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInstancesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInstancesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListInstancesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListInstancesReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListInstancesReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListInstancesReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListInstancesReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrder

`func (o *ListInstancesReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListInstancesReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListInstancesReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListInstancesReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


