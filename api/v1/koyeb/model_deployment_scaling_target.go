/*
Koyeb Rest API

Welcome to the Koyeb API documentation.  The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. The API has predictable resource-oriented URLs, allowing you to request, retrieve and execute actions on resources efficiently.  The API returns JSON-encoded responses and uses standard HTTP response codes, and verbs.  All of the functionality available in the Koyeb Control Panel is also available via the API, allowing you to automate repetitive and complex actions as required.  All production API requests are made to:  - https://app.koyeb.com/<version>  All data is sent and received as JSON.  ## Constructing Requests  Requests are made of three components:  - Base URL: `https://app.koyeb.com` - API Version: `v1` - Resource path: `apps`  To construct a proper request, you need to format the URL as follows:  `https://app.koyeb.com/{version}/{resource}`  Example: `https://app.koyeb.com/v1/apps`  The following code is an example request to retrieve the list of Apps.  ``` $ curl -H 'Authorization: Bearer 1111-1111-1111-1111-1111' https://app.koyeb.com/v1/apps {   \"apps\": [     {       \"id\": \"e267e09d-a2ef-4672-95c7-bf88d787a274\",       \"name\": \"demo\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-04-26T08:28:59.221274Z\",       \"created_at\": \"2021-04-26T08:28:59.221274Z\",       \"domains\": [         {           \"name\": \"demo-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     },     {       \"id\": \"e38caea5-5acb-4fec-aa65-81749546b470\",       \"name\": \"www\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-05-04T08:56:40.557792Z\",       \"created_at\": \"2021-05-04T08:56:40.557792Z\",       \"domains\": [         {           \"name\": \"www-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     }   ],   \"limit\": 20,   \"offset\": 0,   \"count\": 2 } ```  # Authentication  The Koyeb API uses API keys to authenticate requests. You can view and manage your API keys in the Koyeb Control Panel.  Your API keys carry many privileges. Take care to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, Slack, client-side code, and so forth.  To make an authenticated request, include a bearer-type Authorization header containing your API key, prefixed with `Bearer `:  ``` Authorization: Bearer <your-api-key> ```  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  # Errors  We use conventional HTTP response codes to indicate the success or failure of an API request.  In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g., a required parameter is missing, resource not found), and codes in the 5xx range indicate an error with our servers.  ## HTTP Status Code Summary  - `200 OK` - Everything worked as expected. - `400 Bad Request` - Often missing a required parameter. - `401 Unauthorized` - No valid API key provided. - `402 Request Failed` - Parameters were valid, but the request failed. - `403 Forbidden` - Insufficient privileges to access the requested resource. - `404 Not Found` - The requested item doesn't exist. - `50x` Server errors - Something went wrong on the API domain's end.  Input and output data must be valid JSON with the appropriate Content-Type header set.  ## Attributes  - code:    - `invalid_argument`: Occur when your request has invalid parameters.   - `unknown_resource`: Occur when the resource doesn't exist.   - `authentication_error`: Arise when there is a problem with authentication.   - `rate_limit_error`: Too many requests hit the API too quickly.   - `authorization_required`: You don't have sufficient right to access the resource.   - `api_error`: API errors use in case of a problem with api domain's servers (extremely uncommon).  - message:    - A human-readable error giving more details about the error  - status:    - The HTTP code returned by the server.  - fields (Optional):    - An array of parameters with a human-readable message giving more details about the error.  ## Errors responses example  Response 400 (application/json)  ``` {   \"status\": 400,   \"code\": \"invalid_argument\",   \"message\": \"Invalid argument\",   \"fields\": [     {       \"field\": \"definition.docker.image\",       \"description\": \"cannot be blank\"     }   ] } ```  Response 401 (application/json)  ``` {   \"status\": 401,   \"code\": \"authentication_error\",   \"message\": \"Request unauthenticated with bearer\" } ```  Response 403 (application/json)  ``` {   \"status\": 403,   \"code\": \"authentication_required\",   \"message\": \"You don't have the requested permission to access to resource.\" } ```   Response 404 (application/json)  ``` {   \"status\": 404,   \"code\": \"unknown_resource\",   \"message\": \"App not found\" } ```  # Pagination  All top-level API resources receive a paginated response. For instance, list Apps, list Services, and others share a standard structure, taking two parameters: limit and offset.  **Paginated request**  Requests against paginated endpoints accept two query string arguments:  * *limit*, a positive integer lower or equal to 100 to select the number of items to display. * *offset*, an integer to indicate the position to start at.  Paginated endpoints usually also accept filters to search and sort results. These filters are documented along with each endpoint documentation. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// DeploymentScalingTarget struct for DeploymentScalingTarget
type DeploymentScalingTarget struct {
	AverageCpu *DeploymentScalingTargetAverageCPU `json:"average_cpu,omitempty"`
	AverageMem *DeploymentScalingTargetAverageMem `json:"average_mem,omitempty"`
	RequestsPerSecond *DeploymentScalingTargetRequestsPerSecond `json:"requests_per_second,omitempty"`
	ConcurrentRequests *DeploymentScalingTargetConcurrentRequests `json:"concurrent_requests,omitempty"`
	RequestsResponseTime *DeploymentScalingTargetRequestsResponseTime `json:"requests_response_time,omitempty"`
	SleepIdleDelay *DeploymentScalingTargetSleepIdleDelay `json:"sleep_idle_delay,omitempty"`
}

// NewDeploymentScalingTarget instantiates a new DeploymentScalingTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentScalingTarget() *DeploymentScalingTarget {
	this := DeploymentScalingTarget{}
	return &this
}

// NewDeploymentScalingTargetWithDefaults instantiates a new DeploymentScalingTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentScalingTargetWithDefaults() *DeploymentScalingTarget {
	this := DeploymentScalingTarget{}
	return &this
}

// GetAverageCpu returns the AverageCpu field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetAverageCpu() DeploymentScalingTargetAverageCPU {
	if o == nil || isNil(o.AverageCpu) {
		var ret DeploymentScalingTargetAverageCPU
		return ret
	}
	return *o.AverageCpu
}

// GetAverageCpuOk returns a tuple with the AverageCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetAverageCpuOk() (*DeploymentScalingTargetAverageCPU, bool) {
	if o == nil || isNil(o.AverageCpu) {
    return nil, false
	}
	return o.AverageCpu, true
}

// HasAverageCpu returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasAverageCpu() bool {
	if o != nil && !isNil(o.AverageCpu) {
		return true
	}

	return false
}

// SetAverageCpu gets a reference to the given DeploymentScalingTargetAverageCPU and assigns it to the AverageCpu field.
func (o *DeploymentScalingTarget) SetAverageCpu(v DeploymentScalingTargetAverageCPU) {
	o.AverageCpu = &v
}

// GetAverageMem returns the AverageMem field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetAverageMem() DeploymentScalingTargetAverageMem {
	if o == nil || isNil(o.AverageMem) {
		var ret DeploymentScalingTargetAverageMem
		return ret
	}
	return *o.AverageMem
}

// GetAverageMemOk returns a tuple with the AverageMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetAverageMemOk() (*DeploymentScalingTargetAverageMem, bool) {
	if o == nil || isNil(o.AverageMem) {
    return nil, false
	}
	return o.AverageMem, true
}

// HasAverageMem returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasAverageMem() bool {
	if o != nil && !isNil(o.AverageMem) {
		return true
	}

	return false
}

// SetAverageMem gets a reference to the given DeploymentScalingTargetAverageMem and assigns it to the AverageMem field.
func (o *DeploymentScalingTarget) SetAverageMem(v DeploymentScalingTargetAverageMem) {
	o.AverageMem = &v
}

// GetRequestsPerSecond returns the RequestsPerSecond field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetRequestsPerSecond() DeploymentScalingTargetRequestsPerSecond {
	if o == nil || isNil(o.RequestsPerSecond) {
		var ret DeploymentScalingTargetRequestsPerSecond
		return ret
	}
	return *o.RequestsPerSecond
}

// GetRequestsPerSecondOk returns a tuple with the RequestsPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetRequestsPerSecondOk() (*DeploymentScalingTargetRequestsPerSecond, bool) {
	if o == nil || isNil(o.RequestsPerSecond) {
    return nil, false
	}
	return o.RequestsPerSecond, true
}

// HasRequestsPerSecond returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasRequestsPerSecond() bool {
	if o != nil && !isNil(o.RequestsPerSecond) {
		return true
	}

	return false
}

// SetRequestsPerSecond gets a reference to the given DeploymentScalingTargetRequestsPerSecond and assigns it to the RequestsPerSecond field.
func (o *DeploymentScalingTarget) SetRequestsPerSecond(v DeploymentScalingTargetRequestsPerSecond) {
	o.RequestsPerSecond = &v
}

// GetConcurrentRequests returns the ConcurrentRequests field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetConcurrentRequests() DeploymentScalingTargetConcurrentRequests {
	if o == nil || isNil(o.ConcurrentRequests) {
		var ret DeploymentScalingTargetConcurrentRequests
		return ret
	}
	return *o.ConcurrentRequests
}

// GetConcurrentRequestsOk returns a tuple with the ConcurrentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetConcurrentRequestsOk() (*DeploymentScalingTargetConcurrentRequests, bool) {
	if o == nil || isNil(o.ConcurrentRequests) {
    return nil, false
	}
	return o.ConcurrentRequests, true
}

// HasConcurrentRequests returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasConcurrentRequests() bool {
	if o != nil && !isNil(o.ConcurrentRequests) {
		return true
	}

	return false
}

// SetConcurrentRequests gets a reference to the given DeploymentScalingTargetConcurrentRequests and assigns it to the ConcurrentRequests field.
func (o *DeploymentScalingTarget) SetConcurrentRequests(v DeploymentScalingTargetConcurrentRequests) {
	o.ConcurrentRequests = &v
}

// GetRequestsResponseTime returns the RequestsResponseTime field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetRequestsResponseTime() DeploymentScalingTargetRequestsResponseTime {
	if o == nil || isNil(o.RequestsResponseTime) {
		var ret DeploymentScalingTargetRequestsResponseTime
		return ret
	}
	return *o.RequestsResponseTime
}

// GetRequestsResponseTimeOk returns a tuple with the RequestsResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetRequestsResponseTimeOk() (*DeploymentScalingTargetRequestsResponseTime, bool) {
	if o == nil || isNil(o.RequestsResponseTime) {
    return nil, false
	}
	return o.RequestsResponseTime, true
}

// HasRequestsResponseTime returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasRequestsResponseTime() bool {
	if o != nil && !isNil(o.RequestsResponseTime) {
		return true
	}

	return false
}

// SetRequestsResponseTime gets a reference to the given DeploymentScalingTargetRequestsResponseTime and assigns it to the RequestsResponseTime field.
func (o *DeploymentScalingTarget) SetRequestsResponseTime(v DeploymentScalingTargetRequestsResponseTime) {
	o.RequestsResponseTime = &v
}

// GetSleepIdleDelay returns the SleepIdleDelay field value if set, zero value otherwise.
func (o *DeploymentScalingTarget) GetSleepIdleDelay() DeploymentScalingTargetSleepIdleDelay {
	if o == nil || isNil(o.SleepIdleDelay) {
		var ret DeploymentScalingTargetSleepIdleDelay
		return ret
	}
	return *o.SleepIdleDelay
}

// GetSleepIdleDelayOk returns a tuple with the SleepIdleDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTarget) GetSleepIdleDelayOk() (*DeploymentScalingTargetSleepIdleDelay, bool) {
	if o == nil || isNil(o.SleepIdleDelay) {
    return nil, false
	}
	return o.SleepIdleDelay, true
}

// HasSleepIdleDelay returns a boolean if a field has been set.
func (o *DeploymentScalingTarget) HasSleepIdleDelay() bool {
	if o != nil && !isNil(o.SleepIdleDelay) {
		return true
	}

	return false
}

// SetSleepIdleDelay gets a reference to the given DeploymentScalingTargetSleepIdleDelay and assigns it to the SleepIdleDelay field.
func (o *DeploymentScalingTarget) SetSleepIdleDelay(v DeploymentScalingTargetSleepIdleDelay) {
	o.SleepIdleDelay = &v
}

func (o DeploymentScalingTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AverageCpu) {
		toSerialize["average_cpu"] = o.AverageCpu
	}
	if !isNil(o.AverageMem) {
		toSerialize["average_mem"] = o.AverageMem
	}
	if !isNil(o.RequestsPerSecond) {
		toSerialize["requests_per_second"] = o.RequestsPerSecond
	}
	if !isNil(o.ConcurrentRequests) {
		toSerialize["concurrent_requests"] = o.ConcurrentRequests
	}
	if !isNil(o.RequestsResponseTime) {
		toSerialize["requests_response_time"] = o.RequestsResponseTime
	}
	if !isNil(o.SleepIdleDelay) {
		toSerialize["sleep_idle_delay"] = o.SleepIdleDelay
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentScalingTarget struct {
	value *DeploymentScalingTarget
	isSet bool
}

func (v NullableDeploymentScalingTarget) Get() *DeploymentScalingTarget {
	return v.value
}

func (v *NullableDeploymentScalingTarget) Set(val *DeploymentScalingTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentScalingTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentScalingTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentScalingTarget(val *DeploymentScalingTarget) *NullableDeploymentScalingTarget {
	return &NullableDeploymentScalingTarget{value: val, isSet: true}
}

func (v NullableDeploymentScalingTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentScalingTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


