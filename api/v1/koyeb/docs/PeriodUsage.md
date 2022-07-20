# PeriodUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingTime** | Pointer to **time.Time** |  | [optional] 
**EndingTime** | Pointer to **time.Time** |  | [optional] 
**Apps** | Pointer to [**[]AppUsage**](AppUsage.md) |  | [optional] 

## Methods

### NewPeriodUsage

`func NewPeriodUsage() *PeriodUsage`

NewPeriodUsage instantiates a new PeriodUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodUsageWithDefaults

`func NewPeriodUsageWithDefaults() *PeriodUsage`

NewPeriodUsageWithDefaults instantiates a new PeriodUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingTime

`func (o *PeriodUsage) GetStartingTime() time.Time`

GetStartingTime returns the StartingTime field if non-nil, zero value otherwise.

### GetStartingTimeOk

`func (o *PeriodUsage) GetStartingTimeOk() (*time.Time, bool)`

GetStartingTimeOk returns a tuple with the StartingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTime

`func (o *PeriodUsage) SetStartingTime(v time.Time)`

SetStartingTime sets StartingTime field to given value.

### HasStartingTime

`func (o *PeriodUsage) HasStartingTime() bool`

HasStartingTime returns a boolean if a field has been set.

### GetEndingTime

`func (o *PeriodUsage) GetEndingTime() time.Time`

GetEndingTime returns the EndingTime field if non-nil, zero value otherwise.

### GetEndingTimeOk

`func (o *PeriodUsage) GetEndingTimeOk() (*time.Time, bool)`

GetEndingTimeOk returns a tuple with the EndingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingTime

`func (o *PeriodUsage) SetEndingTime(v time.Time)`

SetEndingTime sets EndingTime field to given value.

### HasEndingTime

`func (o *PeriodUsage) HasEndingTime() bool`

HasEndingTime returns a boolean if a field has been set.

### GetApps

`func (o *PeriodUsage) GetApps() []AppUsage`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PeriodUsage) GetAppsOk() (*[]AppUsage, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PeriodUsage) SetApps(v []AppUsage)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PeriodUsage) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


