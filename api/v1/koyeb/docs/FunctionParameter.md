# FunctionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**FunctionParameterType**](FunctionParameterType.md) |  | [optional] [default to FUNCTIONPARAMETERTYPE_UNKNOWN]
**Default** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFunctionParameter

`func NewFunctionParameter() *FunctionParameter`

NewFunctionParameter instantiates a new FunctionParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionParameterWithDefaults

`func NewFunctionParameterWithDefaults() *FunctionParameter`

NewFunctionParameterWithDefaults instantiates a new FunctionParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *FunctionParameter) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FunctionParameter) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FunctionParameter) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FunctionParameter) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *FunctionParameter) GetType() FunctionParameterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionParameter) GetTypeOk() (*FunctionParameterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionParameter) SetType(v FunctionParameterType)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefault

`func (o *FunctionParameter) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FunctionParameter) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FunctionParameter) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FunctionParameter) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


