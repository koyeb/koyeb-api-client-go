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

// GetMetricsReplyMetric struct for GetMetricsReplyMetric
type GetMetricsReplyMetric struct {
	Labels *map[string]string `json:"labels,omitempty"`
	Samples []Sample `json:"samples,omitempty"`
}

// NewGetMetricsReplyMetric instantiates a new GetMetricsReplyMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricsReplyMetric() *GetMetricsReplyMetric {
	this := GetMetricsReplyMetric{}
	return &this
}

// NewGetMetricsReplyMetricWithDefaults instantiates a new GetMetricsReplyMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricsReplyMetricWithDefaults() *GetMetricsReplyMetric {
	this := GetMetricsReplyMetric{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *GetMetricsReplyMetric) GetLabels() map[string]string {
	if o == nil || isNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsReplyMetric) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Labels) {
    return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *GetMetricsReplyMetric) HasLabels() bool {
	if o != nil && !isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *GetMetricsReplyMetric) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *GetMetricsReplyMetric) GetSamples() []Sample {
	if o == nil || isNil(o.Samples) {
		var ret []Sample
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsReplyMetric) GetSamplesOk() ([]Sample, bool) {
	if o == nil || isNil(o.Samples) {
    return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *GetMetricsReplyMetric) HasSamples() bool {
	if o != nil && !isNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []Sample and assigns it to the Samples field.
func (o *GetMetricsReplyMetric) SetSamples(v []Sample) {
	o.Samples = v
}

func (o GetMetricsReplyMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !isNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return json.Marshal(toSerialize)
}

type NullableGetMetricsReplyMetric struct {
	value *GetMetricsReplyMetric
	isSet bool
}

func (v NullableGetMetricsReplyMetric) Get() *GetMetricsReplyMetric {
	return v.value
}

func (v *NullableGetMetricsReplyMetric) Set(val *GetMetricsReplyMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricsReplyMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricsReplyMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricsReplyMetric(val *GetMetricsReplyMetric) *NullableGetMetricsReplyMetric {
	return &NullableGetMetricsReplyMetric{value: val, isSet: true}
}

func (v NullableGetMetricsReplyMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricsReplyMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


