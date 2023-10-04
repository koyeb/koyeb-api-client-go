# DeploymentNeonPostgresDatabaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTimeSeconds** | Pointer to **string** |  | [optional] 
**ComputeTimeSeconds** | Pointer to **string** |  | [optional] 
**WrittenDataBytes** | Pointer to **string** |  | [optional] 
**DataTransferBytes** | Pointer to **string** |  | [optional] 
**DataStorageBytesHour** | Pointer to **string** |  | [optional] 
**ServerHost** | Pointer to **string** |  | [optional] 
**ServerPort** | Pointer to **int64** |  | [optional] 
**EndpointState** | Pointer to **string** |  | [optional] 
**EndpointLastActive** | Pointer to **time.Time** |  | [optional] 
**DefaultBranchId** | Pointer to **string** |  | [optional] 
**DefaultBranchName** | Pointer to **string** |  | [optional] 
**DefaultBranchState** | Pointer to **string** |  | [optional] 
**DefaultBranchLogicalSize** | Pointer to **string** |  | [optional] 
**XyzRoles** | Pointer to [**[]XyzDeploymentNeonPostgresDatabaseInfoRoleObject**](XyzDeploymentNeonPostgresDatabaseInfoRoleObject.md) |  | [optional] 

## Methods

### NewDeploymentNeonPostgresDatabaseInfo

`func NewDeploymentNeonPostgresDatabaseInfo() *DeploymentNeonPostgresDatabaseInfo`

NewDeploymentNeonPostgresDatabaseInfo instantiates a new DeploymentNeonPostgresDatabaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentNeonPostgresDatabaseInfoWithDefaults

`func NewDeploymentNeonPostgresDatabaseInfoWithDefaults() *DeploymentNeonPostgresDatabaseInfo`

NewDeploymentNeonPostgresDatabaseInfoWithDefaults instantiates a new DeploymentNeonPostgresDatabaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) GetActiveTimeSeconds() string`

GetActiveTimeSeconds returns the ActiveTimeSeconds field if non-nil, zero value otherwise.

### GetActiveTimeSecondsOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetActiveTimeSecondsOk() (*string, bool)`

GetActiveTimeSecondsOk returns a tuple with the ActiveTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) SetActiveTimeSeconds(v string)`

SetActiveTimeSeconds sets ActiveTimeSeconds field to given value.

### HasActiveTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) HasActiveTimeSeconds() bool`

HasActiveTimeSeconds returns a boolean if a field has been set.

### GetComputeTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) GetComputeTimeSeconds() string`

GetComputeTimeSeconds returns the ComputeTimeSeconds field if non-nil, zero value otherwise.

### GetComputeTimeSecondsOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetComputeTimeSecondsOk() (*string, bool)`

GetComputeTimeSecondsOk returns a tuple with the ComputeTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) SetComputeTimeSeconds(v string)`

SetComputeTimeSeconds sets ComputeTimeSeconds field to given value.

### HasComputeTimeSeconds

`func (o *DeploymentNeonPostgresDatabaseInfo) HasComputeTimeSeconds() bool`

HasComputeTimeSeconds returns a boolean if a field has been set.

### GetWrittenDataBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) GetWrittenDataBytes() string`

GetWrittenDataBytes returns the WrittenDataBytes field if non-nil, zero value otherwise.

### GetWrittenDataBytesOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetWrittenDataBytesOk() (*string, bool)`

GetWrittenDataBytesOk returns a tuple with the WrittenDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenDataBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) SetWrittenDataBytes(v string)`

SetWrittenDataBytes sets WrittenDataBytes field to given value.

### HasWrittenDataBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) HasWrittenDataBytes() bool`

HasWrittenDataBytes returns a boolean if a field has been set.

### GetDataTransferBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDataTransferBytes() string`

GetDataTransferBytes returns the DataTransferBytes field if non-nil, zero value otherwise.

### GetDataTransferBytesOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDataTransferBytesOk() (*string, bool)`

GetDataTransferBytesOk returns a tuple with the DataTransferBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDataTransferBytes(v string)`

SetDataTransferBytes sets DataTransferBytes field to given value.

### HasDataTransferBytes

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDataTransferBytes() bool`

HasDataTransferBytes returns a boolean if a field has been set.

### GetDataStorageBytesHour

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDataStorageBytesHour() string`

GetDataStorageBytesHour returns the DataStorageBytesHour field if non-nil, zero value otherwise.

### GetDataStorageBytesHourOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDataStorageBytesHourOk() (*string, bool)`

GetDataStorageBytesHourOk returns a tuple with the DataStorageBytesHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorageBytesHour

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDataStorageBytesHour(v string)`

SetDataStorageBytesHour sets DataStorageBytesHour field to given value.

### HasDataStorageBytesHour

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDataStorageBytesHour() bool`

HasDataStorageBytesHour returns a boolean if a field has been set.

### GetServerHost

