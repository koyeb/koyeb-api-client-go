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

// GithubInstallationReply struct for GithubInstallationReply
type GithubInstallationReply struct {
	AppName *string `json:"app_name,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	Url *string `json:"url,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewGithubInstallationReply instantiates a new GithubInstallationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubInstallationReply() *GithubInstallationReply {
	this := GithubInstallationReply{}
	return &this
}

// NewGithubInstallationReplyWithDefaults instantiates a new GithubInstallationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubInstallationReplyWithDefaults() *GithubInstallationReply {
	this := GithubInstallationReply{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *GithubInstallationReply) GetAppName() string {
	if o == nil || isNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubInstallationReply) GetAppNameOk() (*string, bool) {
	if o == nil || isNil(o.AppName) {
    return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *GithubInstallationReply) HasAppName() bool {
	if o != nil && !isNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *GithubInstallationReply) SetAppName(v string) {
	o.AppName = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *GithubInstallationReply) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubInstallationReply) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *GithubInstallationReply) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *GithubInstallationReply) SetAppId(v string) {
	o.AppId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GithubInstallationReply) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubInstallationReply) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GithubInstallationReply) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GithubInstallationReply) SetUrl(v string) {
	o.Url = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GithubInstallationReply) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubInstallationReply) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GithubInstallationReply) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GithubInstallationReply) SetState(v string) {
	o.State = &v
}

func (o GithubInstallationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableGithubInstallationReply struct {
	value *GithubInstallationReply
	isSet bool
}

func (v NullableGithubInstallationReply) Get() *GithubInstallationReply {
	return v.value
}

func (v *NullableGithubInstallationReply) Set(val *GithubInstallationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubInstallationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubInstallationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubInstallationReply(val *GithubInstallationReply) *NullableGithubInstallationReply {
	return &NullableGithubInstallationReply{value: val, isSet: true}
}

func (v NullableGithubInstallationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubInstallationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


