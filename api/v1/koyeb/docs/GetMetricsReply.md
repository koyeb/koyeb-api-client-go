# GetMetricsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]GetMetricsReplyMetric**](GetMetricsReplyMetric.md) |  | [optional] 

## Methods

### NewGetMetricsReply

`func NewGetMetricsReply() *GetMetricsReply`

NewGetMetricsReply instantiates a new GetMetricsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricsReplyWithDefaults

`func NewGetMetricsReplyWithDefaults() *GetMetricsReply`

NewGetMetricsReplyWithDefaults instantiates a new GetMetricsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetMetricsReply) GetMetrics() []GetMetricsReplyMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetMetricsReply) GetMetricsOk() (*[]GetMetricsReplyMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetMetricsReply) SetMetrics(v []GetMetricsReplyMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetMetricsReply) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


