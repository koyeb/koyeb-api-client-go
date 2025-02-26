# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**TwoFactorAuthentication** | Pointer to **bool** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 
**LastLoginIp** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**NewsletterSubscribed** | Pointer to **bool** |  | [optional] 
**GithubId** | Pointer to **string** |  | [optional] 
**GithubUser** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to [**[]UserFlags**](UserFlags.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EmailValidated** | Pointer to **bool** |  | [optional] 
**Trialed** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *User) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *User) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *User) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *User) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetTwoFactorAuthentication

`func (o *User) GetTwoFactorAuthentication() bool`

GetTwoFactorAuthentication returns the TwoFactorAuthentication field if non-nil, zero value otherwise.

### GetTwoFactorAuthenticationOk

`func (o *User) GetTwoFactorAuthenticationOk() (*bool, bool)`

GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthentication

`func (o *User) SetTwoFactorAuthentication(v bool)`

SetTwoFactorAuthentication sets TwoFactorAuthentication field to given value.

### HasTwoFactorAuthentication

`func (o *User) HasTwoFactorAuthentication() bool`

HasTwoFactorAuthentication returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastLoginIp

`func (o *User) GetLastLoginIp() string`

GetLastLoginIp returns the LastLoginIp field if non-nil, zero value otherwise.

### GetLastLoginIpOk

`func (o *User) GetLastLoginIpOk() (*string, bool)`

GetLastLoginIpOk returns a tuple with the LastLoginIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginIp

`func (o *User) SetLastLoginIp(v string)`

SetLastLoginIp sets LastLoginIp field to given value.

### HasLastLoginIp

`func (o *User) HasLastLoginIp() bool`

HasLastLoginIp returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNewsletterSubscribed

`func (o *User) GetNewsletterSubscribed() bool`

GetNewsletterSubscribed returns the NewsletterSubscribed field if non-nil, zero value otherwise.

### GetNewsletterSubscribedOk

`func (o *User) GetNewsletterSubscribedOk() (*bool, bool)`

GetNewsletterSubscribedOk returns a tuple with the NewsletterSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletterSubscribed

`func (o *User) SetNewsletterSubscribed(v bool)`

SetNewsletterSubscribed sets NewsletterSubscribed field to given value.

### HasNewsletterSubscribed

`func (o *User) HasNewsletterSubscribed() bool`

HasNewsletterSubscribed returns a boolean if a field has been set.

### GetGithubId

`func (o *User) GetGithubId() string`

GetGithubId returns the GithubId field if non-nil, zero value otherwise.

### GetGithubIdOk

`func (o *User) GetGithubIdOk() (*string, bool)`

GetGithubIdOk returns a tuple with the GithubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubId

`func (o *User) SetGithubId(v string)`

SetGithubId sets GithubId field to given value.

### HasGithubId

`func (o *User) HasGithubId() bool`

HasGithubId returns a boolean if a field has been set.

### GetGithubUser

`func (o *User) GetGithubUser() string`

GetGithubUser returns the GithubUser field if non-nil, zero value otherwise.

### GetGithubUserOk

`func (o *User) GetGithubUserOk() (*string, bool)`

GetGithubUserOk returns a tuple with the GithubUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUser

`func (o *User) SetGithubUser(v string)`

SetGithubUser sets GithubUser field to given value.

### HasGithubUser

`func (o *User) HasGithubUser() bool`

HasGithubUser returns a boolean if a field has been set.

### GetFlags

`func (o *User) GetFlags() []UserFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *User) GetFlagsOk() (*[]UserFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *User) SetFlags(v []UserFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *User) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailValidated

`func (o *User) GetEmailValidated() bool`

GetEmailValidated returns the EmailValidated field if non-nil, zero value otherwise.

### GetEmailValidatedOk

`func (o *User) GetEmailValidatedOk() (*bool, bool)`

GetEmailValidatedOk returns a tuple with the EmailValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidated

`func (o *User) SetEmailValidated(v bool)`

SetEmailValidated sets EmailValidated field to given value.

### HasEmailValidated

`func (o *User) HasEmailValidated() bool`

HasEmailValidated returns a boolean if a field has been set.

### GetTrialed

`func (o *User) GetTrialed() bool`

GetTrialed returns the Trialed field if non-nil, zero value otherwise.

### GetTrialedOk

`func (o *User) GetTrialedOk() (*bool, bool)`

GetTrialedOk returns a tuple with the Trialed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialed

`func (o *User) SetTrialed(v bool)`

SetTrialed sets Trialed field to given value.

### HasTrialed

`func (o *User) HasTrialed() bool`

HasTrialed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


