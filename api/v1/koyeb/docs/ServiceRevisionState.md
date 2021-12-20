# ServiceRevisionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ServiceRevisionStateStatus**](ServiceRevisionStateStatus.md) |  | [optional] [default to SERVICEREVISIONSTATESTATUS_UNKNOWN]
**StatusMessage** | Pointer to **string** |  | [optional] 
**Datacenters** | Pointer to **[]string** |  | [optional] 
**BuildInfo** | Pointer to [**ServiceRevisionStateBuildInfo**](ServiceRevisionStateBuildInfo.md) |  | [optional] 

## Methods

### NewServiceRevisionState

`func NewServiceRevisionState() *ServiceRevisionState`

NewServiceRevisionState instantiates a new ServiceRevisionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevisionStateWithDefaults

`func NewServiceRevisionStateWithDefaults() *ServiceRevisionState`

NewServiceRevisionStateWithDefaults instantiates a new ServiceRevisionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ServiceRevisionState) GetStatus() ServiceRevisionStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceRevisionState) GetStatusOk() (*ServiceRevisionStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceRevisionState) SetStatus(v ServiceRevisionStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceRevisionState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ServiceRevisionState) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ServiceRevisionState) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ServiceRevisionState) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ServiceRevisionState) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDatacenters

`func (o *ServiceRevisionState) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *ServiceRevisionState) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *ServiceRevisionState) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *ServiceRevisionState) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetBuildInfo

`func (o *ServiceRevisionState) GetBuildInfo() ServiceRevisionStateBuildInfo`

GetBuildInfo returns the BuildInfo field if non-nil, zero value otherwise.

### GetBuildInfoOk

`func (o *ServiceRevisionState) GetBuildInfoOk() (*ServiceRevisionStateBuildInfo, bool)`

GetBuildInfoOk returns a tuple with the BuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfo

`func (o *ServiceRevisionState) SetBuildInfo(v ServiceRevisionStateBuildInfo)`

SetBuildInfo sets BuildInfo field to given value.

### HasBuildInfo

`func (o *ServiceRevisionState) HasBuildInfo() bool`

HasBuildInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


