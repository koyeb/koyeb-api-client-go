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

// QuotaUsage QuotaUsage carries an organization's current resource consumption alongside the corresponding plan limits (from the caller's InternalToken.Quotas). Paused services are included in services_used but excluded from memory / instances_by_type / proxy_ports totals, matching enforcement semantics.
type QuotaUsage struct {
	AppsUsed *int64 `json:"apps_used,omitempty"`
	AppsLimit *int64 `json:"apps_limit,omitempty"`
	ServicesUsed *int64 `json:"services_used,omitempty"`
	ServicesLimit *int64 `json:"services_limit,omitempty"`
	MemoryMbUsed *int64 `json:"memory_mb_used,omitempty"`
	MemoryMbLimit *int64 `json:"memory_mb_limit,omitempty"`
	CustomDomainsUsed *int64 `json:"custom_domains_used,omitempty"`
	CustomDomainsLimit *int64 `json:"custom_domains_limit,omitempty"`
	KoyebLbDomainsUsed *int64 `json:"koyeb_lb_domains_used,omitempty"`
	KoyebLbDomainsLimit *int64 `json:"koyeb_lb_domains_limit,omitempty"`
	ProxyPortsUsed *int64 `json:"proxy_ports_used,omitempty"`
	ProxyPortsLimit *int64 `json:"proxy_ports_limit,omitempty"`
	InstancesByType []InstanceTypeUsage `json:"instances_by_type,omitempty"`
	PersistentVolumesByRegion []PersistentVolumeRegionUsage `json:"persistent_volumes_by_region,omitempty"`
	InstanceSnapshotsByType []InstanceSnapshotTypeUsage `json:"instance_snapshots_by_type,omitempty"`
}

// NewQuotaUsage instantiates a new QuotaUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaUsage() *QuotaUsage {
	this := QuotaUsage{}
	return &this
}

// NewQuotaUsageWithDefaults instantiates a new QuotaUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaUsageWithDefaults() *QuotaUsage {
	this := QuotaUsage{}
	return &this
}

// GetAppsUsed returns the AppsUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetAppsUsed() int64 {
	if o == nil || isNil(o.AppsUsed) {
		var ret int64
		return ret
	}
	return *o.AppsUsed
}

// GetAppsUsedOk returns a tuple with the AppsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetAppsUsedOk() (*int64, bool) {
	if o == nil || isNil(o.AppsUsed) {
    return nil, false
	}
	return o.AppsUsed, true
}

// HasAppsUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasAppsUsed() bool {
	if o != nil && !isNil(o.AppsUsed) {
		return true
	}

	return false
}

// SetAppsUsed gets a reference to the given int64 and assigns it to the AppsUsed field.
func (o *QuotaUsage) SetAppsUsed(v int64) {
	o.AppsUsed = &v
}

// GetAppsLimit returns the AppsLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetAppsLimit() int64 {
	if o == nil || isNil(o.AppsLimit) {
		var ret int64
		return ret
	}
	return *o.AppsLimit
}

// GetAppsLimitOk returns a tuple with the AppsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetAppsLimitOk() (*int64, bool) {
	if o == nil || isNil(o.AppsLimit) {
    return nil, false
	}
	return o.AppsLimit, true
}

// HasAppsLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasAppsLimit() bool {
	if o != nil && !isNil(o.AppsLimit) {
		return true
	}

	return false
}

// SetAppsLimit gets a reference to the given int64 and assigns it to the AppsLimit field.
func (o *QuotaUsage) SetAppsLimit(v int64) {
	o.AppsLimit = &v
}

// GetServicesUsed returns the ServicesUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetServicesUsed() int64 {
	if o == nil || isNil(o.ServicesUsed) {
		var ret int64
		return ret
	}
	return *o.ServicesUsed
}

// GetServicesUsedOk returns a tuple with the ServicesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetServicesUsedOk() (*int64, bool) {
	if o == nil || isNil(o.ServicesUsed) {
    return nil, false
	}
	return o.ServicesUsed, true
}

// HasServicesUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasServicesUsed() bool {
	if o != nil && !isNil(o.ServicesUsed) {
		return true
	}

	return false
}

// SetServicesUsed gets a reference to the given int64 and assigns it to the ServicesUsed field.
func (o *QuotaUsage) SetServicesUsed(v int64) {
	o.ServicesUsed = &v
}

// GetServicesLimit returns the ServicesLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetServicesLimit() int64 {
	if o == nil || isNil(o.ServicesLimit) {
		var ret int64
		return ret
	}
	return *o.ServicesLimit
}

// GetServicesLimitOk returns a tuple with the ServicesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetServicesLimitOk() (*int64, bool) {
	if o == nil || isNil(o.ServicesLimit) {
    return nil, false
	}
	return o.ServicesLimit, true
}

// HasServicesLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasServicesLimit() bool {
	if o != nil && !isNil(o.ServicesLimit) {
		return true
	}

	return false
}

// SetServicesLimit gets a reference to the given int64 and assigns it to the ServicesLimit field.
func (o *QuotaUsage) SetServicesLimit(v int64) {
	o.ServicesLimit = &v
}

// GetMemoryMbUsed returns the MemoryMbUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetMemoryMbUsed() int64 {
	if o == nil || isNil(o.MemoryMbUsed) {
		var ret int64
		return ret
	}
	return *o.MemoryMbUsed
}

// GetMemoryMbUsedOk returns a tuple with the MemoryMbUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetMemoryMbUsedOk() (*int64, bool) {
	if o == nil || isNil(o.MemoryMbUsed) {
    return nil, false
	}
	return o.MemoryMbUsed, true
}

// HasMemoryMbUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasMemoryMbUsed() bool {
	if o != nil && !isNil(o.MemoryMbUsed) {
		return true
	}

	return false
}

// SetMemoryMbUsed gets a reference to the given int64 and assigns it to the MemoryMbUsed field.
func (o *QuotaUsage) SetMemoryMbUsed(v int64) {
	o.MemoryMbUsed = &v
}

// GetMemoryMbLimit returns the MemoryMbLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetMemoryMbLimit() int64 {
	if o == nil || isNil(o.MemoryMbLimit) {
		var ret int64
		return ret
	}
	return *o.MemoryMbLimit
}

// GetMemoryMbLimitOk returns a tuple with the MemoryMbLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetMemoryMbLimitOk() (*int64, bool) {
	if o == nil || isNil(o.MemoryMbLimit) {
    return nil, false
	}
	return o.MemoryMbLimit, true
}

// HasMemoryMbLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasMemoryMbLimit() bool {
	if o != nil && !isNil(o.MemoryMbLimit) {
		return true
	}

	return false
}

// SetMemoryMbLimit gets a reference to the given int64 and assigns it to the MemoryMbLimit field.
func (o *QuotaUsage) SetMemoryMbLimit(v int64) {
	o.MemoryMbLimit = &v
}

// GetCustomDomainsUsed returns the CustomDomainsUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetCustomDomainsUsed() int64 {
	if o == nil || isNil(o.CustomDomainsUsed) {
		var ret int64
		return ret
	}
	return *o.CustomDomainsUsed
}

// GetCustomDomainsUsedOk returns a tuple with the CustomDomainsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetCustomDomainsUsedOk() (*int64, bool) {
	if o == nil || isNil(o.CustomDomainsUsed) {
    return nil, false
	}
	return o.CustomDomainsUsed, true
}

// HasCustomDomainsUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasCustomDomainsUsed() bool {
	if o != nil && !isNil(o.CustomDomainsUsed) {
		return true
	}

	return false
}

// SetCustomDomainsUsed gets a reference to the given int64 and assigns it to the CustomDomainsUsed field.
func (o *QuotaUsage) SetCustomDomainsUsed(v int64) {
	o.CustomDomainsUsed = &v
}

// GetCustomDomainsLimit returns the CustomDomainsLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetCustomDomainsLimit() int64 {
	if o == nil || isNil(o.CustomDomainsLimit) {
		var ret int64
		return ret
	}
	return *o.CustomDomainsLimit
}

// GetCustomDomainsLimitOk returns a tuple with the CustomDomainsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetCustomDomainsLimitOk() (*int64, bool) {
	if o == nil || isNil(o.CustomDomainsLimit) {
    return nil, false
	}
	return o.CustomDomainsLimit, true
}

// HasCustomDomainsLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasCustomDomainsLimit() bool {
	if o != nil && !isNil(o.CustomDomainsLimit) {
		return true
	}

	return false
}

// SetCustomDomainsLimit gets a reference to the given int64 and assigns it to the CustomDomainsLimit field.
func (o *QuotaUsage) SetCustomDomainsLimit(v int64) {
	o.CustomDomainsLimit = &v
}

// GetKoyebLbDomainsUsed returns the KoyebLbDomainsUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetKoyebLbDomainsUsed() int64 {
	if o == nil || isNil(o.KoyebLbDomainsUsed) {
		var ret int64
		return ret
	}
	return *o.KoyebLbDomainsUsed
}

// GetKoyebLbDomainsUsedOk returns a tuple with the KoyebLbDomainsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetKoyebLbDomainsUsedOk() (*int64, bool) {
	if o == nil || isNil(o.KoyebLbDomainsUsed) {
    return nil, false
	}
	return o.KoyebLbDomainsUsed, true
}

// HasKoyebLbDomainsUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasKoyebLbDomainsUsed() bool {
	if o != nil && !isNil(o.KoyebLbDomainsUsed) {
		return true
	}

	return false
}

