# ServiceViewActiveDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**ServiceViewActiveDeploymentState**](ServiceViewActiveDeploymentState.md) |  | [optional] [default to SERVICEVIEWACTIVEDEPLOYMENTSTATE_NONE]
**Id** | Pointer to **string** |  | [optional] 
**Missing** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**DeploymentStatus**](DeploymentStatus.md) |  | [optional] [default to DEPLOYMENTSTATUS_PENDING]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceViewActiveDeployment

`func NewServiceViewActiveDeployment() *ServiceViewActiveDeployment`

NewServiceViewActiveDeployment instantiates a new ServiceViewActiveDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewActiveDeploymentWithDefaults

`func NewServiceViewActiveDeploymentWithDefaults() *ServiceViewActiveDeployment`

NewServiceViewActiveDeploymentWithDefaults instantiates a new ServiceViewActiveDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ServiceViewActiveDeployment) GetState() ServiceViewActiveDeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceViewActiveDeployment) GetStateOk() (*ServiceViewActiveDeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceViewActiveDeployment) SetState(v ServiceViewActiveDeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *ServiceViewActiveDeployment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetId

`func (o *ServiceViewActiveDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceViewActiveDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceViewActiveDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceViewActiveDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMissing

`func (o *ServiceViewActiveDeployment) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *ServiceViewActiveDeployment) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *ServiceViewActiveDeployment) SetMissing(v bool)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *ServiceViewActiveDeployment) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceViewActiveDeployment) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceViewActiveDeployment) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceViewActiveDeployment) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceViewActiveDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceViewActiveDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceViewActiveDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceViewActiveDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceViewActiveDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceViewActiveDeployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceViewActiveDeployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceViewActiveDeployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceViewActiveDeployment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


