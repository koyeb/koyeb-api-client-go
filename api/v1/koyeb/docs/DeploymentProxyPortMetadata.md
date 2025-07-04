# DeploymentProxyPortMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**PublicPort** | Pointer to **int64** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to [**ProxyPortProtocol**](ProxyPortProtocol.md) |  | [optional] [default to PROXYPORTPROTOCOL_TCP]

## Methods

### NewDeploymentProxyPortMetadata

`func NewDeploymentProxyPortMetadata() *DeploymentProxyPortMetadata`

NewDeploymentProxyPortMetadata instantiates a new DeploymentProxyPortMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProxyPortMetadataWithDefaults

`func NewDeploymentProxyPortMetadataWithDefaults() *DeploymentProxyPortMetadata`

NewDeploymentProxyPortMetadataWithDefaults instantiates a new DeploymentProxyPortMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DeploymentProxyPortMetadata) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DeploymentProxyPortMetadata) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DeploymentProxyPortMetadata) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DeploymentProxyPortMetadata) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPublicPort

`func (o *DeploymentProxyPortMetadata) GetPublicPort() int64`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *DeploymentProxyPortMetadata) GetPublicPortOk() (*int64, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *DeploymentProxyPortMetadata) SetPublicPort(v int64)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *DeploymentProxyPortMetadata) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetPort

`func (o *DeploymentProxyPortMetadata) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeploymentProxyPortMetadata) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeploymentProxyPortMetadata) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeploymentProxyPortMetadata) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *DeploymentProxyPortMetadata) GetProtocol() ProxyPortProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DeploymentProxyPortMetadata) GetProtocolOk() (*ProxyPortProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DeploymentProxyPortMetadata) SetProtocol(v ProxyPortProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DeploymentProxyPortMetadata) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


