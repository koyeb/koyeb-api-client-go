# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**StripeSubscriptionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] [default to SUBSCRIPTIONSTATUS_INVALID]
**Messages** | Pointer to **[]string** |  | [optional] 
**HasPendingUpdate** | Pointer to **bool** |  | [optional] 
**StripePendingInvoiceId** | Pointer to **string** |  | [optional] 
**TerminateAt** | Pointer to **time.Time** |  | [optional] 
**CanceledAt** | Pointer to **time.Time** |  | [optional] 
**TerminatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrentPeriodStart** | Pointer to **time.Time** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**AmountPayable** | Pointer to **string** |  | [optional] 
**AmountPaid** | Pointer to **string** |  | [optional] 
**AmountRemaining** | Pointer to **string** |  | [optional] 
**PaymentFailure** | Pointer to [**SubscriptionPaymentFailure**](SubscriptionPaymentFailure.md) |  | [optional] 
**Trialing** | Pointer to **bool** |  | [optional] 
**TrialEndsAt** | Pointer to **time.Time** |  | [optional] 
**TrialMaxSpend** | Pointer to **string** |  | [optional] 
**CurrentSpend** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Subscription) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Subscription) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Subscription) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Subscription) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Subscription) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Subscription) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Subscription) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Subscription) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *Subscription) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *Subscription) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *Subscription) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *Subscription) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *Subscription) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Subscription) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Subscription) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Subscription) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetHasPendingUpdate

`func (o *Subscription) GetHasPendingUpdate() bool`

GetHasPendingUpdate returns the HasPendingUpdate field if non-nil, zero value otherwise.

### GetHasPendingUpdateOk

`func (o *Subscription) GetHasPendingUpdateOk() (*bool, bool)`

GetHasPendingUpdateOk returns a tuple with the HasPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingUpdate

`func (o *Subscription) SetHasPendingUpdate(v bool)`

SetHasPendingUpdate sets HasPendingUpdate field to given value.

### HasHasPendingUpdate

`func (o *Subscription) HasHasPendingUpdate() bool`

HasHasPendingUpdate returns a boolean if a field has been set.

### GetStripePendingInvoiceId

`func (o *Subscription) GetStripePendingInvoiceId() string`

GetStripePendingInvoiceId returns the StripePendingInvoiceId field if non-nil, zero value otherwise.

### GetStripePendingInvoiceIdOk

`func (o *Subscription) GetStripePendingInvoiceIdOk() (*string, bool)`

GetStripePendingInvoiceIdOk returns a tuple with the StripePendingInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePendingInvoiceId

`func (o *Subscription) SetStripePendingInvoiceId(v string)`

SetStripePendingInvoiceId sets StripePendingInvoiceId field to given value.

### HasStripePendingInvoiceId

`func (o *Subscription) HasStripePendingInvoiceId() bool`

HasStripePendingInvoiceId returns a boolean if a field has been set.

### GetTerminateAt

`func (o *Subscription) GetTerminateAt() time.Time`

GetTerminateAt returns the TerminateAt field if non-nil, zero value otherwise.

### GetTerminateAtOk

`func (o *Subscription) GetTerminateAtOk() (*time.Time, bool)`

GetTerminateAtOk returns a tuple with the TerminateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateAt

`func (o *Subscription) SetTerminateAt(v time.Time)`

SetTerminateAt sets TerminateAt field to given value.

### HasTerminateAt

`func (o *Subscription) HasTerminateAt() bool`

HasTerminateAt returns a boolean if a field has been set.

### GetCanceledAt

`func (o *Subscription) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *Subscription) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *Subscription) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *Subscription) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *Subscription) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *Subscription) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *Subscription) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *Subscription) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *Subscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *Subscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrency

`func (o *Subscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Subscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Subscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Subscription) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmountPayable

`func (o *Subscription) GetAmountPayable() string`

GetAmountPayable returns the AmountPayable field if non-nil, zero value otherwise.

### GetAmountPayableOk

