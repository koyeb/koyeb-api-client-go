# UpdateBudgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | In cents. | [optional] 

## Methods

### NewUpdateBudgetRequest

`func NewUpdateBudgetRequest() *UpdateBudgetRequest`

NewUpdateBudgetRequest instantiates a new UpdateBudgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBudgetRequestWithDefaults

`func NewUpdateBudgetRequestWithDefaults() *UpdateBudgetRequest`

NewUpdateBudgetRequestWithDefaults instantiates a new UpdateBudgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdateBudgetRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateBudgetRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateBudgetRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateBudgetRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


