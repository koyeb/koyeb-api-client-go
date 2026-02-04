# ManualServiceScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]string** |  | [optional] 
**Instances** | Pointer to **int64** |  | [optional] 

## Methods

### NewManualServiceScaling

`func NewManualServiceScaling() *ManualServiceScaling`

NewManualServiceScaling instantiates a new ManualServiceScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualServiceScalingWithDefaults

`func NewManualServiceScalingWithDefaults() *ManualServiceScaling`

NewManualServiceScalingWithDefaults instantiates a new ManualServiceScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ManualServiceScaling) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ManualServiceScaling) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ManualServiceScaling) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ManualServiceScaling) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetInstances

`func (o *ManualServiceScaling) GetInstances() int64`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ManualServiceScaling) GetInstancesOk() (*int64, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ManualServiceScaling) SetInstances(v int64)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ManualServiceScaling) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


