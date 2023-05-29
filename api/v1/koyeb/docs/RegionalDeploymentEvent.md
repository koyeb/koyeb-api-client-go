# RegionalDeploymentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**RegionalDeploymentId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRegionalDeploymentEvent

`func NewRegionalDeploymentEvent() *RegionalDeploymentEvent`

NewRegionalDeploymentEvent instantiates a new RegionalDeploymentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionalDeploymentEventWithDefaults

`func NewRegionalDeploymentEventWithDefaults() *RegionalDeploymentEvent`

NewRegionalDeploymentEventWithDefaults instantiates a new RegionalDeploymentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionalDeploymentEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionalDeploymentEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionalDeploymentEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegionalDeploymentEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *RegionalDeploymentEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *RegionalDeploymentEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *RegionalDeploymentEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *RegionalDeploymentEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *RegionalDeploymentEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RegionalDeploymentEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RegionalDeploymentEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RegionalDeploymentEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRegionalDeploymentId

`func (o *RegionalDeploymentEvent) GetRegionalDeploymentId() string`

GetRegionalDeploymentId returns the RegionalDeploymentId field if non-nil, zero value otherwise.

### GetRegionalDeploymentIdOk

`func (o *RegionalDeploymentEvent) GetRegionalDeploymentIdOk() (*string, bool)`

GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeploymentId

`func (o *RegionalDeploymentEvent) SetRegionalDeploymentId(v string)`

SetRegionalDeploymentId sets RegionalDeploymentId field to given value.

### HasRegionalDeploymentId

`func (o *RegionalDeploymentEvent) HasRegionalDeploymentId() bool`

HasRegionalDeploymentId returns a boolean if a field has been set.

### GetType

`func (o *RegionalDeploymentEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegionalDeploymentEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegionalDeploymentEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegionalDeploymentEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *RegionalDeploymentEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegionalDeploymentEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegionalDeploymentEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegionalDeploymentEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *RegionalDeploymentEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegionalDeploymentEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegionalDeploymentEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegionalDeploymentEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


