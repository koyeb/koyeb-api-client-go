# QueryServicesViewFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]ServiceType**](ServiceType.md) |  | [optional] 
**Statuses** | Pointer to [**[]ServiceStatus**](ServiceStatus.md) |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**ActiveDeploymentStates** | Pointer to [**[]ServiceViewActiveDeploymentState**](ServiceViewActiveDeploymentState.md) |  | [optional] 
**LatestDeploymentStates** | Pointer to [**[]ServiceViewLatestDeploymentState**](ServiceViewLatestDeploymentState.md) |  | [optional] 
**NeedsAttention** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueryServicesViewFilters

`func NewQueryServicesViewFilters() *QueryServicesViewFilters`

NewQueryServicesViewFilters instantiates a new QueryServicesViewFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesViewFiltersWithDefaults

`func NewQueryServicesViewFiltersWithDefaults() *QueryServicesViewFilters`

NewQueryServicesViewFiltersWithDefaults instantiates a new QueryServicesViewFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *QueryServicesViewFilters) GetTypes() []ServiceType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryServicesViewFilters) GetTypesOk() (*[]ServiceType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryServicesViewFilters) SetTypes(v []ServiceType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *QueryServicesViewFilters) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetStatuses

`func (o *QueryServicesViewFilters) GetStatuses() []ServiceStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *QueryServicesViewFilters) GetStatusesOk() (*[]ServiceStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *QueryServicesViewFilters) SetStatuses(v []ServiceStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *QueryServicesViewFilters) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetRegions

`func (o *QueryServicesViewFilters) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *QueryServicesViewFilters) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *QueryServicesViewFilters) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *QueryServicesViewFilters) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAppIds

`func (o *QueryServicesViewFilters) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *QueryServicesViewFilters) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *QueryServicesViewFilters) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *QueryServicesViewFilters) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetActiveDeploymentStates

`func (o *QueryServicesViewFilters) GetActiveDeploymentStates() []ServiceViewActiveDeploymentState`

GetActiveDeploymentStates returns the ActiveDeploymentStates field if non-nil, zero value otherwise.

### GetActiveDeploymentStatesOk

`func (o *QueryServicesViewFilters) GetActiveDeploymentStatesOk() (*[]ServiceViewActiveDeploymentState, bool)`

GetActiveDeploymentStatesOk returns a tuple with the ActiveDeploymentStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeploymentStates

`func (o *QueryServicesViewFilters) SetActiveDeploymentStates(v []ServiceViewActiveDeploymentState)`

SetActiveDeploymentStates sets ActiveDeploymentStates field to given value.

### HasActiveDeploymentStates

`func (o *QueryServicesViewFilters) HasActiveDeploymentStates() bool`

HasActiveDeploymentStates returns a boolean if a field has been set.

### GetLatestDeploymentStates

`func (o *QueryServicesViewFilters) GetLatestDeploymentStates() []ServiceViewLatestDeploymentState`

GetLatestDeploymentStates returns the LatestDeploymentStates field if non-nil, zero value otherwise.

### GetLatestDeploymentStatesOk

`func (o *QueryServicesViewFilters) GetLatestDeploymentStatesOk() (*[]ServiceViewLatestDeploymentState, bool)`

GetLatestDeploymentStatesOk returns a tuple with the LatestDeploymentStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeploymentStates

`func (o *QueryServicesViewFilters) SetLatestDeploymentStates(v []ServiceViewLatestDeploymentState)`

SetLatestDeploymentStates sets LatestDeploymentStates field to given value.

### HasLatestDeploymentStates

`func (o *QueryServicesViewFilters) HasLatestDeploymentStates() bool`

HasLatestDeploymentStates returns a boolean if a field has been set.

### GetNeedsAttention

`func (o *QueryServicesViewFilters) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *QueryServicesViewFilters) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *QueryServicesViewFilters) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.

### HasNeedsAttention

`func (o *QueryServicesViewFilters) HasNeedsAttention() bool`

HasNeedsAttention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


