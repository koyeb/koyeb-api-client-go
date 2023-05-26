# ListAppEventsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]AppEvent**](AppEvent.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewListAppEventsReply

`func NewListAppEventsReply() *ListAppEventsReply`

NewListAppEventsReply instantiates a new ListAppEventsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppEventsReplyWithDefaults

`func NewListAppEventsReplyWithDefaults() *ListAppEventsReply`

NewListAppEventsReplyWithDefaults instantiates a new ListAppEventsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListAppEventsReply) GetEvents() []AppEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListAppEventsReply) GetEventsOk() (*[]AppEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListAppEventsReply) SetEvents(v []AppEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListAppEventsReply) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLimit

`func (o *ListAppEventsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListAppEventsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListAppEventsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListAppEventsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListAppEventsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListAppEventsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListAppEventsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListAppEventsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListAppEventsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAppEventsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAppEventsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListAppEventsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrder

`func (o *ListAppEventsReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListAppEventsReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListAppEventsReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListAppEventsReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


