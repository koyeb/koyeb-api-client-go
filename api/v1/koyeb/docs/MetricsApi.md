# \MetricsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /v1/streams/metrics | 



## GetMetrics

> GetMetricsReply GetMetrics(ctx).ServiceId(serviceId).InstanceId(instanceId).Name(name).Start(start).End(end).Step(step).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | ID of the service to query instances metrics for. Ignored if instance_id is set. (optional)
    instanceId := "instanceId_example" // string | ID of the instance to query metrics for. (optional)
    name := "name_example" // string | Metric to query. (optional) (default to "UNKNOWN")
    start := time.Now() // time.Time | (Optional) Defaults to an hour prior to end. (optional)
    end := time.Now() // time.Time | (Optional) Defaults to now. (optional)
    step := "step_example" // string | (Optional) Must be a valid duration in hours (h) or minutes (m). Defaulst to 5m. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetrics(context.Background()).ServiceId(serviceId).InstanceId(instanceId).Name(name).Start(start).End(end).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: GetMetricsReply
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | ID of the service to query instances metrics for. Ignored if instance_id is set. | 
 **instanceId** | **string** | ID of the instance to query metrics for. | 
 **name** | **string** | Metric to query. | [default to &quot;UNKNOWN&quot;]
 **start** | **time.Time** | (Optional) Defaults to an hour prior to end. | 
 **end** | **time.Time** | (Optional) Defaults to now. | 
 **step** | **string** | (Optional) Must be a valid duration in hours (h) or minutes (m). Defaulst to 5m. | 

### Return type

[**GetMetricsReply**](GetMetricsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

