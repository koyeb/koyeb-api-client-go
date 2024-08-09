# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AllocatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**SucceededAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ChildId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentStatus**](DeploymentStatus.md) |  | [optional] [default to DEPLOYMENTSTATUS_PENDING]
**Metadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 
**Definition** | Pointer to [**DeploymentDefinition**](DeploymentDefinition.md) |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**ProvisioningInfo** | Pointer to [**DeploymentProvisioningInfo**](DeploymentProvisioningInfo.md) |  | [optional] 
**DatabaseInfo** | Pointer to [**DeploymentDatabaseInfo**](DeploymentDatabaseInfo.md) |  | [optional] 
**SkipBuild** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to [**DeploymentRole**](DeploymentRole.md) |  | [optional] [default to DEPLOYMENTROLE_INVALID]
**Version** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Deployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Deployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Deployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Deployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Deployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Deployment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAllocatedAt

`func (o *Deployment) GetAllocatedAt() time.Time`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *Deployment) GetAllocatedAtOk() (*time.Time, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *Deployment) SetAllocatedAt(v time.Time)`

SetAllocatedAt sets AllocatedAt field to given value.

### HasAllocatedAt

`func (o *Deployment) HasAllocatedAt() bool`

HasAllocatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Deployment) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Deployment) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Deployment) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Deployment) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSucceededAt

`func (o *Deployment) GetSucceededAt() time.Time`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *Deployment) GetSucceededAtOk() (*time.Time, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *Deployment) SetSucceededAt(v time.Time)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *Deployment) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *Deployment) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *Deployment) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *Deployment) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *Deployment) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Deployment) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Deployment) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Deployment) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Deployment) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *Deployment) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Deployment) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Deployment) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Deployment) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *Deployment) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Deployment) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Deployment) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Deployment) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetParentId

`func (o *Deployment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Deployment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Deployment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Deployment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildId

`func (o *Deployment) GetChildId() string`

GetChildId returns the ChildId field if non-nil, zero value otherwise.

### GetChildIdOk

`func (o *Deployment) GetChildIdOk() (*string, bool)`

GetChildIdOk returns a tuple with the ChildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildId

`func (o *Deployment) SetChildId(v string)`

SetChildId sets ChildId field to given value.

### HasChildId

`func (o *Deployment) HasChildId() bool`

HasChildId returns a boolean if a field has been set.

### GetStatus

`func (o *Deployment) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *Deployment) GetMetadata() DeploymentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Deployment) GetMetadataOk() (*DeploymentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Deployment) SetMetadata(v DeploymentMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Deployment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDefinition

`func (o *Deployment) GetDefinition() DeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Deployment) GetDefinitionOk() (*DeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Deployment) SetDefinition(v DeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *Deployment) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetMessages

`func (o *Deployment) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Deployment) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Deployment) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Deployment) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetProvisioningInfo

`func (o *Deployment) GetProvisioningInfo() DeploymentProvisioningInfo`

GetProvisioningInfo returns the ProvisioningInfo field if non-nil, zero value otherwise.

### GetProvisioningInfoOk

`func (o *Deployment) GetProvisioningInfoOk() (*DeploymentProvisioningInfo, bool)`

GetProvisioningInfoOk returns a tuple with the ProvisioningInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningInfo

`func (o *Deployment) SetProvisioningInfo(v DeploymentProvisioningInfo)`

SetProvisioningInfo sets ProvisioningInfo field to given value.

### HasProvisioningInfo

`func (o *Deployment) HasProvisioningInfo() bool`

HasProvisioningInfo returns a boolean if a field has been set.

### GetDatabaseInfo

`func (o *Deployment) GetDatabaseInfo() DeploymentDatabaseInfo`

GetDatabaseInfo returns the DatabaseInfo field if non-nil, zero value otherwise.

### GetDatabaseInfoOk

`func (o *Deployment) GetDatabaseInfoOk() (*DeploymentDatabaseInfo, bool)`

GetDatabaseInfoOk returns a tuple with the DatabaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseInfo

`func (o *Deployment) SetDatabaseInfo(v DeploymentDatabaseInfo)`

SetDatabaseInfo sets DatabaseInfo field to given value.

### HasDatabaseInfo

`func (o *Deployment) HasDatabaseInfo() bool`

HasDatabaseInfo returns a boolean if a field has been set.

### GetSkipBuild

`func (o *Deployment) GetSkipBuild() bool`

GetSkipBuild returns the SkipBuild field if non-nil, zero value otherwise.

### GetSkipBuildOk

`func (o *Deployment) GetSkipBuildOk() (*bool, bool)`

GetSkipBuildOk returns a tuple with the SkipBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBuild

`func (o *Deployment) SetSkipBuild(v bool)`

SetSkipBuild sets SkipBuild field to given value.

### HasSkipBuild

`func (o *Deployment) HasSkipBuild() bool`

HasSkipBuild returns a boolean if a field has been set.

### GetRole

`func (o *Deployment) GetRole() DeploymentRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Deployment) GetRoleOk() (*DeploymentRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Deployment) SetRole(v DeploymentRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Deployment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVersion

`func (o *Deployment) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Deployment) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Deployment) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Deployment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *Deployment) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *Deployment) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *Deployment) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *Deployment) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