// SetKoyebLbDomainsUsed gets a reference to the given int64 and assigns it to the KoyebLbDomainsUsed field.
func (o *QuotaUsage) SetKoyebLbDomainsUsed(v int64) {
	o.KoyebLbDomainsUsed = &v
}

// GetKoyebLbDomainsLimit returns the KoyebLbDomainsLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetKoyebLbDomainsLimit() int64 {
	if o == nil || isNil(o.KoyebLbDomainsLimit) {
		var ret int64
		return ret
	}
	return *o.KoyebLbDomainsLimit
}

// GetKoyebLbDomainsLimitOk returns a tuple with the KoyebLbDomainsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetKoyebLbDomainsLimitOk() (*int64, bool) {
	if o == nil || isNil(o.KoyebLbDomainsLimit) {
    return nil, false
	}
	return o.KoyebLbDomainsLimit, true
}

// HasKoyebLbDomainsLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasKoyebLbDomainsLimit() bool {
	if o != nil && !isNil(o.KoyebLbDomainsLimit) {
		return true
	}

	return false
}

// SetKoyebLbDomainsLimit gets a reference to the given int64 and assigns it to the KoyebLbDomainsLimit field.
func (o *QuotaUsage) SetKoyebLbDomainsLimit(v int64) {
	o.KoyebLbDomainsLimit = &v
}

// GetProxyPortsUsed returns the ProxyPortsUsed field value if set, zero value otherwise.
func (o *QuotaUsage) GetProxyPortsUsed() int64 {
	if o == nil || isNil(o.ProxyPortsUsed) {
		var ret int64
		return ret
	}
	return *o.ProxyPortsUsed
}

// GetProxyPortsUsedOk returns a tuple with the ProxyPortsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetProxyPortsUsedOk() (*int64, bool) {
	if o == nil || isNil(o.ProxyPortsUsed) {
    return nil, false
	}
	return o.ProxyPortsUsed, true
}

// HasProxyPortsUsed returns a boolean if a field has been set.
func (o *QuotaUsage) HasProxyPortsUsed() bool {
	if o != nil && !isNil(o.ProxyPortsUsed) {
		return true
	}

	return false
}

// SetProxyPortsUsed gets a reference to the given int64 and assigns it to the ProxyPortsUsed field.
func (o *QuotaUsage) SetProxyPortsUsed(v int64) {
	o.ProxyPortsUsed = &v
}

// GetProxyPortsLimit returns the ProxyPortsLimit field value if set, zero value otherwise.
func (o *QuotaUsage) GetProxyPortsLimit() int64 {
	if o == nil || isNil(o.ProxyPortsLimit) {
		var ret int64
		return ret
	}
	return *o.ProxyPortsLimit
}

// GetProxyPortsLimitOk returns a tuple with the ProxyPortsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetProxyPortsLimitOk() (*int64, bool) {
	if o == nil || isNil(o.ProxyPortsLimit) {
    return nil, false
	}
	return o.ProxyPortsLimit, true
}

// HasProxyPortsLimit returns a boolean if a field has been set.
func (o *QuotaUsage) HasProxyPortsLimit() bool {
	if o != nil && !isNil(o.ProxyPortsLimit) {
		return true
	}

	return false
}

// SetProxyPortsLimit gets a reference to the given int64 and assigns it to the ProxyPortsLimit field.
func (o *QuotaUsage) SetProxyPortsLimit(v int64) {
	o.ProxyPortsLimit = &v
}

// GetInstancesByType returns the InstancesByType field value if set, zero value otherwise.
func (o *QuotaUsage) GetInstancesByType() []InstanceTypeUsage {
	if o == nil || isNil(o.InstancesByType) {
		var ret []InstanceTypeUsage
		return ret
	}
	return o.InstancesByType
}

// GetInstancesByTypeOk returns a tuple with the InstancesByType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetInstancesByTypeOk() ([]InstanceTypeUsage, bool) {
	if o == nil || isNil(o.InstancesByType) {
    return nil, false
	}
	return o.InstancesByType, true
}

// HasInstancesByType returns a boolean if a field has been set.
func (o *QuotaUsage) HasInstancesByType() bool {
	if o != nil && !isNil(o.InstancesByType) {
		return true
	}

	return false
}

// SetInstancesByType gets a reference to the given []InstanceTypeUsage and assigns it to the InstancesByType field.
func (o *QuotaUsage) SetInstancesByType(v []InstanceTypeUsage) {
	o.InstancesByType = v
}

// GetPersistentVolumesByRegion returns the PersistentVolumesByRegion field value if set, zero value otherwise.
func (o *QuotaUsage) GetPersistentVolumesByRegion() []PersistentVolumeRegionUsage {
	if o == nil || isNil(o.PersistentVolumesByRegion) {
		var ret []PersistentVolumeRegionUsage
		return ret
	}
	return o.PersistentVolumesByRegion
}

