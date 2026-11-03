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

// GitSource struct for GitSource
type GitSource struct {
	// A url to a git repository (contains the provider as well) .e.g: github.com/koyeb/test.
	Repository *string `json:"repository,omitempty"`
	Branch *string `json:"branch,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Sha *string `json:"sha,omitempty"`
	BuildCommand *string `json:"build_command,omitempty"`
	RunCommand *string `json:"run_command,omitempty"`
	NoDeployOnPush *bool `json:"no_deploy_on_push,omitempty"`
	Workdir *string `json:"workdir,omitempty"`
	Buildpack *BuildpackBuilder `json:"buildpack,omitempty"`
	Docker *DockerBuilder `json:"docker,omitempty"`
	// Credentials can come from a specific source or default to using Github (when unset or set to \"github\").
	CredentialSource *string `json:"credential_source,omitempty"`
	// A flag to skip initializing git submodules when cloning the repository.
	SkipSubmodules *bool `json:"skip_submodules,omitempty"`
}

// NewGitSource instantiates a new GitSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitSource() *GitSource {
	this := GitSource{}
	return &this
}

// NewGitSourceWithDefaults instantiates a new GitSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitSourceWithDefaults() *GitSource {
	this := GitSource{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GitSource) GetRepository() string {
	if o == nil || isNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetRepositoryOk() (*string, bool) {
	if o == nil || isNil(o.Repository) {
    return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GitSource) HasRepository() bool {
	if o != nil && !isNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *GitSource) SetRepository(v string) {
	o.Repository = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *GitSource) GetBranch() string {
	if o == nil || isNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBranchOk() (*string, bool) {
	if o == nil || isNil(o.Branch) {
    return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GitSource) HasBranch() bool {
	if o != nil && !isNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *GitSource) SetBranch(v string) {
	o.Branch = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GitSource) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GitSource) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GitSource) SetTag(v string) {
	o.Tag = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitSource) GetSha() string {
	if o == nil || isNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetShaOk() (*string, bool) {
	if o == nil || isNil(o.Sha) {
    return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitSource) HasSha() bool {
	if o != nil && !isNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitSource) SetSha(v string) {
	o.Sha = &v
}

// GetBuildCommand returns the BuildCommand field value if set, zero value otherwise.
func (o *GitSource) GetBuildCommand() string {
	if o == nil || isNil(o.BuildCommand) {
		var ret string
		return ret
	}
	return *o.BuildCommand
}

// GetBuildCommandOk returns a tuple with the BuildCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBuildCommandOk() (*string, bool) {
	if o == nil || isNil(o.BuildCommand) {
    return nil, false
	}
	return o.BuildCommand, true
}

// HasBuildCommand returns a boolean if a field has been set.
func (o *GitSource) HasBuildCommand() bool {
	if o != nil && !isNil(o.BuildCommand) {
		return true
	}

	return false
}

// SetBuildCommand gets a reference to the given string and assigns it to the BuildCommand field.
func (o *GitSource) SetBuildCommand(v string) {
	o.BuildCommand = &v
}

// GetRunCommand returns the RunCommand field value if set, zero value otherwise.
func (o *GitSource) GetRunCommand() string {
	if o == nil || isNil(o.RunCommand) {
		var ret string
		return ret
	}
	return *o.RunCommand
}

// GetRunCommandOk returns a tuple with the RunCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetRunCommandOk() (*string, bool) {
	if o == nil || isNil(o.RunCommand) {
    return nil, false
	}
	return o.RunCommand, true
}

// HasRunCommand returns a boolean if a field has been set.
func (o *GitSource) HasRunCommand() bool {
	if o != nil && !isNil(o.RunCommand) {
		return true
	}

	return false
}

// SetRunCommand gets a reference to the given string and assigns it to the RunCommand field.
func (o *GitSource) SetRunCommand(v string) {
	o.RunCommand = &v
}

// GetNoDeployOnPush returns the NoDeployOnPush field value if set, zero value otherwise.
func (o *GitSource) GetNoDeployOnPush() bool {
	if o == nil || isNil(o.NoDeployOnPush) {
		var ret bool
		return ret
	}
	return *o.NoDeployOnPush
}

// GetNoDeployOnPushOk returns a tuple with the NoDeployOnPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetNoDeployOnPushOk() (*bool, bool) {
	if o == nil || isNil(o.NoDeployOnPush) {
    return nil, false
	}
	return o.NoDeployOnPush, true
}

// HasNoDeployOnPush returns a boolean if a field has been set.
func (o *GitSource) HasNoDeployOnPush() bool {
	if o != nil && !isNil(o.NoDeployOnPush) {
		return true
	}

	return false
}

// SetNoDeployOnPush gets a reference to the given bool and assigns it to the NoDeployOnPush field.
func (o *GitSource) SetNoDeployOnPush(v bool) {
	o.NoDeployOnPush = &v
}

// GetWorkdir returns the Workdir field value if set, zero value otherwise.
func (o *GitSource) GetWorkdir() string {
	if o == nil || isNil(o.Workdir) {
		var ret string
		return ret
	}
	return *o.Workdir
}

// GetWorkdirOk returns a tuple with the Workdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetWorkdirOk() (*string, bool) {
	if o == nil || isNil(o.Workdir) {
    return nil, false
	}
	return o.Workdir, true
}

// HasWorkdir returns a boolean if a field has been set.
func (o *GitSource) HasWorkdir() bool {
	if o != nil && !isNil(o.Workdir) {
		return true
	}

	return false
}

// SetWorkdir gets a reference to the given string and assigns it to the Workdir field.
func (o *GitSource) SetWorkdir(v string) {
	o.Workdir = &v
}

// GetBuildpack returns the Buildpack field value if set, zero value otherwise.
func (o *GitSource) GetBuildpack() BuildpackBuilder {
	if o == nil || isNil(o.Buildpack) {
		var ret BuildpackBuilder
		return ret
	}
	return *o.Buildpack
}

// GetBuildpackOk returns a tuple with the Buildpack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetBuildpackOk() (*BuildpackBuilder, bool) {
	if o == nil || isNil(o.Buildpack) {
    return nil, false
	}
	return o.Buildpack, true
}

// HasBuildpack returns a boolean if a field has been set.
func (o *GitSource) HasBuildpack() bool {
	if o != nil && !isNil(o.Buildpack) {
		return true
	}

	return false
}

// SetBuildpack gets a reference to the given BuildpackBuilder and assigns it to the Buildpack field.
func (o *GitSource) SetBuildpack(v BuildpackBuilder) {
	o.Buildpack = &v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *GitSource) GetDocker() DockerBuilder {
	if o == nil || isNil(o.Docker) {
		var ret DockerBuilder
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetDockerOk() (*DockerBuilder, bool) {
	if o == nil || isNil(o.Docker) {
    return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *GitSource) HasDocker() bool {
	if o != nil && !isNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerBuilder and assigns it to the Docker field.
func (o *GitSource) SetDocker(v DockerBuilder) {
	o.Docker = &v
}

// GetCredentialSource returns the CredentialSource field value if set, zero value otherwise.
func (o *GitSource) GetCredentialSource() string {
	if o == nil || isNil(o.CredentialSource) {
		var ret string
		return ret
	}
	return *o.CredentialSource
}

// GetCredentialSourceOk returns a tuple with the CredentialSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetCredentialSourceOk() (*string, bool) {
	if o == nil || isNil(o.CredentialSource) {
    return nil, false
	}
	return o.CredentialSource, true
}

// HasCredentialSource returns a boolean if a field has been set.
func (o *GitSource) HasCredentialSource() bool {
	if o != nil && !isNil(o.CredentialSource) {
		return true
	}

	return false
}

// SetCredentialSource gets a reference to the given string and assigns it to the CredentialSource field.
func (o *GitSource) SetCredentialSource(v string) {
	o.CredentialSource = &v
}

// GetSkipSubmodules returns the SkipSubmodules field value if set, zero value otherwise.
func (o *GitSource) GetSkipSubmodules() bool {
	if o == nil || isNil(o.SkipSubmodules) {
		var ret bool
		return ret
	}
	return *o.SkipSubmodules
}

// GetSkipSubmodulesOk returns a tuple with the SkipSubmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitSource) GetSkipSubmodulesOk() (*bool, bool) {
	if o == nil || isNil(o.SkipSubmodules) {
    return nil, false
	}
	return o.SkipSubmodules, true
}

// HasSkipSubmodules returns a boolean if a field has been set.
func (o *GitSource) HasSkipSubmodules() bool {
	if o != nil && !isNil(o.SkipSubmodules) {
		return true
	}

	return false
}

// SetSkipSubmodules gets a reference to the given bool and assigns it to the SkipSubmodules field.
func (o *GitSource) SetSkipSubmodules(v bool) {
	o.SkipSubmodules = &v
}

func (o GitSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !isNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !isNil(o.BuildCommand) {
		toSerialize["build_command"] = o.BuildCommand
	}
	if !isNil(o.RunCommand) {
		toSerialize["run_command"] = o.RunCommand
	}
	if !isNil(o.NoDeployOnPush) {
		toSerialize["no_deploy_on_push"] = o.NoDeployOnPush
	}
	if !isNil(o.Workdir) {
		toSerialize["workdir"] = o.Workdir
	}
	if !isNil(o.Buildpack) {
		toSerialize["buildpack"] = o.Buildpack
	}
	if !isNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	if !isNil(o.CredentialSource) {
		toSerialize["credential_source"] = o.CredentialSource
	}
	if !isNil(o.SkipSubmodules) {
		toSerialize["skip_submodules"] = o.SkipSubmodules
	}
	return json.Marshal(toSerialize)
}

type NullableGitSource struct {
	value *GitSource
	isSet bool
}

func (v NullableGitSource) Get() *GitSource {
	return v.value
}

func (v *NullableGitSource) Set(val *GitSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGitSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGitSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitSource(val *GitSource) *NullableGitSource {
	return &NullableGitSource{value: val, isSet: true}
}

func (v NullableGitSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


