# KgitproxyRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**KgitproxyRepositoryProvider**](KgitproxyRepositoryProvider.md) |  | [optional] [default to KGITPROXYREPOSITORYPROVIDER_UNKNOWN]
**Github** | Pointer to [**KgitproxyGitHubRepository**](KgitproxyGitHubRepository.md) |  | [optional] 

## Methods

### NewKgitproxyRepository

`func NewKgitproxyRepository() *KgitproxyRepository`

NewKgitproxyRepository instantiates a new KgitproxyRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKgitproxyRepositoryWithDefaults

`func NewKgitproxyRepositoryWithDefaults() *KgitproxyRepository`

NewKgitproxyRepositoryWithDefaults instantiates a new KgitproxyRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KgitproxyRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KgitproxyRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KgitproxyRepository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KgitproxyRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *KgitproxyRepository) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KgitproxyRepository) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KgitproxyRepository) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *KgitproxyRepository) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *KgitproxyRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KgitproxyRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KgitproxyRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KgitproxyRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *KgitproxyRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KgitproxyRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KgitproxyRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KgitproxyRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *KgitproxyRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KgitproxyRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KgitproxyRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KgitproxyRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPrivate

`func (o *KgitproxyRepository) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *KgitproxyRepository) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *KgitproxyRepository) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *KgitproxyRepository) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *KgitproxyRepository) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *KgitproxyRepository) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *KgitproxyRepository) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *KgitproxyRepository) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetProvider

`func (o *KgitproxyRepository) GetProvider() KgitproxyRepositoryProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KgitproxyRepository) GetProviderOk() (*KgitproxyRepositoryProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KgitproxyRepository) SetProvider(v KgitproxyRepositoryProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KgitproxyRepository) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetGithub

`func (o *KgitproxyRepository) GetGithub() KgitproxyGitHubRepository`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *KgitproxyRepository) GetGithubOk() (*KgitproxyGitHubRepository, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *KgitproxyRepository) SetGithub(v KgitproxyGitHubRepository)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *KgitproxyRepository) HasGithub() bool`

HasGithub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


