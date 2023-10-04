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

// DeploymentNeonPostgresDatabaseInfo struct for DeploymentNeonPostgresDatabaseInfo
type DeploymentNeonPostgresDatabaseInfo struct {
	ActiveTimeSeconds *string `json:"active_time_seconds,omitempty"`
	ComputeTimeSeconds *string `json:"compute_time_seconds,omitempty"`
	WrittenDataBytes *string `json:"written_data_bytes,omitempty"`
	DataTransferBytes *string `json:"data_transfer_bytes,omitempty"`
	DataStorageBytesHour *string `json:"data_storage_bytes_hour,omitempty"`
	ServerHost *string `json:"server_host,omitempty"`
	ServerPort *int64 `json:"server_port,omitempty"`
	EndpointState *string `json:"endpoint_state,omitempty"`
	EndpointLastActive *time.Time `json:"endpoint_last_active,omitempty"`
	DefaultBranchId *string `json:"default_branch_id,omitempty"`
	DefaultBranchName *string `json:"default_branch_name,omitempty"`
	DefaultBranchState *string `json:"default_branch_state,omitempty"`
	DefaultBranchLogicalSize *string `json:"default_branch_logical_size,omitempty"`
	XyzRoles []XyzDeploymentNeonPostgresDatabaseInfoRoleObject `json:"xyz_roles,omitempty"`
}

// NewDeploymentNeonPostgresDatabaseInfo instantiates a new DeploymentNeonPostgresDatabaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentNeonPostgresDatabaseInfo() *DeploymentNeonPostgresDatabaseInfo {
	this := DeploymentNeonPostgresDatabaseInfo{}
	return &this
}

// NewDeploymentNeonPostgresDatabaseInfoWithDefaults instantiates a new DeploymentNeonPostgresDatabaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentNeonPostgresDatabaseInfoWithDefaults() *DeploymentNeonPostgresDatabaseInfo {
	this := DeploymentNeonPostgresDatabaseInfo{}
	return &this
}

// GetActiveTimeSeconds returns the ActiveTimeSeconds field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetActiveTimeSeconds() string {
	if o == nil || isNil(o.ActiveTimeSeconds) {
		var ret string
		return ret
	}
	return *o.ActiveTimeSeconds
}

// GetActiveTimeSecondsOk returns a tuple with the ActiveTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetActiveTimeSecondsOk() (*string, bool) {
	if o == nil || isNil(o.ActiveTimeSeconds) {
    return nil, false
	}
	return o.ActiveTimeSeconds, true
}

// HasActiveTimeSeconds returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasActiveTimeSeconds() bool {
	if o != nil && !isNil(o.ActiveTimeSeconds) {
		return true
	}

	return false
}

// SetActiveTimeSeconds gets a reference to the given string and assigns it to the ActiveTimeSeconds field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetActiveTimeSeconds(v string) {
	o.ActiveTimeSeconds = &v
}

// GetComputeTimeSeconds returns the ComputeTimeSeconds field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetComputeTimeSeconds() string {
	if o == nil || isNil(o.ComputeTimeSeconds) {
		var ret string
		return ret
	}
	return *o.ComputeTimeSeconds
}

// GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetComputeTimeSecondsOk() (*string, bool) {
	if o == nil || isNil(o.ComputeTimeSeconds) {
    return nil, false
	}
	return o.ComputeTimeSeconds, true
}

// HasComputeTimeSeconds returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasComputeTimeSeconds() bool {
	if o != nil && !isNil(o.ComputeTimeSeconds) {
		return true
	}

	return false
}

// SetComputeTimeSeconds gets a reference to the given string and assigns it to the ComputeTimeSeconds field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetComputeTimeSeconds(v string) {
	o.ComputeTimeSeconds = &v
}

// GetWrittenDataBytes returns the WrittenDataBytes field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetWrittenDataBytes() string {
	if o == nil || isNil(o.WrittenDataBytes) {
		var ret string
		return ret
	}
	return *o.WrittenDataBytes
}

