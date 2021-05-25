# CatalogStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BaseCatalogStatus**](BaseCatalogStatus.md) |  | [optional] [default to BASECATALOGSTATUS_COMING_SOON]
**Tags** | Pointer to **[]string** |  | [optional] 
**SourceControlRef** | Pointer to [**SCMReference**](SCMReference.md) |  | [optional] 
**IsYamlOnly** | Pointer to **bool** | Whether this stack is only a yaml or not (in which case it can be created as a non git managed stack). | [optional] 
**Yaml** | Pointer to **string** | The yaml that is represented the stack. | [optional] 

## Methods

### NewCatalogStack

`func NewCatalogStack() *CatalogStack`

NewCatalogStack instantiates a new CatalogStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogStackWithDefaults

`func NewCatalogStackWithDefaults() *CatalogStack`

NewCatalogStackWithDefaults instantiates a new CatalogStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogStack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogStack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogStack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogStack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogStack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogStack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogStack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogStack) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *CatalogStack) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogStack) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogStack) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogStack) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogStack) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogStack) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogStack) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogStack) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CatalogStack) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogStack) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogStack) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogStack) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogStack) GetStatus() BaseCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogStack) GetStatusOk() (*BaseCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogStack) SetStatus(v BaseCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogStack) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CatalogStack) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogStack) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogStack) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogStack) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSourceControlRef

`func (o *CatalogStack) GetSourceControlRef() SCMReference`

GetSourceControlRef returns the SourceControlRef field if non-nil, zero value otherwise.

### GetSourceControlRefOk

`func (o *CatalogStack) GetSourceControlRefOk() (*SCMReference, bool)`

GetSourceControlRefOk returns a tuple with the SourceControlRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlRef

`func (o *CatalogStack) SetSourceControlRef(v SCMReference)`

SetSourceControlRef sets SourceControlRef field to given value.

### HasSourceControlRef

`func (o *CatalogStack) HasSourceControlRef() bool`

HasSourceControlRef returns a boolean if a field has been set.

### GetIsYamlOnly

`func (o *CatalogStack) GetIsYamlOnly() bool`

GetIsYamlOnly returns the IsYamlOnly field if non-nil, zero value otherwise.

### GetIsYamlOnlyOk

`func (o *CatalogStack) GetIsYamlOnlyOk() (*bool, bool)`

GetIsYamlOnlyOk returns a tuple with the IsYamlOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYamlOnly

`func (o *CatalogStack) SetIsYamlOnly(v bool)`

SetIsYamlOnly sets IsYamlOnly field to given value.

### HasIsYamlOnly

`func (o *CatalogStack) HasIsYamlOnly() bool`

HasIsYamlOnly returns a boolean if a field has been set.

### GetYaml

`func (o *CatalogStack) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *CatalogStack) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *CatalogStack) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *CatalogStack) HasYaml() bool`

HasYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


