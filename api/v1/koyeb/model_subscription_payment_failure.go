/*
Koyeb Rest API

Welcome to the Koyeb API documentation.  The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. The API has predictable resource-oriented URLs, allowing you to request, retrieve and execute actions on resources efficiently.  The API returns JSON-encoded responses and uses standard HTTP response codes, and verbs.  All of the functionality available in the Koyeb Control Panel is also available via the API, allowing you to automate repetitive and complex actions as required.  All production API requests are made to:  - https://app.koyeb.com/<version>  All data is sent and received as JSON.  ## Constructing Requests  Requests are made of three components:  - Base URL: `https://app.koyeb.com` - API Version: `v1` - Resource path: `apps`  To construct a proper request, you need to format the URL as follows:  `https://app.koyeb.com/{version}/{resource}`  Example: `https://app.koyeb.com/v1/apps`  The following code is an example request to retrieve the list of Apps.  ``` $ curl -H 'Authorization: Bearer 1111-1111-1111-1111-1111' https://app.koyeb.com/v1/apps {   \"apps\": [     {       \"id\": \"e267e09d-a2ef-4672-95c7-bf88d787a274\",       \"name\": \"demo\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-04-26T08:28:59.221274Z\",       \"created_at\": \"2021-04-26T08:28:59.221274Z\",       \"domains\": [         {           \"name\": \"demo-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     },     {       \"id\": \"e38caea5-5acb-4fec-aa65-81749546b470\",       \"name\": \"www\",       \"organization_id\": \"c402bcad-d06d-4fd6-8af2-08b790c0e249\",       \"updated_at\": \"2021-05-04T08:56:40.557792Z\",       \"created_at\": \"2021-05-04T08:56:40.557792Z\",       \"domains\": [         {           \"name\": \"www-koyeb.koyeb.app\",           \"deployment_group\": \"prod\"         }       ]     }   ],   \"limit\": 20,   \"offset\": 0,   \"count\": 2 } ```  # Authentication  The Koyeb API uses API keys to authenticate requests. You can view and manage your API keys in the Koyeb Control Panel.  Your API keys carry many privileges. Take care to keep them secure! Do not share your secret API keys in publicly accessible areas such as GitHub, Slack, client-side code, and so forth.  To make an authenticated request, include a bearer-type Authorization header containing your API key, prefixed with `Bearer `:  ``` Authorization: Bearer <your-api-key> ```  All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.  # Errors  We use conventional HTTP response codes to indicate the success or failure of an API request.  In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error that resulted from the provided information (e.g., a required parameter is missing, resource not found), and codes in the 5xx range indicate an error with our servers.  ## HTTP Status Code Summary  - `200 OK` - Everything worked as expected. - `400 Bad Request` - Often missing a required parameter. - `401 Unauthorized` - No valid API key provided. - `402 Request Failed` - Parameters were valid, but the request failed. - `403 Forbidden` - Insufficient privileges to access the requested resource. - `404 Not Found` - The requested item doesn't exist. - `50x` Server errors - Something went wrong on the API domain's end.  Input and output data must be valid JSON with the appropriate Content-Type header set.  ## Attributes  - code:    - `invalid_argument`: Occur when your request has invalid parameters.   - `unknown_resource`: Occur when the resource doesn't exist.   - `authentication_error`: Arise when there is a problem with authentication.   - `rate_limit_error`: Too many requests hit the API too quickly.   - `authorization_required`: You don't have sufficient right to access the resource.   - `api_error`: API errors use in case of a problem with api domain's servers (extremely uncommon).  - message:    - A human-readable error giving more details about the error  - status:    - The HTTP code returned by the server.  - fields (Optional):    - An array of parameters with a human-readable message giving more details about the error.  ## Errors responses example  Response 400 (application/json)  ``` {   \"status\": 400,   \"code\": \"invalid_argument\",   \"message\": \"Invalid argument\",   \"fields\": [     {       \"field\": \"definition.docker.image\",       \"description\": \"cannot be blank\"     }   ] } ```  Response 401 (application/json)  ``` {   \"status\": 401,   \"code\": \"authentication_error\",   \"message\": \"Request unauthenticated with bearer\" } ```  Response 403 (application/json)  ``` {   \"status\": 403,   \"code\": \"authentication_required\",   \"message\": \"You don't have the requested permission to access to resource.\" } ```   Response 404 (application/json)  ``` {   \"status\": 404,   \"code\": \"unknown_resource\",   \"message\": \"App not found\" } ```  # Pagination  All top-level API resources receive a paginated response. For instance, list Apps, list Services, and others share a standard structure, taking two parameters: limit and offset.  **Paginated request**  Requests against paginated endpoints accept two query string arguments:  * *limit*, a positive integer lower or equal to 100 to select the number of items to display. * *offset*, an integer to indicate the position to start at.  Paginated endpoints usually also accept filters to search and sort results. These filters are documented along with each endpoint documentation. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// SubscriptionPaymentFailure struct for SubscriptionPaymentFailure
type SubscriptionPaymentFailure struct {
	FailedAt *time.Time `json:"failed_at,omitempty"`
	NextAttempt *time.Time `json:"next_attempt,omitempty"`
	AttemptCount *string `json:"attempt_count,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	ErrorReason *string `json:"error_reason,omitempty"`
	ErrorType *string `json:"error_type,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	PaymentMethodRequired *bool `json:"payment_method_required,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	StripeSdk *SubscriptionPaymentFailureStripeSDK `json:"stripe_sdk,omitempty"`
}

// NewSubscriptionPaymentFailure instantiates a new SubscriptionPaymentFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPaymentFailure() *SubscriptionPaymentFailure {
	this := SubscriptionPaymentFailure{}
	return &this
}

// NewSubscriptionPaymentFailureWithDefaults instantiates a new SubscriptionPaymentFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPaymentFailureWithDefaults() *SubscriptionPaymentFailure {
	this := SubscriptionPaymentFailure{}
	return &this
}

// GetFailedAt returns the FailedAt field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetFailedAt() time.Time {
	if o == nil || isNil(o.FailedAt) {
		var ret time.Time
		return ret
	}
	return *o.FailedAt
}

// GetFailedAtOk returns a tuple with the FailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetFailedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FailedAt) {
    return nil, false
	}
	return o.FailedAt, true
}

// HasFailedAt returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasFailedAt() bool {
	if o != nil && !isNil(o.FailedAt) {
		return true
	}

	return false
}

// SetFailedAt gets a reference to the given time.Time and assigns it to the FailedAt field.
func (o *SubscriptionPaymentFailure) SetFailedAt(v time.Time) {
	o.FailedAt = &v
}

// GetNextAttempt returns the NextAttempt field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetNextAttempt() time.Time {
	if o == nil || isNil(o.NextAttempt) {
		var ret time.Time
		return ret
	}
	return *o.NextAttempt
}

// GetNextAttemptOk returns a tuple with the NextAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetNextAttemptOk() (*time.Time, bool) {
	if o == nil || isNil(o.NextAttempt) {
    return nil, false
	}
	return o.NextAttempt, true
}

// HasNextAttempt returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasNextAttempt() bool {
	if o != nil && !isNil(o.NextAttempt) {
		return true
	}

	return false
}

// SetNextAttempt gets a reference to the given time.Time and assigns it to the NextAttempt field.
func (o *SubscriptionPaymentFailure) SetNextAttempt(v time.Time) {
	o.NextAttempt = &v
}

// GetAttemptCount returns the AttemptCount field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetAttemptCount() string {
	if o == nil || isNil(o.AttemptCount) {
		var ret string
		return ret
	}
	return *o.AttemptCount
}

// GetAttemptCountOk returns a tuple with the AttemptCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetAttemptCountOk() (*string, bool) {
	if o == nil || isNil(o.AttemptCount) {
    return nil, false
	}
	return o.AttemptCount, true
}

// HasAttemptCount returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasAttemptCount() bool {
	if o != nil && !isNil(o.AttemptCount) {
		return true
	}

	return false
}

// SetAttemptCount gets a reference to the given string and assigns it to the AttemptCount field.
func (o *SubscriptionPaymentFailure) SetAttemptCount(v string) {
	o.AttemptCount = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetErrorCode() string {
	if o == nil || isNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetErrorCodeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorCode) {
    return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasErrorCode() bool {
	if o != nil && !isNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *SubscriptionPaymentFailure) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetErrorReason() string {
	if o == nil || isNil(o.ErrorReason) {
		var ret string
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetErrorReasonOk() (*string, bool) {
	if o == nil || isNil(o.ErrorReason) {
    return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasErrorReason() bool {
	if o != nil && !isNil(o.ErrorReason) {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given string and assigns it to the ErrorReason field.
func (o *SubscriptionPaymentFailure) SetErrorReason(v string) {
	o.ErrorReason = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetErrorType() string {
	if o == nil || isNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetErrorTypeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorType) {
    return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasErrorType() bool {
	if o != nil && !isNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *SubscriptionPaymentFailure) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SubscriptionPaymentFailure) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPaymentMethodRequired returns the PaymentMethodRequired field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetPaymentMethodRequired() bool {
	if o == nil || isNil(o.PaymentMethodRequired) {
		var ret bool
		return ret
	}
	return *o.PaymentMethodRequired
}

// GetPaymentMethodRequiredOk returns a tuple with the PaymentMethodRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetPaymentMethodRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.PaymentMethodRequired) {
    return nil, false
	}
	return o.PaymentMethodRequired, true
}

// HasPaymentMethodRequired returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasPaymentMethodRequired() bool {
	if o != nil && !isNil(o.PaymentMethodRequired) {
		return true
	}

	return false
}

// SetPaymentMethodRequired gets a reference to the given bool and assigns it to the PaymentMethodRequired field.
func (o *SubscriptionPaymentFailure) SetPaymentMethodRequired(v bool) {
	o.PaymentMethodRequired = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetRedirectUrl() string {
	if o == nil || isNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetRedirectUrlOk() (*string, bool) {
	if o == nil || isNil(o.RedirectUrl) {
    return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasRedirectUrl() bool {
	if o != nil && !isNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *SubscriptionPaymentFailure) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetStripeSdk returns the StripeSdk field value if set, zero value otherwise.
func (o *SubscriptionPaymentFailure) GetStripeSdk() SubscriptionPaymentFailureStripeSDK {
	if o == nil || isNil(o.StripeSdk) {
		var ret SubscriptionPaymentFailureStripeSDK
		return ret
	}
	return *o.StripeSdk
}

// GetStripeSdkOk returns a tuple with the StripeSdk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPaymentFailure) GetStripeSdkOk() (*SubscriptionPaymentFailureStripeSDK, bool) {
	if o == nil || isNil(o.StripeSdk) {
    return nil, false
	}
	return o.StripeSdk, true
}

// HasStripeSdk returns a boolean if a field has been set.
func (o *SubscriptionPaymentFailure) HasStripeSdk() bool {
	if o != nil && !isNil(o.StripeSdk) {
		return true
	}

	return false
}

// SetStripeSdk gets a reference to the given SubscriptionPaymentFailureStripeSDK and assigns it to the StripeSdk field.
func (o *SubscriptionPaymentFailure) SetStripeSdk(v SubscriptionPaymentFailureStripeSDK) {
	o.StripeSdk = &v
}

func (o SubscriptionPaymentFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FailedAt) {
		toSerialize["failed_at"] = o.FailedAt
	}
	if !isNil(o.NextAttempt) {
		toSerialize["next_attempt"] = o.NextAttempt
	}
	if !isNil(o.AttemptCount) {
		toSerialize["attempt_count"] = o.AttemptCount
	}
	if !isNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !isNil(o.ErrorReason) {
		toSerialize["error_reason"] = o.ErrorReason
	}
	if !isNil(o.ErrorType) {
		toSerialize["error_type"] = o.ErrorType
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !isNil(o.PaymentMethodRequired) {
		toSerialize["payment_method_required"] = o.PaymentMethodRequired
	}
	if !isNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !isNil(o.StripeSdk) {
		toSerialize["stripe_sdk"] = o.StripeSdk
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionPaymentFailure struct {
	value *SubscriptionPaymentFailure
	isSet bool
}

func (v NullableSubscriptionPaymentFailure) Get() *SubscriptionPaymentFailure {
	return v.value
}

func (v *NullableSubscriptionPaymentFailure) Set(val *SubscriptionPaymentFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPaymentFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPaymentFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPaymentFailure(val *SubscriptionPaymentFailure) *NullableSubscriptionPaymentFailure {
	return &NullableSubscriptionPaymentFailure{value: val, isSet: true}
}

func (v NullableSubscriptionPaymentFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPaymentFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


