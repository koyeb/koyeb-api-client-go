# ExecCommandIO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is base64 encoded | [optional] 
**Close** | Pointer to **bool** | Indicate last data frame | [optional] 

## Methods

### NewExecCommandIO

`func NewExecCommandIO() *ExecCommandIO`

NewExecCommandIO instantiates a new ExecCommandIO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCommandIOWithDefaults

`func NewExecCommandIOWithDefaults() *ExecCommandIO`

NewExecCommandIOWithDefaults instantiates a new ExecCommandIO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExecCommandIO) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecCommandIO) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecCommandIO) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ExecCommandIO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetClose

`func (o *ExecCommandIO) GetClose() bool`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *ExecCommandIO) GetCloseOk() (*bool, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *ExecCommandIO) SetClose(v bool)`

SetClose sets Close field to given value.

### HasClose

`func (o *ExecCommandIO) HasClose() bool`

HasClose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


