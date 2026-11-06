# ServiceViewLatestDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**ServiceViewLatestDeploymentState**](ServiceViewLatestDeploymentState.md) |  | [optional] [default to SERVICEVIEWLATESTDEPLOYMENTSTATE_NONE]
**Id** | Pointer to **string** |  | [optional] 
**Missing** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**DeploymentStatus**](DeploymentStatus.md) |  | [optional] [default to DEPLOYMENTSTATUS_PENDING]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceViewLatestDeployment

`func NewServiceViewLatestDeployment() *ServiceViewLatestDeployment`

NewServiceViewLatestDeployment instantiates a new ServiceViewLatestDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewLatestDeploymentWithDefaults

`func NewServiceViewLatestDeploymentWithDefaults() *ServiceViewLatestDeployment`

NewServiceViewLatestDeploymentWithDefaults instantiates a new ServiceViewLatestDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ServiceViewLatestDeployment) GetState() ServiceViewLatestDeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceViewLatestDeployment) GetStateOk() (*ServiceViewLatestDeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceViewLatestDeployment) SetState(v ServiceViewLatestDeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *ServiceViewLatestDeployment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetId

`func (o *ServiceViewLatestDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceViewLatestDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceViewLatestDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceViewLatestDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMissing

`func (o *ServiceViewLatestDeployment) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *ServiceViewLatestDeployment) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *ServiceViewLatestDeployment) SetMissing(v bool)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *ServiceViewLatestDeployment) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceViewLatestDeployment) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceViewLatestDeployment) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceViewLatestDeployment) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceViewLatestDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceViewLatestDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceViewLatestDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceViewLatestDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceViewLatestDeployment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceViewLatestDeployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceViewLatestDeployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceViewLatestDeployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceViewLatestDeployment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


