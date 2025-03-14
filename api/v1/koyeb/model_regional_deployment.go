/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// RegionalDeployment struct for RegionalDeployment
type RegionalDeployment struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	AllocatedAt *time.Time `json:"allocated_at,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	Region *string `json:"region,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	ChildId *string `json:"child_id,omitempty"`
	Status *RegionalDeploymentStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Definition *RegionalDeploymentDefinition `json:"definition,omitempty"`
	Datacenters []string `json:"datacenters,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ProvisioningInfo *DeploymentProvisioningInfo `json:"provisioning_info,omitempty"`
	Role *RegionalDeploymentRole `json:"role,omitempty"`
	Version *string `json:"version,omitempty"`
	DeploymentGroup *string `json:"deployment_group,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
}

// NewRegionalDeployment instantiates a new RegionalDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeployment() *RegionalDeployment {
	this := RegionalDeployment{}
	var status RegionalDeploymentStatus = REGIONALDEPLOYMENTSTATUS_PENDING
	this.Status = &status
	var role RegionalDeploymentRole = REGIONALDEPLOYMENTROLE_INVALID
	this.Role = &role
	return &this
}

// NewRegionalDeploymentWithDefaults instantiates a new RegionalDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentWithDefaults() *RegionalDeployment {
	this := RegionalDeployment{}
	var status RegionalDeploymentStatus = REGIONALDEPLOYMENTSTATUS_PENDING
	this.Status = &status
	var role RegionalDeploymentRole = REGIONALDEPLOYMENTROLE_INVALID
	this.Role = &role
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegionalDeployment) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegionalDeployment) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RegionalDeployment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RegionalDeployment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetScheduledAt() time.Time {
	if o == nil || isNil(o.ScheduledAt) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledAt) {
    return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasScheduledAt() bool {
	if o != nil && !isNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *RegionalDeployment) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

// GetAllocatedAt returns the AllocatedAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetAllocatedAt() time.Time {
	if o == nil || isNil(o.AllocatedAt) {
		var ret time.Time
		return ret
	}
	return *o.AllocatedAt
}

// GetAllocatedAtOk returns a tuple with the AllocatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetAllocatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AllocatedAt) {
    return nil, false
	}
	return o.AllocatedAt, true
}

// HasAllocatedAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasAllocatedAt() bool {
	if o != nil && !isNil(o.AllocatedAt) {
		return true
	}

	return false
}

// SetAllocatedAt gets a reference to the given time.Time and assigns it to the AllocatedAt field.
func (o *RegionalDeployment) SetAllocatedAt(v time.Time) {
	o.AllocatedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *RegionalDeployment) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetSucceededAt returns the SucceededAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetSucceededAt() time.Time {
	if o == nil || isNil(o.SucceededAt) {
		var ret time.Time
		return ret
	}
	return *o.SucceededAt
}

// GetSucceededAtOk returns a tuple with the SucceededAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetSucceededAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SucceededAt) {
    return nil, false
	}
	return o.SucceededAt, true
}

// HasSucceededAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasSucceededAt() bool {
	if o != nil && !isNil(o.SucceededAt) {
		return true
	}

	return false
}

// SetSucceededAt gets a reference to the given time.Time and assigns it to the SucceededAt field.
func (o *RegionalDeployment) SetSucceededAt(v time.Time) {
	o.SucceededAt = &v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *RegionalDeployment) GetTerminatedAt() time.Time {
	if o == nil || isNil(o.TerminatedAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.TerminatedAt) {
    return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *RegionalDeployment) HasTerminatedAt() bool {
	if o != nil && !isNil(o.TerminatedAt) {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *RegionalDeployment) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *RegionalDeployment) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *RegionalDeployment) SetAppId(v string) {
	o.AppId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *RegionalDeployment) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegionalDeployment) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegionalDeployment) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RegionalDeployment) SetRegion(v string) {
	o.Region = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetParentId() string {
	if o == nil || isNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetParentIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentId) {
    return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *RegionalDeployment) SetParentId(v string) {
	o.ParentId = &v
}

// GetChildId returns the ChildId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetChildId() string {
	if o == nil || isNil(o.ChildId) {
		var ret string
		return ret
	}
	return *o.ChildId
}

// GetChildIdOk returns a tuple with the ChildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetChildIdOk() (*string, bool) {
	if o == nil || isNil(o.ChildId) {
    return nil, false
	}
	return o.ChildId, true
}

// HasChildId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasChildId() bool {
	if o != nil && !isNil(o.ChildId) {
		return true
	}

	return false
}

