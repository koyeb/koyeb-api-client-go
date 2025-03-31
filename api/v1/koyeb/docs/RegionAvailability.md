# RegionAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**AvailabilityLevel**](AvailabilityLevel.md) |  | [optional] [default to AVAILABILITYLEVEL_UNKNOWN]

## Methods

### NewRegionAvailability

`func NewRegionAvailability() *RegionAvailability`

NewRegionAvailability instantiates a new RegionAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionAvailabilityWithDefaults

`func NewRegionAvailabilityWithDefaults() *RegionAvailability`

NewRegionAvailabilityWithDefaults instantiates a new RegionAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *RegionAvailability) GetAvailability() AvailabilityLevel`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *RegionAvailability) GetAvailabilityOk() (*AvailabilityLevel, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *RegionAvailability) SetAvailability(v AvailabilityLevel)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *RegionAvailability) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


