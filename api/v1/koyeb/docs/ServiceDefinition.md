# ServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Ports** | Pointer to [**[]Port**](Port.md) |  | [optional] 
**Env** | Pointer to [**[]Env**](Env.md) |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Scaling** | Pointer to [**Scaling**](Scaling.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**Docker** | Pointer to [**DockerSource**](DockerSource.md) |  | [optional] 
**Git** | Pointer to [**GitSource**](GitSource.md) |  | [optional] 

## Methods

### NewServiceDefinition

`func NewServiceDefinition() *ServiceDefinition`

NewServiceDefinition instantiates a new ServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDefinitionWithDefaults

`func NewServiceDefinitionWithDefaults() *ServiceDefinition`

NewServiceDefinitionWithDefaults instantiates a new ServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *ServiceDefinition) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServiceDefinition) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServiceDefinition) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServiceDefinition) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetPorts

`func (o *ServiceDefinition) GetPorts() []Port`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceDefinition) GetPortsOk() (*[]Port, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceDefinition) SetPorts(v []Port)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ServiceDefinition) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnv

`func (o *ServiceDefinition) GetEnv() []Env`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServiceDefinition) GetEnvOk() (*[]Env, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServiceDefinition) SetEnv(v []Env)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ServiceDefinition) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetRegions

`func (o *ServiceDefinition) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ServiceDefinition) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ServiceDefinition) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ServiceDefinition) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetScaling

`func (o *ServiceDefinition) GetScaling() Scaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *ServiceDefinition) GetScalingOk() (*Scaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *ServiceDefinition) SetScaling(v Scaling)`

SetScaling sets Scaling field to given value.

### HasScaling

`func (o *ServiceDefinition) HasScaling() bool`

HasScaling returns a boolean if a field has been set.

### GetInstanceType

`func (o *ServiceDefinition) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ServiceDefinition) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ServiceDefinition) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ServiceDefinition) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *ServiceDefinition) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *ServiceDefinition) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *ServiceDefinition) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *ServiceDefinition) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetDocker

`func (o *ServiceDefinition) GetDocker() DockerSource`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *ServiceDefinition) GetDockerOk() (*DockerSource, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *ServiceDefinition) SetDocker(v DockerSource)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *ServiceDefinition) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGit

`func (o *ServiceDefinition) GetGit() GitSource`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ServiceDefinition) GetGitOk() (*GitSource, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ServiceDefinition) SetGit(v GitSource)`

SetGit sets Git field to given value.

### HasGit

`func (o *ServiceDefinition) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


