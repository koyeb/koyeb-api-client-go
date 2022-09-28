# HTTPHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHTTPHeader

`func NewHTTPHeader() *HTTPHeader`

NewHTTPHeader instantiates a new HTTPHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPHeaderWithDefaults

`func NewHTTPHeaderWithDefaults() *HTTPHeader`

NewHTTPHeaderWithDefaults instantiates a new HTTPHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *HTTPHeader) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HTTPHeader) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HTTPHeader) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HTTPHeader) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValues

`func (o *HTTPHeader) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HTTPHeader) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HTTPHeader) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *HTTPHeader) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


