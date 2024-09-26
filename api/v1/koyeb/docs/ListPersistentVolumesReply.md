# ListPersistentVolumesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]PersistentVolume**](PersistentVolume.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListPersistentVolumesReply

`func NewListPersistentVolumesReply() *ListPersistentVolumesReply`

NewListPersistentVolumesReply instantiates a new ListPersistentVolumesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPersistentVolumesReplyWithDefaults

`func NewListPersistentVolumesReplyWithDefaults() *ListPersistentVolumesReply`

NewListPersistentVolumesReplyWithDefaults instantiates a new ListPersistentVolumesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *ListPersistentVolumesReply) GetVolumes() []PersistentVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListPersistentVolumesReply) GetVolumesOk() (*[]PersistentVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListPersistentVolumesReply) SetVolumes(v []PersistentVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListPersistentVolumesReply) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetLimit

`func (o *ListPersistentVolumesReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPersistentVolumesReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPersistentVolumesReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPersistentVolumesReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListPersistentVolumesReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPersistentVolumesReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPersistentVolumesReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPersistentVolumesReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetHasNext

`func (o *ListPersistentVolumesReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListPersistentVolumesReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListPersistentVolumesReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListPersistentVolumesReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


