# GetPersistentVolumeReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | Pointer to [**PersistentVolume**](PersistentVolume.md) |  | [optional] 

## Methods

### NewGetPersistentVolumeReply

`func NewGetPersistentVolumeReply() *GetPersistentVolumeReply`

NewGetPersistentVolumeReply instantiates a new GetPersistentVolumeReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPersistentVolumeReplyWithDefaults

`func NewGetPersistentVolumeReplyWithDefaults() *GetPersistentVolumeReply`

NewGetPersistentVolumeReplyWithDefaults instantiates a new GetPersistentVolumeReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *GetPersistentVolumeReply) GetVolume() PersistentVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetPersistentVolumeReply) GetVolumeOk() (*PersistentVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetPersistentVolumeReply) SetVolume(v PersistentVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetPersistentVolumeReply) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


