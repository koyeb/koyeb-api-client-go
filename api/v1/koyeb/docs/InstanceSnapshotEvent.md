# InstanceSnapshotEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**InstanceSnapshotId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceSnapshotEvent

`func NewInstanceSnapshotEvent() *InstanceSnapshotEvent`

NewInstanceSnapshotEvent instantiates a new InstanceSnapshotEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSnapshotEventWithDefaults

`func NewInstanceSnapshotEventWithDefaults() *InstanceSnapshotEvent`

NewInstanceSnapshotEventWithDefaults instantiates a new InstanceSnapshotEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceSnapshotEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceSnapshotEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceSnapshotEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceSnapshotEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *InstanceSnapshotEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *InstanceSnapshotEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *InstanceSnapshotEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *InstanceSnapshotEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InstanceSnapshotEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InstanceSnapshotEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InstanceSnapshotEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InstanceSnapshotEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetInstanceSnapshotId

`func (o *InstanceSnapshotEvent) GetInstanceSnapshotId() string`

GetInstanceSnapshotId returns the InstanceSnapshotId field if non-nil, zero value otherwise.

### GetInstanceSnapshotIdOk

`func (o *InstanceSnapshotEvent) GetInstanceSnapshotIdOk() (*string, bool)`

GetInstanceSnapshotIdOk returns a tuple with the InstanceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSnapshotId

`func (o *InstanceSnapshotEvent) SetInstanceSnapshotId(v string)`

SetInstanceSnapshotId sets InstanceSnapshotId field to given value.

### HasInstanceSnapshotId

`func (o *InstanceSnapshotEvent) HasInstanceSnapshotId() bool`

HasInstanceSnapshotId returns a boolean if a field has been set.

### GetType

`func (o *InstanceSnapshotEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceSnapshotEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceSnapshotEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceSnapshotEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *InstanceSnapshotEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InstanceSnapshotEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InstanceSnapshotEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InstanceSnapshotEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceSnapshotEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceSnapshotEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceSnapshotEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceSnapshotEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


