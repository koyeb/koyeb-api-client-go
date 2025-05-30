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

// Domain struct for Domain
type Domain struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Status *DomainStatus `json:"status,omitempty"`
	Type *DomainType `json:"type,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	DeploymentGroup *string `json:"deployment_group,omitempty"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	IntendedCname *string `json:"intended_cname,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Version *string `json:"version,omitempty"`
	Cloudflare map[string]interface{} `json:"cloudflare,omitempty"`
	Koyeb *DomainLoadBalancerKoyeb `json:"koyeb,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	var status DomainStatus = DOMAINSTATUS_PENDING
	this.Status = &status
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	var status DomainStatus = DOMAINSTATUS_PENDING
	this.Status = &status
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Domain) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Domain) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Domain) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Domain) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Domain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Domain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Domain) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Domain) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Domain) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Domain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Domain) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Domain) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Domain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Domain) GetStatus() DomainStatus {
	if o == nil || isNil(o.Status) {
		var ret DomainStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetStatusOk() (*DomainStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Domain) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DomainStatus and assigns it to the Status field.
func (o *Domain) SetStatus(v DomainStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Domain) GetType() DomainType {
	if o == nil || isNil(o.Type) {
		var ret DomainType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTypeOk() (*DomainType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Domain) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DomainType and assigns it to the Type field.
func (o *Domain) SetType(v DomainType) {
	o.Type = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Domain) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Domain) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Domain) SetAppId(v string) {
	o.AppId = &v
}

// GetDeploymentGroup returns the DeploymentGroup field value if set, zero value otherwise.
func (o *Domain) GetDeploymentGroup() string {
	if o == nil || isNil(o.DeploymentGroup) {
		var ret string
		return ret
	}
	return *o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDeploymentGroupOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentGroup) {
    return nil, false
	}
	return o.DeploymentGroup, true
}

// HasDeploymentGroup returns a boolean if a field has been set.
func (o *Domain) HasDeploymentGroup() bool {
	if o != nil && !isNil(o.DeploymentGroup) {
		return true
	}

	return false
}

// SetDeploymentGroup gets a reference to the given string and assigns it to the DeploymentGroup field.
func (o *Domain) SetDeploymentGroup(v string) {
	o.DeploymentGroup = &v
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise.
func (o *Domain) GetVerifiedAt() time.Time {
	if o == nil || isNil(o.VerifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.VerifiedAt
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetVerifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.VerifiedAt) {
    return nil, false
	}
	return o.VerifiedAt, true
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *Domain) HasVerifiedAt() bool {
	if o != nil && !isNil(o.VerifiedAt) {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given time.Time and assigns it to the VerifiedAt field.
func (o *Domain) SetVerifiedAt(v time.Time) {
	o.VerifiedAt = &v
}

// GetIntendedCname returns the IntendedCname field value if set, zero value otherwise.
func (o *Domain) GetIntendedCname() string {
	if o == nil || isNil(o.IntendedCname) {
		var ret string
		return ret
	}
	return *o.IntendedCname
}

// GetIntendedCnameOk returns a tuple with the IntendedCname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIntendedCnameOk() (*string, bool) {
	if o == nil || isNil(o.IntendedCname) {
    return nil, false
	}
	return o.IntendedCname, true
}

// HasIntendedCname returns a boolean if a field has been set.
func (o *Domain) HasIntendedCname() bool {
	if o != nil && !isNil(o.IntendedCname) {
		return true
	}

	return false
}

// SetIntendedCname gets a reference to the given string and assigns it to the IntendedCname field.
func (o *Domain) SetIntendedCname(v string) {
	o.IntendedCname = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Domain) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Domain) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *Domain) SetMessages(v []string) {
	o.Messages = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Domain) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Domain) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Domain) SetVersion(v string) {
	o.Version = &v
}

// GetCloudflare returns the Cloudflare field value if set, zero value otherwise.
func (o *Domain) GetCloudflare() map[string]interface{} {
	if o == nil || isNil(o.Cloudflare) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cloudflare
}

// GetCloudflareOk returns a tuple with the Cloudflare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCloudflareOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Cloudflare) {
    return map[string]interface{}{}, false
	}
	return o.Cloudflare, true
}

// HasCloudflare returns a boolean if a field has been set.
func (o *Domain) HasCloudflare() bool {
	if o != nil && !isNil(o.Cloudflare) {
		return true
	}

	return false
}

// SetCloudflare gets a reference to the given map[string]interface{} and assigns it to the Cloudflare field.
func (o *Domain) SetCloudflare(v map[string]interface{}) {
	o.Cloudflare = v
}

// GetKoyeb returns the Koyeb field value if set, zero value otherwise.
func (o *Domain) GetKoyeb() DomainLoadBalancerKoyeb {
	if o == nil || isNil(o.Koyeb) {
		var ret DomainLoadBalancerKoyeb
		return ret
	}
	return *o.Koyeb
}

// GetKoyebOk returns a tuple with the Koyeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetKoyebOk() (*DomainLoadBalancerKoyeb, bool) {
	if o == nil || isNil(o.Koyeb) {
    return nil, false
	}
	return o.Koyeb, true
}

// HasKoyeb returns a boolean if a field has been set.
func (o *Domain) HasKoyeb() bool {
	if o != nil && !isNil(o.Koyeb) {
		return true
	}

	return false
}

// SetKoyeb gets a reference to the given DomainLoadBalancerKoyeb and assigns it to the Koyeb field.
func (o *Domain) SetKoyeb(v DomainLoadBalancerKoyeb) {
	o.Koyeb = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.DeploymentGroup) {
		toSerialize["deployment_group"] = o.DeploymentGroup
	}
	if !isNil(o.VerifiedAt) {
		toSerialize["verified_at"] = o.VerifiedAt
	}
	if !isNil(o.IntendedCname) {
		toSerialize["intended_cname"] = o.IntendedCname
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Cloudflare) {
		toSerialize["cloudflare"] = o.Cloudflare
	}
	if !isNil(o.Koyeb) {
		toSerialize["koyeb"] = o.Koyeb
	}
	return json.Marshal(toSerialize)
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


