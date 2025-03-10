# QueryLogsReplyPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**NextStart** | Pointer to **time.Time** |  | [optional] 
**NextEnd** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewQueryLogsReplyPagination

`func NewQueryLogsReplyPagination() *QueryLogsReplyPagination`

NewQueryLogsReplyPagination instantiates a new QueryLogsReplyPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLogsReplyPaginationWithDefaults

`func NewQueryLogsReplyPaginationWithDefaults() *QueryLogsReplyPagination`

NewQueryLogsReplyPaginationWithDefaults instantiates a new QueryLogsReplyPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *QueryLogsReplyPagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QueryLogsReplyPagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QueryLogsReplyPagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *QueryLogsReplyPagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextStart

`func (o *QueryLogsReplyPagination) GetNextStart() time.Time`

GetNextStart returns the NextStart field if non-nil, zero value otherwise.

### GetNextStartOk

`func (o *QueryLogsReplyPagination) GetNextStartOk() (*time.Time, bool)`

GetNextStartOk returns a tuple with the NextStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStart

`func (o *QueryLogsReplyPagination) SetNextStart(v time.Time)`

SetNextStart sets NextStart field to given value.

### HasNextStart

`func (o *QueryLogsReplyPagination) HasNextStart() bool`

HasNextStart returns a boolean if a field has been set.

### GetNextEnd

`func (o *QueryLogsReplyPagination) GetNextEnd() time.Time`

GetNextEnd returns the NextEnd field if non-nil, zero value otherwise.

### GetNextEndOk

`func (o *QueryLogsReplyPagination) GetNextEndOk() (*time.Time, bool)`

GetNextEndOk returns a tuple with the NextEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEnd

`func (o *QueryLogsReplyPagination) SetNextEnd(v time.Time)`

SetNextEnd sets NextEnd field to given value.

### HasNextEnd

`func (o *QueryLogsReplyPagination) HasNextEnd() bool`

HasNextEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


