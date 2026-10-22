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

// CreateSecret struct for CreateSecret
type CreateSecret struct {
	Name *string `json:"name,omitempty"`
	Type *SecretType `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
	DockerHubRegistry *DockerHubRegistryConfiguration `json:"docker_hub_registry,omitempty"`
	PrivateRegistry *PrivateRegistryConfiguration `json:"private_registry,omitempty"`
	DigitalOceanRegistry *DigitalOceanRegistryConfiguration `json:"digital_ocean_registry,omitempty"`
	GithubRegistry *GitHubRegistryConfiguration `json:"github_registry,omitempty"`
	GitlabRegistry *GitLabRegistryConfiguration `json:"gitlab_registry,omitempty"`
	GcpContainerRegistry *GCPContainerRegistryConfiguration `json:"gcp_container_registry,omitempty"`
	AzureContainerRegistry *AzureContainerRegistryConfiguration `json:"azure_container_registry,omitempty"`
}

// NewCreateSecret instantiates a new CreateSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecret() *CreateSecret {
	this := CreateSecret{}
	var type_ SecretType = SECRETTYPE_SIMPLE
	this.Type = &type_
	return &this
}

// NewCreateSecretWithDefaults instantiates a new CreateSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecretWithDefaults() *CreateSecret {
	this := CreateSecret{}
	var type_ SecretType = SECRETTYPE_SIMPLE
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSecret) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSecret) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSecret) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateSecret) GetType() SecretType {
	if o == nil || isNil(o.Type) {
		var ret SecretType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetTypeOk() (*SecretType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateSecret) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SecretType and assigns it to the Type field.
func (o *CreateSecret) SetType(v SecretType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateSecret) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateSecret) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateSecret) SetValue(v string) {
	o.Value = &v
}

// GetDockerHubRegistry returns the DockerHubRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetDockerHubRegistry() DockerHubRegistryConfiguration {
	if o == nil || isNil(o.DockerHubRegistry) {
		var ret DockerHubRegistryConfiguration
		return ret
	}
	return *o.DockerHubRegistry
}

// GetDockerHubRegistryOk returns a tuple with the DockerHubRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetDockerHubRegistryOk() (*DockerHubRegistryConfiguration, bool) {
	if o == nil || isNil(o.DockerHubRegistry) {
    return nil, false
	}
	return o.DockerHubRegistry, true
}

// HasDockerHubRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasDockerHubRegistry() bool {
	if o != nil && !isNil(o.DockerHubRegistry) {
		return true
	}

	return false
}

// SetDockerHubRegistry gets a reference to the given DockerHubRegistryConfiguration and assigns it to the DockerHubRegistry field.
func (o *CreateSecret) SetDockerHubRegistry(v DockerHubRegistryConfiguration) {
	o.DockerHubRegistry = &v
}

// GetPrivateRegistry returns the PrivateRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetPrivateRegistry() PrivateRegistryConfiguration {
	if o == nil || isNil(o.PrivateRegistry) {
		var ret PrivateRegistryConfiguration
		return ret
	}
	return *o.PrivateRegistry
}

// GetPrivateRegistryOk returns a tuple with the PrivateRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetPrivateRegistryOk() (*PrivateRegistryConfiguration, bool) {
	if o == nil || isNil(o.PrivateRegistry) {
    return nil, false
	}
	return o.PrivateRegistry, true
}

// HasPrivateRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasPrivateRegistry() bool {
	if o != nil && !isNil(o.PrivateRegistry) {
		return true
	}

	return false
}

// SetPrivateRegistry gets a reference to the given PrivateRegistryConfiguration and assigns it to the PrivateRegistry field.
func (o *CreateSecret) SetPrivateRegistry(v PrivateRegistryConfiguration) {
	o.PrivateRegistry = &v
}

// GetDigitalOceanRegistry returns the DigitalOceanRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetDigitalOceanRegistry() DigitalOceanRegistryConfiguration {
	if o == nil || isNil(o.DigitalOceanRegistry) {
		var ret DigitalOceanRegistryConfiguration
		return ret
	}
	return *o.DigitalOceanRegistry
}

// GetDigitalOceanRegistryOk returns a tuple with the DigitalOceanRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetDigitalOceanRegistryOk() (*DigitalOceanRegistryConfiguration, bool) {
	if o == nil || isNil(o.DigitalOceanRegistry) {
    return nil, false
	}
	return o.DigitalOceanRegistry, true
}

// HasDigitalOceanRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasDigitalOceanRegistry() bool {
	if o != nil && !isNil(o.DigitalOceanRegistry) {
		return true
	}

	return false
}

// SetDigitalOceanRegistry gets a reference to the given DigitalOceanRegistryConfiguration and assigns it to the DigitalOceanRegistry field.
func (o *CreateSecret) SetDigitalOceanRegistry(v DigitalOceanRegistryConfiguration) {
	o.DigitalOceanRegistry = &v
}

// GetGithubRegistry returns the GithubRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetGithubRegistry() GitHubRegistryConfiguration {
	if o == nil || isNil(o.GithubRegistry) {
		var ret GitHubRegistryConfiguration
		return ret
	}
	return *o.GithubRegistry
}

// GetGithubRegistryOk returns a tuple with the GithubRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetGithubRegistryOk() (*GitHubRegistryConfiguration, bool) {
	if o == nil || isNil(o.GithubRegistry) {
    return nil, false
	}
	return o.GithubRegistry, true
}

// HasGithubRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasGithubRegistry() bool {
	if o != nil && !isNil(o.GithubRegistry) {
		return true
	}

	return false
}

// SetGithubRegistry gets a reference to the given GitHubRegistryConfiguration and assigns it to the GithubRegistry field.
func (o *CreateSecret) SetGithubRegistry(v GitHubRegistryConfiguration) {
	o.GithubRegistry = &v
}

// GetGitlabRegistry returns the GitlabRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetGitlabRegistry() GitLabRegistryConfiguration {
	if o == nil || isNil(o.GitlabRegistry) {
		var ret GitLabRegistryConfiguration
		return ret
	}
	return *o.GitlabRegistry
}

// GetGitlabRegistryOk returns a tuple with the GitlabRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetGitlabRegistryOk() (*GitLabRegistryConfiguration, bool) {
	if o == nil || isNil(o.GitlabRegistry) {
    return nil, false
	}
	return o.GitlabRegistry, true
}

// HasGitlabRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasGitlabRegistry() bool {
	if o != nil && !isNil(o.GitlabRegistry) {
		return true
	}

	return false
}

// SetGitlabRegistry gets a reference to the given GitLabRegistryConfiguration and assigns it to the GitlabRegistry field.
func (o *CreateSecret) SetGitlabRegistry(v GitLabRegistryConfiguration) {
	o.GitlabRegistry = &v
}

// GetGcpContainerRegistry returns the GcpContainerRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetGcpContainerRegistry() GCPContainerRegistryConfiguration {
	if o == nil || isNil(o.GcpContainerRegistry) {
		var ret GCPContainerRegistryConfiguration
		return ret
	}
	return *o.GcpContainerRegistry
}

// GetGcpContainerRegistryOk returns a tuple with the GcpContainerRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetGcpContainerRegistryOk() (*GCPContainerRegistryConfiguration, bool) {
	if o == nil || isNil(o.GcpContainerRegistry) {
    return nil, false
	}
	return o.GcpContainerRegistry, true
}

// HasGcpContainerRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasGcpContainerRegistry() bool {
	if o != nil && !isNil(o.GcpContainerRegistry) {
		return true
	}

	return false
}

// SetGcpContainerRegistry gets a reference to the given GCPContainerRegistryConfiguration and assigns it to the GcpContainerRegistry field.
func (o *CreateSecret) SetGcpContainerRegistry(v GCPContainerRegistryConfiguration) {
	o.GcpContainerRegistry = &v
}

// GetAzureContainerRegistry returns the AzureContainerRegistry field value if set, zero value otherwise.
func (o *CreateSecret) GetAzureContainerRegistry() AzureContainerRegistryConfiguration {
	if o == nil || isNil(o.AzureContainerRegistry) {
		var ret AzureContainerRegistryConfiguration
		return ret
	}
	return *o.AzureContainerRegistry
}

// GetAzureContainerRegistryOk returns a tuple with the AzureContainerRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecret) GetAzureContainerRegistryOk() (*AzureContainerRegistryConfiguration, bool) {
	if o == nil || isNil(o.AzureContainerRegistry) {
    return nil, false
	}
	return o.AzureContainerRegistry, true
}

// HasAzureContainerRegistry returns a boolean if a field has been set.
func (o *CreateSecret) HasAzureContainerRegistry() bool {
	if o != nil && !isNil(o.AzureContainerRegistry) {
		return true
	}

	return false
}

// SetAzureContainerRegistry gets a reference to the given AzureContainerRegistryConfiguration and assigns it to the AzureContainerRegistry field.
func (o *CreateSecret) SetAzureContainerRegistry(v AzureContainerRegistryConfiguration) {
	o.AzureContainerRegistry = &v
}

func (o CreateSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.DockerHubRegistry) {
		toSerialize["docker_hub_registry"] = o.DockerHubRegistry
	}
	if !isNil(o.PrivateRegistry) {
		toSerialize["private_registry"] = o.PrivateRegistry
	}
	if !isNil(o.DigitalOceanRegistry) {
		toSerialize["digital_ocean_registry"] = o.DigitalOceanRegistry
	}
	if !isNil(o.GithubRegistry) {
		toSerialize["github_registry"] = o.GithubRegistry
	}
	if !isNil(o.GitlabRegistry) {
		toSerialize["gitlab_registry"] = o.GitlabRegistry
	}
	if !isNil(o.GcpContainerRegistry) {
		toSerialize["gcp_container_registry"] = o.GcpContainerRegistry
	}
	if !isNil(o.AzureContainerRegistry) {
		toSerialize["azure_container_registry"] = o.AzureContainerRegistry
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSecret struct {
	value *CreateSecret
	isSet bool
}

func (v NullableCreateSecret) Get() *CreateSecret {
	return v.value
}

func (v *NullableCreateSecret) Set(val *CreateSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecret(val *CreateSecret) *NullableCreateSecret {
	return &NullableCreateSecret{value: val, isSet: true}
}

func (v NullableCreateSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


