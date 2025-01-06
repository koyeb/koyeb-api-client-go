# NextInvoiceReplyDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NextInvoiceReplyDiscountType**](NextInvoiceReplyDiscountType.md) |  | [optional] [default to NEXTINVOICEREPLYDISCOUNTTYPE_PERCENT_OFF]
**Name** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 

## Methods

### NewNextInvoiceReplyDiscount

`func NewNextInvoiceReplyDiscount() *NextInvoiceReplyDiscount`

NewNextInvoiceReplyDiscount instantiates a new NextInvoiceReplyDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextInvoiceReplyDiscountWithDefaults

`func NewNextInvoiceReplyDiscountWithDefaults() *NextInvoiceReplyDiscount`

NewNextInvoiceReplyDiscountWithDefaults instantiates a new NextInvoiceReplyDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NextInvoiceReplyDiscount) GetType() NextInvoiceReplyDiscountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NextInvoiceReplyDiscount) GetTypeOk() (*NextInvoiceReplyDiscountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NextInvoiceReplyDiscount) SetType(v NextInvoiceReplyDiscountType)`

SetType sets Type field to given value.

### HasType

`func (o *NextInvoiceReplyDiscount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NextInvoiceReplyDiscount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextInvoiceReplyDiscount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextInvoiceReplyDiscount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextInvoiceReplyDiscount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAmount

`func (o *NextInvoiceReplyDiscount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NextInvoiceReplyDiscount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NextInvoiceReplyDiscount) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *NextInvoiceReplyDiscount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


