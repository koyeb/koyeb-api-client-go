# AppListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Domains** | Pointer to [**[]Domain**](Domain.md) |  | [optional] 
**Status** | Pointer to [**AppStatus**](AppStatus.md) |  | [optional] [default to APPSTATUS_STARTING]
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppListItem

`func NewAppListItem() *AppListItem`

NewAppListItem instantiates a new AppListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppListItemWithDefaults

`func NewAppListItemWithDefaults() *AppListItem`

NewAppListItemWithDefaults instantiates a new AppListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AppListItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AppListItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AppListItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AppListItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDomains

`func (o *AppListItem) GetDomains() []Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AppListItem) GetDomainsOk() (*[]Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AppListItem) SetDomains(v []Domain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AppListItem) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetStatus

`func (o *AppListItem) GetStatus() AppStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppListItem) GetStatusOk() (*AppStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppListItem) SetStatus(v AppStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *AppListItem) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AppListItem) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AppListItem) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AppListItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


