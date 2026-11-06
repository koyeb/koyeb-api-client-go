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

// QueryServicesViewFilters struct for QueryServicesViewFilters
type QueryServicesViewFilters struct {
	Types []ServiceType `json:"types,omitempty"`
	Statuses []ServiceStatus `json:"statuses,omitempty"`
	Regions []string `json:"regions,omitempty"`
	AppIds []string `json:"app_ids,omitempty"`
	ActiveDeploymentStates []ServiceViewActiveDeploymentState `json:"active_deployment_states,omitempty"`
	LatestDeploymentStates []ServiceViewLatestDeploymentState `json:"latest_deployment_states,omitempty"`
	NeedsAttention *bool `json:"needs_attention,omitempty"`
}

// NewQueryServicesViewFilters instantiates a new QueryServicesViewFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryServicesViewFilters() *QueryServicesViewFilters {
	this := QueryServicesViewFilters{}
	return &this
}

// NewQueryServicesViewFiltersWithDefaults instantiates a new QueryServicesViewFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryServicesViewFiltersWithDefaults() *QueryServicesViewFilters {
	this := QueryServicesViewFilters{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetTypes() []ServiceType {
	if o == nil || isNil(o.Types) {
		var ret []ServiceType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetTypesOk() ([]ServiceType, bool) {
	if o == nil || isNil(o.Types) {
    return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasTypes() bool {
	if o != nil && !isNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ServiceType and assigns it to the Types field.
func (o *QueryServicesViewFilters) SetTypes(v []ServiceType) {
	o.Types = v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetStatuses() []ServiceStatus {
	if o == nil || isNil(o.Statuses) {
		var ret []ServiceStatus
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetStatusesOk() ([]ServiceStatus, bool) {
	if o == nil || isNil(o.Statuses) {
    return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasStatuses() bool {
	if o != nil && !isNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []ServiceStatus and assigns it to the Statuses field.
func (o *QueryServicesViewFilters) SetStatuses(v []ServiceStatus) {
	o.Statuses = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *QueryServicesViewFilters) SetRegions(v []string) {
	o.Regions = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *QueryServicesViewFilters) SetAppIds(v []string) {
	o.AppIds = v
}

// GetActiveDeploymentStates returns the ActiveDeploymentStates field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetActiveDeploymentStates() []ServiceViewActiveDeploymentState {
	if o == nil || isNil(o.ActiveDeploymentStates) {
		var ret []ServiceViewActiveDeploymentState
		return ret
	}
	return o.ActiveDeploymentStates
}

// GetActiveDeploymentStatesOk returns a tuple with the ActiveDeploymentStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetActiveDeploymentStatesOk() ([]ServiceViewActiveDeploymentState, bool) {
	if o == nil || isNil(o.ActiveDeploymentStates) {
    return nil, false
	}
	return o.ActiveDeploymentStates, true
}

// HasActiveDeploymentStates returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasActiveDeploymentStates() bool {
	if o != nil && !isNil(o.ActiveDeploymentStates) {
		return true
	}

	return false
}

// SetActiveDeploymentStates gets a reference to the given []ServiceViewActiveDeploymentState and assigns it to the ActiveDeploymentStates field.
func (o *QueryServicesViewFilters) SetActiveDeploymentStates(v []ServiceViewActiveDeploymentState) {
	o.ActiveDeploymentStates = v
}

// GetLatestDeploymentStates returns the LatestDeploymentStates field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetLatestDeploymentStates() []ServiceViewLatestDeploymentState {
	if o == nil || isNil(o.LatestDeploymentStates) {
		var ret []ServiceViewLatestDeploymentState
		return ret
	}
	return o.LatestDeploymentStates
}

// GetLatestDeploymentStatesOk returns a tuple with the LatestDeploymentStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetLatestDeploymentStatesOk() ([]ServiceViewLatestDeploymentState, bool) {
	if o == nil || isNil(o.LatestDeploymentStates) {
    return nil, false
	}
	return o.LatestDeploymentStates, true
}

// HasLatestDeploymentStates returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasLatestDeploymentStates() bool {
	if o != nil && !isNil(o.LatestDeploymentStates) {
		return true
	}

	return false
}

// SetLatestDeploymentStates gets a reference to the given []ServiceViewLatestDeploymentState and assigns it to the LatestDeploymentStates field.
func (o *QueryServicesViewFilters) SetLatestDeploymentStates(v []ServiceViewLatestDeploymentState) {
	o.LatestDeploymentStates = v
}

// GetNeedsAttention returns the NeedsAttention field value if set, zero value otherwise.
func (o *QueryServicesViewFilters) GetNeedsAttention() bool {
	if o == nil || isNil(o.NeedsAttention) {
		var ret bool
		return ret
	}
	return *o.NeedsAttention
}

// GetNeedsAttentionOk returns a tuple with the NeedsAttention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryServicesViewFilters) GetNeedsAttentionOk() (*bool, bool) {
	if o == nil || isNil(o.NeedsAttention) {
    return nil, false
	}
	return o.NeedsAttention, true
}

// HasNeedsAttention returns a boolean if a field has been set.
func (o *QueryServicesViewFilters) HasNeedsAttention() bool {
	if o != nil && !isNil(o.NeedsAttention) {
		return true
	}

	return false
}

// SetNeedsAttention gets a reference to the given bool and assigns it to the NeedsAttention field.
func (o *QueryServicesViewFilters) SetNeedsAttention(v bool) {
	o.NeedsAttention = &v
}

func (o QueryServicesViewFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !isNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.AppIds) {
		toSerialize["app_ids"] = o.AppIds
	}
	if !isNil(o.ActiveDeploymentStates) {
		toSerialize["active_deployment_states"] = o.ActiveDeploymentStates
	}
	if !isNil(o.LatestDeploymentStates) {
		toSerialize["latest_deployment_states"] = o.LatestDeploymentStates
	}
	if !isNil(o.NeedsAttention) {
		toSerialize["needs_attention"] = o.NeedsAttention
	}
	return json.Marshal(toSerialize)
}

type NullableQueryServicesViewFilters struct {
	value *QueryServicesViewFilters
	isSet bool
}

func (v NullableQueryServicesViewFilters) Get() *QueryServicesViewFilters {
	return v.value
}

func (v *NullableQueryServicesViewFilters) Set(val *QueryServicesViewFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryServicesViewFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryServicesViewFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryServicesViewFilters(val *QueryServicesViewFilters) *NullableQueryServicesViewFilters {
	return &NullableQueryServicesViewFilters{value: val, isSet: true}
}

func (v NullableQueryServicesViewFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryServicesViewFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


