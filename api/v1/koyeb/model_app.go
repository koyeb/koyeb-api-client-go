/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// App struct for App
type App struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`
	PausedAt *time.Time `json:"paused_at,omitempty"`
	ResumedAt *time.Time `json:"resumed_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	Status *AppStatus `json:"status,omitempty"`
	Messages *[]string `json:"messages,omitempty"`
	Version *string `json:"version,omitempty"`
	Domains *[]Domain `json:"domains,omitempty"`
}

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp() *App {
	this := App{}
	var status AppStatus = APPSTATUS_STARTING
	this.Status = &status
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	var status AppStatus = APPSTATUS_STARTING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *App) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *App) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *App) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *App) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *App) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *App) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *App) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *App) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *App) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *App) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *App) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *App) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *App) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *App) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *App) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *App) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *App) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *App) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetSucceededAt returns the SucceededAt field value if set, zero value otherwise.
func (o *App) GetSucceededAt() time.Time {
	if o == nil || o.SucceededAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SucceededAt
}

// GetSucceededAtOk returns a tuple with the SucceededAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSucceededAtOk() (*time.Time, bool) {
	if o == nil || o.SucceededAt == nil {
		return nil, false
	}
	return o.SucceededAt, true
}

// HasSucceededAt returns a boolean if a field has been set.
func (o *App) HasSucceededAt() bool {
	if o != nil && o.SucceededAt != nil {
		return true
	}

	return false
}

// SetSucceededAt gets a reference to the given time.Time and assigns it to the SucceededAt field.
func (o *App) SetSucceededAt(v time.Time) {
	o.SucceededAt = &v
}

// GetPausedAt returns the PausedAt field value if set, zero value otherwise.
func (o *App) GetPausedAt() time.Time {
	if o == nil || o.PausedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PausedAt
}

// GetPausedAtOk returns a tuple with the PausedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetPausedAtOk() (*time.Time, bool) {
	if o == nil || o.PausedAt == nil {
		return nil, false
	}
	return o.PausedAt, true
}

// HasPausedAt returns a boolean if a field has been set.
func (o *App) HasPausedAt() bool {
	if o != nil && o.PausedAt != nil {
		return true
	}

	return false
}

// SetPausedAt gets a reference to the given time.Time and assigns it to the PausedAt field.
func (o *App) SetPausedAt(v time.Time) {
	o.PausedAt = &v
}

// GetResumedAt returns the ResumedAt field value if set, zero value otherwise.
func (o *App) GetResumedAt() time.Time {
	if o == nil || o.ResumedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ResumedAt
}

// GetResumedAtOk returns a tuple with the ResumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetResumedAtOk() (*time.Time, bool) {
	if o == nil || o.ResumedAt == nil {
		return nil, false
	}
	return o.ResumedAt, true
}

// HasResumedAt returns a boolean if a field has been set.
func (o *App) HasResumedAt() bool {
	if o != nil && o.ResumedAt != nil {
		return true
	}

	return false
}

// SetResumedAt gets a reference to the given time.Time and assigns it to the ResumedAt field.
func (o *App) SetResumedAt(v time.Time) {
	o.ResumedAt = &v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *App) GetTerminatedAt() time.Time {
	if o == nil || o.TerminatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || o.TerminatedAt == nil {
		return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *App) HasTerminatedAt() bool {
	if o != nil && o.TerminatedAt != nil {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *App) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *App) GetStatus() AppStatus {
	if o == nil || o.Status == nil {
		var ret AppStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetStatusOk() (*AppStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *App) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AppStatus and assigns it to the Status field.
func (o *App) SetStatus(v AppStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *App) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *App) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *App) SetMessages(v []string) {
	o.Messages = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *App) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *App) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *App) SetVersion(v string) {
	o.Version = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *App) GetDomains() []Domain {
	if o == nil || o.Domains == nil {
		var ret []Domain
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetDomainsOk() (*[]Domain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *App) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []Domain and assigns it to the Domains field.
func (o *App) SetDomains(v []Domain) {
	o.Domains = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.SucceededAt != nil {
		toSerialize["succeeded_at"] = o.SucceededAt
	}
	if o.PausedAt != nil {
		toSerialize["paused_at"] = o.PausedAt
	}
	if o.ResumedAt != nil {
		toSerialize["resumed_at"] = o.ResumedAt
	}
	if o.TerminatedAt != nil {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	return json.Marshal(toSerialize)
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