`func (o *Subscription) GetAmountPayableOk() (*string, bool)`

GetAmountPayableOk returns a tuple with the AmountPayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPayable

`func (o *Subscription) SetAmountPayable(v string)`

SetAmountPayable sets AmountPayable field to given value.

### HasAmountPayable

`func (o *Subscription) HasAmountPayable() bool`

HasAmountPayable returns a boolean if a field has been set.

### GetAmountPaid

`func (o *Subscription) GetAmountPaid() string`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *Subscription) GetAmountPaidOk() (*string, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *Subscription) SetAmountPaid(v string)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *Subscription) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetAmountRemaining

`func (o *Subscription) GetAmountRemaining() string`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *Subscription) GetAmountRemainingOk() (*string, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *Subscription) SetAmountRemaining(v string)`

SetAmountRemaining sets AmountRemaining field to given value.

### HasAmountRemaining

`func (o *Subscription) HasAmountRemaining() bool`

HasAmountRemaining returns a boolean if a field has been set.

### GetPaymentFailure

`func (o *Subscription) GetPaymentFailure() SubscriptionPaymentFailure`

GetPaymentFailure returns the PaymentFailure field if non-nil, zero value otherwise.

### GetPaymentFailureOk

`func (o *Subscription) GetPaymentFailureOk() (*SubscriptionPaymentFailure, bool)`

GetPaymentFailureOk returns a tuple with the PaymentFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFailure

`func (o *Subscription) SetPaymentFailure(v SubscriptionPaymentFailure)`

SetPaymentFailure sets PaymentFailure field to given value.

### HasPaymentFailure

`func (o *Subscription) HasPaymentFailure() bool`

HasPaymentFailure returns a boolean if a field has been set.

### GetTrialing

`func (o *Subscription) GetTrialing() bool`

GetTrialing returns the Trialing field if non-nil, zero value otherwise.

### GetTrialingOk

`func (o *Subscription) GetTrialingOk() (*bool, bool)`

GetTrialingOk returns a tuple with the Trialing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialing

`func (o *Subscription) SetTrialing(v bool)`

SetTrialing sets Trialing field to given value.

### HasTrialing

`func (o *Subscription) HasTrialing() bool`

HasTrialing returns a boolean if a field has been set.

### GetTrialEndsAt

`func (o *Subscription) GetTrialEndsAt() time.Time`

GetTrialEndsAt returns the TrialEndsAt field if non-nil, zero value otherwise.

### GetTrialEndsAtOk

`func (o *Subscription) GetTrialEndsAtOk() (*time.Time, bool)`

GetTrialEndsAtOk returns a tuple with the TrialEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndsAt

`func (o *Subscription) SetTrialEndsAt(v time.Time)`

SetTrialEndsAt sets TrialEndsAt field to given value.

### HasTrialEndsAt

`func (o *Subscription) HasTrialEndsAt() bool`

HasTrialEndsAt returns a boolean if a field has been set.

### GetTrialMaxSpend

`func (o *Subscription) GetTrialMaxSpend() string`

GetTrialMaxSpend returns the TrialMaxSpend field if non-nil, zero value otherwise.

### GetTrialMaxSpendOk

`func (o *Subscription) GetTrialMaxSpendOk() (*string, bool)`

GetTrialMaxSpendOk returns a tuple with the TrialMaxSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialMaxSpend

`func (o *Subscription) SetTrialMaxSpend(v string)`

SetTrialMaxSpend sets TrialMaxSpend field to given value.

### HasTrialMaxSpend

`func (o *Subscription) HasTrialMaxSpend() bool`

HasTrialMaxSpend returns a boolean if a field has been set.

### GetCurrentSpend

`func (o *Subscription) GetCurrentSpend() string`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *Subscription) GetCurrentSpendOk() (*string, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *Subscription) SetCurrentSpend(v string)`

SetCurrentSpend sets CurrentSpend field to given value.

### HasCurrentSpend

`func (o *Subscription) HasCurrentSpend() bool`

HasCurrentSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


