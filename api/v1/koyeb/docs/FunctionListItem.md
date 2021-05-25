# FunctionListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**FunctionType**](FunctionType.md) |  | [optional] [default to FUNCTIONTYPE_DOCKER]

## Methods

### NewFunctionListItem

`func NewFunctionListItem() *FunctionListItem`

NewFunctionListItem instantiates a new FunctionListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionListItemWithDefaults

`func NewFunctionListItemWithDefaults() *FunctionListItem`

NewFunctionListItemWithDefaults instantiates a new FunctionListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FunctionListItem) GetType() FunctionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionListItem) GetTypeOk() (*FunctionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionListItem) SetType(v FunctionType)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionListItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


