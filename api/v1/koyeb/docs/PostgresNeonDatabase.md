# PostgresNeonDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgVersion** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to [**[]PostgresNeonDatabaseNeonRole**](PostgresNeonDatabaseNeonRole.md) |  | [optional] 
**Databases** | Pointer to [**[]PostgresNeonDatabaseNeonDatabase**](PostgresNeonDatabaseNeonDatabase.md) |  | [optional] 

## Methods

### NewPostgresNeonDatabase

`func NewPostgresNeonDatabase() *PostgresNeonDatabase`

NewPostgresNeonDatabase instantiates a new PostgresNeonDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresNeonDatabaseWithDefaults

`func NewPostgresNeonDatabaseWithDefaults() *PostgresNeonDatabase`

NewPostgresNeonDatabaseWithDefaults instantiates a new PostgresNeonDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgVersion

`func (o *PostgresNeonDatabase) GetPgVersion() int64`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *PostgresNeonDatabase) GetPgVersionOk() (*int64, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *PostgresNeonDatabase) SetPgVersion(v int64)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *PostgresNeonDatabase) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetRoles

`func (o *PostgresNeonDatabase) GetRoles() []PostgresNeonDatabaseNeonRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PostgresNeonDatabase) GetRolesOk() (*[]PostgresNeonDatabaseNeonRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PostgresNeonDatabase) SetRoles(v []PostgresNeonDatabaseNeonRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PostgresNeonDatabase) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDatabases

`func (o *PostgresNeonDatabase) GetDatabases() []PostgresNeonDatabaseNeonDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *PostgresNeonDatabase) GetDatabasesOk() (*[]PostgresNeonDatabaseNeonDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *PostgresNeonDatabase) SetDatabases(v []PostgresNeonDatabaseNeonDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *PostgresNeonDatabase) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


