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

// NeonPostgresDatabase struct for NeonPostgresDatabase
type NeonPostgresDatabase struct {
	PgVersion *int64 `json:"pg_version,omitempty"`
	Region *string `json:"region,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	Roles []NeonPostgresDatabaseNeonRole `json:"roles,omitempty"`
	Databases []NeonPostgresDatabaseNeonDatabase `json:"databases,omitempty"`
}

// NewNeonPostgresDatabase instantiates a new NeonPostgresDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeonPostgresDatabase() *NeonPostgresDatabase {
	this := NeonPostgresDatabase{}
	return &this
}

// NewNeonPostgresDatabaseWithDefaults instantiates a new NeonPostgresDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeonPostgresDatabaseWithDefaults() *NeonPostgresDatabase {
	this := NeonPostgresDatabase{}
	return &this
}

// GetPgVersion returns the PgVersion field value if set, zero value otherwise.
func (o *NeonPostgresDatabase) GetPgVersion() int64 {
	if o == nil || isNil(o.PgVersion) {
		var ret int64
		return ret
	}
	return *o.PgVersion
}

// GetPgVersionOk returns a tuple with the PgVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeonPostgresDatabase) GetPgVersionOk() (*int64, bool) {
	if o == nil || isNil(o.PgVersion) {
    return nil, false
	}
	return o.PgVersion, true
}

// HasPgVersion returns a boolean if a field has been set.
func (o *NeonPostgresDatabase) HasPgVersion() bool {
	if o != nil && !isNil(o.PgVersion) {
		return true
	}

	return false
}

// SetPgVersion gets a reference to the given int64 and assigns it to the PgVersion field.
func (o *NeonPostgresDatabase) SetPgVersion(v int64) {
	o.PgVersion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NeonPostgresDatabase) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeonPostgresDatabase) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NeonPostgresDatabase) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NeonPostgresDatabase) SetRegion(v string) {
	o.Region = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *NeonPostgresDatabase) GetInstanceType() string {
	if o == nil || isNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeonPostgresDatabase) GetInstanceTypeOk() (*string, bool) {
	if o == nil || isNil(o.InstanceType) {
    return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *NeonPostgresDatabase) HasInstanceType() bool {
	if o != nil && !isNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *NeonPostgresDatabase) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *NeonPostgresDatabase) GetRoles() []NeonPostgresDatabaseNeonRole {
	if o == nil || isNil(o.Roles) {
		var ret []NeonPostgresDatabaseNeonRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeonPostgresDatabase) GetRolesOk() ([]NeonPostgresDatabaseNeonRole, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *NeonPostgresDatabase) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []NeonPostgresDatabaseNeonRole and assigns it to the Roles field.
func (o *NeonPostgresDatabase) SetRoles(v []NeonPostgresDatabaseNeonRole) {
	o.Roles = v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *NeonPostgresDatabase) GetDatabases() []NeonPostgresDatabaseNeonDatabase {
	if o == nil || isNil(o.Databases) {
		var ret []NeonPostgresDatabaseNeonDatabase
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeonPostgresDatabase) GetDatabasesOk() ([]NeonPostgresDatabaseNeonDatabase, bool) {
	if o == nil || isNil(o.Databases) {
    return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *NeonPostgresDatabase) HasDatabases() bool {
	if o != nil && !isNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []NeonPostgresDatabaseNeonDatabase and assigns it to the Databases field.
func (o *NeonPostgresDatabase) SetDatabases(v []NeonPostgresDatabaseNeonDatabase) {
	o.Databases = v
}

func (o NeonPostgresDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PgVersion) {
		toSerialize["pg_version"] = o.PgVersion
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.Databases) {
		toSerialize["databases"] = o.Databases
	}
	return json.Marshal(toSerialize)
}

type NullableNeonPostgresDatabase struct {
	value *NeonPostgresDatabase
	isSet bool
}

func (v NullableNeonPostgresDatabase) Get() *NeonPostgresDatabase {
	return v.value
}

func (v *NullableNeonPostgresDatabase) Set(val *NeonPostgresDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableNeonPostgresDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableNeonPostgresDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeonPostgresDatabase(val *NeonPostgresDatabase) *NullableNeonPostgresDatabase {
	return &NullableNeonPostgresDatabase{value: val, isSet: true}
}

func (v NullableNeonPostgresDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeonPostgresDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


