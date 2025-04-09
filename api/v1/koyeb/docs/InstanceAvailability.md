# InstanceAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**map[string]RegionAvailability**](RegionAvailability.md) |  | [optional] 
**Availability** | Pointer to [**AvailabilityLevel**](AvailabilityLevel.md) |  | [optional] [default to AVAILABILITYLEVEL_UNKNOWN]

## Methods

### NewInstanceAvailability

`func NewInstanceAvailability() *InstanceAvailability`

NewInstanceAvailability instantiates a new InstanceAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAvailabilityWithDefaults

`func NewInstanceAvailabilityWithDefaults() *InstanceAvailability`

NewInstanceAvailabilityWithDefaults instantiates a new InstanceAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *InstanceAvailability) GetRegions() map[string]RegionAvailability`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *InstanceAvailability) GetRegionsOk() (*map[string]RegionAvailability, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *InstanceAvailability) SetRegions(v map[string]RegionAvailability)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *InstanceAvailability) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAvailability

`func (o *InstanceAvailability) GetAvailability() AvailabilityLevel`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceAvailability) GetAvailabilityOk() (*AvailabilityLevel, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceAvailability) SetAvailability(v AvailabilityLevel)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InstanceAvailability) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


