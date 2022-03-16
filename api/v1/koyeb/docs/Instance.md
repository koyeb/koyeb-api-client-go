# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**RegionalDeploymentId** | Pointer to **string** |  | [optional] 
**AllocationId** | Pointer to **string** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**InstanceStatus**](InstanceStatus.md) |  | [optional] [default to INSTANCESTATUS_ALLOCATING]
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Instance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Instance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Instance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Instance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Instance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Instance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Instance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Instance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Instance) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Instance) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Instance) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Instance) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *Instance) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Instance) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Instance) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Instance) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *Instance) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Instance) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Instance) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Instance) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *Instance) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *Instance) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *Instance) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *Instance) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetRegionalDeploymentId

`func (o *Instance) GetRegionalDeploymentId() string`

GetRegionalDeploymentId returns the RegionalDeploymentId field if non-nil, zero value otherwise.

### GetRegionalDeploymentIdOk

`func (o *Instance) GetRegionalDeploymentIdOk() (*string, bool)`

GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeploymentId

`func (o *Instance) SetRegionalDeploymentId(v string)`

SetRegionalDeploymentId sets RegionalDeploymentId field to given value.

### HasRegionalDeploymentId

`func (o *Instance) HasRegionalDeploymentId() bool`

HasRegionalDeploymentId returns a boolean if a field has been set.

### GetAllocationId

`func (o *Instance) GetAllocationId() string`

GetAllocationId returns the AllocationId field if non-nil, zero value otherwise.

### GetAllocationIdOk

`func (o *Instance) GetAllocationIdOk() (*string, bool)`

GetAllocationIdOk returns a tuple with the AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationId

`func (o *Instance) SetAllocationId(v string)`

SetAllocationId sets AllocationId field to given value.

### HasAllocationId

`func (o *Instance) HasAllocationId() bool`

HasAllocationId returns a boolean if a field has been set.

### GetDatacenter

`func (o *Instance) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *Instance) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *Instance) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *Instance) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetStatus

`func (o *Instance) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Instance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *Instance) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Instance) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Instance) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Instance) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


