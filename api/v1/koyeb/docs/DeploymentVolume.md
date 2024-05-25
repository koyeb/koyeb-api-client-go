# DeploymentVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ReplicaIndex** | Pointer to **int64** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeploymentVolume

`func NewDeploymentVolume() *DeploymentVolume`

NewDeploymentVolume instantiates a new DeploymentVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentVolumeWithDefaults

`func NewDeploymentVolumeWithDefaults() *DeploymentVolume`

NewDeploymentVolumeWithDefaults instantiates a new DeploymentVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DeploymentVolume) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeploymentVolume) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeploymentVolume) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeploymentVolume) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReplicaIndex

`func (o *DeploymentVolume) GetReplicaIndex() int64`

GetReplicaIndex returns the ReplicaIndex field if non-nil, zero value otherwise.

### GetReplicaIndexOk

`func (o *DeploymentVolume) GetReplicaIndexOk() (*int64, bool)`

GetReplicaIndexOk returns a tuple with the ReplicaIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaIndex

`func (o *DeploymentVolume) SetReplicaIndex(v int64)`

SetReplicaIndex sets ReplicaIndex field to given value.

### HasReplicaIndex

`func (o *DeploymentVolume) HasReplicaIndex() bool`

HasReplicaIndex returns a boolean if a field has been set.

### GetScopes

`func (o *DeploymentVolume) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeploymentVolume) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeploymentVolume) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeploymentVolume) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