// GetWrittenDataBytesOk returns a tuple with the WrittenDataBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetWrittenDataBytesOk() (*string, bool) {
	if o == nil || isNil(o.WrittenDataBytes) {
    return nil, false
	}
	return o.WrittenDataBytes, true
}

// HasWrittenDataBytes returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasWrittenDataBytes() bool {
	if o != nil && !isNil(o.WrittenDataBytes) {
		return true
	}

	return false
}

// SetWrittenDataBytes gets a reference to the given string and assigns it to the WrittenDataBytes field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetWrittenDataBytes(v string) {
	o.WrittenDataBytes = &v
}

// GetDataTransferBytes returns the DataTransferBytes field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDataTransferBytes() string {
	if o == nil || isNil(o.DataTransferBytes) {
		var ret string
		return ret
	}
	return *o.DataTransferBytes
}

// GetDataTransferBytesOk returns a tuple with the DataTransferBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDataTransferBytesOk() (*string, bool) {
	if o == nil || isNil(o.DataTransferBytes) {
    return nil, false
	}
	return o.DataTransferBytes, true
}

// HasDataTransferBytes returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDataTransferBytes() bool {
	if o != nil && !isNil(o.DataTransferBytes) {
		return true
	}

	return false
}

// SetDataTransferBytes gets a reference to the given string and assigns it to the DataTransferBytes field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDataTransferBytes(v string) {
	o.DataTransferBytes = &v
}

// GetDataStorageBytesHour returns the DataStorageBytesHour field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDataStorageBytesHour() string {
	if o == nil || isNil(o.DataStorageBytesHour) {
		var ret string
		return ret
	}
	return *o.DataStorageBytesHour
}

// GetDataStorageBytesHourOk returns a tuple with the DataStorageBytesHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDataStorageBytesHourOk() (*string, bool) {
	if o == nil || isNil(o.DataStorageBytesHour) {
    return nil, false
	}
	return o.DataStorageBytesHour, true
}

// HasDataStorageBytesHour returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDataStorageBytesHour() bool {
	if o != nil && !isNil(o.DataStorageBytesHour) {
		return true
	}

	return false
}

// SetDataStorageBytesHour gets a reference to the given string and assigns it to the DataStorageBytesHour field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDataStorageBytesHour(v string) {
	o.DataStorageBytesHour = &v
}

// GetServerHost returns the ServerHost field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetServerHost() string {
	if o == nil || isNil(o.ServerHost) {
		var ret string
		return ret
	}
	return *o.ServerHost
}

// GetServerHostOk returns a tuple with the ServerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetServerHostOk() (*string, bool) {
	if o == nil || isNil(o.ServerHost) {
    return nil, false
	}
	return o.ServerHost, true
}

// HasServerHost returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasServerHost() bool {
	if o != nil && !isNil(o.ServerHost) {
		return true
	}

	return false
}

// SetServerHost gets a reference to the given string and assigns it to the ServerHost field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetServerHost(v string) {
	o.ServerHost = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetServerPort() int64 {
	if o == nil || isNil(o.ServerPort) {
		var ret int64
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetServerPortOk() (*int64, bool) {
	if o == nil || isNil(o.ServerPort) {
    return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasServerPort() bool {
	if o != nil && !isNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int64 and assigns it to the ServerPort field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetServerPort(v int64) {
	o.ServerPort = &v
}

// GetEndpointState returns the EndpointState field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointState() string {
	if o == nil || isNil(o.EndpointState) {
		var ret string
		return ret
	}
	return *o.EndpointState
}

// GetEndpointStateOk returns a tuple with the EndpointState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointStateOk() (*string, bool) {
	if o == nil || isNil(o.EndpointState) {
    return nil, false
	}
	return o.EndpointState, true
}

// HasEndpointState returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasEndpointState() bool {
	if o != nil && !isNil(o.EndpointState) {
		return true
	}

	return false
}

// SetEndpointState gets a reference to the given string and assigns it to the EndpointState field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetEndpointState(v string) {
	o.EndpointState = &v
}

// GetEndpointLastActive returns the EndpointLastActive field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointLastActive() time.Time {
	if o == nil || isNil(o.EndpointLastActive) {
		var ret time.Time
		return ret
	}
	return *o.EndpointLastActive
}

