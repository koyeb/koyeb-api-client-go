# DeploymentEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**ValueFromSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentEnv

`func NewDeploymentEnv() *DeploymentEnv`

NewDeploymentEnv instantiates a new DeploymentEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentEnvWithDefaults

`func NewDeploymentEnvWithDefaults() *DeploymentEnv`

NewDeploymentEnvWithDefaults instantiates a new DeploymentEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *DeploymentEnv) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeploymentEnv) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeploymentEnv) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeploymentEnv) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetKey

`func (o *DeploymentEnv) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeploymentEnv) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeploymentEnv) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DeploymentEnv) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DeploymentEnv) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeploymentEnv) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeploymentEnv) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeploymentEnv) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSecret

`func (o *DeploymentEnv) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeploymentEnv) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeploymentEnv) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DeploymentEnv) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetValueFromSecret

`func (o *DeploymentEnv) GetValueFromSecret() string`

GetValueFromSecret returns the ValueFromSecret field if non-nil, zero value otherwise.

### GetValueFromSecretOk

`func (o *DeploymentEnv) GetValueFromSecretOk() (*string, bool)`

GetValueFromSecretOk returns a tuple with the ValueFromSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFromSecret

`func (o *DeploymentEnv) SetValueFromSecret(v string)`

SetValueFromSecret sets ValueFromSecret field to given value.

### HasValueFromSecret

`func (o *DeploymentEnv) HasValueFromSecret() bool`

HasValueFromSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


