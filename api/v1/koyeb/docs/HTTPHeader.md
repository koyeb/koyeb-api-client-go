# HTTPHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

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

### GetValue

`func (o *HTTPHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HTTPHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HTTPHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HTTPHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


