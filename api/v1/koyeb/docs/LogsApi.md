# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryLogs**](LogsApi.md#QueryLogs) | **Get** /v1/streams/logs/query | Query logs
[**TailLogs**](LogsApi.md#TailLogs) | **Get** /v1/streams/logs/tail | Tails logs



## QueryLogs

> QueryLogsReply QueryLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Stream(stream).RegionalDeploymentId(regionalDeploymentId).Start(start).End(end).Order(order).Limit(limit).Regex(regex).Text(text).Execute()

Query logs

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
    type_ := "type__example" // string |  (optional)
    appId := "appId_example" // string |  (optional)
    serviceId := "serviceId_example" // string |  (optional)
    deploymentId := "deploymentId_example" // string |  (optional)
    instanceId := "instanceId_example" // string |  (optional)
    stream := "stream_example" // string |  (optional)
    regionalDeploymentId := "regionalDeploymentId_example" // string |  (optional)
    start := time.Now() // time.Time | (Optional) Must always be before `end`. Defaults to 15 minutes ago. (optional)
    end := time.Now() // time.Time | (Optional) Must always be after `start`. Defaults to now. (optional)
    order := "order_example" // string | (Optional) `asc` or `desc`. Defaults to `desc`. (optional)
    limit := "limit_example" // string | (Optional) Defaults to 100. Maximum of 1000. (optional)
    regex := "regex_example" // string | (Optional) Apply a regex to filter logs. Can't be used with `text`. (optional)
    text := "text_example" // string | (Optional) Looks for this string in logs. Can't be used with `regex`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.QueryLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Stream(stream).RegionalDeploymentId(regionalDeploymentId).Start(start).End(end).Order(order).Limit(limit).Regex(regex).Text(text).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.QueryLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryLogs`: QueryLogsReply
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.QueryLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **appId** | **string** |  | 
 **serviceId** | **string** |  | 
 **deploymentId** | **string** |  | 
 **instanceId** | **string** |  | 
 **stream** | **string** |  | 
 **regionalDeploymentId** | **string** |  | 
 **start** | **time.Time** | (Optional) Must always be before &#x60;end&#x60;. Defaults to 15 minutes ago. | 
 **end** | **time.Time** | (Optional) Must always be after &#x60;start&#x60;. Defaults to now. | 
 **order** | **string** | (Optional) &#x60;asc&#x60; or &#x60;desc&#x60;. Defaults to &#x60;desc&#x60;. | 
 **limit** | **string** | (Optional) Defaults to 100. Maximum of 1000. | 
 **regex** | **string** | (Optional) Apply a regex to filter logs. Can&#39;t be used with &#x60;text&#x60;. | 
 **text** | **string** | (Optional) Looks for this string in logs. Can&#39;t be used with &#x60;regex&#x60;. | 

### Return type

[**QueryLogsReply**](QueryLogsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TailLogs

> StreamResultOfLogEntry TailLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()

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
    regionalDeploymentId := "regionalDeploymentId_example" // string |  (optional)
    instanceId := "instanceId_example" // string |  (optional)
    stream := "stream_example" // string |  (optional)
    start := "start_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.TailLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()
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
 **regionalDeploymentId** | **string** |  | 
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

