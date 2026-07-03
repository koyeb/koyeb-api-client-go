# ListInstanceSnapshotsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSnapshots** | Pointer to [**[]InstanceSnapshot**](InstanceSnapshot.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListInstanceSnapshotsReply

`func NewListInstanceSnapshotsReply() *ListInstanceSnapshotsReply`

NewListInstanceSnapshotsReply instantiates a new ListInstanceSnapshotsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceSnapshotsReplyWithDefaults

`func NewListInstanceSnapshotsReplyWithDefaults() *ListInstanceSnapshotsReply`

NewListInstanceSnapshotsReplyWithDefaults instantiates a new ListInstanceSnapshotsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSnapshots

`func (o *ListInstanceSnapshotsReply) GetInstanceSnapshots() []InstanceSnapshot`

GetInstanceSnapshots returns the InstanceSnapshots field if non-nil, zero value otherwise.

### GetInstanceSnapshotsOk

`func (o *ListInstanceSnapshotsReply) GetInstanceSnapshotsOk() (*[]InstanceSnapshot, bool)`

GetInstanceSnapshotsOk returns a tuple with the InstanceSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSnapshots

`func (o *ListInstanceSnapshotsReply) SetInstanceSnapshots(v []InstanceSnapshot)`

SetInstanceSnapshots sets InstanceSnapshots field to given value.

### HasInstanceSnapshots

`func (o *ListInstanceSnapshotsReply) HasInstanceSnapshots() bool`

HasInstanceSnapshots returns a boolean if a field has been set.

### GetLimit

`func (o *ListInstanceSnapshotsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInstanceSnapshotsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInstanceSnapshotsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListInstanceSnapshotsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListInstanceSnapshotsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInstanceSnapshotsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInstanceSnapshotsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListInstanceSnapshotsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListInstanceSnapshotsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListInstanceSnapshotsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListInstanceSnapshotsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListInstanceSnapshotsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *ListInstanceSnapshotsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListInstanceSnapshotsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListInstanceSnapshotsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListInstanceSnapshotsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


