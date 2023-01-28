# ListOrganizationMembersReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]OrganizationMember**](OrganizationMember.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListOrganizationMembersReply

`func NewListOrganizationMembersReply() *ListOrganizationMembersReply`

NewListOrganizationMembersReply instantiates a new ListOrganizationMembersReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationMembersReplyWithDefaults

`func NewListOrganizationMembersReplyWithDefaults() *ListOrganizationMembersReply`

NewListOrganizationMembersReplyWithDefaults instantiates a new ListOrganizationMembersReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListOrganizationMembersReply) GetMembers() []OrganizationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListOrganizationMembersReply) GetMembersOk() (*[]OrganizationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListOrganizationMembersReply) SetMembers(v []OrganizationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ListOrganizationMembersReply) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetLimit

`func (o *ListOrganizationMembersReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOrganizationMembersReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOrganizationMembersReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListOrganizationMembersReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListOrganizationMembersReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOrganizationMembersReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOrganizationMembersReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListOrganizationMembersReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListOrganizationMembersReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOrganizationMembersReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOrganizationMembersReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListOrganizationMembersReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


