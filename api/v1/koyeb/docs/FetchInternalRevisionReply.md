# FetchInternalRevisionReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**ServiceRevision**](ServiceRevision.md) |  | [optional] 

## Methods

### NewFetchInternalRevisionReply

`func NewFetchInternalRevisionReply() *FetchInternalRevisionReply`

NewFetchInternalRevisionReply instantiates a new FetchInternalRevisionReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchInternalRevisionReplyWithDefaults

`func NewFetchInternalRevisionReplyWithDefaults() *FetchInternalRevisionReply`

NewFetchInternalRevisionReplyWithDefaults instantiates a new FetchInternalRevisionReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *FetchInternalRevisionReply) GetRevision() ServiceRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FetchInternalRevisionReply) GetRevisionOk() (*ServiceRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FetchInternalRevisionReply) SetRevision(v ServiceRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FetchInternalRevisionReply) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


