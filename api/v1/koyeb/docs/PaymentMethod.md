# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  | [optional] [default to PAYMENTMETHODSTATUS_INVALID]
**Messages** | Pointer to **[]string** |  | [optional] 
**StripePaymentMethodId** | Pointer to **string** |  | [optional] 
**AuthorizationVerifiedAt** | Pointer to **time.Time** |  | [optional] 
**AuthorizationCanceledAt** | Pointer to **time.Time** |  | [optional] 
**AuthorizationStripePaymentIntentId** | Pointer to **string** |  | [optional] 
**AuthorizationStripePaymentIntentClientSecret** | Pointer to **string** |  | [optional] 
**CardBrand** | Pointer to **string** |  | [optional] 
**CardCountry** | Pointer to **string** |  | [optional] 
**CardFunding** | Pointer to **string** |  | [optional] 
**CardFingerprint** | Pointer to **string** |  | [optional] 
**CardLastDigits** | Pointer to **string** |  | [optional] 
**CardExpirationMonth** | Pointer to **int64** |  | [optional] 
**CardExpirationYear** | Pointer to **int64** |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentMethod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentMethod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentMethod) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethod) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentMethod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *PaymentMethod) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PaymentMethod) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PaymentMethod) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PaymentMethod) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PaymentMethod) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PaymentMethod) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PaymentMethod) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PaymentMethod) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentMethod) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethod) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethod) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentMethod) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *PaymentMethod) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PaymentMethod) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PaymentMethod) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *PaymentMethod) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetStripePaymentMethodId

`func (o *PaymentMethod) GetStripePaymentMethodId() string`

GetStripePaymentMethodId returns the StripePaymentMethodId field if non-nil, zero value otherwise.

### GetStripePaymentMethodIdOk

`func (o *PaymentMethod) GetStripePaymentMethodIdOk() (*string, bool)`

GetStripePaymentMethodIdOk returns a tuple with the StripePaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePaymentMethodId

`func (o *PaymentMethod) SetStripePaymentMethodId(v string)`

SetStripePaymentMethodId sets StripePaymentMethodId field to given value.

### HasStripePaymentMethodId

`func (o *PaymentMethod) HasStripePaymentMethodId() bool`

HasStripePaymentMethodId returns a boolean if a field has been set.

### GetAuthorizationVerifiedAt

`func (o *PaymentMethod) GetAuthorizationVerifiedAt() time.Time`

GetAuthorizationVerifiedAt returns the AuthorizationVerifiedAt field if non-nil, zero value otherwise.

### GetAuthorizationVerifiedAtOk

`func (o *PaymentMethod) GetAuthorizationVerifiedAtOk() (*time.Time, bool)`

GetAuthorizationVerifiedAtOk returns a tuple with the AuthorizationVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationVerifiedAt

`func (o *PaymentMethod) SetAuthorizationVerifiedAt(v time.Time)`

SetAuthorizationVerifiedAt sets AuthorizationVerifiedAt field to given value.

### HasAuthorizationVerifiedAt

`func (o *PaymentMethod) HasAuthorizationVerifiedAt() bool`

HasAuthorizationVerifiedAt returns a boolean if a field has been set.

### GetAuthorizationCanceledAt

`func (o *PaymentMethod) GetAuthorizationCanceledAt() time.Time`

GetAuthorizationCanceledAt returns the AuthorizationCanceledAt field if non-nil, zero value otherwise.

### GetAuthorizationCanceledAtOk

`func (o *PaymentMethod) GetAuthorizationCanceledAtOk() (*time.Time, bool)`

GetAuthorizationCanceledAtOk returns a tuple with the AuthorizationCanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCanceledAt

`func (o *PaymentMethod) SetAuthorizationCanceledAt(v time.Time)`

SetAuthorizationCanceledAt sets AuthorizationCanceledAt field to given value.

### HasAuthorizationCanceledAt

`func (o *PaymentMethod) HasAuthorizationCanceledAt() bool`

HasAuthorizationCanceledAt returns a boolean if a field has been set.

### GetAuthorizationStripePaymentIntentId

`func (o *PaymentMethod) GetAuthorizationStripePaymentIntentId() string`

GetAuthorizationStripePaymentIntentId returns the AuthorizationStripePaymentIntentId field if non-nil, zero value otherwise.

### GetAuthorizationStripePaymentIntentIdOk

`func (o *PaymentMethod) GetAuthorizationStripePaymentIntentIdOk() (*string, bool)`

GetAuthorizationStripePaymentIntentIdOk returns a tuple with the AuthorizationStripePaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStripePaymentIntentId

`func (o *PaymentMethod) SetAuthorizationStripePaymentIntentId(v string)`

SetAuthorizationStripePaymentIntentId sets AuthorizationStripePaymentIntentId field to given value.

### HasAuthorizationStripePaymentIntentId

