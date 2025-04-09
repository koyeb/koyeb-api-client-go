# Budget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Thresholds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBudget

`func NewBudget() *Budget`

NewBudget instantiates a new Budget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithDefaults

`func NewBudgetWithDefaults() *Budget`

NewBudgetWithDefaults instantiates a new Budget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Budget) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Budget) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Budget) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Budget) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetThresholds

`func (o *Budget) GetThresholds() []string`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Budget) GetThresholdsOk() (*[]string, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Budget) SetThresholds(v []string)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Budget) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


