# GetOrganizationUsageDetailsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageDetails** | Pointer to [**[]UsageDetails**](UsageDetails.md) |  | [optional] 
**DatabaseDetails** | Pointer to [**[]DatabaseUsageDetails**](DatabaseUsageDetails.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOrganizationUsageDetailsReply

`func NewGetOrganizationUsageDetailsReply() *GetOrganizationUsageDetailsReply`

NewGetOrganizationUsageDetailsReply instantiates a new GetOrganizationUsageDetailsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationUsageDetailsReplyWithDefaults

`func NewGetOrganizationUsageDetailsReplyWithDefaults() *GetOrganizationUsageDetailsReply`

NewGetOrganizationUsageDetailsReplyWithDefaults instantiates a new GetOrganizationUsageDetailsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageDetails

`func (o *GetOrganizationUsageDetailsReply) GetUsageDetails() []UsageDetails`

GetUsageDetails returns the UsageDetails field if non-nil, zero value otherwise.

### GetUsageDetailsOk

`func (o *GetOrganizationUsageDetailsReply) GetUsageDetailsOk() (*[]UsageDetails, bool)`

GetUsageDetailsOk returns a tuple with the UsageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDetails

`func (o *GetOrganizationUsageDetailsReply) SetUsageDetails(v []UsageDetails)`

SetUsageDetails sets UsageDetails field to given value.

### HasUsageDetails

`func (o *GetOrganizationUsageDetailsReply) HasUsageDetails() bool`

HasUsageDetails returns a boolean if a field has been set.

### GetDatabaseDetails

`func (o *GetOrganizationUsageDetailsReply) GetDatabaseDetails() []DatabaseUsageDetails`

GetDatabaseDetails returns the DatabaseDetails field if non-nil, zero value otherwise.

### GetDatabaseDetailsOk

`func (o *GetOrganizationUsageDetailsReply) GetDatabaseDetailsOk() (*[]DatabaseUsageDetails, bool)`

GetDatabaseDetailsOk returns a tuple with the DatabaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseDetails

`func (o *GetOrganizationUsageDetailsReply) SetDatabaseDetails(v []DatabaseUsageDetails)`

SetDatabaseDetails sets DatabaseDetails field to given value.

### HasDatabaseDetails

`func (o *GetOrganizationUsageDetailsReply) HasDatabaseDetails() bool`

HasDatabaseDetails returns a boolean if a field has been set.

### GetLimit

`func (o *GetOrganizationUsageDetailsReply) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetOrganizationUsageDetailsReply) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetOrganizationUsageDetailsReply) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetOrganizationUsageDetailsReply) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetOrganizationUsageDetailsReply) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetOrganizationUsageDetailsReply) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetOrganizationUsageDetailsReply) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetOrganizationUsageDetailsReply) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *GetOrganizationUsageDetailsReply) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetOrganizationUsageDetailsReply) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetOrganizationUsageDetailsReply) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetOrganizationUsageDetailsReply) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrder

`func (o *GetOrganizationUsageDetailsReply) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetOrganizationUsageDetailsReply) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetOrganizationUsageDetailsReply) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetOrganizationUsageDetailsReply) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


