# DeploymentFileMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**InterpolationEnabled** | Pointer to **bool** |  | [optional] 
**Secret** | Pointer to [**SecretSource**](SecretSource.md) |  | [optional] 
**Raw** | Pointer to [**RawSource**](RawSource.md) |  | [optional] 

## Methods

### NewDeploymentFileMount

`func NewDeploymentFileMount() *DeploymentFileMount`

NewDeploymentFileMount instantiates a new DeploymentFileMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentFileMountWithDefaults

`func NewDeploymentFileMountWithDefaults() *DeploymentFileMount`

NewDeploymentFileMountWithDefaults instantiates a new DeploymentFileMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *DeploymentFileMount) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeploymentFileMount) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeploymentFileMount) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeploymentFileMount) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *DeploymentFileMount) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DeploymentFileMount) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DeploymentFileMount) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DeploymentFileMount) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetInterpolationEnabled

`func (o *DeploymentFileMount) GetInterpolationEnabled() bool`

GetInterpolationEnabled returns the InterpolationEnabled field if non-nil, zero value otherwise.

### GetInterpolationEnabledOk

`func (o *DeploymentFileMount) GetInterpolationEnabledOk() (*bool, bool)`

GetInterpolationEnabledOk returns a tuple with the InterpolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolationEnabled

`func (o *DeploymentFileMount) SetInterpolationEnabled(v bool)`

SetInterpolationEnabled sets InterpolationEnabled field to given value.

### HasInterpolationEnabled

`func (o *DeploymentFileMount) HasInterpolationEnabled() bool`

HasInterpolationEnabled returns a boolean if a field has been set.

### GetSecret

`func (o *DeploymentFileMount) GetSecret() SecretSource`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeploymentFileMount) GetSecretOk() (*SecretSource, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeploymentFileMount) SetSecret(v SecretSource)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DeploymentFileMount) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRaw

`func (o *DeploymentFileMount) GetRaw() RawSource`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *DeploymentFileMount) GetRawOk() (*RawSource, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *DeploymentFileMount) SetRaw(v RawSource)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *DeploymentFileMount) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


