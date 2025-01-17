# ConfigFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigFile

`func NewConfigFile() *ConfigFile`

NewConfigFile instantiates a new ConfigFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFileWithDefaults

`func NewConfigFileWithDefaults() *ConfigFile`

NewConfigFileWithDefaults instantiates a new ConfigFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ConfigFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConfigFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConfigFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConfigFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *ConfigFile) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ConfigFile) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ConfigFile) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ConfigFile) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetContent

`func (o *ConfigFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConfigFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConfigFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConfigFile) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


