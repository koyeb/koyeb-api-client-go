# NeonPostgresSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** |  | [optional] 
**ByInstanceType** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewNeonPostgresSummary

`func NewNeonPostgresSummary() *NeonPostgresSummary`

NewNeonPostgresSummary instantiates a new NeonPostgresSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeonPostgresSummaryWithDefaults

`func NewNeonPostgresSummaryWithDefaults() *NeonPostgresSummary`

NewNeonPostgresSummaryWithDefaults instantiates a new NeonPostgresSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *NeonPostgresSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NeonPostgresSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NeonPostgresSummary) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NeonPostgresSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByInstanceType

`func (o *NeonPostgresSummary) GetByInstanceType() map[string]string`

GetByInstanceType returns the ByInstanceType field if non-nil, zero value otherwise.

### GetByInstanceTypeOk

`func (o *NeonPostgresSummary) GetByInstanceTypeOk() (*map[string]string, bool)`

GetByInstanceTypeOk returns a tuple with the ByInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInstanceType

`func (o *NeonPostgresSummary) SetByInstanceType(v map[string]string)`

SetByInstanceType sets ByInstanceType field to given value.

### HasByInstanceType

`func (o *NeonPostgresSummary) HasByInstanceType() bool`

HasByInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


