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

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notification{}

// Notification struct for Notification
type Notification struct {
	Id *string `json:"id,omitempty"`
	Activity *Activity `json:"activity,omitempty"`
	IsRead *bool `json:"is_read,omitempty"`
	IsSeen *bool `json:"is_seen,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification() *Notification {
	this := Notification{}
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Notification) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Notification) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Notification) SetId(v string) {
	o.Id = &v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *Notification) GetActivity() Activity {
	if o == nil || IsNil(o.Activity) {
		var ret Activity
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetActivityOk() (*Activity, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *Notification) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given Activity and assigns it to the Activity field.
func (o *Notification) SetActivity(v Activity) {
	o.Activity = &v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *Notification) GetIsRead() bool {
	if o == nil || IsNil(o.IsRead) {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIsReadOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRead) {
		return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *Notification) HasIsRead() bool {
	if o != nil && !IsNil(o.IsRead) {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *Notification) SetIsRead(v bool) {
	o.IsRead = &v
}

// GetIsSeen returns the IsSeen field value if set, zero value otherwise.
func (o *Notification) GetIsSeen() bool {
	if o == nil || IsNil(o.IsSeen) {
		var ret bool
		return ret
	}
	return *o.IsSeen
}

// GetIsSeenOk returns a tuple with the IsSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIsSeenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSeen) {
		return nil, false
	}
	return o.IsSeen, true
}

// HasIsSeen returns a boolean if a field has been set.
func (o *Notification) HasIsSeen() bool {
	if o != nil && !IsNil(o.IsSeen) {
		return true
	}

	return false
}

// SetIsSeen gets a reference to the given bool and assigns it to the IsSeen field.
func (o *Notification) SetIsSeen(v bool) {
	o.IsSeen = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Notification) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Notification) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Notification) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	if !IsNil(o.IsRead) {
		toSerialize["is_read"] = o.IsRead
	}
	if !IsNil(o.IsSeen) {
		toSerialize["is_seen"] = o.IsSeen
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


