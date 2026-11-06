# ServiceViewGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to [**ServiceViewGroupBy**](ServiceViewGroupBy.md) |  | [optional] [default to SERVICEVIEWGROUPBY_NO_GROUP]
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewServiceViewGroup

`func NewServiceViewGroup() *ServiceViewGroup`

NewServiceViewGroup instantiates a new ServiceViewGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewGroupWithDefaults

`func NewServiceViewGroupWithDefaults() *ServiceViewGroup`

NewServiceViewGroupWithDefaults instantiates a new ServiceViewGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *ServiceViewGroup) GetBy() ServiceViewGroupBy`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *ServiceViewGroup) GetByOk() (*ServiceViewGroupBy, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *ServiceViewGroup) SetBy(v ServiceViewGroupBy)`

SetBy sets By field to given value.

### HasBy

`func (o *ServiceViewGroup) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetKey

`func (o *ServiceViewGroup) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ServiceViewGroup) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ServiceViewGroup) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ServiceViewGroup) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceViewGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceViewGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceViewGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceViewGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCount

`func (o *ServiceViewGroup) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServiceViewGroup) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServiceViewGroup) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ServiceViewGroup) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


