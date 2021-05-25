# StackListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LatestRevisionSha** | Pointer to **string** |  | [optional] 
**DeployedRevisionSha** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StackStatus**](StackStatus.md) |  | [optional] [default to STACKSTATUS_UNKNOWN]
**Repository** | Pointer to [**SCMRepository**](SCMRepository.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 

## Methods

### NewStackListItem

`func NewStackListItem() *StackListItem`

NewStackListItem instantiates a new StackListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackListItemWithDefaults

`func NewStackListItemWithDefaults() *StackListItem`

NewStackListItemWithDefaults instantiates a new StackListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StackListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StackListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StackListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLatestRevisionSha

`func (o *StackListItem) GetLatestRevisionSha() string`

GetLatestRevisionSha returns the LatestRevisionSha field if non-nil, zero value otherwise.

### GetLatestRevisionShaOk

`func (o *StackListItem) GetLatestRevisionShaOk() (*string, bool)`

GetLatestRevisionShaOk returns a tuple with the LatestRevisionSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionSha

`func (o *StackListItem) SetLatestRevisionSha(v string)`

SetLatestRevisionSha sets LatestRevisionSha field to given value.

### HasLatestRevisionSha

`func (o *StackListItem) HasLatestRevisionSha() bool`

HasLatestRevisionSha returns a boolean if a field has been set.

### GetDeployedRevisionSha

`func (o *StackListItem) GetDeployedRevisionSha() string`

GetDeployedRevisionSha returns the DeployedRevisionSha field if non-nil, zero value otherwise.

### GetDeployedRevisionShaOk

`func (o *StackListItem) GetDeployedRevisionShaOk() (*string, bool)`

GetDeployedRevisionShaOk returns a tuple with the DeployedRevisionSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedRevisionSha

`func (o *StackListItem) SetDeployedRevisionSha(v string)`

SetDeployedRevisionSha sets DeployedRevisionSha field to given value.

### HasDeployedRevisionSha

`func (o *StackListItem) HasDeployedRevisionSha() bool`

HasDeployedRevisionSha returns a boolean if a field has been set.

### GetStatus

`func (o *StackListItem) GetStatus() StackStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StackListItem) GetStatusOk() (*StackStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StackListItem) SetStatus(v StackStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StackListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRepository

`func (o *StackListItem) GetRepository() SCMRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *StackListItem) GetRepositoryOk() (*SCMRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *StackListItem) SetRepository(v SCMRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *StackListItem) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StackListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StackListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StackListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StackListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StackListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StackListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StackListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StackListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *StackListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *StackListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *StackListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *StackListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


