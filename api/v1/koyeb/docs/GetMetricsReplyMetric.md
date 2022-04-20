# GetMetricsReplyMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Values** | Pointer to **[][]interface{}** |  | [optional] 

## Methods

### NewGetMetricsReplyMetric

`func NewGetMetricsReplyMetric() *GetMetricsReplyMetric`

NewGetMetricsReplyMetric instantiates a new GetMetricsReplyMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricsReplyMetricWithDefaults

`func NewGetMetricsReplyMetricWithDefaults() *GetMetricsReplyMetric`

NewGetMetricsReplyMetricWithDefaults instantiates a new GetMetricsReplyMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *GetMetricsReplyMetric) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetMetricsReplyMetric) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetMetricsReplyMetric) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetMetricsReplyMetric) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetValues

`func (o *GetMetricsReplyMetric) GetValues() [][]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetMetricsReplyMetric) GetValuesOk() (*[][]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetMetricsReplyMetric) SetValues(v [][]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *GetMetricsReplyMetric) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


