# DeploymentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**When** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeploymentEvent

`func NewDeploymentEvent() *DeploymentEvent`

NewDeploymentEvent instantiates a new DeploymentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentEventWithDefaults

`func NewDeploymentEventWithDefaults() *DeploymentEvent`

NewDeploymentEventWithDefaults instantiates a new DeploymentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhen

`func (o *DeploymentEvent) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *DeploymentEvent) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *DeploymentEvent) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *DeploymentEvent) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DeploymentEvent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeploymentEvent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeploymentEvent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeploymentEvent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *DeploymentEvent) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentEvent) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentEvent) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentEvent) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetType

`func (o *DeploymentEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *DeploymentEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeploymentEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *DeploymentEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeploymentEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeploymentEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeploymentEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


