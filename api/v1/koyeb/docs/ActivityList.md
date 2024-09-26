# ActivityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**[]Activity**](Activity.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewActivityList

`func NewActivityList() *ActivityList`

NewActivityList instantiates a new ActivityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityListWithDefaults

`func NewActivityListWithDefaults() *ActivityList`

NewActivityListWithDefaults instantiates a new ActivityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *ActivityList) GetActivities() []Activity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ActivityList) GetActivitiesOk() (*[]Activity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ActivityList) SetActivities(v []Activity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ActivityList) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetLimit

`func (o *ActivityList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ActivityList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ActivityList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ActivityList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ActivityList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ActivityList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ActivityList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ActivityList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetHasNext

`func (o *ActivityList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ActivityList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ActivityList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ActivityList) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


