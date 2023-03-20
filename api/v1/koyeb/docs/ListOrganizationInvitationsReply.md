# ListOrganizationInvitationsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]OrganizationInvitation**](OrganizationInvitation.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListOrganizationInvitationsReply

`func NewListOrganizationInvitationsReply() *ListOrganizationInvitationsReply`

NewListOrganizationInvitationsReply instantiates a new ListOrganizationInvitationsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationInvitationsReplyWithDefaults

`func NewListOrganizationInvitationsReplyWithDefaults() *ListOrganizationInvitationsReply`

NewListOrganizationInvitationsReplyWithDefaults instantiates a new ListOrganizationInvitationsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *ListOrganizationInvitationsReply) GetInvitations() []OrganizationInvitation`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *ListOrganizationInvitationsReply) GetInvitationsOk() (*[]OrganizationInvitation, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *ListOrganizationInvitationsReply) SetInvitations(v []OrganizationInvitation)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *ListOrganizationInvitationsReply) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetLimit

`func (o *ListOrganizationInvitationsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOrganizationInvitationsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOrganizationInvitationsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListOrganizationInvitationsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListOrganizationInvitationsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOrganizationInvitationsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOrganizationInvitationsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListOrganizationInvitationsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListOrganizationInvitationsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOrganizationInvitationsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOrganizationInvitationsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListOrganizationInvitationsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