// GetEndpointLastActiveOk returns a tuple with the EndpointLastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointLastActiveOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndpointLastActive) {
    return nil, false
	}
	return o.EndpointLastActive, true
}

// HasEndpointLastActive returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasEndpointLastActive() bool {
	if o != nil && !isNil(o.EndpointLastActive) {
		return true
	}

	return false
}

// SetEndpointLastActive gets a reference to the given time.Time and assigns it to the EndpointLastActive field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetEndpointLastActive(v time.Time) {
	o.EndpointLastActive = &v
}

// GetDefaultBranchId returns the DefaultBranchId field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchId() string {
	if o == nil || isNil(o.DefaultBranchId) {
		var ret string
		return ret
	}
	return *o.DefaultBranchId
}

// GetDefaultBranchIdOk returns a tuple with the DefaultBranchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchIdOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBranchId) {
    return nil, false
	}
	return o.DefaultBranchId, true
}

// HasDefaultBranchId returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchId() bool {
	if o != nil && !isNil(o.DefaultBranchId) {
		return true
	}

	return false
}

// SetDefaultBranchId gets a reference to the given string and assigns it to the DefaultBranchId field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchId(v string) {
	o.DefaultBranchId = &v
}

// GetDefaultBranchName returns the DefaultBranchName field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchName() string {
	if o == nil || isNil(o.DefaultBranchName) {
		var ret string
		return ret
	}
	return *o.DefaultBranchName
}

// GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchNameOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBranchName) {
    return nil, false
	}
	return o.DefaultBranchName, true
}

// HasDefaultBranchName returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchName() bool {
	if o != nil && !isNil(o.DefaultBranchName) {
		return true
	}

	return false
}

// SetDefaultBranchName gets a reference to the given string and assigns it to the DefaultBranchName field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchName(v string) {
	o.DefaultBranchName = &v
}

// GetDefaultBranchState returns the DefaultBranchState field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchState() string {
	if o == nil || isNil(o.DefaultBranchState) {
		var ret string
		return ret
	}
	return *o.DefaultBranchState
}

// GetDefaultBranchStateOk returns a tuple with the DefaultBranchState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchStateOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBranchState) {
    return nil, false
	}
	return o.DefaultBranchState, true
}

// HasDefaultBranchState returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchState() bool {
	if o != nil && !isNil(o.DefaultBranchState) {
		return true
	}

	return false
}

// SetDefaultBranchState gets a reference to the given string and assigns it to the DefaultBranchState field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchState(v string) {
	o.DefaultBranchState = &v
}

// GetDefaultBranchLogicalSize returns the DefaultBranchLogicalSize field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchLogicalSize() string {
	if o == nil || isNil(o.DefaultBranchLogicalSize) {
		var ret string
		return ret
	}
	return *o.DefaultBranchLogicalSize
}

// GetDefaultBranchLogicalSizeOk returns a tuple with the DefaultBranchLogicalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchLogicalSizeOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBranchLogicalSize) {
    return nil, false
	}
	return o.DefaultBranchLogicalSize, true
}

// HasDefaultBranchLogicalSize returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchLogicalSize() bool {
	if o != nil && !isNil(o.DefaultBranchLogicalSize) {
		return true
	}

	return false
}

// SetDefaultBranchLogicalSize gets a reference to the given string and assigns it to the DefaultBranchLogicalSize field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchLogicalSize(v string) {
	o.DefaultBranchLogicalSize = &v
}

