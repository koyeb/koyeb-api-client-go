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

// PostgresNeonDatabaseNeonDatabase struct for PostgresNeonDatabaseNeonDatabase
type PostgresNeonDatabaseNeonDatabase struct {
	Name *string `json:"name,omitempty"`
	Owner *string `json:"owner,omitempty"`
}

// NewPostgresNeonDatabaseNeonDatabase instantiates a new PostgresNeonDatabaseNeonDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresNeonDatabaseNeonDatabase() *PostgresNeonDatabaseNeonDatabase {
	this := PostgresNeonDatabaseNeonDatabase{}
	return &this
}

// NewPostgresNeonDatabaseNeonDatabaseWithDefaults instantiates a new PostgresNeonDatabaseNeonDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresNeonDatabaseNeonDatabaseWithDefaults() *PostgresNeonDatabaseNeonDatabase {
	this := PostgresNeonDatabaseNeonDatabase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresNeonDatabaseNeonDatabase) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresNeonDatabaseNeonDatabase) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresNeonDatabaseNeonDatabase) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresNeonDatabaseNeonDatabase) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PostgresNeonDatabaseNeonDatabase) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresNeonDatabaseNeonDatabase) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PostgresNeonDatabaseNeonDatabase) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *PostgresNeonDatabaseNeonDatabase) SetOwner(v string) {
	o.Owner = &v
}

func (o PostgresNeonDatabaseNeonDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresNeonDatabaseNeonDatabase struct {
	value *PostgresNeonDatabaseNeonDatabase
	isSet bool
}

func (v NullablePostgresNeonDatabaseNeonDatabase) Get() *PostgresNeonDatabaseNeonDatabase {
	return v.value
}

func (v *NullablePostgresNeonDatabaseNeonDatabase) Set(val *PostgresNeonDatabaseNeonDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresNeonDatabaseNeonDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresNeonDatabaseNeonDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresNeonDatabaseNeonDatabase(val *PostgresNeonDatabaseNeonDatabase) *NullablePostgresNeonDatabaseNeonDatabase {
	return &NullablePostgresNeonDatabaseNeonDatabase{value: val, isSet: true}
}

func (v NullablePostgresNeonDatabaseNeonDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresNeonDatabaseNeonDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


