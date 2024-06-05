# UpdatePersistentVolumeReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | Pointer to [**PersistentVolume**](PersistentVolume.md) |  | [optional] 

## Methods

### NewUpdatePersistentVolumeReply

`func NewUpdatePersistentVolumeReply() *UpdatePersistentVolumeReply`

NewUpdatePersistentVolumeReply instantiates a new UpdatePersistentVolumeReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePersistentVolumeReplyWithDefaults

`func NewUpdatePersistentVolumeReplyWithDefaults() *UpdatePersistentVolumeReply`

NewUpdatePersistentVolumeReplyWithDefaults instantiates a new UpdatePersistentVolumeReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *UpdatePersistentVolumeReply) GetVolume() PersistentVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UpdatePersistentVolumeReply) GetVolumeOk() (*PersistentVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UpdatePersistentVolumeReply) SetVolume(v PersistentVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *UpdatePersistentVolumeReply) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


