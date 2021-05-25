# GetCatalogFunctionsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogFunction** | Pointer to [**CatalogFunction**](CatalogFunction.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetCatalogFunctionsReply

`func NewGetCatalogFunctionsReply() *GetCatalogFunctionsReply`

NewGetCatalogFunctionsReply instantiates a new GetCatalogFunctionsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogFunctionsReplyWithDefaults

`func NewGetCatalogFunctionsReplyWithDefaults() *GetCatalogFunctionsReply`

NewGetCatalogFunctionsReplyWithDefaults instantiates a new GetCatalogFunctionsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogFunction

`func (o *GetCatalogFunctionsReply) GetCatalogFunction() CatalogFunction`

GetCatalogFunction returns the CatalogFunction field if non-nil, zero value otherwise.

### GetCatalogFunctionOk

`func (o *GetCatalogFunctionsReply) GetCatalogFunctionOk() (*CatalogFunction, bool)`

GetCatalogFunctionOk returns a tuple with the CatalogFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogFunction

`func (o *GetCatalogFunctionsReply) SetCatalogFunction(v CatalogFunction)`

SetCatalogFunction sets CatalogFunction field to given value.

### HasCatalogFunction

`func (o *GetCatalogFunctionsReply) HasCatalogFunction() bool`

HasCatalogFunction returns a boolean if a field has been set.

### GetVersions

`func (o *GetCatalogFunctionsReply) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetCatalogFunctionsReply) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetCatalogFunctionsReply) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetCatalogFunctionsReply) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