`func (o *DeploymentNeonPostgresDatabaseInfo) GetServerHost() string`

GetServerHost returns the ServerHost field if non-nil, zero value otherwise.

### GetServerHostOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetServerHostOk() (*string, bool)`

GetServerHostOk returns a tuple with the ServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHost

`func (o *DeploymentNeonPostgresDatabaseInfo) SetServerHost(v string)`

SetServerHost sets ServerHost field to given value.

### HasServerHost

`func (o *DeploymentNeonPostgresDatabaseInfo) HasServerHost() bool`

HasServerHost returns a boolean if a field has been set.

### GetServerPort

`func (o *DeploymentNeonPostgresDatabaseInfo) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *DeploymentNeonPostgresDatabaseInfo) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *DeploymentNeonPostgresDatabaseInfo) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetEndpointState

`func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointState() string`

GetEndpointState returns the EndpointState field if non-nil, zero value otherwise.

### GetEndpointStateOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointStateOk() (*string, bool)`

GetEndpointStateOk returns a tuple with the EndpointState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointState

`func (o *DeploymentNeonPostgresDatabaseInfo) SetEndpointState(v string)`

SetEndpointState sets EndpointState field to given value.

### HasEndpointState

`func (o *DeploymentNeonPostgresDatabaseInfo) HasEndpointState() bool`

HasEndpointState returns a boolean if a field has been set.

### GetEndpointLastActive

`func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointLastActive() time.Time`

GetEndpointLastActive returns the EndpointLastActive field if non-nil, zero value otherwise.

### GetEndpointLastActiveOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetEndpointLastActiveOk() (*time.Time, bool)`

GetEndpointLastActiveOk returns a tuple with the EndpointLastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointLastActive

`func (o *DeploymentNeonPostgresDatabaseInfo) SetEndpointLastActive(v time.Time)`

SetEndpointLastActive sets EndpointLastActive field to given value.

### HasEndpointLastActive

`func (o *DeploymentNeonPostgresDatabaseInfo) HasEndpointLastActive() bool`

HasEndpointLastActive returns a boolean if a field has been set.

### GetDefaultBranchId

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchId() string`

GetDefaultBranchId returns the DefaultBranchId field if non-nil, zero value otherwise.

### GetDefaultBranchIdOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchIdOk() (*string, bool)`

GetDefaultBranchIdOk returns a tuple with the DefaultBranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchId

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchId(v string)`

SetDefaultBranchId sets DefaultBranchId field to given value.

### HasDefaultBranchId

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchId() bool`

HasDefaultBranchId returns a boolean if a field has been set.

### GetDefaultBranchName

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetDefaultBranchState

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchState() string`

GetDefaultBranchState returns the DefaultBranchState field if non-nil, zero value otherwise.

### GetDefaultBranchStateOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchStateOk() (*string, bool)`

GetDefaultBranchStateOk returns a tuple with the DefaultBranchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchState

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchState(v string)`

SetDefaultBranchState sets DefaultBranchState field to given value.

### HasDefaultBranchState

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchState() bool`

HasDefaultBranchState returns a boolean if a field has been set.

### GetDefaultBranchLogicalSize

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchLogicalSize() string`

GetDefaultBranchLogicalSize returns the DefaultBranchLogicalSize field if non-nil, zero value otherwise.

### GetDefaultBranchLogicalSizeOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetDefaultBranchLogicalSizeOk() (*string, bool)`

GetDefaultBranchLogicalSizeOk returns a tuple with the DefaultBranchLogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchLogicalSize

`func (o *DeploymentNeonPostgresDatabaseInfo) SetDefaultBranchLogicalSize(v string)`

SetDefaultBranchLogicalSize sets DefaultBranchLogicalSize field to given value.

### HasDefaultBranchLogicalSize

`func (o *DeploymentNeonPostgresDatabaseInfo) HasDefaultBranchLogicalSize() bool`

HasDefaultBranchLogicalSize returns a boolean if a field has been set.

### GetXyzRoles

`func (o *DeploymentNeonPostgresDatabaseInfo) GetXyzRoles() []XyzDeploymentNeonPostgresDatabaseInfoRoleObject`

GetXyzRoles returns the XyzRoles field if non-nil, zero value otherwise.

### GetXyzRolesOk

`func (o *DeploymentNeonPostgresDatabaseInfo) GetXyzRolesOk() (*[]XyzDeploymentNeonPostgresDatabaseInfoRoleObject, bool)`

GetXyzRolesOk returns a tuple with the XyzRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXyzRoles

`func (o *DeploymentNeonPostgresDatabaseInfo) SetXyzRoles(v []XyzDeploymentNeonPostgresDatabaseInfoRoleObject)`

SetXyzRoles sets XyzRoles field to given value.

### HasXyzRoles

`func (o *DeploymentNeonPostgresDatabaseInfo) HasXyzRoles() bool`

HasXyzRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


