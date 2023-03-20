# UpdateOrganizationPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] [default to PLAN_HOBBY]

## Methods

### NewUpdateOrganizationPlanRequest

`func NewUpdateOrganizationPlanRequest() *UpdateOrganizationPlanRequest`

NewUpdateOrganizationPlanRequest instantiates a new UpdateOrganizationPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationPlanRequestWithDefaults

`func NewUpdateOrganizationPlanRequestWithDefaults() *UpdateOrganizationPlanRequest`

NewUpdateOrganizationPlanRequestWithDefaults instantiates a new UpdateOrganizationPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *UpdateOrganizationPlanRequest) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateOrganizationPlanRequest) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateOrganizationPlanRequest) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateOrganizationPlanRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


