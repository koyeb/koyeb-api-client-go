# NeonPostgresDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgVersion** | Pointer to **int64** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]NeonPostgresDatabaseNeonRole**](NeonPostgresDatabaseNeonRole.md) |  | [optional] 
**Databases** | Pointer to [**[]NeonPostgresDatabaseNeonDatabase**](NeonPostgresDatabaseNeonDatabase.md) |  | [optional] 

## Methods

### NewNeonPostgresDatabase

`func NewNeonPostgresDatabase() *NeonPostgresDatabase`

NewNeonPostgresDatabase instantiates a new NeonPostgresDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeonPostgresDatabaseWithDefaults

`func NewNeonPostgresDatabaseWithDefaults() *NeonPostgresDatabase`

NewNeonPostgresDatabaseWithDefaults instantiates a new NeonPostgresDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgVersion

`func (o *NeonPostgresDatabase) GetPgVersion() int64`

GetPgVersion returns the PgVersion field if non-nil, zero value otherwise.

### GetPgVersionOk

`func (o *NeonPostgresDatabase) GetPgVersionOk() (*int64, bool)`

GetPgVersionOk returns a tuple with the PgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgVersion

`func (o *NeonPostgresDatabase) SetPgVersion(v int64)`

SetPgVersion sets PgVersion field to given value.

### HasPgVersion

`func (o *NeonPostgresDatabase) HasPgVersion() bool`

HasPgVersion returns a boolean if a field has been set.

### GetRegion

`func (o *NeonPostgresDatabase) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NeonPostgresDatabase) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NeonPostgresDatabase) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NeonPostgresDatabase) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoles

`func (o *NeonPostgresDatabase) GetRoles() []NeonPostgresDatabaseNeonRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NeonPostgresDatabase) GetRolesOk() (*[]NeonPostgresDatabaseNeonRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NeonPostgresDatabase) SetRoles(v []NeonPostgresDatabaseNeonRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NeonPostgresDatabase) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDatabases

`func (o *NeonPostgresDatabase) GetDatabases() []NeonPostgresDatabaseNeonDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *NeonPostgresDatabase) GetDatabasesOk() (*[]NeonPostgresDatabaseNeonDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *NeonPostgresDatabase) SetDatabases(v []NeonPostgresDatabaseNeonDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *NeonPostgresDatabase) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


