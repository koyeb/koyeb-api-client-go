# ConnectorUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ChangeUrl** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] [default to CONNECTORTYPE_UNKNOWN]
**With** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConnectorUpsert

`func NewConnectorUpsert() *ConnectorUpsert`

NewConnectorUpsert instantiates a new ConnectorUpsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorUpsertWithDefaults

`func NewConnectorUpsertWithDefaults() *ConnectorUpsert`

NewConnectorUpsertWithDefaults instantiates a new ConnectorUpsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorUpsert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorUpsert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorUpsert) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorUpsert) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChangeUrl

`func (o *ConnectorUpsert) GetChangeUrl() bool`

GetChangeUrl returns the ChangeUrl field if non-nil, zero value otherwise.

### GetChangeUrlOk

`func (o *ConnectorUpsert) GetChangeUrlOk() (*bool, bool)`

GetChangeUrlOk returns a tuple with the ChangeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeUrl

`func (o *ConnectorUpsert) SetChangeUrl(v bool)`

SetChangeUrl sets ChangeUrl field to given value.

### HasChangeUrl

`func (o *ConnectorUpsert) HasChangeUrl() bool`

HasChangeUrl returns a boolean if a field has been set.

### GetType

`func (o *ConnectorUpsert) GetType() ConnectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorUpsert) GetTypeOk() (*ConnectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorUpsert) SetType(v ConnectorType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectorUpsert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWith

`func (o *ConnectorUpsert) GetWith() map[string]interface{}`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *ConnectorUpsert) GetWithOk() (*map[string]interface{}, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *ConnectorUpsert) SetWith(v map[string]interface{})`

SetWith sets With field to given value.

### HasWith

`func (o *ConnectorUpsert) HasWith() bool`

HasWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


