# DeploymentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | Pointer to [**TriggerDeploymentMetadata**](TriggerDeploymentMetadata.md) |  | [optional] 
**Database** | Pointer to [**DatabaseDeploymentMetadata**](DatabaseDeploymentMetadata.md) |  | [optional] 
**Git** | Pointer to [**GitDeploymentMetadata**](GitDeploymentMetadata.md) |  | [optional] 
**Archive** | Pointer to [**ArchiveDeploymentMetadata**](ArchiveDeploymentMetadata.md) |  | [optional] 

## Methods

### NewDeploymentMetadata

`func NewDeploymentMetadata() *DeploymentMetadata`

NewDeploymentMetadata instantiates a new DeploymentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentMetadataWithDefaults

`func NewDeploymentMetadataWithDefaults() *DeploymentMetadata`

NewDeploymentMetadataWithDefaults instantiates a new DeploymentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *DeploymentMetadata) GetTrigger() TriggerDeploymentMetadata`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DeploymentMetadata) GetTriggerOk() (*TriggerDeploymentMetadata, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DeploymentMetadata) SetTrigger(v TriggerDeploymentMetadata)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DeploymentMetadata) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetDatabase

`func (o *DeploymentMetadata) GetDatabase() DatabaseDeploymentMetadata`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DeploymentMetadata) GetDatabaseOk() (*DatabaseDeploymentMetadata, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DeploymentMetadata) SetDatabase(v DatabaseDeploymentMetadata)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DeploymentMetadata) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetGit

`func (o *DeploymentMetadata) GetGit() GitDeploymentMetadata`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *DeploymentMetadata) GetGitOk() (*GitDeploymentMetadata, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *DeploymentMetadata) SetGit(v GitDeploymentMetadata)`

SetGit sets Git field to given value.

### HasGit

`func (o *DeploymentMetadata) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetArchive

`func (o *DeploymentMetadata) GetArchive() ArchiveDeploymentMetadata`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *DeploymentMetadata) GetArchiveOk() (*ArchiveDeploymentMetadata, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *DeploymentMetadata) SetArchive(v ArchiveDeploymentMetadata)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *DeploymentMetadata) HasArchive() bool`

HasArchive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


