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

// KgitproxyRepository struct for KgitproxyRepository
type KgitproxyRepository struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
	Provider *KgitproxyRepositoryProvider `json:"provider,omitempty"`
	LastPushDate *time.Time `json:"last_push_date,omitempty"`
	Github *KgitproxyGitHubRepository `json:"github,omitempty"`
}

// NewKgitproxyRepository instantiates a new KgitproxyRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKgitproxyRepository() *KgitproxyRepository {
	this := KgitproxyRepository{}
	var provider KgitproxyRepositoryProvider = KGITPROXYREPOSITORYPROVIDER_INVALID_PROVIDER
	this.Provider = &provider
	return &this
}

// NewKgitproxyRepositoryWithDefaults instantiates a new KgitproxyRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKgitproxyRepositoryWithDefaults() *KgitproxyRepository {
	this := KgitproxyRepository{}
	var provider KgitproxyRepositoryProvider = KGITPROXYREPOSITORYPROVIDER_INVALID_PROVIDER
	this.Provider = &provider
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KgitproxyRepository) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *KgitproxyRepository) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KgitproxyRepository) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *KgitproxyRepository) SetUrl(v string) {
	o.Url = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KgitproxyRepository) SetDescription(v string) {
	o.Description = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetIsPrivate() bool {
	if o == nil || isNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetIsPrivateOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrivate) {
    return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasIsPrivate() bool {
	if o != nil && !isNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *KgitproxyRepository) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetIsDisabled() bool {
	if o == nil || isNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetIsDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsDisabled) {
    return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasIsDisabled() bool {
	if o != nil && !isNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *KgitproxyRepository) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetDefaultBranch() string {
	if o == nil || isNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetDefaultBranchOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBranch) {
    return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasDefaultBranch() bool {
	if o != nil && !isNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *KgitproxyRepository) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetProvider() KgitproxyRepositoryProvider {
	if o == nil || isNil(o.Provider) {
		var ret KgitproxyRepositoryProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetProviderOk() (*KgitproxyRepositoryProvider, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given KgitproxyRepositoryProvider and assigns it to the Provider field.
func (o *KgitproxyRepository) SetProvider(v KgitproxyRepositoryProvider) {
	o.Provider = &v
}

// GetLastPushDate returns the LastPushDate field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetLastPushDate() time.Time {
	if o == nil || isNil(o.LastPushDate) {
		var ret time.Time
		return ret
	}
	return *o.LastPushDate
}

// GetLastPushDateOk returns a tuple with the LastPushDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetLastPushDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastPushDate) {
    return nil, false
	}
	return o.LastPushDate, true
}

// HasLastPushDate returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasLastPushDate() bool {
	if o != nil && !isNil(o.LastPushDate) {
		return true
	}

	return false
}

// SetLastPushDate gets a reference to the given time.Time and assigns it to the LastPushDate field.
func (o *KgitproxyRepository) SetLastPushDate(v time.Time) {
	o.LastPushDate = &v
}

// GetGithub returns the Github field value if set, zero value otherwise.
func (o *KgitproxyRepository) GetGithub() KgitproxyGitHubRepository {
	if o == nil || isNil(o.Github) {
		var ret KgitproxyGitHubRepository
		return ret
	}
	return *o.Github
}

// GetGithubOk returns a tuple with the Github field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KgitproxyRepository) GetGithubOk() (*KgitproxyGitHubRepository, bool) {
	if o == nil || isNil(o.Github) {
    return nil, false
	}
	return o.Github, true
}

// HasGithub returns a boolean if a field has been set.
func (o *KgitproxyRepository) HasGithub() bool {
	if o != nil && !isNil(o.Github) {
		return true
	}

	return false
}

// SetGithub gets a reference to the given KgitproxyGitHubRepository and assigns it to the Github field.
func (o *KgitproxyRepository) SetGithub(v KgitproxyGitHubRepository) {
	o.Github = &v
}

func (o KgitproxyRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}
	if !isNil(o.IsDisabled) {
		toSerialize["is_disabled"] = o.IsDisabled
	}
	if !isNil(o.DefaultBranch) {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.LastPushDate) {
		toSerialize["last_push_date"] = o.LastPushDate
	}
	if !isNil(o.Github) {
		toSerialize["github"] = o.Github
	}
	return json.Marshal(toSerialize)
}

type NullableKgitproxyRepository struct {
	value *KgitproxyRepository
	isSet bool
}

func (v NullableKgitproxyRepository) Get() *KgitproxyRepository {
	return v.value
}

func (v *NullableKgitproxyRepository) Set(val *KgitproxyRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableKgitproxyRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableKgitproxyRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKgitproxyRepository(val *KgitproxyRepository) *NullableKgitproxyRepository {
	return &NullableKgitproxyRepository{value: val, isSet: true}
}

func (v NullableKgitproxyRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKgitproxyRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


