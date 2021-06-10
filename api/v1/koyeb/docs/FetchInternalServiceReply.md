# FetchInternalServiceReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**Expand** | Pointer to **bool** |  | [optional] 

## Methods

### NewFetchInternalServiceReply

`func NewFetchInternalServiceReply() *FetchInternalServiceReply`

NewFetchInternalServiceReply instantiates a new FetchInternalServiceReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchInternalServiceReplyWithDefaults

`func NewFetchInternalServiceReplyWithDefaults() *FetchInternalServiceReply`

NewFetchInternalServiceReplyWithDefaults instantiates a new FetchInternalServiceReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *FetchInternalServiceReply) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FetchInternalServiceReply) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FetchInternalServiceReply) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *FetchInternalServiceReply) HasService() bool`

HasService returns a boolean if a field has been set.

### GetExpand

`func (o *FetchInternalServiceReply) GetExpand() bool`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *FetchInternalServiceReply) GetExpandOk() (*bool, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *FetchInternalServiceReply) SetExpand(v bool)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *FetchInternalServiceReply) HasExpand() bool`

HasExpand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


