# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TailLogs**](LogsApi.md#TailLogs) | **Get** /v1/streams/logs/tail | Tails logs



## TailLogs

> StreamResultOfLogEntry TailLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()

Tails logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string |  (optional)
    appId := "appId_example" // string |  (optional)
    serviceId := "serviceId_example" // string |  (optional)
    deploymentId := "deploymentId_example" // string |  (optional)
    instanceId := "instanceId_example" // string |  (optional)
    stream := "stream_example" // string |  (optional)
    start := "start_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTailLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **appId** | **string** |  | 
 **serviceId** | **string** |  | 
 **deploymentId** | **string** |  | 
 **instanceId** | **string** |  | 
 **stream** | **string** |  | 
 **start** | **string** |  | 
 **limit** | **string** |  | 

### Return type

[**StreamResultOfLogEntry**](StreamResultOfLogEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

