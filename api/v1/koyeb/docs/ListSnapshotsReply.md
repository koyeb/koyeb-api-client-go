# ListSnapshotsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]Snapshot**](Snapshot.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListSnapshotsReply

`func NewListSnapshotsReply() *ListSnapshotsReply`

NewListSnapshotsReply instantiates a new ListSnapshotsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSnapshotsReplyWithDefaults

`func NewListSnapshotsReplyWithDefaults() *ListSnapshotsReply`

NewListSnapshotsReplyWithDefaults instantiates a new ListSnapshotsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *ListSnapshotsReply) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListSnapshotsReply) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListSnapshotsReply) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ListSnapshotsReply) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetLimit

`func (o *ListSnapshotsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListSnapshotsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListSnapshotsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListSnapshotsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListSnapshotsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListSnapshotsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListSnapshotsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListSnapshotsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetHasNext

`func (o *ListSnapshotsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListSnapshotsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListSnapshotsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListSnapshotsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


