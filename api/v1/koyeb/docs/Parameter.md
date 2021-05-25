# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ParameterParameterType**](ParameterParameterType.md) |  | [optional] [default to PARAMETERPARAMETERTYPE_STRING]
**Sensitive** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **interface{}** |  | [optional] 
**Default** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewParameter

`func NewParameter() *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Parameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Parameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *Parameter) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Parameter) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Parameter) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Parameter) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDescription

`func (o *Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Parameter) GetType() ParameterParameterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Parameter) GetTypeOk() (*ParameterParameterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Parameter) SetType(v ParameterParameterType)`

SetType sets Type field to given value.

### HasType

`func (o *Parameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSensitive

`func (o *Parameter) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *Parameter) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *Parameter) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *Parameter) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetOptions

`func (o *Parameter) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Parameter) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Parameter) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Parameter) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefault

`func (o *Parameter) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Parameter) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Parameter) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Parameter) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


