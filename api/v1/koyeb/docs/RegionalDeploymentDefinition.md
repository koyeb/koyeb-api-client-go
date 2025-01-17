# RegionalDeploymentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**RegionalDeploymentDefinitionType**](RegionalDeploymentDefinitionType.md) |  | [optional] [default to REGIONALDEPLOYMENTDEFINITIONTYPE_INVALID]
**Strategy** | Pointer to [**DeploymentStrategy**](DeploymentStrategy.md) |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Ports** | Pointer to [**[]Port**](Port.md) |  | [optional] 
**Env** | Pointer to [**[]Env**](Env.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Scaling** | Pointer to [**Scaling**](Scaling.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**HealthChecks** | Pointer to [**[]DeploymentHealthCheck**](DeploymentHealthCheck.md) |  | [optional] 
**Volumes** | Pointer to [**[]RegionalDeploymentVolume**](RegionalDeploymentVolume.md) |  | [optional] 
**ConfigFiles** | Pointer to [**[]ConfigFile**](ConfigFile.md) |  | [optional] 
**SkipCache** | Pointer to **bool** |  | [optional] 
**Docker** | Pointer to [**DockerSource**](DockerSource.md) |  | [optional] 
**Git** | Pointer to [**GitSource**](GitSource.md) |  | [optional] 
**Archive** | Pointer to [**ArchiveSource**](ArchiveSource.md) |  | [optional] 

## Methods

### NewRegionalDeploymentDefinition

`func NewRegionalDeploymentDefinition() *RegionalDeploymentDefinition`

NewRegionalDeploymentDefinition instantiates a new RegionalDeploymentDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionalDeploymentDefinitionWithDefaults

`func NewRegionalDeploymentDefinitionWithDefaults() *RegionalDeploymentDefinition`

NewRegionalDeploymentDefinitionWithDefaults instantiates a new RegionalDeploymentDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegionalDeploymentDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionalDeploymentDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionalDeploymentDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionalDeploymentDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RegionalDeploymentDefinition) GetType() RegionalDeploymentDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegionalDeploymentDefinition) GetTypeOk() (*RegionalDeploymentDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegionalDeploymentDefinition) SetType(v RegionalDeploymentDefinitionType)`

SetType sets Type field to given value.

### HasType

`func (o *RegionalDeploymentDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStrategy

`func (o *RegionalDeploymentDefinition) GetStrategy() DeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *RegionalDeploymentDefinition) GetStrategyOk() (*DeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *RegionalDeploymentDefinition) SetStrategy(v DeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *RegionalDeploymentDefinition) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetRoutes

`func (o *RegionalDeploymentDefinition) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *RegionalDeploymentDefinition) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *RegionalDeploymentDefinition) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *RegionalDeploymentDefinition) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetPorts

`func (o *RegionalDeploymentDefinition) GetPorts() []Port`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *RegionalDeploymentDefinition) GetPortsOk() (*[]Port, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *RegionalDeploymentDefinition) SetPorts(v []Port)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *RegionalDeploymentDefinition) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetEnv

`func (o *RegionalDeploymentDefinition) GetEnv() []Env`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RegionalDeploymentDefinition) GetEnvOk() (*[]Env, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RegionalDeploymentDefinition) SetEnv(v []Env)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RegionalDeploymentDefinition) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetRegion

`func (o *RegionalDeploymentDefinition) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionalDeploymentDefinition) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionalDeploymentDefinition) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RegionalDeploymentDefinition) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScaling

`func (o *RegionalDeploymentDefinition) GetScaling() Scaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *RegionalDeploymentDefinition) GetScalingOk() (*Scaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *RegionalDeploymentDefinition) SetScaling(v Scaling)`

SetScaling sets Scaling field to given value.

### HasScaling

`func (o *RegionalDeploymentDefinition) HasScaling() bool`

HasScaling returns a boolean if a field has been set.

### GetInstanceType

`func (o *RegionalDeploymentDefinition) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RegionalDeploymentDefinition) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RegionalDeploymentDefinition) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *RegionalDeploymentDefinition) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *RegionalDeploymentDefinition) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *RegionalDeploymentDefinition) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *RegionalDeploymentDefinition) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *RegionalDeploymentDefinition) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetHealthChecks

`func (o *RegionalDeploymentDefinition) GetHealthChecks() []DeploymentHealthCheck`

GetHealthChecks returns the HealthChecks field if non-nil, zero value otherwise.

### GetHealthChecksOk

`func (o *RegionalDeploymentDefinition) GetHealthChecksOk() (*[]DeploymentHealthCheck, bool)`

GetHealthChecksOk returns a tuple with the HealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthChecks

`func (o *RegionalDeploymentDefinition) SetHealthChecks(v []DeploymentHealthCheck)`

SetHealthChecks sets HealthChecks field to given value.

### HasHealthChecks

`func (o *RegionalDeploymentDefinition) HasHealthChecks() bool`

HasHealthChecks returns a boolean if a field has been set.

### GetVolumes

`func (o *RegionalDeploymentDefinition) GetVolumes() []RegionalDeploymentVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *RegionalDeploymentDefinition) GetVolumesOk() (*[]RegionalDeploymentVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *RegionalDeploymentDefinition) SetVolumes(v []RegionalDeploymentVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *RegionalDeploymentDefinition) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetConfigFiles

`func (o *RegionalDeploymentDefinition) GetConfigFiles() []ConfigFile`

GetConfigFiles returns the ConfigFiles field if non-nil, zero value otherwise.

### GetConfigFilesOk

`func (o *RegionalDeploymentDefinition) GetConfigFilesOk() (*[]ConfigFile, bool)`

GetConfigFilesOk returns a tuple with the ConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFiles

`func (o *RegionalDeploymentDefinition) SetConfigFiles(v []ConfigFile)`

SetConfigFiles sets ConfigFiles field to given value.

### HasConfigFiles

`func (o *RegionalDeploymentDefinition) HasConfigFiles() bool`

HasConfigFiles returns a boolean if a field has been set.

### GetSkipCache

`func (o *RegionalDeploymentDefinition) GetSkipCache() bool`

GetSkipCache returns the SkipCache field if non-nil, zero value otherwise.

### GetSkipCacheOk

`func (o *RegionalDeploymentDefinition) GetSkipCacheOk() (*bool, bool)`

GetSkipCacheOk returns a tuple with the SkipCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCache

`func (o *RegionalDeploymentDefinition) SetSkipCache(v bool)`

SetSkipCache sets SkipCache field to given value.

### HasSkipCache

`func (o *RegionalDeploymentDefinition) HasSkipCache() bool`

HasSkipCache returns a boolean if a field has been set.

### GetDocker

`func (o *RegionalDeploymentDefinition) GetDocker() DockerSource`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *RegionalDeploymentDefinition) GetDockerOk() (*DockerSource, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *RegionalDeploymentDefinition) SetDocker(v DockerSource)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *RegionalDeploymentDefinition) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGit

`func (o *RegionalDeploymentDefinition) GetGit() GitSource`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *RegionalDeploymentDefinition) GetGitOk() (*GitSource, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *RegionalDeploymentDefinition) SetGit(v GitSource)`

SetGit sets Git field to given value.

### HasGit

`func (o *RegionalDeploymentDefinition) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetArchive

`func (o *RegionalDeploymentDefinition) GetArchive() ArchiveSource`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RegionalDeploymentDefinition) GetArchiveOk() (*ArchiveSource, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RegionalDeploymentDefinition) SetArchive(v ArchiveSource)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RegionalDeploymentDefinition) HasArchive() bool`

HasArchive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


