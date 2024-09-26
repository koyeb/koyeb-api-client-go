# PersistentVolumeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PersistentVolumeId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPersistentVolumeEvent

`func NewPersistentVolumeEvent() *PersistentVolumeEvent`

NewPersistentVolumeEvent instantiates a new PersistentVolumeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentVolumeEventWithDefaults

`func NewPersistentVolumeEventWithDefaults() *PersistentVolumeEvent`

NewPersistentVolumeEventWithDefaults instantiates a new PersistentVolumeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersistentVolumeEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersistentVolumeEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersistentVolumeEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersistentVolumeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *PersistentVolumeEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *PersistentVolumeEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *PersistentVolumeEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *PersistentVolumeEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PersistentVolumeEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PersistentVolumeEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PersistentVolumeEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PersistentVolumeEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPersistentVolumeId

`func (o *PersistentVolumeEvent) GetPersistentVolumeId() string`

GetPersistentVolumeId returns the PersistentVolumeId field if non-nil, zero value otherwise.

### GetPersistentVolumeIdOk

`func (o *PersistentVolumeEvent) GetPersistentVolumeIdOk() (*string, bool)`

GetPersistentVolumeIdOk returns a tuple with the PersistentVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeId

`func (o *PersistentVolumeEvent) SetPersistentVolumeId(v string)`

SetPersistentVolumeId sets PersistentVolumeId field to given value.

### HasPersistentVolumeId

`func (o *PersistentVolumeEvent) HasPersistentVolumeId() bool`

HasPersistentVolumeId returns a boolean if a field has been set.

### GetType

`func (o *PersistentVolumeEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PersistentVolumeEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PersistentVolumeEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PersistentVolumeEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *PersistentVolumeEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PersistentVolumeEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PersistentVolumeEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PersistentVolumeEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *PersistentVolumeEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PersistentVolumeEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PersistentVolumeEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PersistentVolumeEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


