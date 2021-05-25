# StackUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**SCMRepository**](SCMRepository.md) |  | [optional] 

## Methods

### NewStackUpsert

`func NewStackUpsert() *StackUpsert`

NewStackUpsert instantiates a new StackUpsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackUpsertWithDefaults

`func NewStackUpsertWithDefaults() *StackUpsert`

NewStackUpsertWithDefaults instantiates a new StackUpsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StackUpsert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackUpsert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackUpsert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackUpsert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *StackUpsert) GetRepository() SCMRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *StackUpsert) GetRepositoryOk() (*SCMRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *StackUpsert) SetRepository(v SCMRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *StackUpsert) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


