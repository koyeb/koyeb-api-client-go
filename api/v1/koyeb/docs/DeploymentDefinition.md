# DeploymentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to [**[]DeploymentRoute**](DeploymentRoute.md) |  | [optional] 
**Ports** | Pointer to [**[]DeploymentPort**](DeploymentPort.md) |  | [optional] 
**Env** | Pointer to [**[]DeploymentEnv**](DeploymentEnv.md) |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Scalings** | Pointer to [**[]DeploymentScaling**](DeploymentScaling.md) |  | [optional] 
**InstanceTypes** | Pointer to [**[]DeploymentInstanceType**](DeploymentInstanceType.md) |  | [optional] 
**Scaling** | Pointer to [**DeploymentScaling**](DeploymentScaling.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Docker** | Pointer to [**DockerSource**](DockerSource.md) |  | [optional] 
**Git** | Pointer to [**GitSource**](GitSource.md) |  | [optional] 

## Methods

### NewDeploymentDefinition

`func NewDeploymentDefinition() *DeploymentDefinition`

NewDeploymentDefinition instantiates a new DeploymentDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDefinitionWithDefaults

`func NewDeploymentDefinitionWithDefaults() *DeploymentDefinition`

NewDeploymentDefinitionWithDefaults instantiates a new DeploymentDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *DeploymentDefinition) GetRoutes() []DeploymentRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *DeploymentDefinition) GetRoutesOk() (*[]DeploymentRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *DeploymentDefinition) SetRoutes(v []DeploymentRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *DeploymentDefinition) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetPorts

`func (o *DeploymentDefinition) GetPorts() []DeploymentPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *DeploymentDefinition) GetPortsOk() (*[]DeploymentPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *DeploymentDefinition) SetPorts(v []DeploymentPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *DeploymentDefinition) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnv

`func (o *DeploymentDefinition) GetEnv() []DeploymentEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *DeploymentDefinition) GetEnvOk() (*[]DeploymentEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *DeploymentDefinition) SetEnv(v []DeploymentEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *DeploymentDefinition) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetRegions

`func (o *DeploymentDefinition) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *DeploymentDefinition) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *DeploymentDefinition) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *DeploymentDefinition) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetScalings

`func (o *DeploymentDefinition) GetScalings() []DeploymentScaling`

GetScalings returns the Scalings field if non-nil, zero value otherwise.

### GetScalingsOk

`func (o *DeploymentDefinition) GetScalingsOk() (*[]DeploymentScaling, bool)`

GetScalingsOk returns a tuple with the Scalings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalings

`func (o *DeploymentDefinition) SetScalings(v []DeploymentScaling)`

SetScalings sets Scalings field to given value.

### HasScalings

`func (o *DeploymentDefinition) HasScalings() bool`

HasScalings returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *DeploymentDefinition) GetInstanceTypes() []DeploymentInstanceType`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *DeploymentDefinition) GetInstanceTypesOk() (*[]DeploymentInstanceType, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *DeploymentDefinition) SetInstanceTypes(v []DeploymentInstanceType)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *DeploymentDefinition) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetScaling

`func (o *DeploymentDefinition) GetScaling() DeploymentScaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *DeploymentDefinition) GetScalingOk() (*DeploymentScaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *DeploymentDefinition) SetScaling(v DeploymentScaling)`

SetScaling sets Scaling field to given value.

### HasScaling

`func (o *DeploymentDefinition) HasScaling() bool`

HasScaling returns a boolean if a field has been set.

### GetInstanceType

`func (o *DeploymentDefinition) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DeploymentDefinition) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DeploymentDefinition) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DeploymentDefinition) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDocker

`func (o *DeploymentDefinition) GetDocker() DockerSource`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *DeploymentDefinition) GetDockerOk() (*DockerSource, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *DeploymentDefinition) SetDocker(v DockerSource)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *DeploymentDefinition) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGit

`func (o *DeploymentDefinition) GetGit() GitSource`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *DeploymentDefinition) GetGitOk() (*GitSource, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *DeploymentDefinition) SetGit(v GitSource)`

SetGit sets Git field to given value.

### HasGit

`func (o *DeploymentDefinition) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


