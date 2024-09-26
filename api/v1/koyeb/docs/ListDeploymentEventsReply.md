# ListDeploymentEventsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]DeploymentEvent**](DeploymentEvent.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListDeploymentEventsReply

`func NewListDeploymentEventsReply() *ListDeploymentEventsReply`

NewListDeploymentEventsReply instantiates a new ListDeploymentEventsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentEventsReplyWithDefaults

`func NewListDeploymentEventsReplyWithDefaults() *ListDeploymentEventsReply`

NewListDeploymentEventsReplyWithDefaults instantiates a new ListDeploymentEventsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListDeploymentEventsReply) GetEvents() []DeploymentEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListDeploymentEventsReply) GetEventsOk() (*[]DeploymentEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListDeploymentEventsReply) SetEvents(v []DeploymentEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListDeploymentEventsReply) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLimit

`func (o *ListDeploymentEventsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListDeploymentEventsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListDeploymentEventsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListDeploymentEventsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListDeploymentEventsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListDeploymentEventsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListDeploymentEventsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListDeploymentEventsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *ListDeploymentEventsReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListDeploymentEventsReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListDeploymentEventsReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListDeploymentEventsReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetHasNext

`func (o *ListDeploymentEventsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListDeploymentEventsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListDeploymentEventsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListDeploymentEventsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


