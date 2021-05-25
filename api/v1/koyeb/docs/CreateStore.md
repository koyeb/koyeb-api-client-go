# CreateStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**With** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateStore

`func NewCreateStore() *CreateStore`

NewCreateStore instantiates a new CreateStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoreWithDefaults

`func NewCreateStoreWithDefaults() *CreateStore`

NewCreateStoreWithDefaults instantiates a new CreateStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CreateStore) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateStore) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateStore) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateStore) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *CreateStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateStore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWith

`func (o *CreateStore) GetWith() map[string]interface{}`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *CreateStore) GetWithOk() (*map[string]interface{}, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *CreateStore) SetWith(v map[string]interface{})`

SetWith sets With field to given value.

### HasWith

`func (o *CreateStore) HasWith() bool`

HasWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


