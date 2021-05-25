# SCMReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SCMReferenceSCMType**](SCMReferenceSCMType.md) |  | [optional] [default to SCMREFERENCESCMTYPE_GITHUB]
**Repo** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 

## Methods

### NewSCMReference

`func NewSCMReference() *SCMReference`

NewSCMReference instantiates a new SCMReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCMReferenceWithDefaults

`func NewSCMReferenceWithDefaults() *SCMReference`

NewSCMReferenceWithDefaults instantiates a new SCMReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SCMReference) GetType() SCMReferenceSCMType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SCMReference) GetTypeOk() (*SCMReferenceSCMType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SCMReference) SetType(v SCMReferenceSCMType)`

SetType sets Type field to given value.

### HasType

`func (o *SCMReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRepo

`func (o *SCMReference) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SCMReference) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SCMReference) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *SCMReference) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRef

`func (o *SCMReference) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SCMReference) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SCMReference) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SCMReference) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


