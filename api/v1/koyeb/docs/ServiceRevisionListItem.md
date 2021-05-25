# ServiceRevisionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ChildId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**DeploymentGroup** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ServiceRevisionStateStatus**](ServiceRevisionStateStatus.md) |  | [optional] [default to SERVICEREVISIONSTATESTATUS_UNKNOWN]

## Methods

### NewServiceRevisionListItem

`func NewServiceRevisionListItem() *ServiceRevisionListItem`

NewServiceRevisionListItem instantiates a new ServiceRevisionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevisionListItemWithDefaults

`func NewServiceRevisionListItemWithDefaults() *ServiceRevisionListItem`

NewServiceRevisionListItemWithDefaults instantiates a new ServiceRevisionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *ServiceRevisionListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceRevisionListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceRevisionListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceRevisionListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceRevisionListItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceRevisionListItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceRevisionListItem) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceRevisionListItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceRevisionListItem) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceRevisionListItem) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceRevisionListItem) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceRevisionListItem) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetId

`func (o *ServiceRevisionListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceRevisionListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceRevisionListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceRevisionListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceRevisionListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceRevisionListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceRevisionListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceRevisionListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceRevisionListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceRevisionListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceRevisionListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceRevisionListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetParentId

`func (o *ServiceRevisionListItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ServiceRevisionListItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ServiceRevisionListItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ServiceRevisionListItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildId

`func (o *ServiceRevisionListItem) GetChildId() string`

GetChildId returns the ChildId field if non-nil, zero value otherwise.

### GetChildIdOk

`func (o *ServiceRevisionListItem) GetChildIdOk() (*string, bool)`

GetChildIdOk returns a tuple with the ChildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildId

`func (o *ServiceRevisionListItem) SetChildId(v string)`

SetChildId sets ChildId field to given value.

### HasChildId

`func (o *ServiceRevisionListItem) HasChildId() bool`

HasChildId returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceRevisionListItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceRevisionListItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceRevisionListItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceRevisionListItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *ServiceRevisionListItem) GetDeploymentGroup() string`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *ServiceRevisionListItem) GetDeploymentGroupOk() (*string, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *ServiceRevisionListItem) SetDeploymentGroup(v string)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *ServiceRevisionListItem) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceRevisionListItem) GetStatus() ServiceRevisionStateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceRevisionListItem) GetStatusOk() (*ServiceRevisionStateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceRevisionListItem) SetStatus(v ServiceRevisionStateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceRevisionListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


