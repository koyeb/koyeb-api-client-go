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

// DeploymentDefinition struct for DeploymentDefinition
type DeploymentDefinition struct {
	// Service name. Deprecated, set it directly in the Service when creating it.
	Name *string `json:"name,omitempty"`
	Type *DeploymentDefinitionType `json:"type,omitempty"`
	Strategy *DeploymentStrategy `json:"strategy,omitempty"`
	Routes []DeploymentRoute `json:"routes,omitempty"`
	Ports []DeploymentPort `json:"ports,omitempty"`
	ProxyPorts []DeploymentProxyPort `json:"proxy_ports,omitempty"`
	Env []DeploymentEnv `json:"env,omitempty"`
	Regions []string `json:"regions,omitempty"`
	Scalings []DeploymentScaling `json:"scalings,omitempty"`
	InstanceTypes []DeploymentInstanceType `json:"instance_types,omitempty"`
	HealthChecks []DeploymentHealthCheck `json:"health_checks,omitempty"`
	Volumes []DeploymentVolume `json:"volumes,omitempty"`
	ConfigFiles []ConfigFile `json:"config_files,omitempty"`
	SkipCache *bool `json:"skip_cache,omitempty"`
	Mesh *DeploymentMesh `json:"mesh,omitempty"`
	NetworkPolicy *NetworkPolicy `json:"network_policy,omitempty"`
	Docker *DockerSource `json:"docker,omitempty"`
	Git *GitSource `json:"git,omitempty"`
	Database *DatabaseSource `json:"database,omitempty"`
	Archive *ArchiveSource `json:"archive,omitempty"`
}

// NewDeploymentDefinition instantiates a new DeploymentDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDefinition() *DeploymentDefinition {
	this := DeploymentDefinition{}
	var type_ DeploymentDefinitionType = DEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	var mesh DeploymentMesh = DEPLOYMENTMESH_AUTO
	this.Mesh = &mesh
	return &this
}

