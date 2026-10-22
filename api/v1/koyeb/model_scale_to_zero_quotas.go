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

// ScaleToZeroQuotas struct for ScaleToZeroQuotas
type ScaleToZeroQuotas struct {
	IsDeepSleepEnabled *bool `json:"is_deep_sleep_enabled,omitempty"`
	DeepSleepIdleDelayMin *int64 `json:"deep_sleep_idle_delay_min,omitempty"`
	DeepSleepIdleDelayMax *int64 `json:"deep_sleep_idle_delay_max,omitempty"`
	IsLightSleepEnabled *bool `json:"is_light_sleep_enabled,omitempty"`
	LightSleepIdleDelayMin *int64 `json:"light_sleep_idle_delay_min,omitempty"`
	LightSleepIdleDelayMax *int64 `json:"light_sleep_idle_delay_max,omitempty"`
}

// NewScaleToZeroQuotas instantiates a new ScaleToZeroQuotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleToZeroQuotas() *ScaleToZeroQuotas {
	this := ScaleToZeroQuotas{}
	return &this
}

// NewScaleToZeroQuotasWithDefaults instantiates a new ScaleToZeroQuotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleToZeroQuotasWithDefaults() *ScaleToZeroQuotas {
	this := ScaleToZeroQuotas{}
	return &this
}

// GetIsDeepSleepEnabled returns the IsDeepSleepEnabled field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetIsDeepSleepEnabled() bool {
	if o == nil || isNil(o.IsDeepSleepEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDeepSleepEnabled
}

// GetIsDeepSleepEnabledOk returns a tuple with the IsDeepSleepEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetIsDeepSleepEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeepSleepEnabled) {
    return nil, false
	}
	return o.IsDeepSleepEnabled, true
}

// HasIsDeepSleepEnabled returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasIsDeepSleepEnabled() bool {
	if o != nil && !isNil(o.IsDeepSleepEnabled) {
		return true
	}

	return false
}

// SetIsDeepSleepEnabled gets a reference to the given bool and assigns it to the IsDeepSleepEnabled field.
func (o *ScaleToZeroQuotas) SetIsDeepSleepEnabled(v bool) {
	o.IsDeepSleepEnabled = &v
}

// GetDeepSleepIdleDelayMin returns the DeepSleepIdleDelayMin field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetDeepSleepIdleDelayMin() int64 {
	if o == nil || isNil(o.DeepSleepIdleDelayMin) {
		var ret int64
		return ret
	}
	return *o.DeepSleepIdleDelayMin
}

// GetDeepSleepIdleDelayMinOk returns a tuple with the DeepSleepIdleDelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetDeepSleepIdleDelayMinOk() (*int64, bool) {
	if o == nil || isNil(o.DeepSleepIdleDelayMin) {
    return nil, false
	}
	return o.DeepSleepIdleDelayMin, true
}

// HasDeepSleepIdleDelayMin returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasDeepSleepIdleDelayMin() bool {
	if o != nil && !isNil(o.DeepSleepIdleDelayMin) {
		return true
	}

	return false
}

// SetDeepSleepIdleDelayMin gets a reference to the given int64 and assigns it to the DeepSleepIdleDelayMin field.
func (o *ScaleToZeroQuotas) SetDeepSleepIdleDelayMin(v int64) {
	o.DeepSleepIdleDelayMin = &v
}

// GetDeepSleepIdleDelayMax returns the DeepSleepIdleDelayMax field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetDeepSleepIdleDelayMax() int64 {
	if o == nil || isNil(o.DeepSleepIdleDelayMax) {
		var ret int64
		return ret
	}
	return *o.DeepSleepIdleDelayMax
}

// GetDeepSleepIdleDelayMaxOk returns a tuple with the DeepSleepIdleDelayMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetDeepSleepIdleDelayMaxOk() (*int64, bool) {
	if o == nil || isNil(o.DeepSleepIdleDelayMax) {
    return nil, false
	}
	return o.DeepSleepIdleDelayMax, true
}

// HasDeepSleepIdleDelayMax returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasDeepSleepIdleDelayMax() bool {
	if o != nil && !isNil(o.DeepSleepIdleDelayMax) {
		return true
	}

	return false
}

// SetDeepSleepIdleDelayMax gets a reference to the given int64 and assigns it to the DeepSleepIdleDelayMax field.
func (o *ScaleToZeroQuotas) SetDeepSleepIdleDelayMax(v int64) {
	o.DeepSleepIdleDelayMax = &v
}

// GetIsLightSleepEnabled returns the IsLightSleepEnabled field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetIsLightSleepEnabled() bool {
	if o == nil || isNil(o.IsLightSleepEnabled) {
		var ret bool
		return ret
	}
	return *o.IsLightSleepEnabled
}

