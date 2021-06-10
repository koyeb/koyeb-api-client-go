# ListInternalServicesReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceIds** | Pointer to [**[]ServiceId**](ServiceId.md) |  | [optional] 
**Services** | Pointer to [**[]ServiceListItem**](ServiceListItem.md) |  | [optional] 

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

### GetServices

`func (o *ListInternalServicesReply) GetServices() []ServiceListItem`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListInternalServicesReply) GetServicesOk() (*[]ServiceListItem, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListInternalServicesReply) SetServices(v []ServiceListItem)`

SetServices sets Services field to given value.

### HasServices

`func (o *ListInternalServicesReply) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


