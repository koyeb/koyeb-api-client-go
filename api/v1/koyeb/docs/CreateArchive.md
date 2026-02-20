# CreateArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | How much space to provision for the archive, in bytes. | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateArchive

`func NewCreateArchive() *CreateArchive`

NewCreateArchive instantiates a new CreateArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArchiveWithDefaults

`func NewCreateArchiveWithDefaults() *CreateArchive`

NewCreateArchiveWithDefaults instantiates a new CreateArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *CreateArchive) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateArchive) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateArchive) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateArchive) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateArchive) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateArchive) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateArchive) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateArchive) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


