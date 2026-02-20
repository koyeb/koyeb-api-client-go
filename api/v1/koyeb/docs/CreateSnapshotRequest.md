# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentVolumeId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSnapshotRequest

`func NewCreateSnapshotRequest() *CreateSnapshotRequest`

NewCreateSnapshotRequest instantiates a new CreateSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotRequestWithDefaults

`func NewCreateSnapshotRequestWithDefaults() *CreateSnapshotRequest`

NewCreateSnapshotRequestWithDefaults instantiates a new CreateSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentVolumeId

`func (o *CreateSnapshotRequest) GetParentVolumeId() string`

GetParentVolumeId returns the ParentVolumeId field if non-nil, zero value otherwise.

### GetParentVolumeIdOk

`func (o *CreateSnapshotRequest) GetParentVolumeIdOk() (*string, bool)`

GetParentVolumeIdOk returns a tuple with the ParentVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVolumeId

`func (o *CreateSnapshotRequest) SetParentVolumeId(v string)`

SetParentVolumeId sets ParentVolumeId field to given value.

### HasParentVolumeId

`func (o *CreateSnapshotRequest) HasParentVolumeId() bool`

HasParentVolumeId returns a boolean if a field has been set.

### GetName

`func (o *CreateSnapshotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSnapshotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateSnapshotRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateSnapshotRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateSnapshotRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateSnapshotRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


