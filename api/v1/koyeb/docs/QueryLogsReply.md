# QueryLogsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LogEntry**](LogEntry.md) |  | [optional] 
**Pagination** | Pointer to [**QueryLogsReplyPagination**](QueryLogsReplyPagination.md) |  | [optional] 

## Methods

### NewQueryLogsReply

`func NewQueryLogsReply() *QueryLogsReply`

NewQueryLogsReply instantiates a new QueryLogsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLogsReplyWithDefaults

`func NewQueryLogsReplyWithDefaults() *QueryLogsReply`

NewQueryLogsReplyWithDefaults instantiates a new QueryLogsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *QueryLogsReply) GetData() []LogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryLogsReply) GetDataOk() (*[]LogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryLogsReply) SetData(v []LogEntry)`

SetData sets Data field to given value.

### HasData

`func (o *QueryLogsReply) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *QueryLogsReply) GetPagination() QueryLogsReplyPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *QueryLogsReply) GetPaginationOk() (*QueryLogsReplyPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *QueryLogsReply) SetPagination(v QueryLogsReplyPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *QueryLogsReply) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


