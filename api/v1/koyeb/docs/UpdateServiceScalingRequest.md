# UpdateServiceScalingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scalings** | Pointer to [**[]ManualServiceScaling**](ManualServiceScaling.md) |  | [optional] 

## Methods

### NewUpdateServiceScalingRequest

`func NewUpdateServiceScalingRequest() *UpdateServiceScalingRequest`

NewUpdateServiceScalingRequest instantiates a new UpdateServiceScalingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceScalingRequestWithDefaults

`func NewUpdateServiceScalingRequestWithDefaults() *UpdateServiceScalingRequest`

NewUpdateServiceScalingRequestWithDefaults instantiates a new UpdateServiceScalingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScalings

`func (o *UpdateServiceScalingRequest) GetScalings() []ManualServiceScaling`

GetScalings returns the Scalings field if non-nil, zero value otherwise.

### GetScalingsOk

`func (o *UpdateServiceScalingRequest) GetScalingsOk() (*[]ManualServiceScaling, bool)`

GetScalingsOk returns a tuple with the Scalings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalings

`func (o *UpdateServiceScalingRequest) SetScalings(v []ManualServiceScaling)`

SetScalings sets Scalings field to given value.

### HasScalings

`func (o *UpdateServiceScalingRequest) HasScalings() bool`

HasScalings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


