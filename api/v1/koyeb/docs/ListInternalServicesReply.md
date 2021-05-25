# ListInternalServicesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** |  | [optional] 
**ServiceIds** | Pointer to [**[]ServiceId**](ServiceId.md) |  | [optional] 

## Methods

### NewListInternalServicesReply

`func NewListInternalServicesReply() *ListInternalServicesReply`

NewListInternalServicesReply instantiates a new ListInternalServicesReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternalServicesReplyWithDefaults

`func NewListInternalServicesReplyWithDefaults() *ListInternalServicesReply`

NewListInternalServicesReplyWithDefaults instantiates a new ListInternalServicesReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListInternalServicesReply) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListInternalServicesReply) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListInternalServicesReply) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListInternalServicesReply) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetServiceIds

`func (o *ListInternalServicesReply) GetServiceIds() []ServiceId`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *ListInternalServicesReply) GetServiceIdsOk() (*[]ServiceId, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *ListInternalServicesReply) SetServiceIds(v []ServiceId)`

SetServiceIds sets ServiceIds field to given value.

### HasServiceIds

`func (o *ListInternalServicesReply) HasServiceIds() bool`

HasServiceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


