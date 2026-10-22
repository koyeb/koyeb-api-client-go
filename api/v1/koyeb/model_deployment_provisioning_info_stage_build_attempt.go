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

// DeploymentProvisioningInfoStageBuildAttempt struct for DeploymentProvisioningInfoStageBuildAttempt
type DeploymentProvisioningInfoStageBuildAttempt struct {
	Id *int64 `json:"id,omitempty"`
	Status *DeploymentProvisioningInfoStageStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Steps []DeploymentProvisioningInfoStageBuildAttemptBuildStep `json:"steps,omitempty"`
	ImagePushed *bool `json:"image_pushed,omitempty"`
	InternalFailure *bool `json:"internal_failure,omitempty"`
	RetryableFailure *bool `json:"retryable_failure,omitempty"`
	// This flag is used to finalize the build, and continue the deployment in case of success, or cancel and potentially retry the build in case of failure.
	WaitCompletion *bool `json:"wait_completion,omitempty"`
}

// NewDeploymentProvisioningInfoStageBuildAttempt instantiates a new DeploymentProvisioningInfoStageBuildAttempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfoStageBuildAttempt() *DeploymentProvisioningInfoStageBuildAttempt {
	this := DeploymentProvisioningInfoStageBuildAttempt{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewDeploymentProvisioningInfoStageBuildAttemptWithDefaults instantiates a new DeploymentProvisioningInfoStageBuildAttempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoStageBuildAttemptWithDefaults() *DeploymentProvisioningInfoStageBuildAttempt {
	this := DeploymentProvisioningInfoStageBuildAttempt{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStatus() DeploymentProvisioningInfoStageStatus {
	if o == nil || isNil(o.Status) {
		var ret DeploymentProvisioningInfoStageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentProvisioningInfoStageStatus and assigns it to the Status field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetStatus(v DeploymentProvisioningInfoStageStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetMessages(v []string) {
	o.Messages = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
    return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetSteps() []DeploymentProvisioningInfoStageBuildAttemptBuildStep {
	if o == nil || isNil(o.Steps) {
		var ret []DeploymentProvisioningInfoStageBuildAttemptBuildStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetStepsOk() ([]DeploymentProvisioningInfoStageBuildAttemptBuildStep, bool) {
	if o == nil || isNil(o.Steps) {
    return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []DeploymentProvisioningInfoStageBuildAttemptBuildStep and assigns it to the Steps field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetSteps(v []DeploymentProvisioningInfoStageBuildAttemptBuildStep) {
	o.Steps = v
}

// GetImagePushed returns the ImagePushed field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetImagePushed() bool {
	if o == nil || isNil(o.ImagePushed) {
		var ret bool
		return ret
	}
	return *o.ImagePushed
}

// GetImagePushedOk returns a tuple with the ImagePushed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetImagePushedOk() (*bool, bool) {
	if o == nil || isNil(o.ImagePushed) {
    return nil, false
	}
	return o.ImagePushed, true
}

// HasImagePushed returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasImagePushed() bool {
	if o != nil && !isNil(o.ImagePushed) {
		return true
	}

	return false
}

// SetImagePushed gets a reference to the given bool and assigns it to the ImagePushed field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetImagePushed(v bool) {
	o.ImagePushed = &v
}

// GetInternalFailure returns the InternalFailure field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetInternalFailure() bool {
	if o == nil || isNil(o.InternalFailure) {
		var ret bool
		return ret
	}
	return *o.InternalFailure
}

// GetInternalFailureOk returns a tuple with the InternalFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetInternalFailureOk() (*bool, bool) {
	if o == nil || isNil(o.InternalFailure) {
    return nil, false
	}
	return o.InternalFailure, true
}

// HasInternalFailure returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasInternalFailure() bool {
	if o != nil && !isNil(o.InternalFailure) {
		return true
	}

	return false
}

// SetInternalFailure gets a reference to the given bool and assigns it to the InternalFailure field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetInternalFailure(v bool) {
	o.InternalFailure = &v
}

// GetRetryableFailure returns the RetryableFailure field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetRetryableFailure() bool {
	if o == nil || isNil(o.RetryableFailure) {
		var ret bool
		return ret
	}
	return *o.RetryableFailure
}

// GetRetryableFailureOk returns a tuple with the RetryableFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetRetryableFailureOk() (*bool, bool) {
	if o == nil || isNil(o.RetryableFailure) {
    return nil, false
	}
	return o.RetryableFailure, true
}

// HasRetryableFailure returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasRetryableFailure() bool {
	if o != nil && !isNil(o.RetryableFailure) {
		return true
	}

	return false
}

// SetRetryableFailure gets a reference to the given bool and assigns it to the RetryableFailure field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetRetryableFailure(v bool) {
	o.RetryableFailure = &v
}

// GetWaitCompletion returns the WaitCompletion field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetWaitCompletion() bool {
	if o == nil || isNil(o.WaitCompletion) {
		var ret bool
		return ret
	}
	return *o.WaitCompletion
}

// GetWaitCompletionOk returns a tuple with the WaitCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) GetWaitCompletionOk() (*bool, bool) {
	if o == nil || isNil(o.WaitCompletion) {
    return nil, false
	}
	return o.WaitCompletion, true
}

// HasWaitCompletion returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStageBuildAttempt) HasWaitCompletion() bool {
	if o != nil && !isNil(o.WaitCompletion) {
		return true
	}

	return false
}

// SetWaitCompletion gets a reference to the given bool and assigns it to the WaitCompletion field.
func (o *DeploymentProvisioningInfoStageBuildAttempt) SetWaitCompletion(v bool) {
	o.WaitCompletion = &v
}

func (o DeploymentProvisioningInfoStageBuildAttempt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !isNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !isNil(o.ImagePushed) {
		toSerialize["image_pushed"] = o.ImagePushed
	}
	if !isNil(o.InternalFailure) {
		toSerialize["internal_failure"] = o.InternalFailure
	}
	if !isNil(o.RetryableFailure) {
		toSerialize["retryable_failure"] = o.RetryableFailure
	}
	if !isNil(o.WaitCompletion) {
		toSerialize["wait_completion"] = o.WaitCompletion
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfoStageBuildAttempt struct {
	value *DeploymentProvisioningInfoStageBuildAttempt
	isSet bool
}

func (v NullableDeploymentProvisioningInfoStageBuildAttempt) Get() *DeploymentProvisioningInfoStageBuildAttempt {
	return v.value
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttempt) Set(val *DeploymentProvisioningInfoStageBuildAttempt) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfoStageBuildAttempt) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfoStageBuildAttempt(val *DeploymentProvisioningInfoStageBuildAttempt) *NullableDeploymentProvisioningInfoStageBuildAttempt {
	return &NullableDeploymentProvisioningInfoStageBuildAttempt{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfoStageBuildAttempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfoStageBuildAttempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


