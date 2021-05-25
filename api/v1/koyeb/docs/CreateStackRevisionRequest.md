# CreateStackRevisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** |  | [optional] 
**Yaml** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStackRevisionRequest

`func NewCreateStackRevisionRequest() *CreateStackRevisionRequest`

NewCreateStackRevisionRequest instantiates a new CreateStackRevisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStackRevisionRequestWithDefaults

`func NewCreateStackRevisionRequestWithDefaults() *CreateStackRevisionRequest`

NewCreateStackRevisionRequestWithDefaults instantiates a new CreateStackRevisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *CreateStackRevisionRequest) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CreateStackRevisionRequest) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CreateStackRevisionRequest) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CreateStackRevisionRequest) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetYaml

`func (o *CreateStackRevisionRequest) GetYaml() interface{}`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *CreateStackRevisionRequest) GetYamlOk() (*interface{}, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *CreateStackRevisionRequest) SetYaml(v interface{})`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *CreateStackRevisionRequest) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetMessage

`func (o *CreateStackRevisionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateStackRevisionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateStackRevisionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateStackRevisionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


