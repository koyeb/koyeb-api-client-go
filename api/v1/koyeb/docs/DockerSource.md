# DockerSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**ImageRegistrySecret** | Pointer to **string** |  | [optional] 
**Entrypoint** | Pointer to **[]string** |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 

## Methods

### NewDockerSource

`func NewDockerSource() *DockerSource`

NewDockerSource instantiates a new DockerSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerSourceWithDefaults

`func NewDockerSourceWithDefaults() *DockerSource`

NewDockerSourceWithDefaults instantiates a new DockerSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *DockerSource) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DockerSource) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DockerSource) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DockerSource) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetCommand

`func (o *DockerSource) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DockerSource) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DockerSource) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *DockerSource) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetArgs

`func (o *DockerSource) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *DockerSource) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *DockerSource) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *DockerSource) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetImageRegistrySecret

`func (o *DockerSource) GetImageRegistrySecret() string`

GetImageRegistrySecret returns the ImageRegistrySecret field if non-nil, zero value otherwise.

### GetImageRegistrySecretOk

`func (o *DockerSource) GetImageRegistrySecretOk() (*string, bool)`

GetImageRegistrySecretOk returns a tuple with the ImageRegistrySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegistrySecret

`func (o *DockerSource) SetImageRegistrySecret(v string)`

SetImageRegistrySecret sets ImageRegistrySecret field to given value.

### HasImageRegistrySecret

`func (o *DockerSource) HasImageRegistrySecret() bool`

HasImageRegistrySecret returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DockerSource) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DockerSource) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DockerSource) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DockerSource) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetPrivileged

`func (o *DockerSource) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *DockerSource) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *DockerSource) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *DockerSource) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


