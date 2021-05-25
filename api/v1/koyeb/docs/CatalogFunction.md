# CatalogFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BaseCatalogStatus**](BaseCatalogStatus.md) |  | [optional] [default to BASECATALOGSTATUS_COMING_SOON]
**Tags** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]FunctionParameter**](FunctionParameter.md) |  | [optional] 

## Methods

### NewCatalogFunction

`func NewCatalogFunction() *CatalogFunction`

NewCatalogFunction instantiates a new CatalogFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogFunctionWithDefaults

`func NewCatalogFunctionWithDefaults() *CatalogFunction`

NewCatalogFunctionWithDefaults instantiates a new CatalogFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogFunction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogFunction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogFunction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogFunction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *CatalogFunction) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogFunction) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogFunction) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogFunction) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CatalogFunction) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogFunction) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogFunction) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogFunction) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetWebsite

`func (o *CatalogFunction) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CatalogFunction) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CatalogFunction) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CatalogFunction) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogFunction) GetStatus() BaseCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogFunction) GetStatusOk() (*BaseCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogFunction) SetStatus(v BaseCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogFunction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CatalogFunction) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogFunction) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogFunction) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogFunction) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogFunction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogFunction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogFunction) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogFunction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImage

`func (o *CatalogFunction) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CatalogFunction) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CatalogFunction) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CatalogFunction) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetTemplate

`func (o *CatalogFunction) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CatalogFunction) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CatalogFunction) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CatalogFunction) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetParameters

`func (o *CatalogFunction) GetParameters() []FunctionParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CatalogFunction) GetParametersOk() (*[]FunctionParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CatalogFunction) SetParameters(v []FunctionParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CatalogFunction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


