# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**With** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StoreStatus**](StoreStatus.md) |  | [optional] [default to STORESTATUS_PROVISIONING]
**StatusMessage** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StoreUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewStore

`func NewStore() *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Store) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Store) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Store) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Store) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Store) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Store) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Store) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Store) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Store) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Store) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Store) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWith

`func (o *Store) GetWith() map[string]interface{}`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *Store) GetWithOk() (*map[string]interface{}, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *Store) SetWith(v map[string]interface{})`

SetWith sets With field to given value.

### HasWith

`func (o *Store) HasWith() bool`

HasWith returns a boolean if a field has been set.

### GetRegion

`func (o *Store) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Store) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Store) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Store) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUrl

`func (o *Store) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Store) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Store) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Store) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *Store) GetStatus() StoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Store) GetStatusOk() (*StoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Store) SetStatus(v StoreStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Store) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Store) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Store) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Store) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Store) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Store) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Store) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Store) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Store) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Store) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Store) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Store) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Store) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Store) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Store) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Store) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Store) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStoreUrl

`func (o *Store) GetStoreUrl() string`

GetStoreUrl returns the StoreUrl field if non-nil, zero value otherwise.

### GetStoreUrlOk

`func (o *Store) GetStoreUrlOk() (*string, bool)`

GetStoreUrlOk returns a tuple with the StoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreUrl

`func (o *Store) SetStoreUrl(v string)`

SetStoreUrl sets StoreUrl field to given value.

### HasStoreUrl

`func (o *Store) HasStoreUrl() bool`

HasStoreUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


