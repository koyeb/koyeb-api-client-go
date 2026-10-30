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

// DatabaseUsageDetails struct for DatabaseUsageDetails
type DatabaseUsageDetails struct {
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	AppName *string `json:"app_name,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	ComputeTimeSeconds *int64 `json:"compute_time_seconds,omitempty"`
	DataStorageMegabytesHour *int64 `json:"data_storage_megabytes_hour,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

// NewDatabaseUsageDetails instantiates a new DatabaseUsageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUsageDetails() *DatabaseUsageDetails {
	this := DatabaseUsageDetails{}
	return &this
}

// NewDatabaseUsageDetailsWithDefaults instantiates a new DatabaseUsageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUsageDetailsWithDefaults() *DatabaseUsageDetails {
	this := DatabaseUsageDetails{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *DatabaseUsageDetails) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *DatabaseUsageDetails) SetAppId(v string) {
	o.AppId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetAppName() string {
	if o == nil || isNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetAppNameOk() (*string, bool) {
	if o == nil || isNil(o.AppName) {
    return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasAppName() bool {
	if o != nil && !isNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *DatabaseUsageDetails) SetAppName(v string) {
	o.AppName = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DatabaseUsageDetails) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetServiceName() string {
	if o == nil || isNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetServiceNameOk() (*string, bool) {
	if o == nil || isNil(o.ServiceName) {
    return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasServiceName() bool {
	if o != nil && !isNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *DatabaseUsageDetails) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetComputeTimeSeconds returns the ComputeTimeSeconds field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetComputeTimeSeconds() int64 {
	if o == nil || isNil(o.ComputeTimeSeconds) {
		var ret int64
		return ret
	}
	return *o.ComputeTimeSeconds
}

// GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetComputeTimeSecondsOk() (*int64, bool) {
	if o == nil || isNil(o.ComputeTimeSeconds) {
    return nil, false
	}
	return o.ComputeTimeSeconds, true
}

// HasComputeTimeSeconds returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasComputeTimeSeconds() bool {
	if o != nil && !isNil(o.ComputeTimeSeconds) {
		return true
	}

	return false
}

// SetComputeTimeSeconds gets a reference to the given int64 and assigns it to the ComputeTimeSeconds field.
func (o *DatabaseUsageDetails) SetComputeTimeSeconds(v int64) {
	o.ComputeTimeSeconds = &v
}

// GetDataStorageMegabytesHour returns the DataStorageMegabytesHour field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetDataStorageMegabytesHour() int64 {
	if o == nil || isNil(o.DataStorageMegabytesHour) {
		var ret int64
		return ret
	}
	return *o.DataStorageMegabytesHour
}

// GetDataStorageMegabytesHourOk returns a tuple with the DataStorageMegabytesHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetDataStorageMegabytesHourOk() (*int64, bool) {
	if o == nil || isNil(o.DataStorageMegabytesHour) {
    return nil, false
	}
	return o.DataStorageMegabytesHour, true
}

// HasDataStorageMegabytesHour returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasDataStorageMegabytesHour() bool {
	if o != nil && !isNil(o.DataStorageMegabytesHour) {
		return true
	}

	return false
}

// SetDataStorageMegabytesHour gets a reference to the given int64 and assigns it to the DataStorageMegabytesHour field.
func (o *DatabaseUsageDetails) SetDataStorageMegabytesHour(v int64) {
	o.DataStorageMegabytesHour = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DatabaseUsageDetails) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *DatabaseUsageDetails) GetTerminatedAt() time.Time {
	if o == nil || isNil(o.TerminatedAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUsageDetails) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.TerminatedAt) {
    return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *DatabaseUsageDetails) HasTerminatedAt() bool {
	if o != nil && !isNil(o.TerminatedAt) {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *DatabaseUsageDetails) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

func (o DatabaseUsageDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !isNil(o.ComputeTimeSeconds) {
		toSerialize["compute_time_seconds"] = o.ComputeTimeSeconds
	}
	if !isNil(o.DataStorageMegabytesHour) {
		toSerialize["data_storage_megabytes_hour"] = o.DataStorageMegabytesHour
	}
	if !isNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !isNil(o.TerminatedAt) {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseUsageDetails struct {
	value *DatabaseUsageDetails
	isSet bool
}

func (v NullableDatabaseUsageDetails) Get() *DatabaseUsageDetails {
	return v.value
}

func (v *NullableDatabaseUsageDetails) Set(val *DatabaseUsageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseUsageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseUsageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseUsageDetails(val *DatabaseUsageDetails) *NullableDatabaseUsageDetails {
	return &NullableDatabaseUsageDetails{value: val, isSet: true}
}

func (v NullableDatabaseUsageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseUsageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


