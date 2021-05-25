# CatalogStackListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BaseCatalogStatus**](BaseCatalogStatus.md) |  | [optional] [default to BASECATALOGSTATUS_COMING_SOON]
**Tags** | Pointer to **[]string** |  | [optional] 
**SourceControlRef** | Pointer to [**SCMReference**](SCMReference.md) |  | [optional] 

## Methods

### NewCatalogStackListItem

`func NewCatalogStackListItem() *CatalogStackListItem`

NewCatalogStackListItem instantiates a new CatalogStackListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogStackListItemWithDefaults

`func NewCatalogStackListItemWithDefaults() *CatalogStackListItem`

NewCatalogStackListItemWithDefaults instantiates a new CatalogStackListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogStackListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogStackListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogStackListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogStackListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogStackListItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogStackListItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogStackListItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogStackListItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *CatalogStackListItem) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogStackListItem) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogStackListItem) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogStackListItem) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CatalogStackListItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogStackListItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogStackListItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogStackListItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogStackListItem) GetStatus() BaseCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogStackListItem) GetStatusOk() (*BaseCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogStackListItem) SetStatus(v BaseCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogStackListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CatalogStackListItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogStackListItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogStackListItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogStackListItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSourceControlRef

`func (o *CatalogStackListItem) GetSourceControlRef() SCMReference`

GetSourceControlRef returns the SourceControlRef field if non-nil, zero value otherwise.

### GetSourceControlRefOk

`func (o *CatalogStackListItem) GetSourceControlRefOk() (*SCMReference, bool)`

GetSourceControlRefOk returns a tuple with the SourceControlRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlRef

`func (o *CatalogStackListItem) SetSourceControlRef(v SCMReference)`

SetSourceControlRef sets SourceControlRef field to given value.

### HasSourceControlRef

`func (o *CatalogStackListItem) HasSourceControlRef() bool`

HasSourceControlRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