// NewDeploymentDefinitionWithDefaults instantiates a new DeploymentDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDefinitionWithDefaults() *DeploymentDefinition {
	this := DeploymentDefinition{}
	var type_ DeploymentDefinitionType = DEPLOYMENTDEFINITIONTYPE_INVALID
	this.Type = &type_
	var mesh DeploymentMesh = DEPLOYMENTMESH_AUTO
	this.Mesh = &mesh
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetType() DeploymentDefinitionType {
	if o == nil || isNil(o.Type) {
		var ret DeploymentDefinitionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetTypeOk() (*DeploymentDefinitionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploymentDefinitionType and assigns it to the Type field.
func (o *DeploymentDefinition) SetType(v DeploymentDefinitionType) {
	o.Type = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetStrategy() DeploymentStrategy {
	if o == nil || isNil(o.Strategy) {
		var ret DeploymentStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetStrategyOk() (*DeploymentStrategy, bool) {
	if o == nil || isNil(o.Strategy) {
    return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasStrategy() bool {
	if o != nil && !isNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeploymentStrategy and assigns it to the Strategy field.
func (o *DeploymentDefinition) SetStrategy(v DeploymentStrategy) {
	o.Strategy = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetRoutes() []DeploymentRoute {
	if o == nil || isNil(o.Routes) {
		var ret []DeploymentRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetRoutesOk() ([]DeploymentRoute, bool) {
	if o == nil || isNil(o.Routes) {
    return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasRoutes() bool {
	if o != nil && !isNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []DeploymentRoute and assigns it to the Routes field.
func (o *DeploymentDefinition) SetRoutes(v []DeploymentRoute) {
	o.Routes = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetPorts() []DeploymentPort {
	if o == nil || isNil(o.Ports) {
		var ret []DeploymentPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetPortsOk() ([]DeploymentPort, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []DeploymentPort and assigns it to the Ports field.
func (o *DeploymentDefinition) SetPorts(v []DeploymentPort) {
	o.Ports = v
}

// GetProxyPorts returns the ProxyPorts field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetProxyPorts() []DeploymentProxyPort {
	if o == nil || isNil(o.ProxyPorts) {
		var ret []DeploymentProxyPort
		return ret
	}
	return o.ProxyPorts
}

// GetProxyPortsOk returns a tuple with the ProxyPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetProxyPortsOk() ([]DeploymentProxyPort, bool) {
	if o == nil || isNil(o.ProxyPorts) {
    return nil, false
	}
	return o.ProxyPorts, true
}

// HasProxyPorts returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasProxyPorts() bool {
	if o != nil && !isNil(o.ProxyPorts) {
		return true
	}

	return false
}

// SetProxyPorts gets a reference to the given []DeploymentProxyPort and assigns it to the ProxyPorts field.
func (o *DeploymentDefinition) SetProxyPorts(v []DeploymentProxyPort) {
	o.ProxyPorts = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetEnv() []DeploymentEnv {
	if o == nil || isNil(o.Env) {
		var ret []DeploymentEnv
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetEnvOk() ([]DeploymentEnv, bool) {
	if o == nil || isNil(o.Env) {
    return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []DeploymentEnv and assigns it to the Env field.
func (o *DeploymentDefinition) SetEnv(v []DeploymentEnv) {
	o.Env = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *DeploymentDefinition) SetRegions(v []string) {
	o.Regions = v
}

// GetScalings returns the Scalings field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetScalings() []DeploymentScaling {
	if o == nil || isNil(o.Scalings) {
		var ret []DeploymentScaling
		return ret
	}
	return o.Scalings
}

// GetScalingsOk returns a tuple with the Scalings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetScalingsOk() ([]DeploymentScaling, bool) {
	if o == nil || isNil(o.Scalings) {
    return nil, false
	}
	return o.Scalings, true
}

// HasScalings returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasScalings() bool {
	if o != nil && !isNil(o.Scalings) {
		return true
	}

	return false
}

// SetScalings gets a reference to the given []DeploymentScaling and assigns it to the Scalings field.
func (o *DeploymentDefinition) SetScalings(v []DeploymentScaling) {
	o.Scalings = v
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetInstanceTypes() []DeploymentInstanceType {
	if o == nil || isNil(o.InstanceTypes) {
		var ret []DeploymentInstanceType
		return ret
	}
	return o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetInstanceTypesOk() ([]DeploymentInstanceType, bool) {
	if o == nil || isNil(o.InstanceTypes) {
    return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasInstanceTypes() bool {
	if o != nil && !isNil(o.InstanceTypes) {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []DeploymentInstanceType and assigns it to the InstanceTypes field.
func (o *DeploymentDefinition) SetInstanceTypes(v []DeploymentInstanceType) {
	o.InstanceTypes = v
}

// GetHealthChecks returns the HealthChecks field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetHealthChecks() []DeploymentHealthCheck {
	if o == nil || isNil(o.HealthChecks) {
		var ret []DeploymentHealthCheck
		return ret
	}
	return o.HealthChecks
}

// GetHealthChecksOk returns a tuple with the HealthChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetHealthChecksOk() ([]DeploymentHealthCheck, bool) {
	if o == nil || isNil(o.HealthChecks) {
    return nil, false
	}
	return o.HealthChecks, true
}

// HasHealthChecks returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasHealthChecks() bool {
	if o != nil && !isNil(o.HealthChecks) {
		return true
	}

	return false
}

// SetHealthChecks gets a reference to the given []DeploymentHealthCheck and assigns it to the HealthChecks field.
func (o *DeploymentDefinition) SetHealthChecks(v []DeploymentHealthCheck) {
	o.HealthChecks = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetVolumes() []DeploymentVolume {
	if o == nil || isNil(o.Volumes) {
		var ret []DeploymentVolume
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetVolumesOk() ([]DeploymentVolume, bool) {
	if o == nil || isNil(o.Volumes) {
    return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasVolumes() bool {
	if o != nil && !isNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []DeploymentVolume and assigns it to the Volumes field.
func (o *DeploymentDefinition) SetVolumes(v []DeploymentVolume) {
	o.Volumes = v
}

// GetConfigFiles returns the ConfigFiles field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetConfigFiles() []ConfigFile {
	if o == nil || isNil(o.ConfigFiles) {
		var ret []ConfigFile
		return ret
	}
	return o.ConfigFiles
}

// GetConfigFilesOk returns a tuple with the ConfigFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetConfigFilesOk() ([]ConfigFile, bool) {
	if o == nil || isNil(o.ConfigFiles) {
    return nil, false
	}
	return o.ConfigFiles, true
}

// HasConfigFiles returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasConfigFiles() bool {
	if o != nil && !isNil(o.ConfigFiles) {
		return true
	}

	return false
}

// SetConfigFiles gets a reference to the given []ConfigFile and assigns it to the ConfigFiles field.
func (o *DeploymentDefinition) SetConfigFiles(v []ConfigFile) {
	o.ConfigFiles = v
}

// GetSkipCache returns the SkipCache field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetSkipCache() bool {
	if o == nil || isNil(o.SkipCache) {
		var ret bool
		return ret
	}
	return *o.SkipCache
}

// GetSkipCacheOk returns a tuple with the SkipCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetSkipCacheOk() (*bool, bool) {
	if o == nil || isNil(o.SkipCache) {
    return nil, false
	}
	return o.SkipCache, true
}

// HasSkipCache returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasSkipCache() bool {
	if o != nil && !isNil(o.SkipCache) {
		return true
	}

	return false
}

// SetSkipCache gets a reference to the given bool and assigns it to the SkipCache field.
func (o *DeploymentDefinition) SetSkipCache(v bool) {
	o.SkipCache = &v
}

// GetMesh returns the Mesh field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetMesh() DeploymentMesh {
	if o == nil || isNil(o.Mesh) {
		var ret DeploymentMesh
		return ret
	}
	return *o.Mesh
}

// GetMeshOk returns a tuple with the Mesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetMeshOk() (*DeploymentMesh, bool) {
	if o == nil || isNil(o.Mesh) {
    return nil, false
	}
	return o.Mesh, true
}

// HasMesh returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasMesh() bool {
	if o != nil && !isNil(o.Mesh) {
		return true
	}

	return false
}

// SetMesh gets a reference to the given DeploymentMesh and assigns it to the Mesh field.
func (o *DeploymentDefinition) SetMesh(v DeploymentMesh) {
	o.Mesh = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetNetworkPolicy() NetworkPolicy {
	if o == nil || isNil(o.NetworkPolicy) {
		var ret NetworkPolicy
		return ret
	}
	return *o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetNetworkPolicyOk() (*NetworkPolicy, bool) {
	if o == nil || isNil(o.NetworkPolicy) {
    return nil, false
	}
	return o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasNetworkPolicy() bool {
	if o != nil && !isNil(o.NetworkPolicy) {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given NetworkPolicy and assigns it to the NetworkPolicy field.
func (o *DeploymentDefinition) SetNetworkPolicy(v NetworkPolicy) {
	o.NetworkPolicy = &v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetDocker() DockerSource {
	if o == nil || isNil(o.Docker) {
		var ret DockerSource
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetDockerOk() (*DockerSource, bool) {
	if o == nil || isNil(o.Docker) {
    return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasDocker() bool {
	if o != nil && !isNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerSource and assigns it to the Docker field.
func (o *DeploymentDefinition) SetDocker(v DockerSource) {
	o.Docker = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetGit() GitSource {
	if o == nil || isNil(o.Git) {
		var ret GitSource
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetGitOk() (*GitSource, bool) {
	if o == nil || isNil(o.Git) {
    return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given GitSource and assigns it to the Git field.
func (o *DeploymentDefinition) SetGit(v GitSource) {
	o.Git = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetDatabase() DatabaseSource {
	if o == nil || isNil(o.Database) {
		var ret DatabaseSource
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetDatabaseOk() (*DatabaseSource, bool) {
	if o == nil || isNil(o.Database) {
    return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasDatabase() bool {
	if o != nil && !isNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given DatabaseSource and assigns it to the Database field.
func (o *DeploymentDefinition) SetDatabase(v DatabaseSource) {
	o.Database = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *DeploymentDefinition) GetArchive() ArchiveSource {
	if o == nil || isNil(o.Archive) {
		var ret ArchiveSource
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDefinition) GetArchiveOk() (*ArchiveSource, bool) {
	if o == nil || isNil(o.Archive) {
    return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *DeploymentDefinition) HasArchive() bool {
	if o != nil && !isNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given ArchiveSource and assigns it to the Archive field.
func (o *DeploymentDefinition) SetArchive(v ArchiveSource) {
	o.Archive = &v
}

func (o DeploymentDefinition) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ProxyPorts) {
		toSerialize["proxy_ports"] = o.ProxyPorts
	}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.Scalings) {
		toSerialize["scalings"] = o.Scalings
	}
	if !isNil(o.InstanceTypes) {
		toSerialize["instance_types"] = o.InstanceTypes
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
	if !isNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !isNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentDefinition struct {
	value *DeploymentDefinition
	isSet bool
}

func (v NullableDeploymentDefinition) Get() *DeploymentDefinition {
	return v.value
}

func (v *NullableDeploymentDefinition) Set(val *DeploymentDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDefinition(val *DeploymentDefinition) *NullableDeploymentDefinition {
	return &NullableDeploymentDefinition{value: val, isSet: true}
}

func (v NullableDeploymentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


