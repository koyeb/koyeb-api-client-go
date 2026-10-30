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

// RegionalDeploymentDefinition struct for RegionalDeploymentDefinition
type RegionalDeploymentDefinition struct {
	Name *string `json:"name,omitempty"`
	Type *RegionalDeploymentDefinitionType `json:"type,omitempty"`
	Strategy *DeploymentStrategy `json:"strategy,omitempty"`
	Routes []Route `json:"routes,omitempty"`
	Ports []Port `json:"ports,omitempty"`
	Env []Env `json:"env,omitempty"`
	Region *string `json:"region,omitempty"`
	Scaling *Scaling `json:"scaling,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	DeploymentGroup *string `json:"deployment_group,omitempty"`
	HealthChecks []DeploymentHealthCheck `json:"health_checks,omitempty"`
	Volumes []RegionalDeploymentVolume `json:"volumes,omitempty"`
	ConfigFiles []ConfigFile `json:"config_files,omitempty"`
	SkipCache *bool `json:"skip_cache,omitempty"`
	Mesh *RegionalDeploymentMesh `json:"mesh,omitempty"`
	NetworkPolicy *NetworkPolicy `json:"network_policy,omitempty"`
	Docker *DockerSource `json:"docker,omitempty"`
	Git *GitSource `json:"git,omitempty"`
	Archive *ArchiveSource `json:"archive,omitempty"`
}

// NewRegionalDeploymentDefinition instantiates a new RegionalDeploymentDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeploymentDefinition() *RegionalDeploymentDefinition {
	this := RegionalDeploymentDefinition{}
	var type_ RegionalDeploymentDefinitionType = REGIONALDEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	var mesh RegionalDeploymentMesh = REGIONALDEPLOYMENTMESH_AUTO
	this.Mesh = &mesh
	return &this
}

// NewRegionalDeploymentDefinitionWithDefaults instantiates a new RegionalDeploymentDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentDefinitionWithDefaults() *RegionalDeploymentDefinition {
	this := RegionalDeploymentDefinition{}
	var type_ RegionalDeploymentDefinitionType = REGIONALDEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	var mesh RegionalDeploymentMesh = REGIONALDEPLOYMENTMESH_AUTO
	this.Mesh = &mesh
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegionalDeploymentDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetType() RegionalDeploymentDefinitionType {
	if o == nil || isNil(o.Type) {
		var ret RegionalDeploymentDefinitionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetTypeOk() (*RegionalDeploymentDefinitionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RegionalDeploymentDefinitionType and assigns it to the Type field.
func (o *RegionalDeploymentDefinition) SetType(v RegionalDeploymentDefinitionType) {
	o.Type = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetStrategy() DeploymentStrategy {
	if o == nil || isNil(o.Strategy) {
		var ret DeploymentStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetStrategyOk() (*DeploymentStrategy, bool) {
	if o == nil || isNil(o.Strategy) {
    return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasStrategy() bool {
	if o != nil && !isNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeploymentStrategy and assigns it to the Strategy field.
func (o *RegionalDeploymentDefinition) SetStrategy(v DeploymentStrategy) {
	o.Strategy = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetRoutes() []Route {
	if o == nil || isNil(o.Routes) {
		var ret []Route
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetRoutesOk() ([]Route, bool) {
	if o == nil || isNil(o.Routes) {
    return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasRoutes() bool {
	if o != nil && !isNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *RegionalDeploymentDefinition) SetRoutes(v []Route) {
	o.Routes = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetPorts() []Port {
	if o == nil || isNil(o.Ports) {
		var ret []Port
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetPortsOk() ([]Port, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []Port and assigns it to the Ports field.
func (o *RegionalDeploymentDefinition) SetPorts(v []Port) {
	o.Ports = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetEnv() []Env {
	if o == nil || isNil(o.Env) {
		var ret []Env
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetEnvOk() ([]Env, bool) {
	if o == nil || isNil(o.Env) {
    return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the Env field.
func (o *RegionalDeploymentDefinition) SetEnv(v []Env) {
	o.Env = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RegionalDeploymentDefinition) SetRegion(v string) {
	o.Region = &v
}

// GetScaling returns the Scaling field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetScaling() Scaling {
	if o == nil || isNil(o.Scaling) {
		var ret Scaling
		return ret
	}
	return *o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetScalingOk() (*Scaling, bool) {
	if o == nil || isNil(o.Scaling) {
    return nil, false
	}
	return o.Scaling, true
}

// HasScaling returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasScaling() bool {
	if o != nil && !isNil(o.Scaling) {
		return true
	}

	return false
}

// SetScaling gets a reference to the given Scaling and assigns it to the Scaling field.
func (o *RegionalDeploymentDefinition) SetScaling(v Scaling) {
	o.Scaling = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetInstanceType() string {
	if o == nil || isNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetInstanceTypeOk() (*string, bool) {
	if o == nil || isNil(o.InstanceType) {
    return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasInstanceType() bool {
	if o != nil && !isNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *RegionalDeploymentDefinition) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetDeploymentGroup returns the DeploymentGroup field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetDeploymentGroup() string {
	if o == nil || isNil(o.DeploymentGroup) {
		var ret string
		return ret
	}
	return *o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetDeploymentGroupOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentGroup) {
    return nil, false
	}
	return o.DeploymentGroup, true
}

// HasDeploymentGroup returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasDeploymentGroup() bool {
	if o != nil && !isNil(o.DeploymentGroup) {
		return true
	}

	return false
}

// SetDeploymentGroup gets a reference to the given string and assigns it to the DeploymentGroup field.
func (o *RegionalDeploymentDefinition) SetDeploymentGroup(v string) {
	o.DeploymentGroup = &v
}

// GetHealthChecks returns the HealthChecks field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetHealthChecks() []DeploymentHealthCheck {
	if o == nil || isNil(o.HealthChecks) {
		var ret []DeploymentHealthCheck
		return ret
	}
	return o.HealthChecks
}

// GetHealthChecksOk returns a tuple with the HealthChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetHealthChecksOk() ([]DeploymentHealthCheck, bool) {
	if o == nil || isNil(o.HealthChecks) {
    return nil, false
	}
	return o.HealthChecks, true
}

// HasHealthChecks returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasHealthChecks() bool {
	if o != nil && !isNil(o.HealthChecks) {
		return true
	}

	return false
}

// SetHealthChecks gets a reference to the given []DeploymentHealthCheck and assigns it to the HealthChecks field.
func (o *RegionalDeploymentDefinition) SetHealthChecks(v []DeploymentHealthCheck) {
	o.HealthChecks = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetVolumes() []RegionalDeploymentVolume {
	if o == nil || isNil(o.Volumes) {
		var ret []RegionalDeploymentVolume
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetVolumesOk() ([]RegionalDeploymentVolume, bool) {
	if o == nil || isNil(o.Volumes) {
    return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasVolumes() bool {
	if o != nil && !isNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []RegionalDeploymentVolume and assigns it to the Volumes field.
func (o *RegionalDeploymentDefinition) SetVolumes(v []RegionalDeploymentVolume) {
	o.Volumes = v
}

// GetConfigFiles returns the ConfigFiles field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetConfigFiles() []ConfigFile {
	if o == nil || isNil(o.ConfigFiles) {
		var ret []ConfigFile
		return ret
	}
	return o.ConfigFiles
}

// GetConfigFilesOk returns a tuple with the ConfigFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetConfigFilesOk() ([]ConfigFile, bool) {
	if o == nil || isNil(o.ConfigFiles) {
    return nil, false
	}
	return o.ConfigFiles, true
}

// HasConfigFiles returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasConfigFiles() bool {
	if o != nil && !isNil(o.ConfigFiles) {
		return true
	}

	return false
}

// SetConfigFiles gets a reference to the given []ConfigFile and assigns it to the ConfigFiles field.
func (o *RegionalDeploymentDefinition) SetConfigFiles(v []ConfigFile) {
	o.ConfigFiles = v
}

// GetSkipCache returns the SkipCache field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetSkipCache() bool {
	if o == nil || isNil(o.SkipCache) {
		var ret bool
		return ret
	}
	return *o.SkipCache
}

// GetSkipCacheOk returns a tuple with the SkipCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetSkipCacheOk() (*bool, bool) {
	if o == nil || isNil(o.SkipCache) {
    return nil, false
	}
	return o.SkipCache, true
}

// HasSkipCache returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasSkipCache() bool {
	if o != nil && !isNil(o.SkipCache) {
		return true
	}

	return false
}

// SetSkipCache gets a reference to the given bool and assigns it to the SkipCache field.
func (o *RegionalDeploymentDefinition) SetSkipCache(v bool) {
	o.SkipCache = &v
}

// GetMesh returns the Mesh field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetMesh() RegionalDeploymentMesh {
	if o == nil || isNil(o.Mesh) {
		var ret RegionalDeploymentMesh
		return ret
	}
	return *o.Mesh
}

// GetMeshOk returns a tuple with the Mesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetMeshOk() (*RegionalDeploymentMesh, bool) {
	if o == nil || isNil(o.Mesh) {
    return nil, false
	}
	return o.Mesh, true
}

// HasMesh returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasMesh() bool {
	if o != nil && !isNil(o.Mesh) {
		return true
	}

	return false
}

// SetMesh gets a reference to the given RegionalDeploymentMesh and assigns it to the Mesh field.
func (o *RegionalDeploymentDefinition) SetMesh(v RegionalDeploymentMesh) {
	o.Mesh = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetNetworkPolicy() NetworkPolicy {
	if o == nil || isNil(o.NetworkPolicy) {
		var ret NetworkPolicy
		return ret
	}
	return *o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetNetworkPolicyOk() (*NetworkPolicy, bool) {
	if o == nil || isNil(o.NetworkPolicy) {
    return nil, false
	}
	return o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasNetworkPolicy() bool {
	if o != nil && !isNil(o.NetworkPolicy) {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given NetworkPolicy and assigns it to the NetworkPolicy field.
func (o *RegionalDeploymentDefinition) SetNetworkPolicy(v NetworkPolicy) {
	o.NetworkPolicy = &v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetDocker() DockerSource {
	if o == nil || isNil(o.Docker) {
		var ret DockerSource
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetDockerOk() (*DockerSource, bool) {
	if o == nil || isNil(o.Docker) {
    return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasDocker() bool {
	if o != nil && !isNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerSource and assigns it to the Docker field.
func (o *RegionalDeploymentDefinition) SetDocker(v DockerSource) {
	o.Docker = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetGit() GitSource {
	if o == nil || isNil(o.Git) {
		var ret GitSource
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetGitOk() (*GitSource, bool) {
	if o == nil || isNil(o.Git) {
    return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given GitSource and assigns it to the Git field.
func (o *RegionalDeploymentDefinition) SetGit(v GitSource) {
	o.Git = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *RegionalDeploymentDefinition) GetArchive() ArchiveSource {
	if o == nil || isNil(o.Archive) {
		var ret ArchiveSource
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentDefinition) GetArchiveOk() (*ArchiveSource, bool) {
	if o == nil || isNil(o.Archive) {
    return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *RegionalDeploymentDefinition) HasArchive() bool {
	if o != nil && !isNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given ArchiveSource and assigns it to the Archive field.
func (o *RegionalDeploymentDefinition) SetArchive(v ArchiveSource) {
	o.Archive = &v
}

func (o RegionalDeploymentDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !isNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Scaling) {
		toSerialize["scaling"] = o.Scaling
	}
	if !isNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !isNil(o.DeploymentGroup) {
		toSerialize["deployment_group"] = o.DeploymentGroup
	}
	if !isNil(o.HealthChecks) {
		toSerialize["health_checks"] = o.HealthChecks
	}
	if !isNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !isNil(o.ConfigFiles) {
		toSerialize["config_files"] = o.ConfigFiles
	}
	if !isNil(o.SkipCache) {
		toSerialize["skip_cache"] = o.SkipCache
	}
	if !isNil(o.Mesh) {
		toSerialize["mesh"] = o.Mesh
	}
	if !isNil(o.NetworkPolicy) {
		toSerialize["network_policy"] = o.NetworkPolicy
	}
	if !isNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	if !isNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeploymentDefinition struct {
	value *RegionalDeploymentDefinition
	isSet bool
}

func (v NullableRegionalDeploymentDefinition) Get() *RegionalDeploymentDefinition {
	return v.value
}

func (v *NullableRegionalDeploymentDefinition) Set(val *RegionalDeploymentDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentDefinition(val *RegionalDeploymentDefinition) *NullableRegionalDeploymentDefinition {
	return &NullableRegionalDeploymentDefinition{value: val, isSet: true}
}

func (v NullableRegionalDeploymentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


