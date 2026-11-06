# ServiceViewFacet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**ServiceViewFacetKind**](ServiceViewFacetKind.md) |  | [optional] [default to SERVICEVIEWFACETKIND_TYPE]
**Values** | Pointer to [**[]ServiceViewFacetValue**](ServiceViewFacetValue.md) |  | [optional] 

## Methods

### NewServiceViewFacet

`func NewServiceViewFacet() *ServiceViewFacet`

NewServiceViewFacet instantiates a new ServiceViewFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewFacetWithDefaults

`func NewServiceViewFacetWithDefaults() *ServiceViewFacet`

NewServiceViewFacetWithDefaults instantiates a new ServiceViewFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ServiceViewFacet) GetKind() ServiceViewFacetKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceViewFacet) GetKindOk() (*ServiceViewFacetKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceViewFacet) SetKind(v ServiceViewFacetKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServiceViewFacet) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetValues

`func (o *ServiceViewFacet) GetValues() []ServiceViewFacetValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceViewFacet) GetValuesOk() (*[]ServiceViewFacetValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceViewFacet) SetValues(v []ServiceViewFacetValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ServiceViewFacet) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


