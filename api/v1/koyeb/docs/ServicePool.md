# ServicePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**ReadyCount** | Pointer to **int64** |  | [optional] 
**Definition** | Pointer to [**DeploymentDefinition**](DeploymentDefinition.md) |  | [optional] 
**Status** | Pointer to [**ServicePoolStatus**](ServicePoolStatus.md) |  | [optional] [default to SERVICEPOOLSTATUS_UNSPECIFIED]
**Messages** | Pointer to **[]string** |  | [optional] 
**Generation** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePool

`func NewServicePool() *ServicePool`

NewServicePool instantiates a new ServicePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePoolWithDefaults

`func NewServicePoolWithDefaults() *ServicePool`

NewServicePoolWithDefaults instantiates a new ServicePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServicePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServicePool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicePool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicePool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServicePool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServicePool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServicePool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServicePool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServicePool) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServicePool) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServicePool) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServicePool) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServicePool) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *ServicePool) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ServicePool) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ServicePool) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ServicePool) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *ServicePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ServicePool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServicePool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServicePool) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ServicePool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetReadyCount

`func (o *ServicePool) GetReadyCount() int64`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *ServicePool) GetReadyCountOk() (*int64, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *ServicePool) SetReadyCount(v int64)`

SetReadyCount sets ReadyCount field to given value.

### HasReadyCount

`func (o *ServicePool) HasReadyCount() bool`

HasReadyCount returns a boolean if a field has been set.

### GetDefinition

`func (o *ServicePool) GetDefinition() DeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ServicePool) GetDefinitionOk() (*DeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ServicePool) SetDefinition(v DeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ServicePool) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetStatus

`func (o *ServicePool) GetStatus() ServicePoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicePool) GetStatusOk() (*ServicePoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicePool) SetStatus(v ServicePoolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicePool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ServicePool) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ServicePool) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ServicePool) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ServicePool) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetGeneration

`func (o *ServicePool) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ServicePool) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ServicePool) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ServicePool) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetCustomerId

`func (o *ServicePool) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ServicePool) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ServicePool) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ServicePool) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


