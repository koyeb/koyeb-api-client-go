# CreateCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] [default to CREDENTIALTYPE_INVALID]
**OrganizationId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCredential

`func NewCreateCredential() *CreateCredential`

NewCreateCredential instantiates a new CreateCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCredentialWithDefaults

`func NewCreateCredentialWithDefaults() *CreateCredential`

NewCreateCredentialWithDefaults instantiates a new CreateCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateCredential) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCredential) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCredential) SetType(v CredentialType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateCredential) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateCredential) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateCredential) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateCredential) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


