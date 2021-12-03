# DesiredDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]DesiredDeploymentGroup**](DesiredDeploymentGroup.md) |  | [optional] 

## Methods

### NewDesiredDeployment

`func NewDesiredDeployment() *DesiredDeployment`

NewDesiredDeployment instantiates a new DesiredDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesiredDeploymentWithDefaults

`func NewDesiredDeploymentWithDefaults() *DesiredDeployment`

NewDesiredDeploymentWithDefaults instantiates a new DesiredDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *DesiredDeployment) GetGroups() []DesiredDeploymentGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DesiredDeployment) GetGroupsOk() (*[]DesiredDeploymentGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DesiredDeployment) SetGroups(v []DesiredDeploymentGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DesiredDeployment) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


