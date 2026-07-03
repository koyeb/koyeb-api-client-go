# CreateInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**InstanceSnapshotType**](InstanceSnapshotType.md) |  | [optional] [default to INSTANCESNAPSHOTTYPE_INVALID]

## Methods

### NewCreateInstanceSnapshotRequest

`func NewCreateInstanceSnapshotRequest() *CreateInstanceSnapshotRequest`

NewCreateInstanceSnapshotRequest instantiates a new CreateInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceSnapshotRequestWithDefaults

`func NewCreateInstanceSnapshotRequestWithDefaults() *CreateInstanceSnapshotRequest`

NewCreateInstanceSnapshotRequestWithDefaults instantiates a new CreateInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *CreateInstanceSnapshotRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateInstanceSnapshotRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateInstanceSnapshotRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CreateInstanceSnapshotRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetName

`func (o *CreateInstanceSnapshotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInstanceSnapshotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInstanceSnapshotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInstanceSnapshotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateInstanceSnapshotRequest) GetType() InstanceSnapshotType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateInstanceSnapshotRequest) GetTypeOk() (*InstanceSnapshotType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateInstanceSnapshotRequest) SetType(v InstanceSnapshotType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateInstanceSnapshotRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