// SetChildId gets a reference to the given string and assigns it to the ChildId field.
func (o *RegionalDeployment) SetChildId(v string) {
	o.ChildId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegionalDeployment) GetStatus() RegionalDeploymentStatus {
	if o == nil || isNil(o.Status) {
		var ret RegionalDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetStatusOk() (*RegionalDeploymentStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegionalDeployment) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RegionalDeploymentStatus and assigns it to the Status field.
func (o *RegionalDeployment) SetStatus(v RegionalDeploymentStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *RegionalDeployment) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *RegionalDeployment) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *RegionalDeployment) SetMessages(v []string) {
	o.Messages = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *RegionalDeployment) GetDefinition() RegionalDeploymentDefinition {
	if o == nil || isNil(o.Definition) {
		var ret RegionalDeploymentDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetDefinitionOk() (*RegionalDeploymentDefinition, bool) {
	if o == nil || isNil(o.Definition) {
    return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *RegionalDeployment) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given RegionalDeploymentDefinition and assigns it to the Definition field.
func (o *RegionalDeployment) SetDefinition(v RegionalDeploymentDefinition) {
	o.Definition = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *RegionalDeployment) GetDatacenters() []string {
	if o == nil || isNil(o.Datacenters) {
		var ret []string
		return ret
	}
	return o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetDatacentersOk() ([]string, bool) {
	if o == nil || isNil(o.Datacenters) {
    return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *RegionalDeployment) HasDatacenters() bool {
	if o != nil && !isNil(o.Datacenters) {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *RegionalDeployment) SetDatacenters(v []string) {
	o.Datacenters = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RegionalDeployment) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegionalDeployment) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RegionalDeployment) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProvisioningInfo returns the ProvisioningInfo field value if set, zero value otherwise.
func (o *RegionalDeployment) GetProvisioningInfo() DeploymentProvisioningInfo {
	if o == nil || isNil(o.ProvisioningInfo) {
		var ret DeploymentProvisioningInfo
		return ret
	}
	return *o.ProvisioningInfo
}

// GetProvisioningInfoOk returns a tuple with the ProvisioningInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetProvisioningInfoOk() (*DeploymentProvisioningInfo, bool) {
	if o == nil || isNil(o.ProvisioningInfo) {
    return nil, false
	}
	return o.ProvisioningInfo, true
}

// HasProvisioningInfo returns a boolean if a field has been set.
func (o *RegionalDeployment) HasProvisioningInfo() bool {
	if o != nil && !isNil(o.ProvisioningInfo) {
		return true
	}

	return false
}

// SetProvisioningInfo gets a reference to the given DeploymentProvisioningInfo and assigns it to the ProvisioningInfo field.
func (o *RegionalDeployment) SetProvisioningInfo(v DeploymentProvisioningInfo) {
	o.ProvisioningInfo = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RegionalDeployment) GetRole() RegionalDeploymentRole {
	if o == nil || isNil(o.Role) {
		var ret RegionalDeploymentRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetRoleOk() (*RegionalDeploymentRole, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RegionalDeployment) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given RegionalDeploymentRole and assigns it to the Role field.
func (o *RegionalDeployment) SetRole(v RegionalDeploymentRole) {
	o.Role = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RegionalDeployment) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RegionalDeployment) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RegionalDeployment) SetVersion(v string) {
	o.Version = &v
}

// GetDeploymentGroup returns the DeploymentGroup field value if set, zero value otherwise.
func (o *RegionalDeployment) GetDeploymentGroup() string {
	if o == nil || isNil(o.DeploymentGroup) {
		var ret string
		return ret
	}
	return *o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetDeploymentGroupOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentGroup) {
    return nil, false
	}
	return o.DeploymentGroup, true
}

// HasDeploymentGroup returns a boolean if a field has been set.
func (o *RegionalDeployment) HasDeploymentGroup() bool {
	if o != nil && !isNil(o.DeploymentGroup) {
		return true
	}

	return false
}

// SetDeploymentGroup gets a reference to the given string and assigns it to the DeploymentGroup field.
func (o *RegionalDeployment) SetDeploymentGroup(v string) {
	o.DeploymentGroup = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *RegionalDeployment) GetDeploymentId() string {
	if o == nil || isNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeployment) GetDeploymentIdOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentId) {
    return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *RegionalDeployment) HasDeploymentId() bool {
	if o != nil && !isNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *RegionalDeployment) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

func (o RegionalDeployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.ScheduledAt) {
		toSerialize["scheduled_at"] = o.ScheduledAt
	}
	if !isNil(o.AllocatedAt) {
		toSerialize["allocated_at"] = o.AllocatedAt
	}
	if !isNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !isNil(o.SucceededAt) {
		toSerialize["succeeded_at"] = o.SucceededAt
	}
	if !isNil(o.TerminatedAt) {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !isNil(o.ChildId) {
		toSerialize["child_id"] = o.ChildId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !isNil(o.Datacenters) {
		toSerialize["datacenters"] = o.Datacenters
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.ProvisioningInfo) {
		toSerialize["provisioning_info"] = o.ProvisioningInfo
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.DeploymentGroup) {
		toSerialize["deployment_group"] = o.DeploymentGroup
	}
	if !isNil(o.DeploymentId) {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeployment struct {
	value *RegionalDeployment
	isSet bool
}

func (v NullableRegionalDeployment) Get() *RegionalDeployment {
	return v.value
}

func (v *NullableRegionalDeployment) Set(val *RegionalDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeployment(val *RegionalDeployment) *NullableRegionalDeployment {
	return &NullableRegionalDeployment{value: val, isSet: true}
}

func (v NullableRegionalDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


