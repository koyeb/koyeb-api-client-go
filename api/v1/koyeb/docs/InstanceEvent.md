# InstanceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceEvent

`func NewInstanceEvent() *InstanceEvent`

NewInstanceEvent instantiates a new InstanceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceEventWithDefaults

`func NewInstanceEventWithDefaults() *InstanceEvent`

NewInstanceEventWithDefaults instantiates a new InstanceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *InstanceEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *InstanceEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *InstanceEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *InstanceEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InstanceEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InstanceEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InstanceEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InstanceEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *InstanceEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InstanceEvent) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetType

`func (o *InstanceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *InstanceEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InstanceEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InstanceEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InstanceEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