// GetXyzRoles returns the XyzRoles field value if set, zero value otherwise.
func (o *DeploymentNeonPostgresDatabaseInfo) GetXyzRoles() []XyzDeploymentNeonPostgresDatabaseInfoRoleObject {
	if o == nil || isNil(o.XyzRoles) {
		var ret []XyzDeploymentNeonPostgresDatabaseInfoRoleObject
		return ret
	}
	return o.XyzRoles
}

// GetXyzRolesOk returns a tuple with the XyzRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) GetXyzRolesOk() ([]XyzDeploymentNeonPostgresDatabaseInfoRoleObject, bool) {
	if o == nil || isNil(o.XyzRoles) {
    return nil, false
	}
	return o.XyzRoles, true
}

// HasXyzRoles returns a boolean if a field has been set.
func (o *DeploymentNeonPostgresDatabaseInfo) HasXyzRoles() bool {
	if o != nil && !isNil(o.XyzRoles) {
		return true
	}

	return false
}

// SetXyzRoles gets a reference to the given []XyzDeploymentNeonPostgresDatabaseInfoRoleObject and assigns it to the XyzRoles field.
func (o *DeploymentNeonPostgresDatabaseInfo) SetXyzRoles(v []XyzDeploymentNeonPostgresDatabaseInfoRoleObject) {
	o.XyzRoles = v
}

func (o DeploymentNeonPostgresDatabaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActiveTimeSeconds) {
		toSerialize["active_time_seconds"] = o.ActiveTimeSeconds
	}
	if !isNil(o.ComputeTimeSeconds) {
		toSerialize["compute_time_seconds"] = o.ComputeTimeSeconds
	}
	if !isNil(o.WrittenDataBytes) {
		toSerialize["written_data_bytes"] = o.WrittenDataBytes
	}
	if !isNil(o.DataTransferBytes) {
		toSerialize["data_transfer_bytes"] = o.DataTransferBytes
	}
	if !isNil(o.DataStorageBytesHour) {
		toSerialize["data_storage_bytes_hour"] = o.DataStorageBytesHour
	}
	if !isNil(o.ServerHost) {
		toSerialize["server_host"] = o.ServerHost
	}
	if !isNil(o.ServerPort) {
		toSerialize["server_port"] = o.ServerPort
	}
	if !isNil(o.EndpointState) {
		toSerialize["endpoint_state"] = o.EndpointState
	}
	if !isNil(o.EndpointLastActive) {
		toSerialize["endpoint_last_active"] = o.EndpointLastActive
	}
	if !isNil(o.DefaultBranchId) {
		toSerialize["default_branch_id"] = o.DefaultBranchId
	}
	if !isNil(o.DefaultBranchName) {
		toSerialize["default_branch_name"] = o.DefaultBranchName
	}
	if !isNil(o.DefaultBranchState) {
		toSerialize["default_branch_state"] = o.DefaultBranchState
	}
	if !isNil(o.DefaultBranchLogicalSize) {
		toSerialize["default_branch_logical_size"] = o.DefaultBranchLogicalSize
	}
	if !isNil(o.XyzRoles) {
		toSerialize["xyz_roles"] = o.XyzRoles
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentNeonPostgresDatabaseInfo struct {
	value *DeploymentNeonPostgresDatabaseInfo
	isSet bool
}

func (v NullableDeploymentNeonPostgresDatabaseInfo) Get() *DeploymentNeonPostgresDatabaseInfo {
	return v.value
}

func (v *NullableDeploymentNeonPostgresDatabaseInfo) Set(val *DeploymentNeonPostgresDatabaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentNeonPostgresDatabaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentNeonPostgresDatabaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentNeonPostgresDatabaseInfo(val *DeploymentNeonPostgresDatabaseInfo) *NullableDeploymentNeonPostgresDatabaseInfo {
	return &NullableDeploymentNeonPostgresDatabaseInfo{value: val, isSet: true}
}

func (v NullableDeploymentNeonPostgresDatabaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentNeonPostgresDatabaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

