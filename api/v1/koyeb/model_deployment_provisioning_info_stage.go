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

// DeploymentProvisioningInfoStage struct for DeploymentProvisioningInfoStage
type DeploymentProvisioningInfoStage struct {
	Name *string `json:"name,omitempty"`
	Status *DeploymentProvisioningInfoStageStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	BuildAttempts []DeploymentProvisioningInfoStageBuildAttempt `json:"build_attempts,omitempty"`
}

// NewDeploymentProvisioningInfoStage instantiates a new DeploymentProvisioningInfoStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfoStage() *DeploymentProvisioningInfoStage {
	this := DeploymentProvisioningInfoStage{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewDeploymentProvisioningInfoStageWithDefaults instantiates a new DeploymentProvisioningInfoStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoStageWithDefaults() *DeploymentProvisioningInfoStage {
	this := DeploymentProvisioningInfoStage{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentProvisioningInfoStage) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetStatus() DeploymentProvisioningInfoStageStatus {
	if o == nil || isNil(o.Status) {
		var ret DeploymentProvisioningInfoStageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentProvisioningInfoStageStatus and assigns it to the Status field.
func (o *DeploymentProvisioningInfoStage) SetStatus(v DeploymentProvisioningInfoStageStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *DeploymentProvisioningInfoStage) SetMessages(v []string) {
	o.Messages = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DeploymentProvisioningInfoStage) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
    return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DeploymentProvisioningInfoStage) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetBuildAttempts returns the BuildAttempts field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetBuildAttempts() []DeploymentProvisioningInfoStageBuildAttempt {
	if o == nil || isNil(o.BuildAttempts) {
		var ret []DeploymentProvisioningInfoStageBuildAttempt
		return ret
	}
	return o.BuildAttempts
}

// GetBuildAttemptsOk returns a tuple with the BuildAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetBuildAttemptsOk() ([]DeploymentProvisioningInfoStageBuildAttempt, bool) {
	if o == nil || isNil(o.BuildAttempts) {
    return nil, false
	}
	return o.BuildAttempts, true
}

// HasBuildAttempts returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasBuildAttempts() bool {
	if o != nil && !isNil(o.BuildAttempts) {
		return true
	}

	return false
}

// SetBuildAttempts gets a reference to the given []DeploymentProvisioningInfoStageBuildAttempt and assigns it to the BuildAttempts field.
func (o *DeploymentProvisioningInfoStage) SetBuildAttempts(v []DeploymentProvisioningInfoStageBuildAttempt) {
	o.BuildAttempts = v
}

func (o DeploymentProvisioningInfoStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !isNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !isNil(o.BuildAttempts) {
		toSerialize["build_attempts"] = o.BuildAttempts
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfoStage struct {
	value *DeploymentProvisioningInfoStage
	isSet bool
}

func (v NullableDeploymentProvisioningInfoStage) Get() *DeploymentProvisioningInfoStage {
	return v.value
}

func (v *NullableDeploymentProvisioningInfoStage) Set(val *DeploymentProvisioningInfoStage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfoStage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfoStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfoStage(val *DeploymentProvisioningInfoStage) *NullableDeploymentProvisioningInfoStage {
	return &NullableDeploymentProvisioningInfoStage{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfoStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfoStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


