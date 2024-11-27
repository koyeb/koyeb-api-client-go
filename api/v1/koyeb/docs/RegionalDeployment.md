# RegionalDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ScheduledAt** | Pointer to **time.Time** |  | [optional] 
**AllocatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**SucceededAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ChildId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**RegionalDeploymentStatus**](RegionalDeploymentStatus.md) |  | [optional] [default to REGIONALDEPLOYMENTSTATUS_PENDING]
**Messages** | Pointer to **[]string** |  | [optional] 
**Definition** | Pointer to [**RegionalDeploymentDefinition**](RegionalDeploymentDefinition.md) |  | [optional] 
**Datacenters** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProvisioningInfo** | Pointer to [**DeploymentProvisioningInfo**](DeploymentProvisioningInfo.md) |  | [optional] 
**Role** | Pointer to [**RegionalDeploymentRole**](RegionalDeploymentRole.md) |  | [optional] [default to REGIONALDEPLOYMENTROLE_INVALID]
**Version** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 

## Methods

### NewRegionalDeployment

`func NewRegionalDeployment() *RegionalDeployment`

NewRegionalDeployment instantiates a new RegionalDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionalDeploymentWithDefaults

`func NewRegionalDeploymentWithDefaults() *RegionalDeployment`

NewRegionalDeploymentWithDefaults instantiates a new RegionalDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionalDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionalDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionalDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegionalDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegionalDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegionalDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegionalDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegionalDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RegionalDeployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegionalDeployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegionalDeployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegionalDeployment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetScheduledAt

`func (o *RegionalDeployment) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *RegionalDeployment) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *RegionalDeployment) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *RegionalDeployment) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetAllocatedAt

`func (o *RegionalDeployment) GetAllocatedAt() time.Time`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *RegionalDeployment) GetAllocatedAtOk() (*time.Time, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *RegionalDeployment) SetAllocatedAt(v time.Time)`

SetAllocatedAt sets AllocatedAt field to given value.

### HasAllocatedAt

`func (o *RegionalDeployment) HasAllocatedAt() bool`

HasAllocatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RegionalDeployment) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RegionalDeployment) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RegionalDeployment) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RegionalDeployment) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSucceededAt

`func (o *RegionalDeployment) GetSucceededAt() time.Time`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *RegionalDeployment) GetSucceededAtOk() (*time.Time, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *RegionalDeployment) SetSucceededAt(v time.Time)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *RegionalDeployment) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *RegionalDeployment) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *RegionalDeployment) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *RegionalDeployment) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *RegionalDeployment) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *RegionalDeployment) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RegionalDeployment) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RegionalDeployment) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RegionalDeployment) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *RegionalDeployment) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *RegionalDeployment) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *RegionalDeployment) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *RegionalDeployment) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *RegionalDeployment) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RegionalDeployment) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RegionalDeployment) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *RegionalDeployment) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRegion

`func (o *RegionalDeployment) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionalDeployment) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionalDeployment) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RegionalDeployment) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetParentId

`func (o *RegionalDeployment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *RegionalDeployment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *RegionalDeployment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *RegionalDeployment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildId

`func (o *RegionalDeployment) GetChildId() string`

GetChildId returns the ChildId field if non-nil, zero value otherwise.

### GetChildIdOk

`func (o *RegionalDeployment) GetChildIdOk() (*string, bool)`

GetChildIdOk returns a tuple with the ChildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildId

`func (o *RegionalDeployment) SetChildId(v string)`

SetChildId sets ChildId field to given value.

### HasChildId

`func (o *RegionalDeployment) HasChildId() bool`

HasChildId returns a boolean if a field has been set.

### GetStatus

`func (o *RegionalDeployment) GetStatus() RegionalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionalDeployment) GetStatusOk() (*RegionalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionalDeployment) SetStatus(v RegionalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegionalDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *RegionalDeployment) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *RegionalDeployment) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *RegionalDeployment) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *RegionalDeployment) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetDefinition

`func (o *RegionalDeployment) GetDefinition() RegionalDeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *RegionalDeployment) GetDefinitionOk() (*RegionalDeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *RegionalDeployment) SetDefinition(v RegionalDeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *RegionalDeployment) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDatacenters

`func (o *RegionalDeployment) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *RegionalDeployment) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *RegionalDeployment) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *RegionalDeployment) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetMetadata

`func (o *RegionalDeployment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegionalDeployment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegionalDeployment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegionalDeployment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProvisioningInfo

`func (o *RegionalDeployment) GetProvisioningInfo() DeploymentProvisioningInfo`

GetProvisioningInfo returns the ProvisioningInfo field if non-nil, zero value otherwise.

### GetProvisioningInfoOk

`func (o *RegionalDeployment) GetProvisioningInfoOk() (*DeploymentProvisioningInfo, bool)`

GetProvisioningInfoOk returns a tuple with the ProvisioningInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningInfo

`func (o *RegionalDeployment) SetProvisioningInfo(v DeploymentProvisioningInfo)`

SetProvisioningInfo sets ProvisioningInfo field to given value.

### HasProvisioningInfo

`func (o *RegionalDeployment) HasProvisioningInfo() bool`

HasProvisioningInfo returns a boolean if a field has been set.

### GetRole

`func (o *RegionalDeployment) GetRole() RegionalDeploymentRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RegionalDeployment) GetRoleOk() (*RegionalDeploymentRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RegionalDeployment) SetRole(v RegionalDeploymentRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *RegionalDeployment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVersion

`func (o *RegionalDeployment) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RegionalDeployment) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RegionalDeployment) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RegionalDeployment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *RegionalDeployment) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *RegionalDeployment) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *RegionalDeployment) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *RegionalDeployment) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetDeploymentId

`func (o *RegionalDeployment) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *RegionalDeployment) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *RegionalDeployment) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *RegionalDeployment) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


