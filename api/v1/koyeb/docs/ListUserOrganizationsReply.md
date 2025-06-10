# ListUserOrganizationsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**HasNext** | Pointer to **bool** |  | [optional] 

## Methods

### NewListUserOrganizationsReply

`func NewListUserOrganizationsReply() *ListUserOrganizationsReply`

NewListUserOrganizationsReply instantiates a new ListUserOrganizationsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserOrganizationsReplyWithDefaults

`func NewListUserOrganizationsReplyWithDefaults() *ListUserOrganizationsReply`

NewListUserOrganizationsReplyWithDefaults instantiates a new ListUserOrganizationsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *ListUserOrganizationsReply) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ListUserOrganizationsReply) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ListUserOrganizationsReply) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ListUserOrganizationsReply) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetLimit

`func (o *ListUserOrganizationsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListUserOrganizationsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListUserOrganizationsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListUserOrganizationsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListUserOrganizationsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListUserOrganizationsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListUserOrganizationsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListUserOrganizationsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetHasNext

`func (o *ListUserOrganizationsReply) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListUserOrganizationsReply) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListUserOrganizationsReply) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *ListUserOrganizationsReply) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


