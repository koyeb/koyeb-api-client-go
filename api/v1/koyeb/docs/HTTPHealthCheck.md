# HTTPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**EnableSsl** | Pointer to **bool** |  | [optional] 
**SkipTls** | Pointer to **bool** |  | [optional] 
**Headers** | Pointer to [**[]HTTPHeader**](HTTPHeader.md) |  | [optional] 

## Methods

### NewHTTPHealthCheck

`func NewHTTPHealthCheck() *HTTPHealthCheck`

NewHTTPHealthCheck instantiates a new HTTPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPHealthCheckWithDefaults

`func NewHTTPHealthCheckWithDefaults() *HTTPHealthCheck`

NewHTTPHealthCheckWithDefaults instantiates a new HTTPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *HTTPHealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HTTPHealthCheck) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HTTPHealthCheck) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *HTTPHealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *HTTPHealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HTTPHealthCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HTTPHealthCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HTTPHealthCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *HTTPHealthCheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HTTPHealthCheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HTTPHealthCheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HTTPHealthCheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetEnableSsl

`func (o *HTTPHealthCheck) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *HTTPHealthCheck) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *HTTPHealthCheck) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *HTTPHealthCheck) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetSkipTls

`func (o *HTTPHealthCheck) GetSkipTls() bool`

GetSkipTls returns the SkipTls field if non-nil, zero value otherwise.

### GetSkipTlsOk

`func (o *HTTPHealthCheck) GetSkipTlsOk() (*bool, bool)`

GetSkipTlsOk returns a tuple with the SkipTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTls

`func (o *HTTPHealthCheck) SetSkipTls(v bool)`

SetSkipTls sets SkipTls field to given value.

### HasSkipTls

`func (o *HTTPHealthCheck) HasSkipTls() bool`

HasSkipTls returns a boolean if a field has been set.

### GetHeaders

`func (o *HTTPHealthCheck) GetHeaders() []HTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HTTPHealthCheck) GetHeadersOk() (*[]HTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HTTPHealthCheck) SetHeaders(v []HTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HTTPHealthCheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


