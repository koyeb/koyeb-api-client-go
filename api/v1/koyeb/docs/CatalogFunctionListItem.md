# CatalogFunctionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BaseCatalogStatus**](BaseCatalogStatus.md) |  | [optional] [default to BASECATALOGSTATUS_COMING_SOON]
**Tags** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]FunctionParameter**](FunctionParameter.md) |  | [optional] 

## Methods

### NewCatalogFunctionListItem

`func NewCatalogFunctionListItem() *CatalogFunctionListItem`

NewCatalogFunctionListItem instantiates a new CatalogFunctionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogFunctionListItemWithDefaults

`func NewCatalogFunctionListItemWithDefaults() *CatalogFunctionListItem`

NewCatalogFunctionListItemWithDefaults instantiates a new CatalogFunctionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogFunctionListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogFunctionListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogFunctionListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogFunctionListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogFunctionListItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogFunctionListItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogFunctionListItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogFunctionListItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *CatalogFunctionListItem) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogFunctionListItem) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogFunctionListItem) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogFunctionListItem) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CatalogFunctionListItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogFunctionListItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogFunctionListItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogFunctionListItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetWebsite

`func (o *CatalogFunctionListItem) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CatalogFunctionListItem) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CatalogFunctionListItem) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CatalogFunctionListItem) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogFunctionListItem) GetStatus() BaseCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogFunctionListItem) GetStatusOk() (*BaseCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogFunctionListItem) SetStatus(v BaseCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogFunctionListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CatalogFunctionListItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogFunctionListItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogFunctionListItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogFunctionListItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogFunctionListItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogFunctionListItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogFunctionListItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogFunctionListItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImage

`func (o *CatalogFunctionListItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CatalogFunctionListItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CatalogFunctionListItem) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CatalogFunctionListItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetParameters

`func (o *CatalogFunctionListItem) GetParameters() []FunctionParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CatalogFunctionListItem) GetParametersOk() (*[]FunctionParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CatalogFunctionListItem) SetParameters(v []FunctionParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CatalogFunctionListItem) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


