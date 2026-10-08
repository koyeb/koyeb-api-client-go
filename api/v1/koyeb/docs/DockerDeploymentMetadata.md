# DockerDeploymentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolvedDigest** | Pointer to **string** | Digest of the resolved image manifest (\&quot;sha256:...\&quot;). | [optional] 
**ResolvedEntrypoint** | Pointer to **[]string** | ENTRYPOINT from the image config. | [optional] 
**ResolvedCmd** | Pointer to **[]string** | CMD from the image config. | [optional] 

## Methods

### NewDockerDeploymentMetadata

`func NewDockerDeploymentMetadata() *DockerDeploymentMetadata`

NewDockerDeploymentMetadata instantiates a new DockerDeploymentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerDeploymentMetadataWithDefaults

`func NewDockerDeploymentMetadataWithDefaults() *DockerDeploymentMetadata`

NewDockerDeploymentMetadataWithDefaults instantiates a new DockerDeploymentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolvedDigest

`func (o *DockerDeploymentMetadata) GetResolvedDigest() string`

GetResolvedDigest returns the ResolvedDigest field if non-nil, zero value otherwise.

### GetResolvedDigestOk

`func (o *DockerDeploymentMetadata) GetResolvedDigestOk() (*string, bool)`

GetResolvedDigestOk returns a tuple with the ResolvedDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDigest

`func (o *DockerDeploymentMetadata) SetResolvedDigest(v string)`

SetResolvedDigest sets ResolvedDigest field to given value.

### HasResolvedDigest

`func (o *DockerDeploymentMetadata) HasResolvedDigest() bool`

HasResolvedDigest returns a boolean if a field has been set.

### GetResolvedEntrypoint

`func (o *DockerDeploymentMetadata) GetResolvedEntrypoint() []string`

GetResolvedEntrypoint returns the ResolvedEntrypoint field if non-nil, zero value otherwise.

### GetResolvedEntrypointOk

`func (o *DockerDeploymentMetadata) GetResolvedEntrypointOk() (*[]string, bool)`

GetResolvedEntrypointOk returns a tuple with the ResolvedEntrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedEntrypoint

`func (o *DockerDeploymentMetadata) SetResolvedEntrypoint(v []string)`

SetResolvedEntrypoint sets ResolvedEntrypoint field to given value.

### HasResolvedEntrypoint

`func (o *DockerDeploymentMetadata) HasResolvedEntrypoint() bool`

HasResolvedEntrypoint returns a boolean if a field has been set.

### GetResolvedCmd

`func (o *DockerDeploymentMetadata) GetResolvedCmd() []string`

GetResolvedCmd returns the ResolvedCmd field if non-nil, zero value otherwise.

### GetResolvedCmdOk

`func (o *DockerDeploymentMetadata) GetResolvedCmdOk() (*[]string, bool)`

GetResolvedCmdOk returns a tuple with the ResolvedCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedCmd

`func (o *DockerDeploymentMetadata) SetResolvedCmd(v []string)`

SetResolvedCmd sets ResolvedCmd field to given value.

### HasResolvedCmd

`func (o *DockerDeploymentMetadata) HasResolvedCmd() bool`

HasResolvedCmd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


