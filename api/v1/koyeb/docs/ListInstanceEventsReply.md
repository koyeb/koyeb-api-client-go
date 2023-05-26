# ListInstanceEventsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]InstanceEvent**](InstanceEvent.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewListInstanceEventsReply

`func NewListInstanceEventsReply() *ListInstanceEventsReply`

NewListInstanceEventsReply instantiates a new ListInstanceEventsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceEventsReplyWithDefaults

`func NewListInstanceEventsReplyWithDefaults() *ListInstanceEventsReply`

NewListInstanceEventsReplyWithDefaults instantiates a new ListInstanceEventsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListInstanceEventsReply) GetEvents() []InstanceEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListInstanceEventsReply) GetEventsOk() (*[]InstanceEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListInstanceEventsReply) SetEvents(v []InstanceEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListInstanceEventsReply) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLimit

`func (o *ListInstanceEventsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInstanceEventsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInstanceEventsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListInstanceEventsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListInstanceEventsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInstanceEventsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInstanceEventsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListInstanceEventsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListInstanceEventsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListInstanceEventsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListInstanceEventsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListInstanceEventsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrder

`func (o *ListInstanceEventsReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListInstanceEventsReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListInstanceEventsReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListInstanceEventsReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


