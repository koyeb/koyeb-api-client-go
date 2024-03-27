# GetMetricsReplyMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Samples** | Pointer to [**[]Sample**](Sample.md) |  | [optional] 

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

### GetSamples

`func (o *GetMetricsReplyMetric) GetSamples() []Sample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *GetMetricsReplyMetric) GetSamplesOk() (*[]Sample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *GetMetricsReplyMetric) SetSamples(v []Sample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *GetMetricsReplyMetric) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


