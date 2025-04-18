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

// CreateDomain struct for CreateDomain
type CreateDomain struct {
	Name *string `json:"name,omitempty"`
	Type *DomainType `json:"type,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	Cloudflare map[string]interface{} `json:"cloudflare,omitempty"`
	Koyeb *DomainLoadBalancerKoyeb `json:"koyeb,omitempty"`
}

// NewCreateDomain instantiates a new CreateDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomain() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// NewCreateDomainWithDefaults instantiates a new CreateDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainWithDefaults() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDomain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDomain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDomain) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateDomain) GetType() DomainType {
	if o == nil || isNil(o.Type) {
		var ret DomainType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetTypeOk() (*DomainType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateDomain) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DomainType and assigns it to the Type field.
func (o *CreateDomain) SetType(v DomainType) {
	o.Type = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateDomain) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateDomain) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateDomain) SetAppId(v string) {
	o.AppId = &v
}

// GetCloudflare returns the Cloudflare field value if set, zero value otherwise.
func (o *CreateDomain) GetCloudflare() map[string]interface{} {
	if o == nil || isNil(o.Cloudflare) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cloudflare
}

// GetCloudflareOk returns a tuple with the Cloudflare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetCloudflareOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Cloudflare) {
    return map[string]interface{}{}, false
	}
	return o.Cloudflare, true
}

// HasCloudflare returns a boolean if a field has been set.
func (o *CreateDomain) HasCloudflare() bool {
	if o != nil && !isNil(o.Cloudflare) {
		return true
	}

	return false
}

// SetCloudflare gets a reference to the given map[string]interface{} and assigns it to the Cloudflare field.
func (o *CreateDomain) SetCloudflare(v map[string]interface{}) {
	o.Cloudflare = v
}

// GetKoyeb returns the Koyeb field value if set, zero value otherwise.
func (o *CreateDomain) GetKoyeb() DomainLoadBalancerKoyeb {
	if o == nil || isNil(o.Koyeb) {
		var ret DomainLoadBalancerKoyeb
		return ret
	}
	return *o.Koyeb
}

// GetKoyebOk returns a tuple with the Koyeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetKoyebOk() (*DomainLoadBalancerKoyeb, bool) {
	if o == nil || isNil(o.Koyeb) {
    return nil, false
	}
	return o.Koyeb, true
}

// HasKoyeb returns a boolean if a field has been set.
func (o *CreateDomain) HasKoyeb() bool {
	if o != nil && !isNil(o.Koyeb) {
		return true
	}

	return false
}

// SetKoyeb gets a reference to the given DomainLoadBalancerKoyeb and assigns it to the Koyeb field.
func (o *CreateDomain) SetKoyeb(v DomainLoadBalancerKoyeb) {
	o.Koyeb = &v
}

func (o CreateDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Cloudflare) {
		toSerialize["cloudflare"] = o.Cloudflare
	}
	if !isNil(o.Koyeb) {
		toSerialize["koyeb"] = o.Koyeb
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDomain struct {
	value *CreateDomain
	isSet bool
}

func (v NullableCreateDomain) Get() *CreateDomain {
	return v.value
}

func (v *NullableCreateDomain) Set(val *CreateDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomain(val *CreateDomain) *NullableCreateDomain {
	return &NullableCreateDomain{value: val, isSet: true}
}

func (v NullableCreateDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


