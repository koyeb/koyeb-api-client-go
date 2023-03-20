# OrganizationInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**UserRoleRole**](UserRoleRole.md) |  | [optional] [default to USERROLEROLE_INVALID]
**Status** | Pointer to [**OrganizationInvitationStatus**](OrganizationInvitationStatus.md) |  | [optional] [default to ORGANIZATIONINVITATIONSTATUS_INVALID]
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**PublicOrganization**](PublicOrganization.md) |  | [optional] 
**User** | Pointer to [**PublicUser**](PublicUser.md) |  | [optional] 

## Methods

### NewOrganizationInvitation

`func NewOrganizationInvitation() *OrganizationInvitation`

NewOrganizationInvitation instantiates a new OrganizationInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationWithDefaults

`func NewOrganizationInvitationWithDefaults() *OrganizationInvitation`

NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationInvitation) GetRole() UserRoleRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInvitation) GetRoleOk() (*UserRoleRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInvitation) SetRole(v UserRoleRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInvitation) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationInvitation) GetStatus() OrganizationInvitationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationInvitation) GetStatusOk() (*OrganizationInvitationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationInvitation) SetStatus(v OrganizationInvitationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationInvitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationInvitation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationInvitation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationInvitation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationInvitation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationInvitation) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationInvitation) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationInvitation) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationInvitation) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationInvitation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationInvitation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationInvitation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationInvitation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationInvitation) GetOrganization() PublicOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationInvitation) GetOrganizationOk() (*PublicOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationInvitation) SetOrganization(v PublicOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationInvitation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationInvitation) GetUser() PublicUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationInvitation) GetUserOk() (*PublicUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationInvitation) SetUser(v PublicUser)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationInvitation) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


