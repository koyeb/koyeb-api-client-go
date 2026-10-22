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

// DeploymentHealthCheck struct for DeploymentHealthCheck
type DeploymentHealthCheck struct {
	GracePeriod *int64 `json:"grace_period,omitempty"`
	Interval *int64 `json:"interval,omitempty"`
	RestartLimit *int64 `json:"restart_limit,omitempty"`
	Timeout *int64 `json:"timeout,omitempty"`
	Tcp *TCPHealthCheck `json:"tcp,omitempty"`
	Http *HTTPHealthCheck `json:"http,omitempty"`
	Grpc *GRPCHealthCheck `json:"grpc,omitempty"`
}

// NewDeploymentHealthCheck instantiates a new DeploymentHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHealthCheck() *DeploymentHealthCheck {
	this := DeploymentHealthCheck{}
	return &this
}

// NewDeploymentHealthCheckWithDefaults instantiates a new DeploymentHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHealthCheckWithDefaults() *DeploymentHealthCheck {
	this := DeploymentHealthCheck{}
	return &this
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetGracePeriod() int64 {
	if o == nil || isNil(o.GracePeriod) {
		var ret int64
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetGracePeriodOk() (*int64, bool) {
	if o == nil || isNil(o.GracePeriod) {
    return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasGracePeriod() bool {
	if o != nil && !isNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int64 and assigns it to the GracePeriod field.
func (o *DeploymentHealthCheck) SetGracePeriod(v int64) {
	o.GracePeriod = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetInterval() int64 {
	if o == nil || isNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetIntervalOk() (*int64, bool) {
	if o == nil || isNil(o.Interval) {
    return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *DeploymentHealthCheck) SetInterval(v int64) {
	o.Interval = &v
}

// GetRestartLimit returns the RestartLimit field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetRestartLimit() int64 {
	if o == nil || isNil(o.RestartLimit) {
		var ret int64
		return ret
	}
	return *o.RestartLimit
}

// GetRestartLimitOk returns a tuple with the RestartLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetRestartLimitOk() (*int64, bool) {
	if o == nil || isNil(o.RestartLimit) {
    return nil, false
	}
	return o.RestartLimit, true
}

// HasRestartLimit returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasRestartLimit() bool {
	if o != nil && !isNil(o.RestartLimit) {
		return true
	}

	return false
}

// SetRestartLimit gets a reference to the given int64 and assigns it to the RestartLimit field.
func (o *DeploymentHealthCheck) SetRestartLimit(v int64) {
	o.RestartLimit = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetTimeout() int64 {
	if o == nil || isNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetTimeoutOk() (*int64, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *DeploymentHealthCheck) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetTcp() TCPHealthCheck {
	if o == nil || isNil(o.Tcp) {
		var ret TCPHealthCheck
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetTcpOk() (*TCPHealthCheck, bool) {
	if o == nil || isNil(o.Tcp) {
    return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasTcp() bool {
	if o != nil && !isNil(o.Tcp) {
		return true
	}

	return false
}

// SetTcp gets a reference to the given TCPHealthCheck and assigns it to the Tcp field.
func (o *DeploymentHealthCheck) SetTcp(v TCPHealthCheck) {
	o.Tcp = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetHttp() HTTPHealthCheck {
	if o == nil || isNil(o.Http) {
		var ret HTTPHealthCheck
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetHttpOk() (*HTTPHealthCheck, bool) {
	if o == nil || isNil(o.Http) {
    return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasHttp() bool {
	if o != nil && !isNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given HTTPHealthCheck and assigns it to the Http field.
func (o *DeploymentHealthCheck) SetHttp(v HTTPHealthCheck) {
	o.Http = &v
}

// GetGrpc returns the Grpc field value if set, zero value otherwise.
func (o *DeploymentHealthCheck) GetGrpc() GRPCHealthCheck {
	if o == nil || isNil(o.Grpc) {
		var ret GRPCHealthCheck
		return ret
	}
	return *o.Grpc
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHealthCheck) GetGrpcOk() (*GRPCHealthCheck, bool) {
	if o == nil || isNil(o.Grpc) {
    return nil, false
	}
	return o.Grpc, true
}

// HasGrpc returns a boolean if a field has been set.
func (o *DeploymentHealthCheck) HasGrpc() bool {
	if o != nil && !isNil(o.Grpc) {
		return true
	}

	return false
}

// SetGrpc gets a reference to the given GRPCHealthCheck and assigns it to the Grpc field.
func (o *DeploymentHealthCheck) SetGrpc(v GRPCHealthCheck) {
	o.Grpc = &v
}

func (o DeploymentHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GracePeriod) {
		toSerialize["grace_period"] = o.GracePeriod
	}
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !isNil(o.RestartLimit) {
		toSerialize["restart_limit"] = o.RestartLimit
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.Tcp) {
		toSerialize["tcp"] = o.Tcp
	}
	if !isNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	if !isNil(o.Grpc) {
		toSerialize["grpc"] = o.Grpc
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHealthCheck struct {
	value *DeploymentHealthCheck
	isSet bool
}

func (v NullableDeploymentHealthCheck) Get() *DeploymentHealthCheck {
	return v.value
}

func (v *NullableDeploymentHealthCheck) Set(val *DeploymentHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHealthCheck(val *DeploymentHealthCheck) *NullableDeploymentHealthCheck {
	return &NullableDeploymentHealthCheck{value: val, isSet: true}
}

func (v NullableDeploymentHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


