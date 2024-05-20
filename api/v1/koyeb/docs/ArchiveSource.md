# ArchiveSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Buildpack** | Pointer to [**BuildpackBuilder**](BuildpackBuilder.md) |  | [optional] 
**Docker** | Pointer to [**DockerBuilder**](DockerBuilder.md) |  | [optional] 

## Methods

### NewArchiveSource

`func NewArchiveSource() *ArchiveSource`

NewArchiveSource instantiates a new ArchiveSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveSourceWithDefaults

`func NewArchiveSourceWithDefaults() *ArchiveSource`

NewArchiveSourceWithDefaults instantiates a new ArchiveSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArchiveSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuildpack

`func (o *ArchiveSource) GetBuildpack() BuildpackBuilder`

GetBuildpack returns the Buildpack field if non-nil, zero value otherwise.

### GetBuildpackOk

`func (o *ArchiveSource) GetBuildpackOk() (*BuildpackBuilder, bool)`

GetBuildpackOk returns a tuple with the Buildpack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpack

`func (o *ArchiveSource) SetBuildpack(v BuildpackBuilder)`

SetBuildpack sets Buildpack field to given value.

### HasBuildpack

`func (o *ArchiveSource) HasBuildpack() bool`

HasBuildpack returns a boolean if a field has been set.

### GetDocker

`func (o *ArchiveSource) GetDocker() DockerBuilder`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *ArchiveSource) GetDockerOk() (*DockerBuilder, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *ArchiveSource) SetDocker(v DockerBuilder)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *ArchiveSource) HasDocker() bool`

HasDocker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


