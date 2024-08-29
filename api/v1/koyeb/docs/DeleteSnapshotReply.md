# DeleteSnapshotReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to [**Snapshot**](Snapshot.md) |  | [optional] 

## Methods

### NewDeleteSnapshotReply

`func NewDeleteSnapshotReply() *DeleteSnapshotReply`

NewDeleteSnapshotReply instantiates a new DeleteSnapshotReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSnapshotReplyWithDefaults

`func NewDeleteSnapshotReplyWithDefaults() *DeleteSnapshotReply`

NewDeleteSnapshotReplyWithDefaults instantiates a new DeleteSnapshotReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *DeleteSnapshotReply) GetSnapshot() Snapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *DeleteSnapshotReply) GetSnapshotOk() (*Snapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *DeleteSnapshotReply) SetSnapshot(v Snapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *DeleteSnapshotReply) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


