# GatewayRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Datacenters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGatewayRoute

`func NewGatewayRoute() *GatewayRoute`

NewGatewayRoute instantiates a new GatewayRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRouteWithDefaults

`func NewGatewayRouteWithDefaults() *GatewayRoute`

NewGatewayRouteWithDefaults instantiates a new GatewayRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GatewayRoute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GatewayRoute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GatewayRoute) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GatewayRoute) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *GatewayRoute) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *GatewayRoute) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *GatewayRoute) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *GatewayRoute) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetRevisionId

`func (o *GatewayRoute) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GatewayRoute) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GatewayRoute) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GatewayRoute) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetServiceId

`func (o *GatewayRoute) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GatewayRoute) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GatewayRoute) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GatewayRoute) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPort

`func (o *GatewayRoute) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayRoute) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayRoute) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayRoute) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDatacenters

`func (o *GatewayRoute) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *GatewayRoute) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *GatewayRoute) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *GatewayRoute) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


