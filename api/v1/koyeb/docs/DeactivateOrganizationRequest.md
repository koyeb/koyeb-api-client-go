# DeactivateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipConfirmation** | Pointer to **bool** | if set to true, skip_confirmation will directly start the deactivation process, without sending a confirmation email beforehand. | [optional] 

## Methods

### NewDeactivateOrganizationRequest

`func NewDeactivateOrganizationRequest() *DeactivateOrganizationRequest`

NewDeactivateOrganizationRequest instantiates a new DeactivateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateOrganizationRequestWithDefaults

`func NewDeactivateOrganizationRequestWithDefaults() *DeactivateOrganizationRequest`

NewDeactivateOrganizationRequestWithDefaults instantiates a new DeactivateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipConfirmation

`func (o *DeactivateOrganizationRequest) GetSkipConfirmation() bool`

GetSkipConfirmation returns the SkipConfirmation field if non-nil, zero value otherwise.

### GetSkipConfirmationOk

`func (o *DeactivateOrganizationRequest) GetSkipConfirmationOk() (*bool, bool)`

GetSkipConfirmationOk returns a tuple with the SkipConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmation

`func (o *DeactivateOrganizationRequest) SetSkipConfirmation(v bool)`

SetSkipConfirmation sets SkipConfirmation field to given value.

### HasSkipConfirmation

`func (o *DeactivateOrganizationRequest) HasSkipConfirmation() bool`

HasSkipConfirmation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


