# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryLogs**](LogsApi.md#QueryLogs) | **Get** /v1/streams/logs/query | Query logs
[**TailLogs**](LogsApi.md#TailLogs) | **Get** /v1/streams/logs/tail | Tails logs



## QueryLogs

> QueryLogsReply QueryLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).InstanceIds(instanceIds).Stream(stream).Streams(streams).Start(start).End(end).Order(order).Limit(limit).Regex(regex).Text(text).Regions(regions).Execute()

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
    type_ := "type__example" // string | Type of logs to retrieve, either \"build\" or \"runtime\". Defaults to \"runtime\". (optional)
    appId := "appId_example" // string | (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    serviceId := "serviceId_example" // string | (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    deploymentId := "deploymentId_example" // string | (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    regionalDeploymentId := "regionalDeploymentId_example" // string | (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    instanceId := "instanceId_example" // string | Deprecated, prefer using instance_ids instead. (optional)
    instanceIds := []string{"Inner_example"} // []string | (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    stream := "stream_example" // string | Deprecated, prefer using streams instead. (optional)
    streams := []string{"Inner_example"} // []string | (Optional) Filter on stream: either \"stdout\", \"stderr\" or \"koyeb\" (for system logs). (optional)
    start := time.Now() // time.Time | (Optional) Must always be before `end`. Defaults to 15 minutes ago. (optional)
    end := time.Now() // time.Time | (Optional) Must always be after `start`. Defaults to now. (optional)
    order := "order_example" // string | (Optional) `asc` or `desc`. Defaults to `desc`. (optional)
    limit := "limit_example" // string | (Optional) Defaults to 100. Maximum of 1000. (optional)
    regex := "regex_example" // string | (Optional) Apply a regex to filter logs. Can't be used with `text`. (optional)
    text := "text_example" // string | (Optional) Looks for this string in logs. Can't be used with `regex`. (optional)
    regions := []string{"Inner_example"} // []string | (Optional) Filter on the provided regions (e.g. [\"fra\", \"was\"]). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.QueryLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).InstanceIds(instanceIds).Stream(stream).Streams(streams).Start(start).End(end).Order(order).Limit(limit).Regex(regex).Text(text).Regions(regions).Execute()
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
 **type_** | **string** | Type of logs to retrieve, either \&quot;build\&quot; or \&quot;runtime\&quot;. Defaults to \&quot;runtime\&quot;. | 
 **appId** | **string** | (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **serviceId** | **string** | (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **deploymentId** | **string** | (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **regionalDeploymentId** | **string** | (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **instanceId** | **string** | Deprecated, prefer using instance_ids instead. | 
 **instanceIds** | **[]string** | (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **stream** | **string** | Deprecated, prefer using streams instead. | 
 **streams** | **[]string** | (Optional) Filter on stream: either \&quot;stdout\&quot;, \&quot;stderr\&quot; or \&quot;koyeb\&quot; (for system logs). | 
 **start** | **time.Time** | (Optional) Must always be before &#x60;end&#x60;. Defaults to 15 minutes ago. | 
 **end** | **time.Time** | (Optional) Must always be after &#x60;start&#x60;. Defaults to now. | 
 **order** | **string** | (Optional) &#x60;asc&#x60; or &#x60;desc&#x60;. Defaults to &#x60;desc&#x60;. | 
 **limit** | **string** | (Optional) Defaults to 100. Maximum of 1000. | 
 **regex** | **string** | (Optional) Apply a regex to filter logs. Can&#39;t be used with &#x60;text&#x60;. | 
 **text** | **string** | (Optional) Looks for this string in logs. Can&#39;t be used with &#x60;regex&#x60;. | 
 **regions** | **[]string** | (Optional) Filter on the provided regions (e.g. [\&quot;fra\&quot;, \&quot;was\&quot;]). | 

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

> StreamResultOfLogEntry TailLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).InstanceIds(instanceIds).Stream(stream).Streams(streams).Start(start).Limit(limit).Regex(regex).Text(text).Regions(regions).Execute()

Tails logs

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
    type_ := "type__example" // string | Type of logs to retrieve, either \"build\" or \"runtime\". Defaults to \"runtime\". (optional)
    appId := "appId_example" // string | (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    serviceId := "serviceId_example" // string | (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    deploymentId := "deploymentId_example" // string | (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    regionalDeploymentId := "regionalDeploymentId_example" // string | (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    instanceId := "instanceId_example" // string | Deprecated, prefer using instance_ids instead. (optional)
    instanceIds := []string{"Inner_example"} // []string | (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. (optional)
    stream := "stream_example" // string | Deprecated, prefer using streams instead. (optional)
    streams := []string{"Inner_example"} // []string | (Optional) Filter on stream: either \"stdout\", \"stderr\" or \"koyeb\" (for system logs). (optional)
    start := time.Now() // time.Time | (Optional) Defaults to 24 hours ago. (optional)
    limit := "limit_example" // string | (Optional) Defaults to 1000. Maximum of 1000. (optional)
    regex := "regex_example" // string | (Optional) Apply a regex to filter logs. Can't be used with `text`. (optional)
    text := "text_example" // string | (Optional) Looks for this string in logs. Can't be used with `regex`. (optional)
    regions := []string{"Inner_example"} // []string | (Optional) Filter on the provided regions (e.g. [\"fra\", \"was\"]). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.TailLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).InstanceId(instanceId).InstanceIds(instanceIds).Stream(stream).Streams(streams).Start(start).Limit(limit).Regex(regex).Text(text).Regions(regions).Execute()
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
 **type_** | **string** | Type of logs to retrieve, either \&quot;build\&quot; or \&quot;runtime\&quot;. Defaults to \&quot;runtime\&quot;. | 
 **appId** | **string** | (Optional) Filter on the provided app_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **serviceId** | **string** | (Optional) Filter on the provided service_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **deploymentId** | **string** | (Optional) Filter on the provided deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **regionalDeploymentId** | **string** | (Optional) Filter on the provided regional_deployment_id. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **instanceId** | **string** | Deprecated, prefer using instance_ids instead. | 
 **instanceIds** | **[]string** | (Optional) Filter on the provided instance_ids. At least one of app_id, service_id, deployment_id, regional_deployment_id or instance_ids must be set. | 
 **stream** | **string** | Deprecated, prefer using streams instead. | 
 **streams** | **[]string** | (Optional) Filter on stream: either \&quot;stdout\&quot;, \&quot;stderr\&quot; or \&quot;koyeb\&quot; (for system logs). | 
 **start** | **time.Time** | (Optional) Defaults to 24 hours ago. | 
 **limit** | **string** | (Optional) Defaults to 1000. Maximum of 1000. | 
 **regex** | **string** | (Optional) Apply a regex to filter logs. Can&#39;t be used with &#x60;text&#x60;. | 
 **text** | **string** | (Optional) Looks for this string in logs. Can&#39;t be used with &#x60;regex&#x60;. | 
 **regions** | **[]string** | (Optional) Filter on the provided regions (e.g. [\&quot;fra\&quot;, \&quot;was\&quot;]). | 

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

