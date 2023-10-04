# DatabaseSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostgresNeon** | Pointer to [**PostgresNeonDatabase**](PostgresNeonDatabase.md) |  | [optional] 

## Methods

### NewDatabaseSource

`func NewDatabaseSource() *DatabaseSource`

NewDatabaseSource instantiates a new DatabaseSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSourceWithDefaults

`func NewDatabaseSourceWithDefaults() *DatabaseSource`

NewDatabaseSourceWithDefaults instantiates a new DatabaseSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostgresNeon

`func (o *DatabaseSource) GetPostgresNeon() PostgresNeonDatabase`

GetPostgresNeon returns the PostgresNeon field if non-nil, zero value otherwise.

### GetPostgresNeonOk

`func (o *DatabaseSource) GetPostgresNeonOk() (*PostgresNeonDatabase, bool)`

GetPostgresNeonOk returns a tuple with the PostgresNeon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresNeon

`func (o *DatabaseSource) SetPostgresNeon(v PostgresNeonDatabase)`

SetPostgresNeon sets PostgresNeon field to given value.

### HasPostgresNeon

`func (o *DatabaseSource) HasPostgresNeon() bool`

HasPostgresNeon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


