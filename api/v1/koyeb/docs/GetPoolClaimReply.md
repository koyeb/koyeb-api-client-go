# GetPoolClaimReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | Pointer to [**PoolClaim**](PoolClaim.md) |  | [optional] 

## Methods

### NewGetPoolClaimReply

`func NewGetPoolClaimReply() *GetPoolClaimReply`

NewGetPoolClaimReply instantiates a new GetPoolClaimReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoolClaimReplyWithDefaults

`func NewGetPoolClaimReplyWithDefaults() *GetPoolClaimReply`

NewGetPoolClaimReplyWithDefaults instantiates a new GetPoolClaimReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *GetPoolClaimReply) GetClaim() PoolClaim`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *GetPoolClaimReply) GetClaimOk() (*PoolClaim, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *GetPoolClaimReply) SetClaim(v PoolClaim)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *GetPoolClaimReply) HasClaim() bool`

HasClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


