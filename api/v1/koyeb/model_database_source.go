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

// DatabaseSource struct for DatabaseSource
type DatabaseSource struct {
	PostgresNeon *PostgresNeonDatabase `json:"postgres_neon,omitempty"`
}

// NewDatabaseSource instantiates a new DatabaseSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseSource() *DatabaseSource {
	this := DatabaseSource{}
	return &this
}

// NewDatabaseSourceWithDefaults instantiates a new DatabaseSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseSourceWithDefaults() *DatabaseSource {
	this := DatabaseSource{}
	return &this
}

// GetPostgresNeon returns the PostgresNeon field value if set, zero value otherwise.
func (o *DatabaseSource) GetPostgresNeon() PostgresNeonDatabase {
	if o == nil || isNil(o.PostgresNeon) {
		var ret PostgresNeonDatabase
		return ret
	}
	return *o.PostgresNeon
}

// GetPostgresNeonOk returns a tuple with the PostgresNeon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSource) GetPostgresNeonOk() (*PostgresNeonDatabase, bool) {
	if o == nil || isNil(o.PostgresNeon) {
    return nil, false
	}
	return o.PostgresNeon, true
}

// HasPostgresNeon returns a boolean if a field has been set.
func (o *DatabaseSource) HasPostgresNeon() bool {
	if o != nil && !isNil(o.PostgresNeon) {
		return true
	}

	return false
}

// SetPostgresNeon gets a reference to the given PostgresNeonDatabase and assigns it to the PostgresNeon field.
func (o *DatabaseSource) SetPostgresNeon(v PostgresNeonDatabase) {
	o.PostgresNeon = &v
}

func (o DatabaseSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PostgresNeon) {
		toSerialize["postgres_neon"] = o.PostgresNeon
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseSource struct {
	value *DatabaseSource
	isSet bool
}

func (v NullableDatabaseSource) Get() *DatabaseSource {
	return v.value
}

func (v *NullableDatabaseSource) Set(val *DatabaseSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseSource(val *DatabaseSource) *NullableDatabaseSource {
	return &NullableDatabaseSource{value: val, isSet: true}
}

func (v NullableDatabaseSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

