# ListDatacentersReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenters** | Pointer to [**[]DatacenterListItem**](DatacenterListItem.md) |  | [optional] 

## Methods

### NewListDatacentersReply

`func NewListDatacentersReply() *ListDatacentersReply`

NewListDatacentersReply instantiates a new ListDatacentersReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatacentersReplyWithDefaults

`func NewListDatacentersReplyWithDefaults() *ListDatacentersReply`

NewListDatacentersReplyWithDefaults instantiates a new ListDatacentersReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenters

`func (o *ListDatacentersReply) GetDatacenters() []DatacenterListItem`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *ListDatacentersReply) GetDatacentersOk() (*[]DatacenterListItem, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *ListDatacentersReply) SetDatacenters(v []DatacenterListItem)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *ListDatacentersReply) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


