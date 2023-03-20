# PublicOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] [default to PLAN_HOBBY]
**Status** | Pointer to [**OrganizationStatus**](OrganizationStatus.md) |  | [optional] [default to ORGANIZATIONSTATUS_WARNING]

## Methods

### NewPublicOrganization

`func NewPublicOrganization() *PublicOrganization`

NewPublicOrganization instantiates a new PublicOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicOrganizationWithDefaults

`func NewPublicOrganizationWithDefaults() *PublicOrganization`

NewPublicOrganizationWithDefaults instantiates a new PublicOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PublicOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *PublicOrganization) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PublicOrganization) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PublicOrganization) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PublicOrganization) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetStatus

`func (o *PublicOrganization) GetStatus() OrganizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicOrganization) GetStatusOk() (*OrganizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicOrganization) SetStatus(v OrganizationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicOrganization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


