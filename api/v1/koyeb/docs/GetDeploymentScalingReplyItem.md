# GetDeploymentScalingReplyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**ReplicaIndex** | Pointer to **int64** |  | [optional] 
**Instances** | Pointer to [**[]Instance**](Instance.md) | An array of &#x60;active&#x60; and &#x60;starting&#x60; instances.  Status of the active instance (and if none the most recent instance)  string status &#x3D; 4;  Status message of the active instance (and if none the most recent instance)  string message &#x3D; 5; | [optional] 

## Methods

### NewGetDeploymentScalingReplyItem

`func NewGetDeploymentScalingReplyItem() *GetDeploymentScalingReplyItem`

NewGetDeploymentScalingReplyItem instantiates a new GetDeploymentScalingReplyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentScalingReplyItemWithDefaults

`func NewGetDeploymentScalingReplyItemWithDefaults() *GetDeploymentScalingReplyItem`

NewGetDeploymentScalingReplyItemWithDefaults instantiates a new GetDeploymentScalingReplyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *GetDeploymentScalingReplyItem) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetDeploymentScalingReplyItem) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetDeploymentScalingReplyItem) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetDeploymentScalingReplyItem) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReplicaIndex

`func (o *GetDeploymentScalingReplyItem) GetReplicaIndex() int64`

GetReplicaIndex returns the ReplicaIndex field if non-nil, zero value otherwise.

### GetReplicaIndexOk

`func (o *GetDeploymentScalingReplyItem) GetReplicaIndexOk() (*int64, bool)`

GetReplicaIndexOk returns a tuple with the ReplicaIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaIndex

`func (o *GetDeploymentScalingReplyItem) SetReplicaIndex(v int64)`

SetReplicaIndex sets ReplicaIndex field to given value.

### HasReplicaIndex

`func (o *GetDeploymentScalingReplyItem) HasReplicaIndex() bool`

HasReplicaIndex returns a boolean if a field has been set.

### GetInstances

`func (o *GetDeploymentScalingReplyItem) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetDeploymentScalingReplyItem) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetDeploymentScalingReplyItem) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetDeploymentScalingReplyItem) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


