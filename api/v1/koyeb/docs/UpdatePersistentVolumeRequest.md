# UpdatePersistentVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**MaxSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdatePersistentVolumeRequest

`func NewUpdatePersistentVolumeRequest() *UpdatePersistentVolumeRequest`

NewUpdatePersistentVolumeRequest instantiates a new UpdatePersistentVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePersistentVolumeRequestWithDefaults

`func NewUpdatePersistentVolumeRequestWithDefaults() *UpdatePersistentVolumeRequest`

NewUpdatePersistentVolumeRequestWithDefaults instantiates a new UpdatePersistentVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePersistentVolumeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePersistentVolumeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePersistentVolumeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePersistentVolumeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxSize

`func (o *UpdatePersistentVolumeRequest) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *UpdatePersistentVolumeRequest) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *UpdatePersistentVolumeRequest) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *UpdatePersistentVolumeRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


