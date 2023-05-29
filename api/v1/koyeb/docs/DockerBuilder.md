# DockerBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dockerfile** | Pointer to **string** |  | [optional] 
**Entrypoint** | Pointer to **[]string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewDockerBuilder

`func NewDockerBuilder() *DockerBuilder`

NewDockerBuilder instantiates a new DockerBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerBuilderWithDefaults

`func NewDockerBuilderWithDefaults() *DockerBuilder`

NewDockerBuilderWithDefaults instantiates a new DockerBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerfile

`func (o *DockerBuilder) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *DockerBuilder) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *DockerBuilder) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *DockerBuilder) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DockerBuilder) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DockerBuilder) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DockerBuilder) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DockerBuilder) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetCommand

`func (o *DockerBuilder) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DockerBuilder) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DockerBuilder) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *DockerBuilder) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetArgs

`func (o *DockerBuilder) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *DockerBuilder) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *DockerBuilder) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *DockerBuilder) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetTarget

`func (o *DockerBuilder) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DockerBuilder) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DockerBuilder) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DockerBuilder) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


