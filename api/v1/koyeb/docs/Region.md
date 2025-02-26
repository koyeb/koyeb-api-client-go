# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to **[]string** |  | [optional] 
**Datacenters** | Pointer to **[]string** |  | [optional] 
**VolumesEnabled** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to **string** | The scope of the region, continent, metropolitan area, etc. | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Region) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCoordinates

`func (o *Region) GetCoordinates() []string`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Region) GetCoordinatesOk() (*[]string, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Region) SetCoordinates(v []string)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Region) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetStatus

`func (o *Region) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Region) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Region) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Region) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInstances

`func (o *Region) GetInstances() []string`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Region) GetInstancesOk() (*[]string, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Region) SetInstances(v []string)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Region) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDatacenters

`func (o *Region) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *Region) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *Region) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *Region) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetVolumesEnabled

`func (o *Region) GetVolumesEnabled() bool`

GetVolumesEnabled returns the VolumesEnabled field if non-nil, zero value otherwise.

### GetVolumesEnabledOk

`func (o *Region) GetVolumesEnabledOk() (*bool, bool)`

GetVolumesEnabledOk returns a tuple with the VolumesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEnabled

`func (o *Region) SetVolumesEnabled(v bool)`

SetVolumesEnabled sets VolumesEnabled field to given value.

### HasVolumesEnabled

`func (o *Region) HasVolumesEnabled() bool`

HasVolumesEnabled returns a boolean if a field has been set.

### GetScope

`func (o *Region) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Region) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Region) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Region) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


