# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ChildId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**ServiceDefinition**](ServiceDefinition.md) |  | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Metadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 

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

### GetDefinition

`func (o *Deployment) GetDefinition() ServiceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Deployment) GetDefinitionOk() (*ServiceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Deployment) SetDefinition(v ServiceDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *Deployment) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetState

`func (o *Deployment) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *Deployment) HasState() bool`

HasState returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


