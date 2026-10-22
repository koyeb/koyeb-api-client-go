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

// DeploymentProvisioningInfo struct for DeploymentProvisioningInfo
type DeploymentProvisioningInfo struct {
	// The git sha for this build (we resolve the reference at the start of the build).
	Sha *string `json:"sha,omitempty"`
	// The docker image built as a result of this build.
	Image *string `json:"image,omitempty"`
	// Some info about the build.
	Stages []DeploymentProvisioningInfoStage `json:"stages,omitempty"`
}

// NewDeploymentProvisioningInfo instantiates a new DeploymentProvisioningInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfo() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// NewDeploymentProvisioningInfoWithDefaults instantiates a new DeploymentProvisioningInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoWithDefaults() *DeploymentProvisioningInfo {
	this := DeploymentProvisioningInfo{}
	return &this
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetSha() string {
	if o == nil || isNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetShaOk() (*string, bool) {
	if o == nil || isNil(o.Sha) {
    return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasSha() bool {
	if o != nil && !isNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *DeploymentProvisioningInfo) SetSha(v string) {
	o.Sha = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DeploymentProvisioningInfo) SetImage(v string) {
	o.Image = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfo) GetStages() []DeploymentProvisioningInfoStage {
	if o == nil || isNil(o.Stages) {
		var ret []DeploymentProvisioningInfoStage
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfo) GetStagesOk() ([]DeploymentProvisioningInfoStage, bool) {
	if o == nil || isNil(o.Stages) {
    return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfo) HasStages() bool {
	if o != nil && !isNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []DeploymentProvisioningInfoStage and assigns it to the Stages field.
func (o *DeploymentProvisioningInfo) SetStages(v []DeploymentProvisioningInfoStage) {
	o.Stages = v
}

func (o DeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfo struct {
	value *DeploymentProvisioningInfo
	isSet bool
}

func (v NullableDeploymentProvisioningInfo) Get() *DeploymentProvisioningInfo {
	return v.value
}

func (v *NullableDeploymentProvisioningInfo) Set(val *DeploymentProvisioningInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfo(val *DeploymentProvisioningInfo) *NullableDeploymentProvisioningInfo {
	return &NullableDeploymentProvisioningInfo{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


