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

// ServiceViewRow struct for ServiceViewRow
type ServiceViewRow struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *ServiceType `json:"type,omitempty"`
	Status *ServiceStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	App *ServiceViewApp `json:"app,omitempty"`
	ActiveDeployment *ServiceViewActiveDeployment `json:"active_deployment,omitempty"`
	LatestDeployment *ServiceViewLatestDeployment `json:"latest_deployment,omitempty"`
	NeedsAttention *bool `json:"needs_attention,omitempty"`
	Regions []string `json:"regions,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	Source *ServiceViewSource `json:"source,omitempty"`
}

// NewServiceViewRow instantiates a new ServiceViewRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceViewRow() *ServiceViewRow {
	this := ServiceViewRow{}
	var type_ ServiceType = SERVICETYPE_INVALID_TYPE
	this.Type = &type_
	var status ServiceStatus = SERVICESTATUS_STARTING
	this.Status = &status
	return &this
}

// NewServiceViewRowWithDefaults instantiates a new ServiceViewRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceViewRowWithDefaults() *ServiceViewRow {
	this := ServiceViewRow{}
	var type_ ServiceType = SERVICETYPE_INVALID_TYPE
	this.Type = &type_
	var status ServiceStatus = SERVICESTATUS_STARTING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceViewRow) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceViewRow) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceViewRow) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceViewRow) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceViewRow) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceViewRow) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceViewRow) GetType() ServiceType {
	if o == nil || isNil(o.Type) {
		var ret ServiceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetTypeOk() (*ServiceType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceViewRow) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceType and assigns it to the Type field.
func (o *ServiceViewRow) SetType(v ServiceType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceViewRow) GetStatus() ServiceStatus {
	if o == nil || isNil(o.Status) {
		var ret ServiceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetStatusOk() (*ServiceStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceViewRow) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceStatus and assigns it to the Status field.
func (o *ServiceViewRow) SetStatus(v ServiceStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ServiceViewRow) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ServiceViewRow) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ServiceViewRow) SetMessages(v []string) {
	o.Messages = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceViewRow) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceViewRow) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceViewRow) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceViewRow) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceViewRow) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ServiceViewRow) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceViewRow) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceViewRow) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceViewRow) SetAppId(v string) {
	o.AppId = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ServiceViewRow) GetApp() ServiceViewApp {
	if o == nil || isNil(o.App) {
		var ret ServiceViewApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetAppOk() (*ServiceViewApp, bool) {
	if o == nil || isNil(o.App) {
    return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ServiceViewRow) HasApp() bool {
	if o != nil && !isNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ServiceViewApp and assigns it to the App field.
func (o *ServiceViewRow) SetApp(v ServiceViewApp) {
	o.App = &v
}

// GetActiveDeployment returns the ActiveDeployment field value if set, zero value otherwise.
func (o *ServiceViewRow) GetActiveDeployment() ServiceViewActiveDeployment {
	if o == nil || isNil(o.ActiveDeployment) {
		var ret ServiceViewActiveDeployment
		return ret
	}
	return *o.ActiveDeployment
}

// GetActiveDeploymentOk returns a tuple with the ActiveDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetActiveDeploymentOk() (*ServiceViewActiveDeployment, bool) {
	if o == nil || isNil(o.ActiveDeployment) {
    return nil, false
	}
	return o.ActiveDeployment, true
}

// HasActiveDeployment returns a boolean if a field has been set.
func (o *ServiceViewRow) HasActiveDeployment() bool {
	if o != nil && !isNil(o.ActiveDeployment) {
		return true
	}

	return false
}

// SetActiveDeployment gets a reference to the given ServiceViewActiveDeployment and assigns it to the ActiveDeployment field.
func (o *ServiceViewRow) SetActiveDeployment(v ServiceViewActiveDeployment) {
	o.ActiveDeployment = &v
}

// GetLatestDeployment returns the LatestDeployment field value if set, zero value otherwise.
func (o *ServiceViewRow) GetLatestDeployment() ServiceViewLatestDeployment {
	if o == nil || isNil(o.LatestDeployment) {
		var ret ServiceViewLatestDeployment
		return ret
	}
	return *o.LatestDeployment
}

// GetLatestDeploymentOk returns a tuple with the LatestDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetLatestDeploymentOk() (*ServiceViewLatestDeployment, bool) {
	if o == nil || isNil(o.LatestDeployment) {
    return nil, false
	}
	return o.LatestDeployment, true
}

// HasLatestDeployment returns a boolean if a field has been set.
func (o *ServiceViewRow) HasLatestDeployment() bool {
	if o != nil && !isNil(o.LatestDeployment) {
		return true
	}

	return false
}

// SetLatestDeployment gets a reference to the given ServiceViewLatestDeployment and assigns it to the LatestDeployment field.
func (o *ServiceViewRow) SetLatestDeployment(v ServiceViewLatestDeployment) {
	o.LatestDeployment = &v
}

// GetNeedsAttention returns the NeedsAttention field value if set, zero value otherwise.
func (o *ServiceViewRow) GetNeedsAttention() bool {
	if o == nil || isNil(o.NeedsAttention) {
		var ret bool
		return ret
	}
	return *o.NeedsAttention
}

// GetNeedsAttentionOk returns a tuple with the NeedsAttention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetNeedsAttentionOk() (*bool, bool) {
	if o == nil || isNil(o.NeedsAttention) {
    return nil, false
	}
	return o.NeedsAttention, true
}

// HasNeedsAttention returns a boolean if a field has been set.
func (o *ServiceViewRow) HasNeedsAttention() bool {
	if o != nil && !isNil(o.NeedsAttention) {
		return true
	}

	return false
}

// SetNeedsAttention gets a reference to the given bool and assigns it to the NeedsAttention field.
func (o *ServiceViewRow) SetNeedsAttention(v bool) {
	o.NeedsAttention = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ServiceViewRow) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *ServiceViewRow) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *ServiceViewRow) SetRegions(v []string) {
	o.Regions = v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *ServiceViewRow) GetInstanceType() string {
	if o == nil || isNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetInstanceTypeOk() (*string, bool) {
	if o == nil || isNil(o.InstanceType) {
    return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *ServiceViewRow) HasInstanceType() bool {
	if o != nil && !isNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *ServiceViewRow) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ServiceViewRow) GetSource() ServiceViewSource {
	if o == nil || isNil(o.Source) {
		var ret ServiceViewSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceViewRow) GetSourceOk() (*ServiceViewSource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ServiceViewRow) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ServiceViewSource and assigns it to the Source field.
func (o *ServiceViewRow) SetSource(v ServiceViewSource) {
	o.Source = &v
}

func (o ServiceViewRow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !isNil(o.ActiveDeployment) {
		toSerialize["active_deployment"] = o.ActiveDeployment
	}
	if !isNil(o.LatestDeployment) {
		toSerialize["latest_deployment"] = o.LatestDeployment
	}
	if !isNil(o.NeedsAttention) {
		toSerialize["needs_attention"] = o.NeedsAttention
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableServiceViewRow struct {
	value *ServiceViewRow
	isSet bool
}

func (v NullableServiceViewRow) Get() *ServiceViewRow {
	return v.value
}

func (v *NullableServiceViewRow) Set(val *ServiceViewRow) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceViewRow) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceViewRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceViewRow(val *ServiceViewRow) *NullableServiceViewRow {
	return &NullableServiceViewRow{value: val, isSet: true}
}

func (v NullableServiceViewRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceViewRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


