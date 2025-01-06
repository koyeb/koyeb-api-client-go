# NextInvoiceReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeInvoice** | Pointer to **map[string]interface{}** |  | [optional] 
**Lines** | Pointer to [**[]NextInvoiceReplyLine**](NextInvoiceReplyLine.md) |  | [optional] 
**Discounts** | Pointer to [**[]NextInvoiceReplyDiscount**](NextInvoiceReplyDiscount.md) |  | [optional] 

## Methods

### NewNextInvoiceReply

`func NewNextInvoiceReply() *NextInvoiceReply`

NewNextInvoiceReply instantiates a new NextInvoiceReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextInvoiceReplyWithDefaults

`func NewNextInvoiceReplyWithDefaults() *NextInvoiceReply`

NewNextInvoiceReplyWithDefaults instantiates a new NextInvoiceReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeInvoice

`func (o *NextInvoiceReply) GetStripeInvoice() map[string]interface{}`

GetStripeInvoice returns the StripeInvoice field if non-nil, zero value otherwise.

### GetStripeInvoiceOk

`func (o *NextInvoiceReply) GetStripeInvoiceOk() (*map[string]interface{}, bool)`

GetStripeInvoiceOk returns a tuple with the StripeInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoice

`func (o *NextInvoiceReply) SetStripeInvoice(v map[string]interface{})`

SetStripeInvoice sets StripeInvoice field to given value.

### HasStripeInvoice

`func (o *NextInvoiceReply) HasStripeInvoice() bool`

HasStripeInvoice returns a boolean if a field has been set.

### GetLines

`func (o *NextInvoiceReply) GetLines() []NextInvoiceReplyLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *NextInvoiceReply) GetLinesOk() (*[]NextInvoiceReplyLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *NextInvoiceReply) SetLines(v []NextInvoiceReplyLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *NextInvoiceReply) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetDiscounts

`func (o *NextInvoiceReply) GetDiscounts() []NextInvoiceReplyDiscount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *NextInvoiceReply) GetDiscountsOk() (*[]NextInvoiceReplyDiscount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *NextInvoiceReply) SetDiscounts(v []NextInvoiceReplyDiscount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *NextInvoiceReply) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


