# LogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Labels** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewLogEntry

`func NewLogEntry() *LogEntry`

NewLogEntry instantiates a new LogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntryWithDefaults

`func NewLogEntryWithDefaults() *LogEntry`

NewLogEntryWithDefaults instantiates a new LogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *LogEntry) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *LogEntry) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *LogEntry) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *LogEntry) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LogEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *LogEntry) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LogEntry) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LogEntry) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LogEntry) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


