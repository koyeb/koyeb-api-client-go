# DeploymentPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentPort

`func NewDeploymentPort() *DeploymentPort`

NewDeploymentPort instantiates a new DeploymentPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPortWithDefaults

`func NewDeploymentPortWithDefaults() *DeploymentPort`

NewDeploymentPortWithDefaults instantiates a new DeploymentPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *DeploymentPort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeploymentPort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeploymentPort) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeploymentPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *DeploymentPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DeploymentPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DeploymentPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DeploymentPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


