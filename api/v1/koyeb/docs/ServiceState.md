# ServiceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredDeployment** | Pointer to [**DesiredDeployment**](DesiredDeployment.md) |  | [optional] 
**AutoRelease** | Pointer to [**AutoRelease**](AutoRelease.md) |  | [optional] 

## Methods

### NewServiceState

`func NewServiceState() *ServiceState`

NewServiceState instantiates a new ServiceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStateWithDefaults

`func NewServiceStateWithDefaults() *ServiceState`

NewServiceStateWithDefaults instantiates a new ServiceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredDeployment

`func (o *ServiceState) GetDesiredDeployment() DesiredDeployment`

GetDesiredDeployment returns the DesiredDeployment field if non-nil, zero value otherwise.

### GetDesiredDeploymentOk

`func (o *ServiceState) GetDesiredDeploymentOk() (*DesiredDeployment, bool)`

GetDesiredDeploymentOk returns a tuple with the DesiredDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredDeployment

`func (o *ServiceState) SetDesiredDeployment(v DesiredDeployment)`

SetDesiredDeployment sets DesiredDeployment field to given value.

### HasDesiredDeployment

`func (o *ServiceState) HasDesiredDeployment() bool`

HasDesiredDeployment returns a boolean if a field has been set.

### GetAutoRelease

`func (o *ServiceState) GetAutoRelease() AutoRelease`

GetAutoRelease returns the AutoRelease field if non-nil, zero value otherwise.

### GetAutoReleaseOk

`func (o *ServiceState) GetAutoReleaseOk() (*AutoRelease, bool)`

GetAutoReleaseOk returns a tuple with the AutoRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRelease

`func (o *ServiceState) SetAutoRelease(v AutoRelease)`

SetAutoRelease sets AutoRelease field to given value.

### HasAutoRelease

`func (o *ServiceState) HasAutoRelease() bool`

HasAutoRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


