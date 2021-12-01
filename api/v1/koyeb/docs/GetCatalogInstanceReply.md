# GetCatalogInstanceReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**CatalogInstance**](CatalogInstance.md) |  | [optional] 

## Methods

### NewGetCatalogInstanceReply

`func NewGetCatalogInstanceReply() *GetCatalogInstanceReply`

NewGetCatalogInstanceReply instantiates a new GetCatalogInstanceReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogInstanceReplyWithDefaults

`func NewGetCatalogInstanceReplyWithDefaults() *GetCatalogInstanceReply`

NewGetCatalogInstanceReplyWithDefaults instantiates a new GetCatalogInstanceReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *GetCatalogInstanceReply) GetInstance() CatalogInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetCatalogInstanceReply) GetInstanceOk() (*CatalogInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetCatalogInstanceReply) SetInstance(v CatalogInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetCatalogInstanceReply) HasInstance() bool`

HasInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


