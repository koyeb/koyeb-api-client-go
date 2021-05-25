# ServiceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredDeployment** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Status** | Pointer to [**ServiceStateStatus**](ServiceStateStatus.md) |  | [optional] [default to SERVICESTATESTATUS_UNKNOWN]

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

`func (o *ServiceState) GetDesiredDeployment() DeploymentState`

GetDesiredDeployment returns the DesiredDeployment field if non-nil, zero value otherwise.

### GetDesiredDeploymentOk

`func (o *ServiceState) GetDesiredDeploymentOk() (*DeploymentState, bool)`

GetDesiredDeploymentOk returns a tuple with the DesiredDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredDeployment

`func (o *ServiceState) SetDesiredDeployment(v DeploymentState)`

SetDesiredDeployment sets DesiredDeployment field to given value.

### HasDesiredDeployment

`func (o *ServiceState) HasDesiredDeployment() bool`

HasDesiredDeployment returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceState) GetStatus() ServiceStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceState) GetStatusOk() (*ServiceStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceState) SetStatus(v ServiceStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


