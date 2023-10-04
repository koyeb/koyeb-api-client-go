# ServiceListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ServiceType**](ServiceType.md) |  | [optional] [default to SERVICETYPE_INVALID_TYPE]
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**ServiceStatus**](ServiceStatus.md) |  | [optional] [default to SERVICESTATUS_STARTING]
**Messages** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ServiceState**](ServiceState.md) |  | [optional] 
**ActiveDeploymentId** | Pointer to **string** |  | [optional] 
**LatestDeploymentId** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceListItem

`func NewServiceListItem() *ServiceListItem`

NewServiceListItem instantiates a new ServiceListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListItemWithDefaults

`func NewServiceListItemWithDefaults() *ServiceListItem`

NewServiceListItemWithDefaults instantiates a new ServiceListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ServiceListItem) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceListItem) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceListItem) SetType(v ServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServiceListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceListItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceListItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceListItem) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceListItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceListItem) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceListItem) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceListItem) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ServiceListItem) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ServiceListItem) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ServiceListItem) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ServiceListItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceListItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceListItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceListItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceListItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetState

`func (o *ServiceListItem) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceListItem) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceListItem) SetState(v ServiceState)`

SetState sets State field to given value.

### HasState

`func (o *ServiceListItem) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActiveDeploymentId

`func (o *ServiceListItem) GetActiveDeploymentId() string`

GetActiveDeploymentId returns the ActiveDeploymentId field if non-nil, zero value otherwise.

### GetActiveDeploymentIdOk

`func (o *ServiceListItem) GetActiveDeploymentIdOk() (*string, bool)`

GetActiveDeploymentIdOk returns a tuple with the ActiveDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeploymentId

`func (o *ServiceListItem) SetActiveDeploymentId(v string)`

SetActiveDeploymentId sets ActiveDeploymentId field to given value.

### HasActiveDeploymentId

`func (o *ServiceListItem) HasActiveDeploymentId() bool`

HasActiveDeploymentId returns a boolean if a field has been set.

### GetLatestDeploymentId

`func (o *ServiceListItem) GetLatestDeploymentId() string`

GetLatestDeploymentId returns the LatestDeploymentId field if non-nil, zero value otherwise.

### GetLatestDeploymentIdOk

`func (o *ServiceListItem) GetLatestDeploymentIdOk() (*string, bool)`

GetLatestDeploymentIdOk returns a tuple with the LatestDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeploymentId

`func (o *ServiceListItem) SetLatestDeploymentId(v string)`

SetLatestDeploymentId sets LatestDeploymentId field to given value.

### HasLatestDeploymentId

`func (o *ServiceListItem) HasLatestDeploymentId() bool`

HasLatestDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


