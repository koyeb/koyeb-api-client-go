# ServiceRevision

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
**Definition** | Pointer to [**ServiceDefinition**](ServiceDefinition.md) |  | [optional] 
**State** | Pointer to [**ServiceRevisionState**](ServiceRevisionState.md) |  | [optional] 

## Methods

### NewServiceRevision

`func NewServiceRevision() *ServiceRevision`

NewServiceRevision instantiates a new ServiceRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevisionWithDefaults

`func NewServiceRevisionWithDefaults() *ServiceRevision`

NewServiceRevisionWithDefaults instantiates a new ServiceRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *ServiceRevision) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceRevision) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceRevision) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceRevision) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceRevision) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceRevision) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceRevision) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceRevision) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceRevision) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceRevision) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceRevision) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceRevision) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetId

`func (o *ServiceRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceRevision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceRevision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceRevision) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceRevision) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceRevision) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceRevision) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceRevision) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetParentId

`func (o *ServiceRevision) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ServiceRevision) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ServiceRevision) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ServiceRevision) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetChildId

`func (o *ServiceRevision) GetChildId() string`

GetChildId returns the ChildId field if non-nil, zero value otherwise.

### GetChildIdOk

`func (o *ServiceRevision) GetChildIdOk() (*string, bool)`

GetChildIdOk returns a tuple with the ChildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildId

`func (o *ServiceRevision) SetChildId(v string)`

SetChildId sets ChildId field to given value.

### HasChildId

`func (o *ServiceRevision) HasChildId() bool`

HasChildId returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceRevision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceRevision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceRevision) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceRevision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefinition

`func (o *ServiceRevision) GetDefinition() ServiceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ServiceRevision) GetDefinitionOk() (*ServiceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ServiceRevision) SetDefinition(v ServiceDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ServiceRevision) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetState

`func (o *ServiceRevision) GetState() ServiceRevisionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceRevision) GetStateOk() (*ServiceRevisionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceRevision) SetState(v ServiceRevisionState)`

SetState sets State field to given value.

### HasState

`func (o *ServiceRevision) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


