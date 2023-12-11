# MembersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** |  | [optional] 
**InvitationsByStatus** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMembersSummary

`func NewMembersSummary() *MembersSummary`

NewMembersSummary instantiates a new MembersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembersSummaryWithDefaults

`func NewMembersSummaryWithDefaults() *MembersSummary`

NewMembersSummaryWithDefaults instantiates a new MembersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MembersSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MembersSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MembersSummary) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MembersSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetInvitationsByStatus

`func (o *MembersSummary) GetInvitationsByStatus() map[string]string`

GetInvitationsByStatus returns the InvitationsByStatus field if non-nil, zero value otherwise.

### GetInvitationsByStatusOk

`func (o *MembersSummary) GetInvitationsByStatusOk() (*map[string]string, bool)`

GetInvitationsByStatusOk returns a tuple with the InvitationsByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationsByStatus

`func (o *MembersSummary) SetInvitationsByStatus(v map[string]string)`

SetInvitationsByStatus sets InvitationsByStatus field to given value.

### HasInvitationsByStatus

`func (o *MembersSummary) HasInvitationsByStatus() bool`

HasInvitationsByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


