# InstanceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AvailableAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**RegionalDeploymentId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**InstanceSnapshotStatus**](InstanceSnapshotStatus.md) |  | [optional] [default to INSTANCESNAPSHOTSTATUS_INVALID]
**Type** | Pointer to [**InstanceSnapshotType**](InstanceSnapshotType.md) |  | [optional] [default to INSTANCESNAPSHOTTYPE_INVALID]
**Version** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInstanceSnapshot

`func NewInstanceSnapshot() *InstanceSnapshot`

NewInstanceSnapshot instantiates a new InstanceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSnapshotWithDefaults

`func NewInstanceSnapshotWithDefaults() *InstanceSnapshot`

NewInstanceSnapshotWithDefaults instantiates a new InstanceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstanceSnapshot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceSnapshot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceSnapshot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstanceSnapshot) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAvailableAt

`func (o *InstanceSnapshot) GetAvailableAt() time.Time`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *InstanceSnapshot) GetAvailableAtOk() (*time.Time, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *InstanceSnapshot) SetAvailableAt(v time.Time)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *InstanceSnapshot) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *InstanceSnapshot) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InstanceSnapshot) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InstanceSnapshot) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InstanceSnapshot) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InstanceSnapshot) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InstanceSnapshot) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InstanceSnapshot) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InstanceSnapshot) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectId

`func (o *InstanceSnapshot) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InstanceSnapshot) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InstanceSnapshot) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InstanceSnapshot) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServiceId

`func (o *InstanceSnapshot) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InstanceSnapshot) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InstanceSnapshot) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *InstanceSnapshot) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *InstanceSnapshot) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *InstanceSnapshot) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *InstanceSnapshot) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *InstanceSnapshot) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetRegionalDeploymentId

`func (o *InstanceSnapshot) GetRegionalDeploymentId() string`

GetRegionalDeploymentId returns the RegionalDeploymentId field if non-nil, zero value otherwise.

### GetRegionalDeploymentIdOk

`func (o *InstanceSnapshot) GetRegionalDeploymentIdOk() (*string, bool)`

GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeploymentId

`func (o *InstanceSnapshot) SetRegionalDeploymentId(v string)`

SetRegionalDeploymentId sets RegionalDeploymentId field to given value.

### HasRegionalDeploymentId

`func (o *InstanceSnapshot) HasRegionalDeploymentId() bool`

HasRegionalDeploymentId returns a boolean if a field has been set.

### GetInstanceId

`func (o *InstanceSnapshot) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceSnapshot) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceSnapshot) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InstanceSnapshot) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceSnapshot) GetStatus() InstanceSnapshotStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceSnapshot) GetStatusOk() (*InstanceSnapshotStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceSnapshot) SetStatus(v InstanceSnapshotStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *InstanceSnapshot) GetType() InstanceSnapshotType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceSnapshot) GetTypeOk() (*InstanceSnapshotType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceSnapshot) SetType(v InstanceSnapshotType)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *InstanceSnapshot) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceSnapshot) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceSnapshot) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceSnapshot) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMessages

`func (o *InstanceSnapshot) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *InstanceSnapshot) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *InstanceSnapshot) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *InstanceSnapshot) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


