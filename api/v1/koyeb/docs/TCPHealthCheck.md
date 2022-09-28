# TCPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 

## Methods

### NewTCPHealthCheck

`func NewTCPHealthCheck() *TCPHealthCheck`

NewTCPHealthCheck instantiates a new TCPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTCPHealthCheckWithDefaults

`func NewTCPHealthCheckWithDefaults() *TCPHealthCheck`

NewTCPHealthCheckWithDefaults instantiates a new TCPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *TCPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TCPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TCPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *TCPHealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


