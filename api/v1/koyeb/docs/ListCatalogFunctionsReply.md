# ListCatalogFunctionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogFunctions** | Pointer to [**[]CatalogFunctionListItem**](CatalogFunctionListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCatalogFunctionsReply

`func NewListCatalogFunctionsReply() *ListCatalogFunctionsReply`

NewListCatalogFunctionsReply instantiates a new ListCatalogFunctionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogFunctionsReplyWithDefaults

`func NewListCatalogFunctionsReplyWithDefaults() *ListCatalogFunctionsReply`

NewListCatalogFunctionsReplyWithDefaults instantiates a new ListCatalogFunctionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogFunctions

`func (o *ListCatalogFunctionsReply) GetCatalogFunctions() []CatalogFunctionListItem`

GetCatalogFunctions returns the CatalogFunctions field if non-nil, zero value otherwise.

### GetCatalogFunctionsOk

`func (o *ListCatalogFunctionsReply) GetCatalogFunctionsOk() (*[]CatalogFunctionListItem, bool)`

GetCatalogFunctionsOk returns a tuple with the CatalogFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogFunctions

`func (o *ListCatalogFunctionsReply) SetCatalogFunctions(v []CatalogFunctionListItem)`

SetCatalogFunctions sets CatalogFunctions field to given value.

### HasCatalogFunctions

`func (o *ListCatalogFunctionsReply) HasCatalogFunctions() bool`

HasCatalogFunctions returns a boolean if a field has been set.

### GetLimit

`func (o *ListCatalogFunctionsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCatalogFunctionsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCatalogFunctionsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCatalogFunctionsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCatalogFunctionsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCatalogFunctionsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCatalogFunctionsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCatalogFunctionsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListCatalogFunctionsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCatalogFunctionsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCatalogFunctionsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCatalogFunctionsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