`func (o *PaymentMethod) HasAuthorizationStripePaymentIntentId() bool`

HasAuthorizationStripePaymentIntentId returns a boolean if a field has been set.

### GetAuthorizationStripePaymentIntentClientSecret

`func (o *PaymentMethod) GetAuthorizationStripePaymentIntentClientSecret() string`

GetAuthorizationStripePaymentIntentClientSecret returns the AuthorizationStripePaymentIntentClientSecret field if non-nil, zero value otherwise.

### GetAuthorizationStripePaymentIntentClientSecretOk

`func (o *PaymentMethod) GetAuthorizationStripePaymentIntentClientSecretOk() (*string, bool)`

GetAuthorizationStripePaymentIntentClientSecretOk returns a tuple with the AuthorizationStripePaymentIntentClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStripePaymentIntentClientSecret

`func (o *PaymentMethod) SetAuthorizationStripePaymentIntentClientSecret(v string)`

SetAuthorizationStripePaymentIntentClientSecret sets AuthorizationStripePaymentIntentClientSecret field to given value.

### HasAuthorizationStripePaymentIntentClientSecret

`func (o *PaymentMethod) HasAuthorizationStripePaymentIntentClientSecret() bool`

HasAuthorizationStripePaymentIntentClientSecret returns a boolean if a field has been set.

### GetCardBrand

`func (o *PaymentMethod) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PaymentMethod) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PaymentMethod) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *PaymentMethod) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCountry

`func (o *PaymentMethod) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *PaymentMethod) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *PaymentMethod) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *PaymentMethod) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardFunding

`func (o *PaymentMethod) GetCardFunding() string`

GetCardFunding returns the CardFunding field if non-nil, zero value otherwise.

### GetCardFundingOk

`func (o *PaymentMethod) GetCardFundingOk() (*string, bool)`

GetCardFundingOk returns a tuple with the CardFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFunding

`func (o *PaymentMethod) SetCardFunding(v string)`

SetCardFunding sets CardFunding field to given value.

### HasCardFunding

`func (o *PaymentMethod) HasCardFunding() bool`

HasCardFunding returns a boolean if a field has been set.

### GetCardFingerprint

`func (o *PaymentMethod) GetCardFingerprint() string`

GetCardFingerprint returns the CardFingerprint field if non-nil, zero value otherwise.

### GetCardFingerprintOk

`func (o *PaymentMethod) GetCardFingerprintOk() (*string, bool)`

GetCardFingerprintOk returns a tuple with the CardFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFingerprint

`func (o *PaymentMethod) SetCardFingerprint(v string)`

SetCardFingerprint sets CardFingerprint field to given value.

### HasCardFingerprint

`func (o *PaymentMethod) HasCardFingerprint() bool`

HasCardFingerprint returns a boolean if a field has been set.

### GetCardLastDigits

`func (o *PaymentMethod) GetCardLastDigits() string`

GetCardLastDigits returns the CardLastDigits field if non-nil, zero value otherwise.

### GetCardLastDigitsOk

`func (o *PaymentMethod) GetCardLastDigitsOk() (*string, bool)`

GetCardLastDigitsOk returns a tuple with the CardLastDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastDigits

`func (o *PaymentMethod) SetCardLastDigits(v string)`

SetCardLastDigits sets CardLastDigits field to given value.

### HasCardLastDigits

`func (o *PaymentMethod) HasCardLastDigits() bool`

HasCardLastDigits returns a boolean if a field has been set.

### GetCardExpirationMonth

`func (o *PaymentMethod) GetCardExpirationMonth() int64`

GetCardExpirationMonth returns the CardExpirationMonth field if non-nil, zero value otherwise.

### GetCardExpirationMonthOk

`func (o *PaymentMethod) GetCardExpirationMonthOk() (*int64, bool)`

GetCardExpirationMonthOk returns a tuple with the CardExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationMonth

`func (o *PaymentMethod) SetCardExpirationMonth(v int64)`

SetCardExpirationMonth sets CardExpirationMonth field to given value.

### HasCardExpirationMonth

`func (o *PaymentMethod) HasCardExpirationMonth() bool`

HasCardExpirationMonth returns a boolean if a field has been set.

### GetCardExpirationYear

`func (o *PaymentMethod) GetCardExpirationYear() int64`

GetCardExpirationYear returns the CardExpirationYear field if non-nil, zero value otherwise.

### GetCardExpirationYearOk

`func (o *PaymentMethod) GetCardExpirationYearOk() (*int64, bool)`

GetCardExpirationYearOk returns a tuple with the CardExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpirationYear

`func (o *PaymentMethod) SetCardExpirationYear(v int64)`

SetCardExpirationYear sets CardExpirationYear field to given value.

### HasCardExpirationYear

`func (o *PaymentMethod) HasCardExpirationYear() bool`

HasCardExpirationYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


