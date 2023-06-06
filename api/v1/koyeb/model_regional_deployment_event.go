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

// RegionalDeploymentEvent struct for RegionalDeploymentEvent
type RegionalDeploymentEvent struct {
	Id *string `json:"id,omitempty"`
	When *time.Time `json:"when,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	RegionalDeploymentId *string `json:"regional_deployment_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewRegionalDeploymentEvent instantiates a new RegionalDeploymentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeploymentEvent() *RegionalDeploymentEvent {
	this := RegionalDeploymentEvent{}
	return &this
}

// NewRegionalDeploymentEventWithDefaults instantiates a new RegionalDeploymentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentEventWithDefaults() *RegionalDeploymentEvent {
	this := RegionalDeploymentEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegionalDeploymentEvent) SetId(v string) {
	o.Id = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetWhen() time.Time {
	if o == nil || isNil(o.When) {
		var ret time.Time
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetWhenOk() (*time.Time, bool) {
	if o == nil || isNil(o.When) {
    return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasWhen() bool {
	if o != nil && !isNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given time.Time and assigns it to the When field.
func (o *RegionalDeploymentEvent) SetWhen(v time.Time) {
	o.When = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *RegionalDeploymentEvent) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetRegionalDeploymentId returns the RegionalDeploymentId field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetRegionalDeploymentId() string {
	if o == nil || isNil(o.RegionalDeploymentId) {
		var ret string
		return ret
	}
	return *o.RegionalDeploymentId
}

// GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetRegionalDeploymentIdOk() (*string, bool) {
	if o == nil || isNil(o.RegionalDeploymentId) {
    return nil, false
	}
	return o.RegionalDeploymentId, true
}

// HasRegionalDeploymentId returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasRegionalDeploymentId() bool {
	if o != nil && !isNil(o.RegionalDeploymentId) {
		return true
	}

	return false
}

// SetRegionalDeploymentId gets a reference to the given string and assigns it to the RegionalDeploymentId field.
func (o *RegionalDeploymentEvent) SetRegionalDeploymentId(v string) {
	o.RegionalDeploymentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegionalDeploymentEvent) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RegionalDeploymentEvent) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RegionalDeploymentEvent) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentEvent) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegionalDeploymentEvent) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RegionalDeploymentEvent) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RegionalDeploymentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.When) {
		toSerialize["when"] = o.When
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.RegionalDeploymentId) {
		toSerialize["regional_deployment_id"] = o.RegionalDeploymentId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeploymentEvent struct {
	value *RegionalDeploymentEvent
	isSet bool
}

func (v NullableRegionalDeploymentEvent) Get() *RegionalDeploymentEvent {
	return v.value
}

func (v *NullableRegionalDeploymentEvent) Set(val *RegionalDeploymentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentEvent(val *RegionalDeploymentEvent) *NullableRegionalDeploymentEvent {
	return &NullableRegionalDeploymentEvent{value: val, isSet: true}
}

func (v NullableRegionalDeploymentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

