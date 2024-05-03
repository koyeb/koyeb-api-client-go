# RegionalDeploymentListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**RegionalDeploymentStatus**](RegionalDeploymentStatus.md) |  | [optional] [default to REGIONALDEPLOYMENTSTATUS_PENDING]
**Messages** | Pointer to **[]string** |  | [optional] 
**Definition** | Pointer to [**RegionalDeploymentDefinition**](RegionalDeploymentDefinition.md) |  | [optional] 
**UseKumaV2** | Pointer to **bool** |  | [optional] 
**UseKata** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegionalDeploymentListItem

`func NewRegionalDeploymentListItem() *RegionalDeploymentListItem`

NewRegionalDeploymentListItem instantiates a new RegionalDeploymentListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionalDeploymentListItemWithDefaults

`func NewRegionalDeploymentListItemWithDefaults() *RegionalDeploymentListItem`

NewRegionalDeploymentListItemWithDefaults instantiates a new RegionalDeploymentListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionalDeploymentListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionalDeploymentListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionalDeploymentListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegionalDeploymentListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegionalDeploymentListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegionalDeploymentListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegionalDeploymentListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegionalDeploymentListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RegionalDeploymentListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegionalDeploymentListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegionalDeploymentListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegionalDeploymentListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRegion

`func (o *RegionalDeploymentListItem) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionalDeploymentListItem) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionalDeploymentListItem) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RegionalDeploymentListItem) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *RegionalDeploymentListItem) GetStatus() RegionalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionalDeploymentListItem) GetStatusOk() (*RegionalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionalDeploymentListItem) SetStatus(v RegionalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegionalDeploymentListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessages

`func (o *RegionalDeploymentListItem) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *RegionalDeploymentListItem) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *RegionalDeploymentListItem) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *RegionalDeploymentListItem) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetDefinition

`func (o *RegionalDeploymentListItem) GetDefinition() RegionalDeploymentDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *RegionalDeploymentListItem) GetDefinitionOk() (*RegionalDeploymentDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *RegionalDeploymentListItem) SetDefinition(v RegionalDeploymentDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *RegionalDeploymentListItem) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetUseKumaV2

`func (o *RegionalDeploymentListItem) GetUseKumaV2() bool`

GetUseKumaV2 returns the UseKumaV2 field if non-nil, zero value otherwise.

### GetUseKumaV2Ok

`func (o *RegionalDeploymentListItem) GetUseKumaV2Ok() (*bool, bool)`

GetUseKumaV2Ok returns a tuple with the UseKumaV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKumaV2

`func (o *RegionalDeploymentListItem) SetUseKumaV2(v bool)`

SetUseKumaV2 sets UseKumaV2 field to given value.

### HasUseKumaV2

`func (o *RegionalDeploymentListItem) HasUseKumaV2() bool`

HasUseKumaV2 returns a boolean if a field has been set.

### GetUseKata

`func (o *RegionalDeploymentListItem) GetUseKata() bool`

GetUseKata returns the UseKata field if non-nil, zero value otherwise.

### GetUseKataOk

`func (o *RegionalDeploymentListItem) GetUseKataOk() (*bool, bool)`

GetUseKataOk returns a tuple with the UseKata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKata

`func (o *RegionalDeploymentListItem) SetUseKata(v bool)`

SetUseKata sets UseKata field to given value.

### HasUseKata

`func (o *RegionalDeploymentListItem) HasUseKata() bool`

HasUseKata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


