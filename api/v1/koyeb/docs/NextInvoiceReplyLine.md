# NextInvoiceReplyLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountExcludingTax** | Pointer to **int32** |  | [optional] 
**Period** | Pointer to [**NextInvoiceReplyLinePeriod**](NextInvoiceReplyLinePeriod.md) |  | [optional] 
**PlanNickname** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**NextInvoiceReplyLinePrice**](NextInvoiceReplyLinePrice.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewNextInvoiceReplyLine

`func NewNextInvoiceReplyLine() *NextInvoiceReplyLine`

NewNextInvoiceReplyLine instantiates a new NextInvoiceReplyLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextInvoiceReplyLineWithDefaults

`func NewNextInvoiceReplyLineWithDefaults() *NextInvoiceReplyLine`

NewNextInvoiceReplyLineWithDefaults instantiates a new NextInvoiceReplyLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountExcludingTax

`func (o *NextInvoiceReplyLine) GetAmountExcludingTax() int32`

GetAmountExcludingTax returns the AmountExcludingTax field if non-nil, zero value otherwise.

### GetAmountExcludingTaxOk

`func (o *NextInvoiceReplyLine) GetAmountExcludingTaxOk() (*int32, bool)`

GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountExcludingTax

`func (o *NextInvoiceReplyLine) SetAmountExcludingTax(v int32)`

SetAmountExcludingTax sets AmountExcludingTax field to given value.

### HasAmountExcludingTax

`func (o *NextInvoiceReplyLine) HasAmountExcludingTax() bool`

HasAmountExcludingTax returns a boolean if a field has been set.

### GetPeriod

`func (o *NextInvoiceReplyLine) GetPeriod() NextInvoiceReplyLinePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *NextInvoiceReplyLine) GetPeriodOk() (*NextInvoiceReplyLinePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *NextInvoiceReplyLine) SetPeriod(v NextInvoiceReplyLinePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *NextInvoiceReplyLine) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPlanNickname

`func (o *NextInvoiceReplyLine) GetPlanNickname() string`

GetPlanNickname returns the PlanNickname field if non-nil, zero value otherwise.

### GetPlanNicknameOk

`func (o *NextInvoiceReplyLine) GetPlanNicknameOk() (*string, bool)`

GetPlanNicknameOk returns a tuple with the PlanNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNickname

`func (o *NextInvoiceReplyLine) SetPlanNickname(v string)`

SetPlanNickname sets PlanNickname field to given value.

### HasPlanNickname

`func (o *NextInvoiceReplyLine) HasPlanNickname() bool`

HasPlanNickname returns a boolean if a field has been set.

### GetPrice

`func (o *NextInvoiceReplyLine) GetPrice() NextInvoiceReplyLinePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NextInvoiceReplyLine) GetPriceOk() (*NextInvoiceReplyLinePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NextInvoiceReplyLine) SetPrice(v NextInvoiceReplyLinePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NextInvoiceReplyLine) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *NextInvoiceReplyLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NextInvoiceReplyLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NextInvoiceReplyLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NextInvoiceReplyLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


