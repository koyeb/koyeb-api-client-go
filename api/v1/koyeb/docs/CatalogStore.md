# CatalogStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**BaseCatalogStatus**](BaseCatalogStatus.md) |  | [optional] [default to BASECATALOGSTATUS_COMING_SOON]
**Parameters** | Pointer to [**[]Parameter**](Parameter.md) |  | [optional] 

## Methods

### NewCatalogStore

`func NewCatalogStore() *CatalogStore`

NewCatalogStore instantiates a new CatalogStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogStoreWithDefaults

`func NewCatalogStoreWithDefaults() *CatalogStore`

NewCatalogStoreWithDefaults instantiates a new CatalogStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogStore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogStore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogStore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogStore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *CatalogStore) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogStore) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogStore) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogStore) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CatalogStore) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogStore) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogStore) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogStore) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetWebsite

`func (o *CatalogStore) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CatalogStore) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CatalogStore) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CatalogStore) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetVendor

`func (o *CatalogStore) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CatalogStore) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CatalogStore) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CatalogStore) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetTags

`func (o *CatalogStore) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogStore) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogStore) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogStore) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogStore) GetStatus() BaseCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogStore) GetStatusOk() (*BaseCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogStore) SetStatus(v BaseCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogStore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParameters

`func (o *CatalogStore) GetParameters() []Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CatalogStore) GetParametersOk() (*[]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CatalogStore) SetParameters(v []Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CatalogStore) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


