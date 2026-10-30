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

// PoolClaim struct for PoolClaim
type PoolClaim struct {
	Id *string `json:"id,omitempty"`
	PoolId *string `json:"pool_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	Status *PoolClaimStatus `json:"status,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	WorkspaceId *string `json:"workspace_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	FulfilledAt *time.Time `json:"fulfilled_at,omitempty"`
	PoolGeneration *string `json:"pool_generation,omitempty"`
}

// NewPoolClaim instantiates a new PoolClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolClaim() *PoolClaim {
	this := PoolClaim{}
	var status PoolClaimStatus = POOLCLAIMSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// NewPoolClaimWithDefaults instantiates a new PoolClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolClaimWithDefaults() *PoolClaim {
	this := PoolClaim{}
	var status PoolClaimStatus = POOLCLAIMSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolClaim) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolClaim) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolClaim) SetId(v string) {
	o.Id = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *PoolClaim) GetPoolId() string {
	if o == nil || isNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetPoolIdOk() (*string, bool) {
	if o == nil || isNil(o.PoolId) {
    return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *PoolClaim) HasPoolId() bool {
	if o != nil && !isNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *PoolClaim) SetPoolId(v string) {
	o.PoolId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PoolClaim) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PoolClaim) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *PoolClaim) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PoolClaim) GetRequestId() string {
	if o == nil || isNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetRequestIdOk() (*string, bool) {
	if o == nil || isNil(o.RequestId) {
    return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PoolClaim) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PoolClaim) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PoolClaim) GetStatus() PoolClaimStatus {
	if o == nil || isNil(o.Status) {
		var ret PoolClaimStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetStatusOk() (*PoolClaimStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PoolClaim) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PoolClaimStatus and assigns it to the Status field.
func (o *PoolClaim) SetStatus(v PoolClaimStatus) {
	o.Status = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PoolClaim) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PoolClaim) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *PoolClaim) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *PoolClaim) GetWorkspaceId() string {
	if o == nil || isNil(o.WorkspaceId) {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkspaceId) {
    return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *PoolClaim) HasWorkspaceId() bool {
	if o != nil && !isNil(o.WorkspaceId) {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *PoolClaim) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PoolClaim) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PoolClaim) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PoolClaim) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFulfilledAt returns the FulfilledAt field value if set, zero value otherwise.
func (o *PoolClaim) GetFulfilledAt() time.Time {
	if o == nil || isNil(o.FulfilledAt) {
		var ret time.Time
		return ret
	}
	return *o.FulfilledAt
}

// GetFulfilledAtOk returns a tuple with the FulfilledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetFulfilledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FulfilledAt) {
    return nil, false
	}
	return o.FulfilledAt, true
}

// HasFulfilledAt returns a boolean if a field has been set.
func (o *PoolClaim) HasFulfilledAt() bool {
	if o != nil && !isNil(o.FulfilledAt) {
		return true
	}

	return false
}

// SetFulfilledAt gets a reference to the given time.Time and assigns it to the FulfilledAt field.
func (o *PoolClaim) SetFulfilledAt(v time.Time) {
	o.FulfilledAt = &v
}

// GetPoolGeneration returns the PoolGeneration field value if set, zero value otherwise.
func (o *PoolClaim) GetPoolGeneration() string {
	if o == nil || isNil(o.PoolGeneration) {
		var ret string
		return ret
	}
	return *o.PoolGeneration
}

// GetPoolGenerationOk returns a tuple with the PoolGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolClaim) GetPoolGenerationOk() (*string, bool) {
	if o == nil || isNil(o.PoolGeneration) {
    return nil, false
	}
	return o.PoolGeneration, true
}

// HasPoolGeneration returns a boolean if a field has been set.
func (o *PoolClaim) HasPoolGeneration() bool {
	if o != nil && !isNil(o.PoolGeneration) {
		return true
	}

	return false
}

// SetPoolGeneration gets a reference to the given string and assigns it to the PoolGeneration field.
func (o *PoolClaim) SetPoolGeneration(v string) {
	o.PoolGeneration = &v
}

func (o PoolClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.PoolId) {
		toSerialize["pool_id"] = o.PoolId
	}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.WorkspaceId) {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.FulfilledAt) {
		toSerialize["fulfilled_at"] = o.FulfilledAt
	}
	if !isNil(o.PoolGeneration) {
		toSerialize["pool_generation"] = o.PoolGeneration
	}
	return json.Marshal(toSerialize)
}

type NullablePoolClaim struct {
	value *PoolClaim
	isSet bool
}

func (v NullablePoolClaim) Get() *PoolClaim {
	return v.value
}

func (v *NullablePoolClaim) Set(val *PoolClaim) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolClaim) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolClaim(val *PoolClaim) *NullablePoolClaim {
	return &NullablePoolClaim{value: val, isSet: true}
}

func (v NullablePoolClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


