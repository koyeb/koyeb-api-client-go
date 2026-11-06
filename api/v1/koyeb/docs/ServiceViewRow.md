# ServiceViewRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ServiceType**](ServiceType.md) |  | [optional] [default to SERVICETYPE_INVALID_TYPE]
**Status** | Pointer to [**ServiceStatus**](ServiceStatus.md) |  | [optional] [default to SERVICESTATUS_STARTING]
**Messages** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**App** | Pointer to [**ServiceViewApp**](ServiceViewApp.md) |  | [optional] 
**ActiveDeployment** | Pointer to [**ServiceViewActiveDeployment**](ServiceViewActiveDeployment.md) |  | [optional] 
**LatestDeployment** | Pointer to [**ServiceViewLatestDeployment**](ServiceViewLatestDeployment.md) |  | [optional] 
**NeedsAttention** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**ServiceViewSource**](ServiceViewSource.md) |  | [optional] 

## Methods

### NewServiceViewRow

`func NewServiceViewRow() *ServiceViewRow`

NewServiceViewRow instantiates a new ServiceViewRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceViewRowWithDefaults

`func NewServiceViewRowWithDefaults() *ServiceViewRow`

NewServiceViewRowWithDefaults instantiates a new ServiceViewRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceViewRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceViewRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceViewRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceViewRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceViewRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceViewRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceViewRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceViewRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ServiceViewRow) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceViewRow) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceViewRow) SetType(v ServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceViewRow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceViewRow) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceViewRow) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceViewRow) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceViewRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ServiceViewRow) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ServiceViewRow) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ServiceViewRow) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ServiceViewRow) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceViewRow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceViewRow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceViewRow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceViewRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceViewRow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceViewRow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceViewRow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceViewRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceViewRow) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceViewRow) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceViewRow) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceViewRow) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetApp

`func (o *ServiceViewRow) GetApp() ServiceViewApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ServiceViewRow) GetAppOk() (*ServiceViewApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ServiceViewRow) SetApp(v ServiceViewApp)`

SetApp sets App field to given value.

### HasApp

`func (o *ServiceViewRow) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetActiveDeployment

`func (o *ServiceViewRow) GetActiveDeployment() ServiceViewActiveDeployment`

GetActiveDeployment returns the ActiveDeployment field if non-nil, zero value otherwise.

### GetActiveDeploymentOk

`func (o *ServiceViewRow) GetActiveDeploymentOk() (*ServiceViewActiveDeployment, bool)`

GetActiveDeploymentOk returns a tuple with the ActiveDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeployment

`func (o *ServiceViewRow) SetActiveDeployment(v ServiceViewActiveDeployment)`

SetActiveDeployment sets ActiveDeployment field to given value.

### HasActiveDeployment

`func (o *ServiceViewRow) HasActiveDeployment() bool`

HasActiveDeployment returns a boolean if a field has been set.

### GetLatestDeployment

`func (o *ServiceViewRow) GetLatestDeployment() ServiceViewLatestDeployment`

GetLatestDeployment returns the LatestDeployment field if non-nil, zero value otherwise.

### GetLatestDeploymentOk

`func (o *ServiceViewRow) GetLatestDeploymentOk() (*ServiceViewLatestDeployment, bool)`

GetLatestDeploymentOk returns a tuple with the LatestDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeployment

`func (o *ServiceViewRow) SetLatestDeployment(v ServiceViewLatestDeployment)`

SetLatestDeployment sets LatestDeployment field to given value.

### HasLatestDeployment

`func (o *ServiceViewRow) HasLatestDeployment() bool`

HasLatestDeployment returns a boolean if a field has been set.

### GetNeedsAttention

`func (o *ServiceViewRow) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *ServiceViewRow) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *ServiceViewRow) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.

### HasNeedsAttention

`func (o *ServiceViewRow) HasNeedsAttention() bool`

HasNeedsAttention returns a boolean if a field has been set.

### GetRegions

`func (o *ServiceViewRow) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ServiceViewRow) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ServiceViewRow) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ServiceViewRow) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetInstanceType

`func (o *ServiceViewRow) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ServiceViewRow) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ServiceViewRow) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ServiceViewRow) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSource

`func (o *ServiceViewRow) GetSource() ServiceViewSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ServiceViewRow) GetSourceOk() (*ServiceViewSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ServiceViewRow) SetSource(v ServiceViewSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ServiceViewRow) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


