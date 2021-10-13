# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SecretType**](SecretType.md) |  | [optional] [default to SECRETTYPE_SIMPLE]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**DockerHubRegistry** | Pointer to [**DockerHubRegistryConfiguration**](DockerHubRegistryConfiguration.md) |  | [optional] 
**PrivateRegistry** | Pointer to [**PrivateRegistryConfiguration**](PrivateRegistryConfiguration.md) |  | [optional] 
**DigitalOceanRegistry** | Pointer to [**DigitalOceanRegistryConfiguration**](DigitalOceanRegistryConfiguration.md) |  | [optional] 
**GithubRegistry** | Pointer to [**GitHubRegistryConfiguration**](GitHubRegistryConfiguration.md) |  | [optional] 
**GitlabRegistry** | Pointer to [**GitLabRegistryConfiguration**](GitLabRegistryConfiguration.md) |  | [optional] 
**GcpContainerRegistry** | Pointer to [**GCPContainerRegistryConfiguration**](GCPContainerRegistryConfiguration.md) |  | [optional] 
**AzureContainerRegistry** | Pointer to [**AzureContainerRegistryConfiguration**](AzureContainerRegistryConfiguration.md) |  | [optional] 

## Methods

### NewSecret

`func NewSecret() *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Secret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Secret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Secret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Secret) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Secret) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Secret) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Secret) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetType

`func (o *Secret) GetType() SecretType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Secret) GetTypeOk() (*SecretType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Secret) SetType(v SecretType)`

SetType sets Type field to given value.

### HasType

`func (o *Secret) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Secret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Secret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Secret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Secret) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Secret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Secret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Secret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Secret) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetValue

`func (o *Secret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Secret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Secret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Secret) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDockerHubRegistry

`func (o *Secret) GetDockerHubRegistry() DockerHubRegistryConfiguration`

GetDockerHubRegistry returns the DockerHubRegistry field if non-nil, zero value otherwise.

### GetDockerHubRegistryOk

`func (o *Secret) GetDockerHubRegistryOk() (*DockerHubRegistryConfiguration, bool)`

GetDockerHubRegistryOk returns a tuple with the DockerHubRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHubRegistry

`func (o *Secret) SetDockerHubRegistry(v DockerHubRegistryConfiguration)`

SetDockerHubRegistry sets DockerHubRegistry field to given value.

### HasDockerHubRegistry

`func (o *Secret) HasDockerHubRegistry() bool`

HasDockerHubRegistry returns a boolean if a field has been set.

### GetPrivateRegistry

`func (o *Secret) GetPrivateRegistry() PrivateRegistryConfiguration`

GetPrivateRegistry returns the PrivateRegistry field if non-nil, zero value otherwise.

### GetPrivateRegistryOk

`func (o *Secret) GetPrivateRegistryOk() (*PrivateRegistryConfiguration, bool)`

GetPrivateRegistryOk returns a tuple with the PrivateRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegistry

`func (o *Secret) SetPrivateRegistry(v PrivateRegistryConfiguration)`

SetPrivateRegistry sets PrivateRegistry field to given value.

### HasPrivateRegistry

`func (o *Secret) HasPrivateRegistry() bool`

HasPrivateRegistry returns a boolean if a field has been set.

### GetDigitalOceanRegistry

`func (o *Secret) GetDigitalOceanRegistry() DigitalOceanRegistryConfiguration`

GetDigitalOceanRegistry returns the DigitalOceanRegistry field if non-nil, zero value otherwise.

### GetDigitalOceanRegistryOk

`func (o *Secret) GetDigitalOceanRegistryOk() (*DigitalOceanRegistryConfiguration, bool)`

GetDigitalOceanRegistryOk returns a tuple with the DigitalOceanRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalOceanRegistry

`func (o *Secret) SetDigitalOceanRegistry(v DigitalOceanRegistryConfiguration)`

SetDigitalOceanRegistry sets DigitalOceanRegistry field to given value.

### HasDigitalOceanRegistry

`func (o *Secret) HasDigitalOceanRegistry() bool`

HasDigitalOceanRegistry returns a boolean if a field has been set.

### GetGithubRegistry

`func (o *Secret) GetGithubRegistry() GitHubRegistryConfiguration`

GetGithubRegistry returns the GithubRegistry field if non-nil, zero value otherwise.

### GetGithubRegistryOk

`func (o *Secret) GetGithubRegistryOk() (*GitHubRegistryConfiguration, bool)`

GetGithubRegistryOk returns a tuple with the GithubRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubRegistry

`func (o *Secret) SetGithubRegistry(v GitHubRegistryConfiguration)`

SetGithubRegistry sets GithubRegistry field to given value.

### HasGithubRegistry

`func (o *Secret) HasGithubRegistry() bool`

HasGithubRegistry returns a boolean if a field has been set.

### GetGitlabRegistry

`func (o *Secret) GetGitlabRegistry() GitLabRegistryConfiguration`

GetGitlabRegistry returns the GitlabRegistry field if non-nil, zero value otherwise.

### GetGitlabRegistryOk

`func (o *Secret) GetGitlabRegistryOk() (*GitLabRegistryConfiguration, bool)`

GetGitlabRegistryOk returns a tuple with the GitlabRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabRegistry

`func (o *Secret) SetGitlabRegistry(v GitLabRegistryConfiguration)`

SetGitlabRegistry sets GitlabRegistry field to given value.

### HasGitlabRegistry

`func (o *Secret) HasGitlabRegistry() bool`

HasGitlabRegistry returns a boolean if a field has been set.

### GetGcpContainerRegistry

`func (o *Secret) GetGcpContainerRegistry() GCPContainerRegistryConfiguration`

GetGcpContainerRegistry returns the GcpContainerRegistry field if non-nil, zero value otherwise.

### GetGcpContainerRegistryOk

`func (o *Secret) GetGcpContainerRegistryOk() (*GCPContainerRegistryConfiguration, bool)`

GetGcpContainerRegistryOk returns a tuple with the GcpContainerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpContainerRegistry

`func (o *Secret) SetGcpContainerRegistry(v GCPContainerRegistryConfiguration)`

SetGcpContainerRegistry sets GcpContainerRegistry field to given value.

### HasGcpContainerRegistry

`func (o *Secret) HasGcpContainerRegistry() bool`

HasGcpContainerRegistry returns a boolean if a field has been set.

### GetAzureContainerRegistry

`func (o *Secret) GetAzureContainerRegistry() AzureContainerRegistryConfiguration`

GetAzureContainerRegistry returns the AzureContainerRegistry field if non-nil, zero value otherwise.

### GetAzureContainerRegistryOk

`func (o *Secret) GetAzureContainerRegistryOk() (*AzureContainerRegistryConfiguration, bool)`

GetAzureContainerRegistryOk returns a tuple with the AzureContainerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureContainerRegistry

`func (o *Secret) SetAzureContainerRegistry(v AzureContainerRegistryConfiguration)`

SetAzureContainerRegistry sets AzureContainerRegistry field to given value.

### HasAzureContainerRegistry

`func (o *Secret) HasAzureContainerRegistry() bool`

HasAzureContainerRegistry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


