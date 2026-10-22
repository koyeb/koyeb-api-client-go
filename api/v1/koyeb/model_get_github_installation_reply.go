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

// GetGithubInstallationReply struct for GetGithubInstallationReply
type GetGithubInstallationReply struct {
	InstallationId *string `json:"installation_id,omitempty"`
	InstallationUrl *string `json:"installation_url,omitempty"`
	Name *string `json:"name,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Status *KgitproxyGithubInstallationStatus `json:"status,omitempty"`
	InstalledAt *time.Time `json:"installed_at,omitempty"`
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	IndexingStatus *KgitproxyIndexingStatus `json:"indexing_status,omitempty"`
	IndexedRepositories *int64 `json:"indexed_repositories,omitempty"`
	TotalRepositories *int64 `json:"total_repositories,omitempty"`
}

// NewGetGithubInstallationReply instantiates a new GetGithubInstallationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGithubInstallationReply() *GetGithubInstallationReply {
	this := GetGithubInstallationReply{}
	var status KgitproxyGithubInstallationStatus = KGITPROXYGITHUBINSTALLATIONSTATUS_INVALID
	this.Status = &status
	var indexingStatus KgitproxyIndexingStatus = KGITPROXYINDEXINGSTATUS_INVALID_INDEXING_STATUS
	this.IndexingStatus = &indexingStatus
	return &this
}

// NewGetGithubInstallationReplyWithDefaults instantiates a new GetGithubInstallationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGithubInstallationReplyWithDefaults() *GetGithubInstallationReply {
	this := GetGithubInstallationReply{}
	var status KgitproxyGithubInstallationStatus = KGITPROXYGITHUBINSTALLATIONSTATUS_INVALID
	this.Status = &status
	var indexingStatus KgitproxyIndexingStatus = KGITPROXYINDEXINGSTATUS_INVALID_INDEXING_STATUS
	this.IndexingStatus = &indexingStatus
	return &this
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetInstallationId() string {
	if o == nil || isNil(o.InstallationId) {
		var ret string
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetInstallationIdOk() (*string, bool) {
	if o == nil || isNil(o.InstallationId) {
    return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasInstallationId() bool {
	if o != nil && !isNil(o.InstallationId) {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given string and assigns it to the InstallationId field.
func (o *GetGithubInstallationReply) SetInstallationId(v string) {
	o.InstallationId = &v
}

// GetInstallationUrl returns the InstallationUrl field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetInstallationUrl() string {
	if o == nil || isNil(o.InstallationUrl) {
		var ret string
		return ret
	}
	return *o.InstallationUrl
}

// GetInstallationUrlOk returns a tuple with the InstallationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetInstallationUrlOk() (*string, bool) {
	if o == nil || isNil(o.InstallationUrl) {
    return nil, false
	}
	return o.InstallationUrl, true
}

// HasInstallationUrl returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasInstallationUrl() bool {
	if o != nil && !isNil(o.InstallationUrl) {
		return true
	}

	return false
}

// SetInstallationUrl gets a reference to the given string and assigns it to the InstallationUrl field.
func (o *GetGithubInstallationReply) SetInstallationUrl(v string) {
	o.InstallationUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetGithubInstallationReply) SetName(v string) {
	o.Name = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetAvatarUrl() string {
	if o == nil || isNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetAvatarUrlOk() (*string, bool) {
	if o == nil || isNil(o.AvatarUrl) {
    return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasAvatarUrl() bool {
	if o != nil && !isNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *GetGithubInstallationReply) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetStatus() KgitproxyGithubInstallationStatus {
	if o == nil || isNil(o.Status) {
		var ret KgitproxyGithubInstallationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetStatusOk() (*KgitproxyGithubInstallationStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given KgitproxyGithubInstallationStatus and assigns it to the Status field.
func (o *GetGithubInstallationReply) SetStatus(v KgitproxyGithubInstallationStatus) {
	o.Status = &v
}

// GetInstalledAt returns the InstalledAt field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetInstalledAt() time.Time {
	if o == nil || isNil(o.InstalledAt) {
		var ret time.Time
		return ret
	}
	return *o.InstalledAt
}

// GetInstalledAtOk returns a tuple with the InstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetInstalledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.InstalledAt) {
    return nil, false
	}
	return o.InstalledAt, true
}

// HasInstalledAt returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasInstalledAt() bool {
	if o != nil && !isNil(o.InstalledAt) {
		return true
	}

	return false
}

// SetInstalledAt gets a reference to the given time.Time and assigns it to the InstalledAt field.
func (o *GetGithubInstallationReply) SetInstalledAt(v time.Time) {
	o.InstalledAt = &v
}

// GetSuspendedAt returns the SuspendedAt field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetSuspendedAt() time.Time {
	if o == nil || isNil(o.SuspendedAt) {
		var ret time.Time
		return ret
	}
	return *o.SuspendedAt
}

// GetSuspendedAtOk returns a tuple with the SuspendedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetSuspendedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SuspendedAt) {
    return nil, false
	}
	return o.SuspendedAt, true
}

// HasSuspendedAt returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasSuspendedAt() bool {
	if o != nil && !isNil(o.SuspendedAt) {
		return true
	}

	return false
}

// SetSuspendedAt gets a reference to the given time.Time and assigns it to the SuspendedAt field.
func (o *GetGithubInstallationReply) SetSuspendedAt(v time.Time) {
	o.SuspendedAt = &v
}

// GetIndexingStatus returns the IndexingStatus field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetIndexingStatus() KgitproxyIndexingStatus {
	if o == nil || isNil(o.IndexingStatus) {
		var ret KgitproxyIndexingStatus
		return ret
	}
	return *o.IndexingStatus
}

// GetIndexingStatusOk returns a tuple with the IndexingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetIndexingStatusOk() (*KgitproxyIndexingStatus, bool) {
	if o == nil || isNil(o.IndexingStatus) {
    return nil, false
	}
	return o.IndexingStatus, true
}

// HasIndexingStatus returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasIndexingStatus() bool {
	if o != nil && !isNil(o.IndexingStatus) {
		return true
	}

	return false
}

// SetIndexingStatus gets a reference to the given KgitproxyIndexingStatus and assigns it to the IndexingStatus field.
func (o *GetGithubInstallationReply) SetIndexingStatus(v KgitproxyIndexingStatus) {
	o.IndexingStatus = &v
}

// GetIndexedRepositories returns the IndexedRepositories field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetIndexedRepositories() int64 {
	if o == nil || isNil(o.IndexedRepositories) {
		var ret int64
		return ret
	}
	return *o.IndexedRepositories
}

// GetIndexedRepositoriesOk returns a tuple with the IndexedRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetIndexedRepositoriesOk() (*int64, bool) {
	if o == nil || isNil(o.IndexedRepositories) {
    return nil, false
	}
	return o.IndexedRepositories, true
}

// HasIndexedRepositories returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasIndexedRepositories() bool {
	if o != nil && !isNil(o.IndexedRepositories) {
		return true
	}

	return false
}

// SetIndexedRepositories gets a reference to the given int64 and assigns it to the IndexedRepositories field.
func (o *GetGithubInstallationReply) SetIndexedRepositories(v int64) {
	o.IndexedRepositories = &v
}

// GetTotalRepositories returns the TotalRepositories field value if set, zero value otherwise.
func (o *GetGithubInstallationReply) GetTotalRepositories() int64 {
	if o == nil || isNil(o.TotalRepositories) {
		var ret int64
		return ret
	}
	return *o.TotalRepositories
}

// GetTotalRepositoriesOk returns a tuple with the TotalRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGithubInstallationReply) GetTotalRepositoriesOk() (*int64, bool) {
	if o == nil || isNil(o.TotalRepositories) {
    return nil, false
	}
	return o.TotalRepositories, true
}

// HasTotalRepositories returns a boolean if a field has been set.
func (o *GetGithubInstallationReply) HasTotalRepositories() bool {
	if o != nil && !isNil(o.TotalRepositories) {
		return true
	}

	return false
}

// SetTotalRepositories gets a reference to the given int64 and assigns it to the TotalRepositories field.
func (o *GetGithubInstallationReply) SetTotalRepositories(v int64) {
	o.TotalRepositories = &v
}

func (o GetGithubInstallationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InstallationId) {
		toSerialize["installation_id"] = o.InstallationId
	}
	if !isNil(o.InstallationUrl) {
		toSerialize["installation_url"] = o.InstallationUrl
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.InstalledAt) {
		toSerialize["installed_at"] = o.InstalledAt
	}
	if !isNil(o.SuspendedAt) {
		toSerialize["suspended_at"] = o.SuspendedAt
	}
	if !isNil(o.IndexingStatus) {
		toSerialize["indexing_status"] = o.IndexingStatus
	}
	if !isNil(o.IndexedRepositories) {
		toSerialize["indexed_repositories"] = o.IndexedRepositories
	}
	if !isNil(o.TotalRepositories) {
		toSerialize["total_repositories"] = o.TotalRepositories
	}
	return json.Marshal(toSerialize)
}

type NullableGetGithubInstallationReply struct {
	value *GetGithubInstallationReply
	isSet bool
}

func (v NullableGetGithubInstallationReply) Get() *GetGithubInstallationReply {
	return v.value
}

func (v *NullableGetGithubInstallationReply) Set(val *GetGithubInstallationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGithubInstallationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGithubInstallationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGithubInstallationReply(val *GetGithubInstallationReply) *NullableGetGithubInstallationReply {
	return &NullableGetGithubInstallationReply{value: val, isSet: true}
}

func (v NullableGetGithubInstallationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGithubInstallationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


