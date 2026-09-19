# ListPoolClaimReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]PoolClaim**](PoolClaim.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListPoolClaimReply

`func NewListPoolClaimReply() *ListPoolClaimReply`

NewListPoolClaimReply instantiates a new ListPoolClaimReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoolClaimReplyWithDefaults

`func NewListPoolClaimReplyWithDefaults() *ListPoolClaimReply`

NewListPoolClaimReplyWithDefaults instantiates a new ListPoolClaimReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ListPoolClaimReply) GetClaims() []PoolClaim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ListPoolClaimReply) GetClaimsOk() (*[]PoolClaim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ListPoolClaimReply) SetClaims(v []PoolClaim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ListPoolClaimReply) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLimit

`func (o *ListPoolClaimReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPoolClaimReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPoolClaimReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPoolClaimReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListPoolClaimReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPoolClaimReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPoolClaimReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPoolClaimReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListPoolClaimReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListPoolClaimReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListPoolClaimReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListPoolClaimReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *ListPoolClaimReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListPoolClaimReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListPoolClaimReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListPoolClaimReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


