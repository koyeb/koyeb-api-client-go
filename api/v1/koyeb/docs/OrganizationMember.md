# OrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**Role** | Pointer to [**UserRoleRole**](UserRoleRole.md) |  | [optional] [default to USERROLEROLE_INVALID]
**Status** | Pointer to [**OrganizationMemberStatus**](OrganizationMemberStatus.md) |  | [optional] [default to ORGANIZATIONMEMBERSTATUS_INVALID]
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewOrganizationMember

`func NewOrganizationMember() *OrganizationMember`

NewOrganizationMember instantiates a new OrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberWithDefaults

`func NewOrganizationMemberWithDefaults() *OrganizationMember`

NewOrganizationMemberWithDefaults instantiates a new OrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationMember) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationMember) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationMember) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationMember) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationMember) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMember) GetRole() UserRoleRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMember) GetRoleOk() (*UserRoleRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMember) SetRole(v UserRoleRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMember) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationMember) GetStatus() OrganizationMemberStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationMember) GetStatusOk() (*OrganizationMemberStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationMember) SetStatus(v OrganizationMemberStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationMember) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationMember) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMember) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMember) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationMember) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationMember) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationMember) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationMember) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationMember) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


