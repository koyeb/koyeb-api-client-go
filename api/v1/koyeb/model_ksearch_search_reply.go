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

// KsearchSearchReply struct for KsearchSearchReply
type KsearchSearchReply struct {
	Organizations []KsearchOrganization `json:"organizations,omitempty"`
	Users []KsearchUser `json:"users,omitempty"`
	Apps []KsearchApp `json:"apps,omitempty"`
	Services []KsearchService `json:"services,omitempty"`
	GlobalDeployments []KsearchGlobalDeployment `json:"global_deployments,omitempty"`
	RegionalDeployments []KsearchRegionalDeployment `json:"regional_deployments,omitempty"`
	Instances []KsearchInstance `json:"instances,omitempty"`
}

// NewKsearchSearchReply instantiates a new KsearchSearchReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsearchSearchReply() *KsearchSearchReply {
	this := KsearchSearchReply{}
	return &this
}

// NewKsearchSearchReplyWithDefaults instantiates a new KsearchSearchReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsearchSearchReplyWithDefaults() *KsearchSearchReply {
	this := KsearchSearchReply{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetOrganizations() []KsearchOrganization {
	if o == nil || isNil(o.Organizations) {
		var ret []KsearchOrganization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetOrganizationsOk() ([]KsearchOrganization, bool) {
	if o == nil || isNil(o.Organizations) {
    return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasOrganizations() bool {
	if o != nil && !isNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []KsearchOrganization and assigns it to the Organizations field.
func (o *KsearchSearchReply) SetOrganizations(v []KsearchOrganization) {
	o.Organizations = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetUsers() []KsearchUser {
	if o == nil || isNil(o.Users) {
		var ret []KsearchUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetUsersOk() ([]KsearchUser, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []KsearchUser and assigns it to the Users field.
func (o *KsearchSearchReply) SetUsers(v []KsearchUser) {
	o.Users = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetApps() []KsearchApp {
	if o == nil || isNil(o.Apps) {
		var ret []KsearchApp
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetAppsOk() ([]KsearchApp, bool) {
	if o == nil || isNil(o.Apps) {
    return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasApps() bool {
	if o != nil && !isNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []KsearchApp and assigns it to the Apps field.
func (o *KsearchSearchReply) SetApps(v []KsearchApp) {
	o.Apps = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetServices() []KsearchService {
	if o == nil || isNil(o.Services) {
		var ret []KsearchService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetServicesOk() ([]KsearchService, bool) {
	if o == nil || isNil(o.Services) {
    return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasServices() bool {
	if o != nil && !isNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []KsearchService and assigns it to the Services field.
func (o *KsearchSearchReply) SetServices(v []KsearchService) {
	o.Services = v
}

// GetGlobalDeployments returns the GlobalDeployments field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetGlobalDeployments() []KsearchGlobalDeployment {
	if o == nil || isNil(o.GlobalDeployments) {
		var ret []KsearchGlobalDeployment
		return ret
	}
	return o.GlobalDeployments
}

// GetGlobalDeploymentsOk returns a tuple with the GlobalDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetGlobalDeploymentsOk() ([]KsearchGlobalDeployment, bool) {
	if o == nil || isNil(o.GlobalDeployments) {
    return nil, false
	}
	return o.GlobalDeployments, true
}

// HasGlobalDeployments returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasGlobalDeployments() bool {
	if o != nil && !isNil(o.GlobalDeployments) {
		return true
	}

	return false
}

// SetGlobalDeployments gets a reference to the given []KsearchGlobalDeployment and assigns it to the GlobalDeployments field.
func (o *KsearchSearchReply) SetGlobalDeployments(v []KsearchGlobalDeployment) {
	o.GlobalDeployments = v
}

// GetRegionalDeployments returns the RegionalDeployments field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetRegionalDeployments() []KsearchRegionalDeployment {
	if o == nil || isNil(o.RegionalDeployments) {
		var ret []KsearchRegionalDeployment
		return ret
	}
	return o.RegionalDeployments
}

// GetRegionalDeploymentsOk returns a tuple with the RegionalDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetRegionalDeploymentsOk() ([]KsearchRegionalDeployment, bool) {
	if o == nil || isNil(o.RegionalDeployments) {
    return nil, false
	}
	return o.RegionalDeployments, true
}

// HasRegionalDeployments returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasRegionalDeployments() bool {
	if o != nil && !isNil(o.RegionalDeployments) {
		return true
	}

	return false
}

// SetRegionalDeployments gets a reference to the given []KsearchRegionalDeployment and assigns it to the RegionalDeployments field.
func (o *KsearchSearchReply) SetRegionalDeployments(v []KsearchRegionalDeployment) {
	o.RegionalDeployments = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *KsearchSearchReply) GetInstances() []KsearchInstance {
	if o == nil || isNil(o.Instances) {
		var ret []KsearchInstance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchSearchReply) GetInstancesOk() ([]KsearchInstance, bool) {
	if o == nil || isNil(o.Instances) {
    return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *KsearchSearchReply) HasInstances() bool {
	if o != nil && !isNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []KsearchInstance and assigns it to the Instances field.
func (o *KsearchSearchReply) SetInstances(v []KsearchInstance) {
	o.Instances = v
}

func (o KsearchSearchReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !isNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !isNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !isNil(o.GlobalDeployments) {
		toSerialize["global_deployments"] = o.GlobalDeployments
	}
	if !isNil(o.RegionalDeployments) {
		toSerialize["regional_deployments"] = o.RegionalDeployments
	}
	if !isNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullableKsearchSearchReply struct {
	value *KsearchSearchReply
	isSet bool
}

func (v NullableKsearchSearchReply) Get() *KsearchSearchReply {
	return v.value
}

func (v *NullableKsearchSearchReply) Set(val *KsearchSearchReply) {
	v.value = val
	v.isSet = true
}

func (v NullableKsearchSearchReply) IsSet() bool {
	return v.isSet
}

func (v *NullableKsearchSearchReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsearchSearchReply(val *KsearchSearchReply) *NullableKsearchSearchReply {
	return &NullableKsearchSearchReply{value: val, isSet: true}
}

func (v NullableKsearchSearchReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsearchSearchReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


