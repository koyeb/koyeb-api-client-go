# CreateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**SecretType**](SecretType.md) |  | [optional] [default to SECRETTYPE_SIMPLE]
**Value** | Pointer to **string** |  | [optional] 
**DockerHubRegistry** | Pointer to [**DockerHubRegistryConfiguration**](DockerHubRegistryConfiguration.md) |  | [optional] 
**PrivateRegistry** | Pointer to [**PrivateRegistryConfiguration**](PrivateRegistryConfiguration.md) |  | [optional] 
**DigitalOceanRegistry** | Pointer to [**DigitalOceanRegistryConfiguration**](DigitalOceanRegistryConfiguration.md) |  | [optional] 
**GithubRegistry** | Pointer to [**GitHubRegistryConfiguration**](GitHubRegistryConfiguration.md) |  | [optional] 
**GitlabRegistry** | Pointer to [**GitLabRegistryConfiguration**](GitLabRegistryConfiguration.md) |  | [optional] 
**GcpContainerRegistry** | Pointer to [**GCPContainerRegistryConfiguration**](GCPContainerRegistryConfiguration.md) |  | [optional] 
**AzureContainerRegistry** | Pointer to [**AzureContainerRegistryConfiguration**](AzureContainerRegistryConfiguration.md) |  | [optional] 

## Methods

### NewCreateSecret

`func NewCreateSecret() *CreateSecret`

NewCreateSecret instantiates a new CreateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretWithDefaults

`func NewCreateSecretWithDefaults() *CreateSecret`

NewCreateSecretWithDefaults instantiates a new CreateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateSecret) GetType() SecretType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSecret) GetTypeOk() (*SecretType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSecret) SetType(v SecretType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSecret) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CreateSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDockerHubRegistry

`func (o *CreateSecret) GetDockerHubRegistry() DockerHubRegistryConfiguration`

GetDockerHubRegistry returns the DockerHubRegistry field if non-nil, zero value otherwise.

### GetDockerHubRegistryOk

`func (o *CreateSecret) GetDockerHubRegistryOk() (*DockerHubRegistryConfiguration, bool)`

GetDockerHubRegistryOk returns a tuple with the DockerHubRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHubRegistry

`func (o *CreateSecret) SetDockerHubRegistry(v DockerHubRegistryConfiguration)`

SetDockerHubRegistry sets DockerHubRegistry field to given value.

### HasDockerHubRegistry

`func (o *CreateSecret) HasDockerHubRegistry() bool`

HasDockerHubRegistry returns a boolean if a field has been set.

### GetPrivateRegistry

`func (o *CreateSecret) GetPrivateRegistry() PrivateRegistryConfiguration`

GetPrivateRegistry returns the PrivateRegistry field if non-nil, zero value otherwise.

### GetPrivateRegistryOk

`func (o *CreateSecret) GetPrivateRegistryOk() (*PrivateRegistryConfiguration, bool)`

GetPrivateRegistryOk returns a tuple with the PrivateRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegistry

`func (o *CreateSecret) SetPrivateRegistry(v PrivateRegistryConfiguration)`

SetPrivateRegistry sets PrivateRegistry field to given value.

### HasPrivateRegistry

`func (o *CreateSecret) HasPrivateRegistry() bool`

HasPrivateRegistry returns a boolean if a field has been set.

### GetDigitalOceanRegistry

`func (o *CreateSecret) GetDigitalOceanRegistry() DigitalOceanRegistryConfiguration`

GetDigitalOceanRegistry returns the DigitalOceanRegistry field if non-nil, zero value otherwise.

### GetDigitalOceanRegistryOk

`func (o *CreateSecret) GetDigitalOceanRegistryOk() (*DigitalOceanRegistryConfiguration, bool)`

GetDigitalOceanRegistryOk returns a tuple with the DigitalOceanRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalOceanRegistry

`func (o *CreateSecret) SetDigitalOceanRegistry(v DigitalOceanRegistryConfiguration)`

SetDigitalOceanRegistry sets DigitalOceanRegistry field to given value.

### HasDigitalOceanRegistry

`func (o *CreateSecret) HasDigitalOceanRegistry() bool`

HasDigitalOceanRegistry returns a boolean if a field has been set.

### GetGithubRegistry

`func (o *CreateSecret) GetGithubRegistry() GitHubRegistryConfiguration`

GetGithubRegistry returns the GithubRegistry field if non-nil, zero value otherwise.

### GetGithubRegistryOk

`func (o *CreateSecret) GetGithubRegistryOk() (*GitHubRegistryConfiguration, bool)`

GetGithubRegistryOk returns a tuple with the GithubRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubRegistry

`func (o *CreateSecret) SetGithubRegistry(v GitHubRegistryConfiguration)`

SetGithubRegistry sets GithubRegistry field to given value.

### HasGithubRegistry

`func (o *CreateSecret) HasGithubRegistry() bool`

HasGithubRegistry returns a boolean if a field has been set.

### GetGitlabRegistry

`func (o *CreateSecret) GetGitlabRegistry() GitLabRegistryConfiguration`

GetGitlabRegistry returns the GitlabRegistry field if non-nil, zero value otherwise.

### GetGitlabRegistryOk

`func (o *CreateSecret) GetGitlabRegistryOk() (*GitLabRegistryConfiguration, bool)`

GetGitlabRegistryOk returns a tuple with the GitlabRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabRegistry

`func (o *CreateSecret) SetGitlabRegistry(v GitLabRegistryConfiguration)`

SetGitlabRegistry sets GitlabRegistry field to given value.

### HasGitlabRegistry

`func (o *CreateSecret) HasGitlabRegistry() bool`

HasGitlabRegistry returns a boolean if a field has been set.

### GetGcpContainerRegistry

`func (o *CreateSecret) GetGcpContainerRegistry() GCPContainerRegistryConfiguration`

GetGcpContainerRegistry returns the GcpContainerRegistry field if non-nil, zero value otherwise.

### GetGcpContainerRegistryOk

`func (o *CreateSecret) GetGcpContainerRegistryOk() (*GCPContainerRegistryConfiguration, bool)`

GetGcpContainerRegistryOk returns a tuple with the GcpContainerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpContainerRegistry

`func (o *CreateSecret) SetGcpContainerRegistry(v GCPContainerRegistryConfiguration)`

SetGcpContainerRegistry sets GcpContainerRegistry field to given value.

### HasGcpContainerRegistry

`func (o *CreateSecret) HasGcpContainerRegistry() bool`

HasGcpContainerRegistry returns a boolean if a field has been set.

### GetAzureContainerRegistry

`func (o *CreateSecret) GetAzureContainerRegistry() AzureContainerRegistryConfiguration`

GetAzureContainerRegistry returns the AzureContainerRegistry field if non-nil, zero value otherwise.

### GetAzureContainerRegistryOk

`func (o *CreateSecret) GetAzureContainerRegistryOk() (*AzureContainerRegistryConfiguration, bool)`

GetAzureContainerRegistryOk returns a tuple with the AzureContainerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureContainerRegistry

`func (o *CreateSecret) SetAzureContainerRegistry(v AzureContainerRegistryConfiguration)`

SetAzureContainerRegistry sets AzureContainerRegistry field to given value.

### HasAzureContainerRegistry

`func (o *CreateSecret) HasAzureContainerRegistry() bool`

HasAzureContainerRegistry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


