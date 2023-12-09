# SecretsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** |  | [optional] 
**ByType** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSecretsSummary

`func NewSecretsSummary() *SecretsSummary`

NewSecretsSummary instantiates a new SecretsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsSummaryWithDefaults

`func NewSecretsSummaryWithDefaults() *SecretsSummary`

NewSecretsSummaryWithDefaults instantiates a new SecretsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SecretsSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SecretsSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SecretsSummary) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SecretsSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByType

`func (o *SecretsSummary) GetByType() map[string]string`

GetByType returns the ByType field if non-nil, zero value otherwise.

### GetByTypeOk

`func (o *SecretsSummary) GetByTypeOk() (*map[string]string, bool)`

GetByTypeOk returns a tuple with the ByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByType

`func (o *SecretsSummary) SetByType(v map[string]string)`

SetByType sets ByType field to given value.

### HasByType

`func (o *SecretsSummary) HasByType() bool`

HasByType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


