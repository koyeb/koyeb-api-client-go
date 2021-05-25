# SCMRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SCMRepositoryType**](SCMRepositoryType.md) |  | [optional] [default to SCMREPOSITORYTYPE_GITHUB]
**Name** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** | The branch to track changes on. | [optional] 

## Methods

### NewSCMRepository

`func NewSCMRepository() *SCMRepository`

NewSCMRepository instantiates a new SCMRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCMRepositoryWithDefaults

`func NewSCMRepositoryWithDefaults() *SCMRepository`

NewSCMRepositoryWithDefaults instantiates a new SCMRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SCMRepository) GetType() SCMRepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SCMRepository) GetTypeOk() (*SCMRepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SCMRepository) SetType(v SCMRepositoryType)`

SetType sets Type field to given value.

### HasType

`func (o *SCMRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SCMRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCMRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCMRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SCMRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBranch

`func (o *SCMRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *SCMRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *SCMRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *SCMRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


