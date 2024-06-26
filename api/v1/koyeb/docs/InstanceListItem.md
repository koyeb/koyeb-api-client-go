# InstanceListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**RegionalDeploymentId** | Pointer to **string** |  | [optional] 
**AllocationId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ReplicaIndex** | Pointer to **int64** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**InstanceStatus**](InstanceStatus.md) |  | [optional] [default to INSTANCESTATUS_ALLOCATING]
**Messages** | Pointer to **[]string** |  | [optional] 
**XyzDeploymentId** | Pointer to **string** | WARNING: Please don&#39;t use the following attribute. Koyeb doesn&#39;t guarantee backwards compatible breaking change and reserve the right to completely drop it without notice. USE AT YOUR OWN RISK. | [optional] 

## Methods

### NewInstanceListItem

`func NewInstanceListItem() *InstanceListItem`

NewInstanceListItem instantiates a new InstanceListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceListItemWithDefaults

`func NewInstanceListItemWithDefaults() *InstanceListItem`

NewInstanceListItemWithDefaults instantiates a new InstanceListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstanceListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstanceListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InstanceListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InstanceListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InstanceListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InstanceListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *InstanceListItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InstanceListItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InstanceListItem) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *InstanceListItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *InstanceListItem) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InstanceListItem) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InstanceListItem) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *InstanceListItem) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRegionalDeploymentId

`func (o *InstanceListItem) GetRegionalDeploymentId() string`

GetRegionalDeploymentId returns the RegionalDeploymentId field if non-nil, zero value otherwise.

### GetRegionalDeploymentIdOk

`func (o *InstanceListItem) GetRegionalDeploymentIdOk() (*string, bool)`

GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeploymentId

`func (o *InstanceListItem) SetRegionalDeploymentId(v string)`

SetRegionalDeploymentId sets RegionalDeploymentId field to given value.

### HasRegionalDeploymentId

`func (o *InstanceListItem) HasRegionalDeploymentId() bool`

HasRegionalDeploymentId returns a boolean if a field has been set.

### GetAllocationId

`func (o *InstanceListItem) GetAllocationId() string`

GetAllocationId returns the AllocationId field if non-nil, zero value otherwise.

### GetAllocationIdOk

`func (o *InstanceListItem) GetAllocationIdOk() (*string, bool)`

GetAllocationIdOk returns a tuple with the AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationId

`func (o *InstanceListItem) SetAllocationId(v string)`

SetAllocationId sets AllocationId field to given value.

### HasAllocationId

`func (o *InstanceListItem) HasAllocationId() bool`

HasAllocationId returns a boolean if a field has been set.

### GetType

`func (o *InstanceListItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceListItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceListItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReplicaIndex

`func (o *InstanceListItem) GetReplicaIndex() int64`

GetReplicaIndex returns the ReplicaIndex field if non-nil, zero value otherwise.

### GetReplicaIndexOk

`func (o *InstanceListItem) GetReplicaIndexOk() (*int64, bool)`

GetReplicaIndexOk returns a tuple with the ReplicaIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaIndex

`func (o *InstanceListItem) SetReplicaIndex(v int64)`

SetReplicaIndex sets ReplicaIndex field to given value.

### HasReplicaIndex

`func (o *InstanceListItem) HasReplicaIndex() bool`

HasReplicaIndex returns a boolean if a field has been set.

### GetRegion

`func (o *InstanceListItem) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceListItem) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceListItem) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceListItem) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDatacenter

`func (o *InstanceListItem) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *InstanceListItem) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *InstanceListItem) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *InstanceListItem) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceListItem) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceListItem) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceListItem) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *InstanceListItem) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *InstanceListItem) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *InstanceListItem) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *InstanceListItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetXyzDeploymentId

`func (o *InstanceListItem) GetXyzDeploymentId() string`

GetXyzDeploymentId returns the XyzDeploymentId field if non-nil, zero value otherwise.

### GetXyzDeploymentIdOk

`func (o *InstanceListItem) GetXyzDeploymentIdOk() (*string, bool)`

GetXyzDeploymentIdOk returns a tuple with the XyzDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXyzDeploymentId

`func (o *InstanceListItem) SetXyzDeploymentId(v string)`

SetXyzDeploymentId sets XyzDeploymentId field to given value.

### HasXyzDeploymentId

`func (o *InstanceListItem) HasXyzDeploymentId() bool`

HasXyzDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


