# ListInstanceSnapshotEventsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]InstanceSnapshotEvent**](InstanceSnapshotEvent.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListInstanceSnapshotEventsReply

`func NewListInstanceSnapshotEventsReply() *ListInstanceSnapshotEventsReply`

NewListInstanceSnapshotEventsReply instantiates a new ListInstanceSnapshotEventsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceSnapshotEventsReplyWithDefaults

`func NewListInstanceSnapshotEventsReplyWithDefaults() *ListInstanceSnapshotEventsReply`

NewListInstanceSnapshotEventsReplyWithDefaults instantiates a new ListInstanceSnapshotEventsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListInstanceSnapshotEventsReply) GetEvents() []InstanceSnapshotEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListInstanceSnapshotEventsReply) GetEventsOk() (*[]InstanceSnapshotEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListInstanceSnapshotEventsReply) SetEvents(v []InstanceSnapshotEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListInstanceSnapshotEventsReply) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLimit

`func (o *ListInstanceSnapshotEventsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInstanceSnapshotEventsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInstanceSnapshotEventsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListInstanceSnapshotEventsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListInstanceSnapshotEventsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInstanceSnapshotEventsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInstanceSnapshotEventsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListInstanceSnapshotEventsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *ListInstanceSnapshotEventsReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListInstanceSnapshotEventsReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListInstanceSnapshotEventsReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListInstanceSnapshotEventsReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetHasNext

`func (o *ListInstanceSnapshotEventsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListInstanceSnapshotEventsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListInstanceSnapshotEventsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListInstanceSnapshotEventsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


