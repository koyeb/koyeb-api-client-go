# SubscriptionPaymentFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAt** | Pointer to **time.Time** |  | [optional] 
**NextAttempt** | Pointer to **time.Time** |  | [optional] 
**AttemptCount** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorReason** | Pointer to **string** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**PaymentMethodRequired** | Pointer to **bool** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**StripeSdk** | Pointer to [**SubscriptionPaymentFailureStripeSDK**](SubscriptionPaymentFailureStripeSDK.md) |  | [optional] 

## Methods

### NewSubscriptionPaymentFailure

`func NewSubscriptionPaymentFailure() *SubscriptionPaymentFailure`

NewSubscriptionPaymentFailure instantiates a new SubscriptionPaymentFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPaymentFailureWithDefaults

`func NewSubscriptionPaymentFailureWithDefaults() *SubscriptionPaymentFailure`

NewSubscriptionPaymentFailureWithDefaults instantiates a new SubscriptionPaymentFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAt

`func (o *SubscriptionPaymentFailure) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *SubscriptionPaymentFailure) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *SubscriptionPaymentFailure) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *SubscriptionPaymentFailure) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetNextAttempt

`func (o *SubscriptionPaymentFailure) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *SubscriptionPaymentFailure) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *SubscriptionPaymentFailure) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *SubscriptionPaymentFailure) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### GetAttemptCount

`func (o *SubscriptionPaymentFailure) GetAttemptCount() string`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *SubscriptionPaymentFailure) GetAttemptCountOk() (*string, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *SubscriptionPaymentFailure) SetAttemptCount(v string)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *SubscriptionPaymentFailure) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *SubscriptionPaymentFailure) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SubscriptionPaymentFailure) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SubscriptionPaymentFailure) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SubscriptionPaymentFailure) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorReason

`func (o *SubscriptionPaymentFailure) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *SubscriptionPaymentFailure) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *SubscriptionPaymentFailure) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *SubscriptionPaymentFailure) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### GetErrorType

`func (o *SubscriptionPaymentFailure) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *SubscriptionPaymentFailure) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *SubscriptionPaymentFailure) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *SubscriptionPaymentFailure) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SubscriptionPaymentFailure) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SubscriptionPaymentFailure) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SubscriptionPaymentFailure) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SubscriptionPaymentFailure) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPaymentMethodRequired

`func (o *SubscriptionPaymentFailure) GetPaymentMethodRequired() bool`

GetPaymentMethodRequired returns the PaymentMethodRequired field if non-nil, zero value otherwise.

### GetPaymentMethodRequiredOk

`func (o *SubscriptionPaymentFailure) GetPaymentMethodRequiredOk() (*bool, bool)`

GetPaymentMethodRequiredOk returns a tuple with the PaymentMethodRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodRequired

`func (o *SubscriptionPaymentFailure) SetPaymentMethodRequired(v bool)`

SetPaymentMethodRequired sets PaymentMethodRequired field to given value.

### HasPaymentMethodRequired

`func (o *SubscriptionPaymentFailure) HasPaymentMethodRequired() bool`

HasPaymentMethodRequired returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SubscriptionPaymentFailure) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SubscriptionPaymentFailure) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SubscriptionPaymentFailure) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SubscriptionPaymentFailure) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetStripeSdk

`func (o *SubscriptionPaymentFailure) GetStripeSdk() SubscriptionPaymentFailureStripeSDK`

GetStripeSdk returns the StripeSdk field if non-nil, zero value otherwise.

### GetStripeSdkOk

`func (o *SubscriptionPaymentFailure) GetStripeSdkOk() (*SubscriptionPaymentFailureStripeSDK, bool)`

GetStripeSdkOk returns a tuple with the StripeSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSdk

`func (o *SubscriptionPaymentFailure) SetStripeSdk(v SubscriptionPaymentFailureStripeSDK)`

SetStripeSdk sets StripeSdk field to given value.

### HasStripeSdk

`func (o *SubscriptionPaymentFailure) HasStripeSdk() bool`

HasStripeSdk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


