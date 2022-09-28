# RegionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**map[string]InstanceUsage**](InstanceUsage.md) |  | [optional] 

## Methods

### NewRegionUsage

`func NewRegionUsage() *RegionUsage`

NewRegionUsage instantiates a new RegionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionUsageWithDefaults

`func NewRegionUsageWithDefaults() *RegionUsage`

NewRegionUsageWithDefaults instantiates a new RegionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *RegionUsage) GetInstances() map[string]InstanceUsage`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *RegionUsage) GetInstancesOk() (*map[string]InstanceUsage, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *RegionUsage) SetInstances(v map[string]InstanceUsage)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *RegionUsage) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


