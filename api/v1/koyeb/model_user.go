/*
Koyeb Rest API

Welcome to the Koyeb API documentation.  The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. The API has predictable resource-oriented URLs, allowing you to request, retrieve and execute actions on resources efficiently.  The API returns JSON-encoded responses and uses standard HTTP response codes, and verbs.  All of the functionality available in the Koyeb Control Panel is also available via the API, allowing you to automate repetitive and complex actions as required.  All production API requests are made to:  - https://app.koyeb.com/<version>  All data is sent and received as JSON.  ## Constructing Requests  Requests are made of three components:  - Base URL: `https://app.koyeb.com` - API Version: `v1` - Resource path: `apps`  To construct a proper request, you need to format the URL as follows:  `https://app.koyeb.com/{version}/{resource}`  Example: `https://app.koyeb.com/v1/apps`  The following code is an example request to retrieve the list of Apps.  ``` $ curl -H 'Authorization: Bearer 1111-1111-1111-1111-1111' https://app.koyeb.com/v1/apps {   \"apps\": [     {       \"id\": \"e267e09d-a2ef-4672-95c7-bf88d787a274\",       \"name\": \"demo\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-04-26T08:28:59.221274Z\",       \"created_at\": \"2021-04-26T08:28:59.221274Z\",       \"domains\": [         {           \"name\": \"demo-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     },     {       \"id\": \"e38caea5-5acb-4fec-aa65-81749546b470\",       \"name\": \"www\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-05-04T08:56:40.557792Z\",       \"created_at\": \"2021-05-04T08:56:40.557792Z\",       \"domains\": [         {           \"name\": \"www-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     }   ],   \"limit\": 20,   \"offset\": 0,   \"count\": 2 } ```  # Authentication  The Koyeb API uses API keys to authenticate requests. You can view and manage your API keys in the Koyeb Control Panel.  Your API keys carry many privileges. Take care to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, Slack, client-side code, and so forth.  To make an authenticated request, include a bearer-type Authorization header containing your API key, prefixed with `Bearer `:  ``` Authorization: Bearer <your-api-key> ```  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  # Errors  We use conventional HTTP response codes to indicate the success or failure of an API request.  In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g., a required parameter is missing, resource not found), and codes in the 5xx range indicate an error with our servers.  ## HTTP Status Code Summary  - `200 OK` - Everything worked as expected. - `400 Bad Request` - Often missing a required parameter. - `401 Unauthorized` - No valid API key provided. - `402 Request Failed` - Parameters were valid, but the request failed. - `403 Forbidden` - Insufficient privileges to access the requested resource. - `404 Not Found` - The requested item doesn't exist. - `50x` Server errors - Something went wrong on the API domain's end.  Input and output data must be valid JSON with the appropriate Content-Type header set.  ## Attributes  - code:    - `invalid_argument`: Occur when your request has invalid parameters.   - `unknown_resource`: Occur when the resource doesn't exist.   - `authentication_error`: Arise when there is a problem with authentication.   - `rate_limit_error`: Too many requests hit the API too quickly.   - `authorization_required`: You don't have sufficient right to access the resource.   - `api_error`: API errors use in case of a problem with api domain's servers (extremely uncommon).  - message:    - A human-readable error giving more details about the error  - status:    - The HTTP code returned by the server.  - fields (Optional):    - An array of parameters with a human-readable message giving more details about the error.  ## Errors responses example  Response 400 (application/json)  ``` {   \"status\": 400,   \"code\": \"invalid_argument\",   \"message\": \"Invalid argument\",   \"fields\": [     {       \"field\": \"definition.docker.image\",       \"description\": \"cannot be blank\"     }   ] } ```  Response 401 (application/json)  ``` {   \"status\": 401,   \"code\": \"authentication_error\",   \"message\": \"Request unauthenticated with bearer\" } ```  Response 403 (application/json)  ``` {   \"status\": 403,   \"code\": \"authentication_required\",   \"message\": \"You don't have the requested permission to access to resource.\" } ```   Response 404 (application/json)  ``` {   \"status\": 404,   \"code\": \"unknown_resource\",   \"message\": \"App not found\" } ```  # Pagination  All top-level API resources receive a paginated response. For instance, list Apps, list Services, and others share a standard structure, taking two parameters: limit and offset.  **Paginated request**  Requests against paginated endpoints accept two query string arguments:  * *limit*, a positive integer lower or equal to 100 to select the number of items to display. * *offset*, an integer to indicate the position to start at.  Paginated endpoints usually also accept filters to search and sort results. These filters are documented along with each endpoint documentation. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// User struct for User
type User struct {
	Id *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	TwoFactorAuthentication *bool `json:"two_factor_authentication,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	LastLoginIp *string `json:"last_login_ip,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	NewsletterSubscribed *bool `json:"newsletter_subscribed,omitempty"`
	GithubId *string `json:"github_id,omitempty"`
	GithubUser *string `json:"github_user,omitempty"`
	Flags []UserFlags `json:"flags,omitempty"`
	Name *string `json:"name,omitempty"`
	EmailValidated *bool `json:"email_validated,omitempty"`
	Trialed *bool `json:"trialed,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *User) GetAvatarUrl() string {
	if o == nil || isNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil || isNil(o.AvatarUrl) {
    return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && !isNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetTwoFactorAuthentication returns the TwoFactorAuthentication field value if set, zero value otherwise.
func (o *User) GetTwoFactorAuthentication() bool {
	if o == nil || isNil(o.TwoFactorAuthentication) {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthentication
}

// GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTwoFactorAuthenticationOk() (*bool, bool) {
	if o == nil || isNil(o.TwoFactorAuthentication) {
    return nil, false
	}
	return o.TwoFactorAuthentication, true
}

// HasTwoFactorAuthentication returns a boolean if a field has been set.
func (o *User) HasTwoFactorAuthentication() bool {
	if o != nil && !isNil(o.TwoFactorAuthentication) {
		return true
	}

	return false
}

// SetTwoFactorAuthentication gets a reference to the given bool and assigns it to the TwoFactorAuthentication field.
func (o *User) SetTwoFactorAuthentication(v bool) {
	o.TwoFactorAuthentication = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() time.Time {
	if o == nil || isNil(o.LastLogin) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastLogin) {
    return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && !isNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetLastLoginIp returns the LastLoginIp field value if set, zero value otherwise.
func (o *User) GetLastLoginIp() string {
	if o == nil || isNil(o.LastLoginIp) {
		var ret string
		return ret
	}
	return *o.LastLoginIp
}

// GetLastLoginIpOk returns a tuple with the LastLoginIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginIpOk() (*string, bool) {
	if o == nil || isNil(o.LastLoginIp) {
    return nil, false
	}
	return o.LastLoginIp, true
}

// HasLastLoginIp returns a boolean if a field has been set.
func (o *User) HasLastLoginIp() bool {
	if o != nil && !isNil(o.LastLoginIp) {
		return true
	}

	return false
}

// SetLastLoginIp gets a reference to the given string and assigns it to the LastLoginIp field.
func (o *User) SetLastLoginIp(v string) {
	o.LastLoginIp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetNewsletterSubscribed returns the NewsletterSubscribed field value if set, zero value otherwise.
func (o *User) GetNewsletterSubscribed() bool {
	if o == nil || isNil(o.NewsletterSubscribed) {
		var ret bool
		return ret
	}
	return *o.NewsletterSubscribed
}

// GetNewsletterSubscribedOk returns a tuple with the NewsletterSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNewsletterSubscribedOk() (*bool, bool) {
	if o == nil || isNil(o.NewsletterSubscribed) {
    return nil, false
	}
	return o.NewsletterSubscribed, true
}

// HasNewsletterSubscribed returns a boolean if a field has been set.
func (o *User) HasNewsletterSubscribed() bool {
	if o != nil && !isNil(o.NewsletterSubscribed) {
		return true
	}

	return false
}

// SetNewsletterSubscribed gets a reference to the given bool and assigns it to the NewsletterSubscribed field.
func (o *User) SetNewsletterSubscribed(v bool) {
	o.NewsletterSubscribed = &v
}

// GetGithubId returns the GithubId field value if set, zero value otherwise.
func (o *User) GetGithubId() string {
	if o == nil || isNil(o.GithubId) {
		var ret string
		return ret
	}
	return *o.GithubId
}

// GetGithubIdOk returns a tuple with the GithubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGithubIdOk() (*string, bool) {
	if o == nil || isNil(o.GithubId) {
    return nil, false
	}
	return o.GithubId, true
}

// HasGithubId returns a boolean if a field has been set.
func (o *User) HasGithubId() bool {
	if o != nil && !isNil(o.GithubId) {
		return true
	}

	return false
}

// SetGithubId gets a reference to the given string and assigns it to the GithubId field.
func (o *User) SetGithubId(v string) {
	o.GithubId = &v
}

// GetGithubUser returns the GithubUser field value if set, zero value otherwise.
func (o *User) GetGithubUser() string {
	if o == nil || isNil(o.GithubUser) {
		var ret string
		return ret
	}
	return *o.GithubUser
}

// GetGithubUserOk returns a tuple with the GithubUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGithubUserOk() (*string, bool) {
	if o == nil || isNil(o.GithubUser) {
    return nil, false
	}
	return o.GithubUser, true
}

// HasGithubUser returns a boolean if a field has been set.
func (o *User) HasGithubUser() bool {
	if o != nil && !isNil(o.GithubUser) {
		return true
	}

	return false
}

// SetGithubUser gets a reference to the given string and assigns it to the GithubUser field.
func (o *User) SetGithubUser(v string) {
	o.GithubUser = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *User) GetFlags() []UserFlags {
	if o == nil || isNil(o.Flags) {
		var ret []UserFlags
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFlagsOk() ([]UserFlags, bool) {
	if o == nil || isNil(o.Flags) {
    return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *User) HasFlags() bool {
	if o != nil && !isNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []UserFlags and assigns it to the Flags field.
func (o *User) SetFlags(v []UserFlags) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetEmailValidated returns the EmailValidated field value if set, zero value otherwise.
func (o *User) GetEmailValidated() bool {
	if o == nil || isNil(o.EmailValidated) {
		var ret bool
		return ret
	}
	return *o.EmailValidated
}

// GetEmailValidatedOk returns a tuple with the EmailValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailValidatedOk() (*bool, bool) {
	if o == nil || isNil(o.EmailValidated) {
    return nil, false
	}
	return o.EmailValidated, true
}

// HasEmailValidated returns a boolean if a field has been set.
func (o *User) HasEmailValidated() bool {
	if o != nil && !isNil(o.EmailValidated) {
		return true
	}

	return false
}

// SetEmailValidated gets a reference to the given bool and assigns it to the EmailValidated field.
func (o *User) SetEmailValidated(v bool) {
	o.EmailValidated = &v
}

// GetTrialed returns the Trialed field value if set, zero value otherwise.
func (o *User) GetTrialed() bool {
	if o == nil || isNil(o.Trialed) {
		var ret bool
		return ret
	}
	return *o.Trialed
}

// GetTrialedOk returns a tuple with the Trialed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTrialedOk() (*bool, bool) {
	if o == nil || isNil(o.Trialed) {
    return nil, false
	}
	return o.Trialed, true
}

// HasTrialed returns a boolean if a field has been set.
func (o *User) HasTrialed() bool {
	if o != nil && !isNil(o.Trialed) {
		return true
	}

	return false
}

// SetTrialed gets a reference to the given bool and assigns it to the Trialed field.
func (o *User) SetTrialed(v bool) {
	o.Trialed = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !isNil(o.TwoFactorAuthentication) {
		toSerialize["two_factor_authentication"] = o.TwoFactorAuthentication
	}
	if !isNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !isNil(o.LastLoginIp) {
		toSerialize["last_login_ip"] = o.LastLoginIp
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.NewsletterSubscribed) {
		toSerialize["newsletter_subscribed"] = o.NewsletterSubscribed
	}
	if !isNil(o.GithubId) {
		toSerialize["github_id"] = o.GithubId
	}
	if !isNil(o.GithubUser) {
		toSerialize["github_user"] = o.GithubUser
	}
	if !isNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.EmailValidated) {
		toSerialize["email_validated"] = o.EmailValidated
	}
	if !isNil(o.Trialed) {
		toSerialize["trialed"] = o.Trialed
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


