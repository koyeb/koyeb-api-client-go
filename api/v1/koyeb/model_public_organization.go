/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// checks if the PublicOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicOrganization{}

// PublicOrganization struct for PublicOrganization
type PublicOrganization struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Plan *Plan `json:"plan,omitempty"`
	Status *OrganizationStatus `json:"status,omitempty"`
}

// NewPublicOrganization instantiates a new PublicOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicOrganization() *PublicOrganization {
	this := PublicOrganization{}
	var plan Plan = PLAN_HOBBY
	this.Plan = &plan
	var status OrganizationStatus = ORGANIZATIONSTATUS_WARNING
	this.Status = &status
	return &this
}

// NewPublicOrganizationWithDefaults instantiates a new PublicOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicOrganizationWithDefaults() *PublicOrganization {
	this := PublicOrganization{}
	var plan Plan = PLAN_HOBBY
	this.Plan = &plan
	var status OrganizationStatus = ORGANIZATIONSTATUS_WARNING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicOrganization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicOrganization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicOrganization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicOrganization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicOrganization) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicOrganization) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicOrganization) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicOrganization) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *PublicOrganization) GetPlan() Plan {
	if o == nil || IsNil(o.Plan) {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicOrganization) GetPlanOk() (*Plan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *PublicOrganization) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *PublicOrganization) SetPlan(v Plan) {
	o.Plan = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicOrganization) GetStatus() OrganizationStatus {
	if o == nil || IsNil(o.Status) {
		var ret OrganizationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicOrganization) GetStatusOk() (*OrganizationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicOrganization) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationStatus and assigns it to the Status field.
func (o *PublicOrganization) SetStatus(v OrganizationStatus) {
	o.Status = &v
}

func (o PublicOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePublicOrganization struct {
	value *PublicOrganization
	isSet bool
}

func (v NullablePublicOrganization) Get() *PublicOrganization {
	return v.value
}

func (v *NullablePublicOrganization) Set(val *PublicOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicOrganization(val *PublicOrganization) *NullablePublicOrganization {
	return &NullablePublicOrganization{value: val, isSet: true}
}

func (v NullablePublicOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


