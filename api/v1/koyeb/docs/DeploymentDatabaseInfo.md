# DeploymentDatabaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeonPostgres** | Pointer to [**DeploymentNeonPostgresDatabaseInfo**](DeploymentNeonPostgresDatabaseInfo.md) |  | [optional] 

## Methods

### NewDeploymentDatabaseInfo

`func NewDeploymentDatabaseInfo() *DeploymentDatabaseInfo`

NewDeploymentDatabaseInfo instantiates a new DeploymentDatabaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDatabaseInfoWithDefaults

`func NewDeploymentDatabaseInfoWithDefaults() *DeploymentDatabaseInfo`

NewDeploymentDatabaseInfoWithDefaults instantiates a new DeploymentDatabaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeonPostgres

`func (o *DeploymentDatabaseInfo) GetNeonPostgres() DeploymentNeonPostgresDatabaseInfo`

GetNeonPostgres returns the NeonPostgres field if non-nil, zero value otherwise.

### GetNeonPostgresOk

`func (o *DeploymentDatabaseInfo) GetNeonPostgresOk() (*DeploymentNeonPostgresDatabaseInfo, bool)`

GetNeonPostgresOk returns a tuple with the NeonPostgres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeonPostgres

`func (o *DeploymentDatabaseInfo) SetNeonPostgres(v DeploymentNeonPostgresDatabaseInfo)`

SetNeonPostgres sets NeonPostgres field to given value.

### HasNeonPostgres

`func (o *DeploymentDatabaseInfo) HasNeonPostgres() bool`

HasNeonPostgres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


