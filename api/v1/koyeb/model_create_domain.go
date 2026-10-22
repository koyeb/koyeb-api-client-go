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

// CreateDomain struct for CreateDomain
type CreateDomain struct {
	Name *string `json:"name,omitempty"`
	Type *DomainType `json:"type,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	Cloudflare map[string]interface{} `json:"cloudflare,omitempty"`
	Koyeb *DomainLoadBalancerKoyeb `json:"koyeb,omitempty"`
}

// NewCreateDomain instantiates a new CreateDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomain() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// NewCreateDomainWithDefaults instantiates a new CreateDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainWithDefaults() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDomain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDomain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDomain) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateDomain) GetType() DomainType {
	if o == nil || isNil(o.Type) {
		var ret DomainType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetTypeOk() (*DomainType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateDomain) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DomainType and assigns it to the Type field.
func (o *CreateDomain) SetType(v DomainType) {
	o.Type = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateDomain) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateDomain) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateDomain) SetAppId(v string) {
	o.AppId = &v
}

// GetCloudflare returns the Cloudflare field value if set, zero value otherwise.
func (o *CreateDomain) GetCloudflare() map[string]interface{} {
	if o == nil || isNil(o.Cloudflare) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cloudflare
}

// GetCloudflareOk returns a tuple with the Cloudflare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetCloudflareOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Cloudflare) {
    return map[string]interface{}{}, false
	}
	return o.Cloudflare, true
}

// HasCloudflare returns a boolean if a field has been set.
func (o *CreateDomain) HasCloudflare() bool {
	if o != nil && !isNil(o.Cloudflare) {
		return true
	}

	return false
}

// SetCloudflare gets a reference to the given map[string]interface{} and assigns it to the Cloudflare field.
func (o *CreateDomain) SetCloudflare(v map[string]interface{}) {
	o.Cloudflare = v
}

// GetKoyeb returns the Koyeb field value if set, zero value otherwise.
func (o *CreateDomain) GetKoyeb() DomainLoadBalancerKoyeb {
	if o == nil || isNil(o.Koyeb) {
		var ret DomainLoadBalancerKoyeb
		return ret
	}
	return *o.Koyeb
}

// GetKoyebOk returns a tuple with the Koyeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetKoyebOk() (*DomainLoadBalancerKoyeb, bool) {
	if o == nil || isNil(o.Koyeb) {
    return nil, false
	}
	return o.Koyeb, true
}

// HasKoyeb returns a boolean if a field has been set.
func (o *CreateDomain) HasKoyeb() bool {
	if o != nil && !isNil(o.Koyeb) {
		return true
	}

	return false
}

// SetKoyeb gets a reference to the given DomainLoadBalancerKoyeb and assigns it to the Koyeb field.
func (o *CreateDomain) SetKoyeb(v DomainLoadBalancerKoyeb) {
	o.Koyeb = &v
}

func (o CreateDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Cloudflare) {
		toSerialize["cloudflare"] = o.Cloudflare
	}
	if !isNil(o.Koyeb) {
		toSerialize["koyeb"] = o.Koyeb
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDomain struct {
	value *CreateDomain
	isSet bool
}

func (v NullableCreateDomain) Get() *CreateDomain {
	return v.value
}

func (v *NullableCreateDomain) Set(val *CreateDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomain(val *CreateDomain) *NullableCreateDomain {
	return &NullableCreateDomain{value: val, isSet: true}
}

func (v NullableCreateDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


