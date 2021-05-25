# Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**LatestRevision** | Pointer to [**StackRevision**](StackRevision.md) |  | [optional] 
**LatestRevisionSha** | Pointer to **string** |  | [optional] 
**DeployedRevisionSha** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**SCMRepository**](SCMRepository.md) |  | [optional] 
**Status** | Pointer to [**StackStatus**](StackStatus.md) |  | [optional] [default to STACKSTATUS_UNKNOWN]
**StatusMessage** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStack

`func NewStack() *Stack`

NewStack instantiates a new Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackWithDefaults

`func NewStackWithDefaults() *Stack`

NewStackWithDefaults instantiates a new Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stack) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Stack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Stack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Stack) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Stack) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Stack) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Stack) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetLatestRevision

`func (o *Stack) GetLatestRevision() StackRevision`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *Stack) GetLatestRevisionOk() (*StackRevision, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *Stack) SetLatestRevision(v StackRevision)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *Stack) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetLatestRevisionSha

`func (o *Stack) GetLatestRevisionSha() string`

GetLatestRevisionSha returns the LatestRevisionSha field if non-nil, zero value otherwise.

### GetLatestRevisionShaOk

`func (o *Stack) GetLatestRevisionShaOk() (*string, bool)`

GetLatestRevisionShaOk returns a tuple with the LatestRevisionSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionSha

`func (o *Stack) SetLatestRevisionSha(v string)`

SetLatestRevisionSha sets LatestRevisionSha field to given value.

### HasLatestRevisionSha

`func (o *Stack) HasLatestRevisionSha() bool`

HasLatestRevisionSha returns a boolean if a field has been set.

### GetDeployedRevisionSha

`func (o *Stack) GetDeployedRevisionSha() string`

GetDeployedRevisionSha returns the DeployedRevisionSha field if non-nil, zero value otherwise.

### GetDeployedRevisionShaOk

`func (o *Stack) GetDeployedRevisionShaOk() (*string, bool)`

GetDeployedRevisionShaOk returns a tuple with the DeployedRevisionSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedRevisionSha

`func (o *Stack) SetDeployedRevisionSha(v string)`

SetDeployedRevisionSha sets DeployedRevisionSha field to given value.

### HasDeployedRevisionSha

`func (o *Stack) HasDeployedRevisionSha() bool`

HasDeployedRevisionSha returns a boolean if a field has been set.

### GetRepository

`func (o *Stack) GetRepository() SCMRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Stack) GetRepositoryOk() (*SCMRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Stack) SetRepository(v SCMRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Stack) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetStatus

`func (o *Stack) GetStatus() StackStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Stack) GetStatusOk() (*StackStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Stack) SetStatus(v StackStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Stack) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Stack) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Stack) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Stack) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Stack) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Stack) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Stack) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Stack) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Stack) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Stack) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Stack) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Stack) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Stack) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


