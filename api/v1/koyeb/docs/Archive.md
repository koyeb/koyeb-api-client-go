# Archive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The archive id, that can be referenced when creating or updating a service. | [optional] 
**OrganizationId** | Pointer to **string** | Organization owning the archive. | [optional] 
**UploadUrl** | Pointer to **string** | The URL where to upload the archive. This URL is signed and can only be used to upload the archive until &#x60;valid_until&#x60;. | [optional] 
**Size** | Pointer to **string** | The provisioned space for the archive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date of creation of the archive. | [optional] 
**DeletedAt** | Pointer to **time.Time** | This field is automatically set by Koyeb when the archive is garbage collected. | [optional] 

## Methods

### NewArchive

`func NewArchive() *Archive`

NewArchive instantiates a new Archive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveWithDefaults

`func NewArchiveWithDefaults() *Archive`

NewArchiveWithDefaults instantiates a new Archive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Archive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Archive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Archive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Archive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Archive) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Archive) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Archive) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Archive) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *Archive) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *Archive) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *Archive) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *Archive) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetSize

`func (o *Archive) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Archive) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Archive) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Archive) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Archive) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Archive) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Archive) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Archive) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Archive) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Archive) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Archive) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Archive) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