// GetIsLightSleepEnabledOk returns a tuple with the IsLightSleepEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetIsLightSleepEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsLightSleepEnabled) {
    return nil, false
	}
	return o.IsLightSleepEnabled, true
}

// HasIsLightSleepEnabled returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasIsLightSleepEnabled() bool {
	if o != nil && !isNil(o.IsLightSleepEnabled) {
		return true
	}

	return false
}

// SetIsLightSleepEnabled gets a reference to the given bool and assigns it to the IsLightSleepEnabled field.
func (o *ScaleToZeroQuotas) SetIsLightSleepEnabled(v bool) {
	o.IsLightSleepEnabled = &v
}

// GetLightSleepIdleDelayMin returns the LightSleepIdleDelayMin field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetLightSleepIdleDelayMin() int64 {
	if o == nil || isNil(o.LightSleepIdleDelayMin) {
		var ret int64
		return ret
	}
	return *o.LightSleepIdleDelayMin
}

// GetLightSleepIdleDelayMinOk returns a tuple with the LightSleepIdleDelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetLightSleepIdleDelayMinOk() (*int64, bool) {
	if o == nil || isNil(o.LightSleepIdleDelayMin) {
    return nil, false
	}
	return o.LightSleepIdleDelayMin, true
}

// HasLightSleepIdleDelayMin returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasLightSleepIdleDelayMin() bool {
	if o != nil && !isNil(o.LightSleepIdleDelayMin) {
		return true
	}

	return false
}

// SetLightSleepIdleDelayMin gets a reference to the given int64 and assigns it to the LightSleepIdleDelayMin field.
func (o *ScaleToZeroQuotas) SetLightSleepIdleDelayMin(v int64) {
	o.LightSleepIdleDelayMin = &v
}

// GetLightSleepIdleDelayMax returns the LightSleepIdleDelayMax field value if set, zero value otherwise.
func (o *ScaleToZeroQuotas) GetLightSleepIdleDelayMax() int64 {
	if o == nil || isNil(o.LightSleepIdleDelayMax) {
		var ret int64
		return ret
	}
	return *o.LightSleepIdleDelayMax
}

// GetLightSleepIdleDelayMaxOk returns a tuple with the LightSleepIdleDelayMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleToZeroQuotas) GetLightSleepIdleDelayMaxOk() (*int64, bool) {
	if o == nil || isNil(o.LightSleepIdleDelayMax) {
    return nil, false
	}
	return o.LightSleepIdleDelayMax, true
}

// HasLightSleepIdleDelayMax returns a boolean if a field has been set.
func (o *ScaleToZeroQuotas) HasLightSleepIdleDelayMax() bool {
	if o != nil && !isNil(o.LightSleepIdleDelayMax) {
		return true
	}

	return false
}

// SetLightSleepIdleDelayMax gets a reference to the given int64 and assigns it to the LightSleepIdleDelayMax field.
func (o *ScaleToZeroQuotas) SetLightSleepIdleDelayMax(v int64) {
	o.LightSleepIdleDelayMax = &v
}

func (o ScaleToZeroQuotas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsDeepSleepEnabled) {
		toSerialize["is_deep_sleep_enabled"] = o.IsDeepSleepEnabled
	}
	if !isNil(o.DeepSleepIdleDelayMin) {
		toSerialize["deep_sleep_idle_delay_min"] = o.DeepSleepIdleDelayMin
	}
	if !isNil(o.DeepSleepIdleDelayMax) {
		toSerialize["deep_sleep_idle_delay_max"] = o.DeepSleepIdleDelayMax
	}
	if !isNil(o.IsLightSleepEnabled) {
		toSerialize["is_light_sleep_enabled"] = o.IsLightSleepEnabled
	}
	if !isNil(o.LightSleepIdleDelayMin) {
		toSerialize["light_sleep_idle_delay_min"] = o.LightSleepIdleDelayMin
	}
	if !isNil(o.LightSleepIdleDelayMax) {
		toSerialize["light_sleep_idle_delay_max"] = o.LightSleepIdleDelayMax
	}
	return json.Marshal(toSerialize)
}

type NullableScaleToZeroQuotas struct {
	value *ScaleToZeroQuotas
	isSet bool
}

func (v NullableScaleToZeroQuotas) Get() *ScaleToZeroQuotas {
	return v.value
}

func (v *NullableScaleToZeroQuotas) Set(val *ScaleToZeroQuotas) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleToZeroQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleToZeroQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleToZeroQuotas(val *ScaleToZeroQuotas) *NullableScaleToZeroQuotas {
	return &NullableScaleToZeroQuotas{value: val, isSet: true}
}

func (v NullableScaleToZeroQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleToZeroQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


