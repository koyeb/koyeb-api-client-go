# FileMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **string** | interpolation_enabled is a flag to enable/disable interpolation in the file content  bool interpolation_enabled &#x3D; 3; | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewFileMount

`func NewFileMount() *FileMount`

NewFileMount instantiates a new FileMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMountWithDefaults

`func NewFileMountWithDefaults() *FileMount`

NewFileMountWithDefaults instantiates a new FileMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileMount) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileMount) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileMount) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileMount) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *FileMount) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FileMount) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FileMount) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FileMount) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetContent

`func (o *FileMount) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileMount) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileMount) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileMount) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

