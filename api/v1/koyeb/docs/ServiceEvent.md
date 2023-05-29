# ServiceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewServiceEvent

`func NewServiceEvent() *ServiceEvent`

NewServiceEvent instantiates a new ServiceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEventWithDefaults

`func NewServiceEventWithDefaults() *ServiceEvent`

NewServiceEventWithDefaults instantiates a new ServiceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *ServiceEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ServiceEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ServiceEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *ServiceEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServiceEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceEvent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceEvent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceEvent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceEvent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetType

`func (o *ServiceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


