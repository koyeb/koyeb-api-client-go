# DeploymentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DeploymentDefinitionType**](DeploymentDefinitionType.md) |  | [optional] [default to DEPLOYMENTDEFINITIONTYPE_INVALID]
**Strategy** | Pointer to [**DeploymentStrategy**](DeploymentStrategy.md) |  | [optional] 
**Routes** | Pointer to [**[]DeploymentRoute**](DeploymentRoute.md) |  | [optional] 
**Ports** | Pointer to [**[]DeploymentPort**](DeploymentPort.md) |  | [optional] 
**Env** | Pointer to [**[]DeploymentEnv**](DeploymentEnv.md) |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**Scalings** | Pointer to [**[]DeploymentScaling**](DeploymentScaling.md) |  | [optional] 
**InstanceTypes** | Pointer to [**[]DeploymentInstanceType**](DeploymentInstanceType.md) |  | [optional] 
**HealthChecks** | Pointer to [**[]DeploymentHealthCheck**](DeploymentHealthCheck.md) |  | [optional] 
**Volumes** | Pointer to [**[]DeploymentVolume**](DeploymentVolume.md) |  | [optional] 
**ConfigFiles** | Pointer to [**[]ConfigFile**](ConfigFile.md) |  | [optional] 
**SkipCache** | Pointer to **bool** |  | [optional] 
**Docker** | Pointer to [**DockerSource**](DockerSource.md) |  | [optional] 
**Git** | Pointer to [**GitSource**](GitSource.md) |  | [optional] 
**Database** | Pointer to [**DatabaseSource**](DatabaseSource.md) |  | [optional] 
**Archive** | Pointer to [**ArchiveSource**](ArchiveSource.md) |  | [optional] 

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

### GetType

`func (o *DeploymentDefinition) GetType() DeploymentDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentDefinition) GetTypeOk() (*DeploymentDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentDefinition) SetType(v DeploymentDefinitionType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStrategy

`func (o *DeploymentDefinition) GetStrategy() DeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *DeploymentDefinition) GetStrategyOk() (*DeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *DeploymentDefinition) SetStrategy(v DeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *DeploymentDefinition) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

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

### GetHealthChecks

`func (o *DeploymentDefinition) GetHealthChecks() []DeploymentHealthCheck`

GetHealthChecks returns the HealthChecks field if non-nil, zero value otherwise.

### GetHealthChecksOk

`func (o *DeploymentDefinition) GetHealthChecksOk() (*[]DeploymentHealthCheck, bool)`

GetHealthChecksOk returns a tuple with the HealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChecks

`func (o *DeploymentDefinition) SetHealthChecks(v []DeploymentHealthCheck)`

SetHealthChecks sets HealthChecks field to given value.

### HasHealthChecks

`func (o *DeploymentDefinition) HasHealthChecks() bool`

HasHealthChecks returns a boolean if a field has been set.

### GetVolumes

`func (o *DeploymentDefinition) GetVolumes() []DeploymentVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DeploymentDefinition) GetVolumesOk() (*[]DeploymentVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DeploymentDefinition) SetVolumes(v []DeploymentVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *DeploymentDefinition) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetConfigFiles

`func (o *DeploymentDefinition) GetConfigFiles() []ConfigFile`

GetConfigFiles returns the ConfigFiles field if non-nil, zero value otherwise.

### GetConfigFilesOk

`func (o *DeploymentDefinition) GetConfigFilesOk() (*[]ConfigFile, bool)`

GetConfigFilesOk returns a tuple with the ConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFiles

`func (o *DeploymentDefinition) SetConfigFiles(v []ConfigFile)`

SetConfigFiles sets ConfigFiles field to given value.

### HasConfigFiles

`func (o *DeploymentDefinition) HasConfigFiles() bool`

HasConfigFiles returns a boolean if a field has been set.

### GetSkipCache

`func (o *DeploymentDefinition) GetSkipCache() bool`

GetSkipCache returns the SkipCache field if non-nil, zero value otherwise.

### GetSkipCacheOk

`func (o *DeploymentDefinition) GetSkipCacheOk() (*bool, bool)`

GetSkipCacheOk returns a tuple with the SkipCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCache

`func (o *DeploymentDefinition) SetSkipCache(v bool)`

SetSkipCache sets SkipCache field to given value.

### HasSkipCache

`func (o *DeploymentDefinition) HasSkipCache() bool`

HasSkipCache returns a boolean if a field has been set.

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

### GetDatabase

`func (o *DeploymentDefinition) GetDatabase() DatabaseSource`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DeploymentDefinition) GetDatabaseOk() (*DatabaseSource, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DeploymentDefinition) SetDatabase(v DatabaseSource)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DeploymentDefinition) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetArchive

`func (o *DeploymentDefinition) GetArchive() ArchiveSource`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *DeploymentDefinition) GetArchiveOk() (*ArchiveSource, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *DeploymentDefinition) SetArchive(v ArchiveSource)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *DeploymentDefinition) HasArchive() bool`

HasArchive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