// GetPersistentVolumesByRegionOk returns a tuple with the PersistentVolumesByRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetPersistentVolumesByRegionOk() ([]PersistentVolumeRegionUsage, bool) {
	if o == nil || isNil(o.PersistentVolumesByRegion) {
    return nil, false
	}
	return o.PersistentVolumesByRegion, true
}

// HasPersistentVolumesByRegion returns a boolean if a field has been set.
func (o *QuotaUsage) HasPersistentVolumesByRegion() bool {
	if o != nil && !isNil(o.PersistentVolumesByRegion) {
		return true
	}

	return false
}

// SetPersistentVolumesByRegion gets a reference to the given []PersistentVolumeRegionUsage and assigns it to the PersistentVolumesByRegion field.
func (o *QuotaUsage) SetPersistentVolumesByRegion(v []PersistentVolumeRegionUsage) {
	o.PersistentVolumesByRegion = v
}

// GetInstanceSnapshotsByType returns the InstanceSnapshotsByType field value if set, zero value otherwise.
func (o *QuotaUsage) GetInstanceSnapshotsByType() []InstanceSnapshotTypeUsage {
	if o == nil || isNil(o.InstanceSnapshotsByType) {
		var ret []InstanceSnapshotTypeUsage
		return ret
	}
	return o.InstanceSnapshotsByType
}

// GetInstanceSnapshotsByTypeOk returns a tuple with the InstanceSnapshotsByType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaUsage) GetInstanceSnapshotsByTypeOk() ([]InstanceSnapshotTypeUsage, bool) {
	if o == nil || isNil(o.InstanceSnapshotsByType) {
    return nil, false
	}
	return o.InstanceSnapshotsByType, true
}

// HasInstanceSnapshotsByType returns a boolean if a field has been set.
func (o *QuotaUsage) HasInstanceSnapshotsByType() bool {
	if o != nil && !isNil(o.InstanceSnapshotsByType) {
		return true
	}

	return false
}

// SetInstanceSnapshotsByType gets a reference to the given []InstanceSnapshotTypeUsage and assigns it to the InstanceSnapshotsByType field.
func (o *QuotaUsage) SetInstanceSnapshotsByType(v []InstanceSnapshotTypeUsage) {
	o.InstanceSnapshotsByType = v
}

func (o QuotaUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppsUsed) {
		toSerialize["apps_used"] = o.AppsUsed
	}
	if !isNil(o.AppsLimit) {
		toSerialize["apps_limit"] = o.AppsLimit
	}
	if !isNil(o.ServicesUsed) {
		toSerialize["services_used"] = o.ServicesUsed
	}
	if !isNil(o.ServicesLimit) {
		toSerialize["services_limit"] = o.ServicesLimit
	}
	if !isNil(o.MemoryMbUsed) {
		toSerialize["memory_mb_used"] = o.MemoryMbUsed
	}
	if !isNil(o.MemoryMbLimit) {
		toSerialize["memory_mb_limit"] = o.MemoryMbLimit
	}
	if !isNil(o.CustomDomainsUsed) {
		toSerialize["custom_domains_used"] = o.CustomDomainsUsed
	}
	if !isNil(o.CustomDomainsLimit) {
		toSerialize["custom_domains_limit"] = o.CustomDomainsLimit
	}
	if !isNil(o.KoyebLbDomainsUsed) {
		toSerialize["koyeb_lb_domains_used"] = o.KoyebLbDomainsUsed
	}
	if !isNil(o.KoyebLbDomainsLimit) {
		toSerialize["koyeb_lb_domains_limit"] = o.KoyebLbDomainsLimit
	}
	if !isNil(o.ProxyPortsUsed) {
		toSerialize["proxy_ports_used"] = o.ProxyPortsUsed
	}
	if !isNil(o.ProxyPortsLimit) {
		toSerialize["proxy_ports_limit"] = o.ProxyPortsLimit
	}
	if !isNil(o.InstancesByType) {
		toSerialize["instances_by_type"] = o.InstancesByType
	}
	if !isNil(o.PersistentVolumesByRegion) {
		toSerialize["persistent_volumes_by_region"] = o.PersistentVolumesByRegion
	}
	if !isNil(o.InstanceSnapshotsByType) {
		toSerialize["instance_snapshots_by_type"] = o.InstanceSnapshotsByType
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaUsage struct {
	value *QuotaUsage
	isSet bool
}

func (v NullableQuotaUsage) Get() *QuotaUsage {
	return v.value
}

func (v *NullableQuotaUsage) Set(val *QuotaUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaUsage(val *QuotaUsage) *NullableQuotaUsage {
	return &NullableQuotaUsage{value: val, isSet: true}
}

func (v NullableQuotaUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


