# ServiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** |  | [optional] 
**ByStatus** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceSummary

`func NewServiceSummary() *ServiceSummary`

NewServiceSummary instantiates a new ServiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSummaryWithDefaults

`func NewServiceSummaryWithDefaults() *ServiceSummary`

NewServiceSummaryWithDefaults instantiates a new ServiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ServiceSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ServiceSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ServiceSummary) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ServiceSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *ServiceSummary) GetByStatus() map[string]string`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *ServiceSummary) GetByStatusOk() (*map[string]string, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *ServiceSummary) SetByStatus(v map[string]string)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *ServiceSummary) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


