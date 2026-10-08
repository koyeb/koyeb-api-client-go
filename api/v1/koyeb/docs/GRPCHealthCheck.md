# GRPCHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewGRPCHealthCheck

`func NewGRPCHealthCheck() *GRPCHealthCheck`

NewGRPCHealthCheck instantiates a new GRPCHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGRPCHealthCheckWithDefaults

`func NewGRPCHealthCheckWithDefaults() *GRPCHealthCheck`

NewGRPCHealthCheckWithDefaults instantiates a new GRPCHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *GRPCHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GRPCHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GRPCHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GRPCHealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetService

`func (o *GRPCHealthCheck) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GRPCHealthCheck) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GRPCHealthCheck) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GRPCHealthCheck) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


