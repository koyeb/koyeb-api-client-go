# ListUserOrganizationInvitationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]OrganizationInvitationStatus**](OrganizationInvitationStatus.md) |  | [optional] 

## Methods

### NewListUserOrganizationInvitationsRequest

`func NewListUserOrganizationInvitationsRequest() *ListUserOrganizationInvitationsRequest`

NewListUserOrganizationInvitationsRequest instantiates a new ListUserOrganizationInvitationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserOrganizationInvitationsRequestWithDefaults

`func NewListUserOrganizationInvitationsRequestWithDefaults() *ListUserOrganizationInvitationsRequest`

NewListUserOrganizationInvitationsRequestWithDefaults instantiates a new ListUserOrganizationInvitationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListUserOrganizationInvitationsRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListUserOrganizationInvitationsRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListUserOrganizationInvitationsRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListUserOrganizationInvitationsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListUserOrganizationInvitationsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListUserOrganizationInvitationsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListUserOrganizationInvitationsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListUserOrganizationInvitationsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatuses

`func (o *ListUserOrganizationInvitationsRequest) GetStatuses() []OrganizationInvitationStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ListUserOrganizationInvitationsRequest) GetStatusesOk() (*[]OrganizationInvitationStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ListUserOrganizationInvitationsRequest) SetStatuses(v []OrganizationInvitationStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ListUserOrganizationInvitationsRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


