# DeletePersistentVolumeReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | Pointer to [**PersistentVolume**](PersistentVolume.md) |  | [optional] 

## Methods

### NewDeletePersistentVolumeReply

`func NewDeletePersistentVolumeReply() *DeletePersistentVolumeReply`

NewDeletePersistentVolumeReply instantiates a new DeletePersistentVolumeReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePersistentVolumeReplyWithDefaults

`func NewDeletePersistentVolumeReplyWithDefaults() *DeletePersistentVolumeReply`

NewDeletePersistentVolumeReplyWithDefaults instantiates a new DeletePersistentVolumeReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *DeletePersistentVolumeReply) GetVolume() PersistentVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *DeletePersistentVolumeReply) GetVolumeOk() (*PersistentVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *DeletePersistentVolumeReply) SetVolume(v PersistentVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *DeletePersistentVolumeReply) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


