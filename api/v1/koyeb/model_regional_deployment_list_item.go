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

// RegionalDeploymentListItem struct for RegionalDeploymentListItem
type RegionalDeploymentListItem struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Region *string `json:"region,omitempty"`
	Status *RegionalDeploymentStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Definition *RegionalDeploymentDefinition `json:"definition,omitempty"`
	UseKumaV2 *bool `json:"use_kuma_v2,omitempty"`
	UseKata *bool `json:"use_kata,omitempty"`
}

// NewRegionalDeploymentListItem instantiates a new RegionalDeploymentListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeploymentListItem() *RegionalDeploymentListItem {
	this := RegionalDeploymentListItem{}
	var status RegionalDeploymentStatus = REGIONALDEPLOYMENTSTATUS_PENDING
	this.Status = &status
	return &this
}

// NewRegionalDeploymentListItemWithDefaults instantiates a new RegionalDeploymentListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentListItemWithDefaults() *RegionalDeploymentListItem {
	this := RegionalDeploymentListItem{}
	var status RegionalDeploymentStatus = REGIONALDEPLOYMENTSTATUS_PENDING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegionalDeploymentListItem) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RegionalDeploymentListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RegionalDeploymentListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RegionalDeploymentListItem) SetRegion(v string) {
	o.Region = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetStatus() RegionalDeploymentStatus {
	if o == nil || isNil(o.Status) {
		var ret RegionalDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetStatusOk() (*RegionalDeploymentStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RegionalDeploymentStatus and assigns it to the Status field.
func (o *RegionalDeploymentListItem) SetStatus(v RegionalDeploymentStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *RegionalDeploymentListItem) SetMessages(v []string) {
	o.Messages = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetDefinition() RegionalDeploymentDefinition {
	if o == nil || isNil(o.Definition) {
		var ret RegionalDeploymentDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetDefinitionOk() (*RegionalDeploymentDefinition, bool) {
	if o == nil || isNil(o.Definition) {
    return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given RegionalDeploymentDefinition and assigns it to the Definition field.
func (o *RegionalDeploymentListItem) SetDefinition(v RegionalDeploymentDefinition) {
	o.Definition = &v
}

// GetUseKumaV2 returns the UseKumaV2 field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetUseKumaV2() bool {
	if o == nil || isNil(o.UseKumaV2) {
		var ret bool
		return ret
	}
	return *o.UseKumaV2
}

// GetUseKumaV2Ok returns a tuple with the UseKumaV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetUseKumaV2Ok() (*bool, bool) {
	if o == nil || isNil(o.UseKumaV2) {
    return nil, false
	}
	return o.UseKumaV2, true
}

// HasUseKumaV2 returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasUseKumaV2() bool {
	if o != nil && !isNil(o.UseKumaV2) {
		return true
	}

	return false
}

// SetUseKumaV2 gets a reference to the given bool and assigns it to the UseKumaV2 field.
func (o *RegionalDeploymentListItem) SetUseKumaV2(v bool) {
	o.UseKumaV2 = &v
}

// GetUseKata returns the UseKata field value if set, zero value otherwise.
func (o *RegionalDeploymentListItem) GetUseKata() bool {
	if o == nil || isNil(o.UseKata) {
		var ret bool
		return ret
	}
	return *o.UseKata
}

// GetUseKataOk returns a tuple with the UseKata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentListItem) GetUseKataOk() (*bool, bool) {
	if o == nil || isNil(o.UseKata) {
    return nil, false
	}
	return o.UseKata, true
}

// HasUseKata returns a boolean if a field has been set.
func (o *RegionalDeploymentListItem) HasUseKata() bool {
	if o != nil && !isNil(o.UseKata) {
		return true
	}

	return false
}

// SetUseKata gets a reference to the given bool and assigns it to the UseKata field.
func (o *RegionalDeploymentListItem) SetUseKata(v bool) {
	o.UseKata = &v
}

func (o RegionalDeploymentListItem) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
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
	if !isNil(o.UseKumaV2) {
		toSerialize["use_kuma_v2"] = o.UseKumaV2
	}
	if !isNil(o.UseKata) {
		toSerialize["use_kata"] = o.UseKata
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeploymentListItem struct {
	value *RegionalDeploymentListItem
	isSet bool
}

func (v NullableRegionalDeploymentListItem) Get() *RegionalDeploymentListItem {
	return v.value
}

func (v *NullableRegionalDeploymentListItem) Set(val *RegionalDeploymentListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentListItem(val *RegionalDeploymentListItem) *NullableRegionalDeploymentListItem {
	return &NullableRegionalDeploymentListItem{value: val, isSet: true}
}

func (v NullableRegionalDeploymentListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


