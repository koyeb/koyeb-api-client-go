# QueryServicesViewSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to [**QueryServicesViewSortField**](QueryServicesViewSortField.md) |  | [optional] [default to QUERYSERVICESVIEWSORTFIELD_CREATED_AT]
**Direction** | Pointer to [**QueryServicesViewSortDirection**](QueryServicesViewSortDirection.md) |  | [optional] [default to QUERYSERVICESVIEWSORTDIRECTION_ASC]

## Methods

### NewQueryServicesViewSort

`func NewQueryServicesViewSort() *QueryServicesViewSort`

NewQueryServicesViewSort instantiates a new QueryServicesViewSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesViewSortWithDefaults

`func NewQueryServicesViewSortWithDefaults() *QueryServicesViewSort`

NewQueryServicesViewSortWithDefaults instantiates a new QueryServicesViewSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *QueryServicesViewSort) GetField() QueryServicesViewSortField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QueryServicesViewSort) GetFieldOk() (*QueryServicesViewSortField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QueryServicesViewSort) SetField(v QueryServicesViewSortField)`

SetField sets Field field to given value.

### HasField

`func (o *QueryServicesViewSort) HasField() bool`

HasField returns a boolean if a field has been set.

### GetDirection

`func (o *QueryServicesViewSort) GetDirection() QueryServicesViewSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QueryServicesViewSort) GetDirectionOk() (*QueryServicesViewSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QueryServicesViewSort) SetDirection(v QueryServicesViewSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *QueryServicesViewSort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


