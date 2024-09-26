# ListRegionalDeploymentsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionalDeployments** | Pointer to [**[]RegionalDeploymentListItem**](RegionalDeploymentListItem.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListRegionalDeploymentsReply

`func NewListRegionalDeploymentsReply() *ListRegionalDeploymentsReply`

NewListRegionalDeploymentsReply instantiates a new ListRegionalDeploymentsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRegionalDeploymentsReplyWithDefaults

`func NewListRegionalDeploymentsReplyWithDefaults() *ListRegionalDeploymentsReply`

NewListRegionalDeploymentsReplyWithDefaults instantiates a new ListRegionalDeploymentsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionalDeployments

`func (o *ListRegionalDeploymentsReply) GetRegionalDeployments() []RegionalDeploymentListItem`

GetRegionalDeployments returns the RegionalDeployments field if non-nil, zero value otherwise.

### GetRegionalDeploymentsOk

`func (o *ListRegionalDeploymentsReply) GetRegionalDeploymentsOk() (*[]RegionalDeploymentListItem, bool)`

GetRegionalDeploymentsOk returns a tuple with the RegionalDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalDeployments

`func (o *ListRegionalDeploymentsReply) SetRegionalDeployments(v []RegionalDeploymentListItem)`

SetRegionalDeployments sets RegionalDeployments field to given value.

### HasRegionalDeployments

`func (o *ListRegionalDeploymentsReply) HasRegionalDeployments() bool`

HasRegionalDeployments returns a boolean if a field has been set.

### GetLimit

`func (o *ListRegionalDeploymentsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListRegionalDeploymentsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListRegionalDeploymentsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListRegionalDeploymentsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListRegionalDeploymentsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListRegionalDeploymentsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListRegionalDeploymentsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListRegionalDeploymentsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *ListRegionalDeploymentsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListRegionalDeploymentsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListRegionalDeploymentsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListRegionalDeploymentsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasNext

`func (o *ListRegionalDeploymentsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListRegionalDeploymentsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListRegionalDeploymentsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListRegionalDeploymentsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


