# DeploymentListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AllocatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ChildId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DeploymentStatus**](DeploymentStatus.md) |  | [optional] [default to DEPLOYMENTSTATUS_PENDING]
**Metadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**ServiceDefinition**](ServiceDefinition.md) |  | [optional] 
**State** | Pointer to [**ServiceRevisionState**](ServiceRevisionState.md) |  | [optional] 

## Methods

### NewDeploymentListItem

`func NewDeploymentListItem() *DeploymentListItem`

NewDeploymentListItem instantiates a new DeploymentListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentListItemWithDefaults

`func NewDeploymentListItemWithDefaults() *DeploymentListItem`

NewDeploymentListItemWithDefaults instantiates a new DeploymentListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeploymentListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeploymentListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeploymentListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAllocatedAt

`func (o *DeploymentListItem) GetAllocatedAt() time.Time`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *DeploymentListItem) GetAllocatedAtOk() (*time.Time, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *DeploymentListItem) SetAllocatedAt(v time.Time)`

SetAllocatedAt sets AllocatedAt field to given value.

### HasAllocatedAt

`func (o *DeploymentListItem) HasAllocatedAt() bool`

HasAllocatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DeploymentListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeploymentListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeploymentListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeploymentListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *DeploymentListItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *DeploymentListItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *DeploymentListItem) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *DeploymentListItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *DeploymentListItem) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeploymentListItem) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeploymentListItem) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeploymentListItem) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetParentId

`func (o *DeploymentListItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DeploymentListItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DeploymentListItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DeploymentListItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildId

`func (o *DeploymentListItem) GetChildId() string`

GetChildId returns the ChildId field if non-nil, zero value otherwise.

### GetChildIdOk

`func (o *DeploymentListItem) GetChildIdOk() (*string, bool)`

GetChildIdOk returns a tuple with the ChildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildId

`func (o *DeploymentListItem) SetChildId(v string)`

SetChildId sets ChildId field to given value.

### HasChildId

`func (o *DeploymentListItem) HasChildId() bool`

HasChildId returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentListItem) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentListItem) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentListItem) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *DeploymentListItem) GetMetadata() DeploymentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeploymentListItem) GetMetadataOk() (*DeploymentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeploymentListItem) SetMetadata(v DeploymentMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeploymentListItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMessages

`func (o *DeploymentListItem) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeploymentListItem) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeploymentListItem) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeploymentListItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentListItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentListItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentListItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentListItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *DeploymentListItem) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *DeploymentListItem) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *DeploymentListItem) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *DeploymentListItem) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetDefinition

`func (o *DeploymentListItem) GetDefinition() ServiceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DeploymentListItem) GetDefinitionOk() (*ServiceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DeploymentListItem) SetDefinition(v ServiceDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DeploymentListItem) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetState

`func (o *DeploymentListItem) GetState() ServiceRevisionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentListItem) GetStateOk() (*ServiceRevisionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentListItem) SetState(v ServiceRevisionState)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentListItem) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


