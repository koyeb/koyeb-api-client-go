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

// Archive struct for Archive
type Archive struct {
	// The archive id, that can be referenced when creating or updating a service.
	Id *string `json:"id,omitempty"`
	// Organization owning the archive.
	OrganizationId *string `json:"organization_id,omitempty"`
	// The URL where to upload the archive. This URL is signed and can only be used to upload the archive until `valid_until`.
	UploadUrl *string `json:"upload_url,omitempty"`
	// The provisioned space for the archive.
	Size *string `json:"size,omitempty"`
	// Date of creation of the archive.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// This field is automatically set by Koyeb when the archive is garbage collected.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A signed URL to download the archive. Getting an archive signs a new one when it has less than 5 minutes left, valid for 24 hours.
	DownloadUrl *string `json:"download_url,omitempty"`
	// When download_url expires.
	DownloadUrlDeadline *time.Time `json:"download_url_deadline,omitempty"`
	Status *ArchiveStatus `json:"status,omitempty"`
}

// NewArchive instantiates a new Archive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchive() *Archive {
	this := Archive{}
	var status ArchiveStatus = ARCHIVESTATUS_INVALID
	this.Status = &status
	return &this
}

// NewArchiveWithDefaults instantiates a new Archive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveWithDefaults() *Archive {
	this := Archive{}
	var status ArchiveStatus = ARCHIVESTATUS_INVALID
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Archive) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Archive) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Archive) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Archive) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Archive) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Archive) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *Archive) GetUploadUrl() string {
	if o == nil || isNil(o.UploadUrl) {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetUploadUrlOk() (*string, bool) {
	if o == nil || isNil(o.UploadUrl) {
    return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *Archive) HasUploadUrl() bool {
	if o != nil && !isNil(o.UploadUrl) {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *Archive) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Archive) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Archive) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Archive) SetSize(v string) {
	o.Size = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Archive) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Archive) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Archive) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Archive) GetDeletedAt() time.Time {
	if o == nil || isNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletedAt) {
    return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Archive) HasDeletedAt() bool {
	if o != nil && !isNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Archive) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *Archive) GetDownloadUrl() string {
	if o == nil || isNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetDownloadUrlOk() (*string, bool) {
	if o == nil || isNil(o.DownloadUrl) {
    return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *Archive) HasDownloadUrl() bool {
	if o != nil && !isNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *Archive) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetDownloadUrlDeadline returns the DownloadUrlDeadline field value if set, zero value otherwise.
func (o *Archive) GetDownloadUrlDeadline() time.Time {
	if o == nil || isNil(o.DownloadUrlDeadline) {
		var ret time.Time
		return ret
	}
	return *o.DownloadUrlDeadline
}

// GetDownloadUrlDeadlineOk returns a tuple with the DownloadUrlDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetDownloadUrlDeadlineOk() (*time.Time, bool) {
	if o == nil || isNil(o.DownloadUrlDeadline) {
    return nil, false
	}
	return o.DownloadUrlDeadline, true
}

// HasDownloadUrlDeadline returns a boolean if a field has been set.
func (o *Archive) HasDownloadUrlDeadline() bool {
	if o != nil && !isNil(o.DownloadUrlDeadline) {
		return true
	}

	return false
}

// SetDownloadUrlDeadline gets a reference to the given time.Time and assigns it to the DownloadUrlDeadline field.
func (o *Archive) SetDownloadUrlDeadline(v time.Time) {
	o.DownloadUrlDeadline = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Archive) GetStatus() ArchiveStatus {
	if o == nil || isNil(o.Status) {
		var ret ArchiveStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Archive) GetStatusOk() (*ArchiveStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Archive) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ArchiveStatus and assigns it to the Status field.
func (o *Archive) SetStatus(v ArchiveStatus) {
	o.Status = &v
}

func (o Archive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.UploadUrl) {
		toSerialize["upload_url"] = o.UploadUrl
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !isNil(o.DownloadUrl) {
		toSerialize["download_url"] = o.DownloadUrl
	}
	if !isNil(o.DownloadUrlDeadline) {
		toSerialize["download_url_deadline"] = o.DownloadUrlDeadline
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableArchive struct {
	value *Archive
	isSet bool
}

func (v NullableArchive) Get() *Archive {
	return v.value
}

func (v *NullableArchive) Set(val *Archive) {
	v.value = val
	v.isSet = true
}

func (v NullableArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchive(val *Archive) *NullableArchive {
	return &NullableArchive{value: val, isSet: true}
}

func (v NullableArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


