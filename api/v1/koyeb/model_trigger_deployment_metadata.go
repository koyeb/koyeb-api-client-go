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

// TriggerDeploymentMetadata struct for TriggerDeploymentMetadata
type TriggerDeploymentMetadata struct {
	Type *TriggerDeploymentMetadataTriggerType `json:"type,omitempty"`
	Actor *TriggerDeploymentMetadataActorType `json:"actor,omitempty"`
	Git *TriggerGitDeploymentMetadata `json:"git,omitempty"`
}

// NewTriggerDeploymentMetadata instantiates a new TriggerDeploymentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerDeploymentMetadata() *TriggerDeploymentMetadata {
	this := TriggerDeploymentMetadata{}
	var type_ TriggerDeploymentMetadataTriggerType = TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN_TYPE
	this.Type = &type_
	var actor TriggerDeploymentMetadataActorType = TRIGGERDEPLOYMENTMETADATAACTORTYPE_UNKNOWN_ACTOR
	this.Actor = &actor
	return &this
}

// NewTriggerDeploymentMetadataWithDefaults instantiates a new TriggerDeploymentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerDeploymentMetadataWithDefaults() *TriggerDeploymentMetadata {
	this := TriggerDeploymentMetadata{}
	var type_ TriggerDeploymentMetadataTriggerType = TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN_TYPE
	this.Type = &type_
	var actor TriggerDeploymentMetadataActorType = TRIGGERDEPLOYMENTMETADATAACTORTYPE_UNKNOWN_ACTOR
	this.Actor = &actor
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetType() TriggerDeploymentMetadataTriggerType {
	if o == nil || isNil(o.Type) {
		var ret TriggerDeploymentMetadataTriggerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetTypeOk() (*TriggerDeploymentMetadataTriggerType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TriggerDeploymentMetadataTriggerType and assigns it to the Type field.
func (o *TriggerDeploymentMetadata) SetType(v TriggerDeploymentMetadataTriggerType) {
	o.Type = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetActor() TriggerDeploymentMetadataActorType {
	if o == nil || isNil(o.Actor) {
		var ret TriggerDeploymentMetadataActorType
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetActorOk() (*TriggerDeploymentMetadataActorType, bool) {
	if o == nil || isNil(o.Actor) {
    return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasActor() bool {
	if o != nil && !isNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given TriggerDeploymentMetadataActorType and assigns it to the Actor field.
func (o *TriggerDeploymentMetadata) SetActor(v TriggerDeploymentMetadataActorType) {
	o.Actor = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *TriggerDeploymentMetadata) GetGit() TriggerGitDeploymentMetadata {
	if o == nil || isNil(o.Git) {
		var ret TriggerGitDeploymentMetadata
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerDeploymentMetadata) GetGitOk() (*TriggerGitDeploymentMetadata, bool) {
	if o == nil || isNil(o.Git) {
    return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *TriggerDeploymentMetadata) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given TriggerGitDeploymentMetadata and assigns it to the Git field.
func (o *TriggerDeploymentMetadata) SetGit(v TriggerGitDeploymentMetadata) {
	o.Git = &v
}

func (o TriggerDeploymentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerDeploymentMetadata struct {
	value *TriggerDeploymentMetadata
	isSet bool
}

func (v NullableTriggerDeploymentMetadata) Get() *TriggerDeploymentMetadata {
	return v.value
}

func (v *NullableTriggerDeploymentMetadata) Set(val *TriggerDeploymentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerDeploymentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerDeploymentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerDeploymentMetadata(val *TriggerDeploymentMetadata) *NullableTriggerDeploymentMetadata {
	return &NullableTriggerDeploymentMetadata{value: val, isSet: true}
}

func (v NullableTriggerDeploymentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerDeploymentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


