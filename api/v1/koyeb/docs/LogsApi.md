# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TailInstanceLogs**](LogsApi.md#TailInstanceLogs) | **Get** /v1/apps/{app_id_or_name}/services/{service_id_or_name}/revisions/{revision_id}/instances/{instance_id}/logs/{stream}/tail | 
[**TailLogs**](LogsApi.md#TailLogs) | **Get** /v1/streams/logs | 
[**TailRevisionBuildLogs**](LogsApi.md#TailRevisionBuildLogs) | **Get** /v1/apps/{app_id_or_name}/services/{service_id_or_name}/revisions/{revision_id}/builds/tail | 
[**TailRuntimeLogs**](LogsApi.md#TailRuntimeLogs) | **Get** /v1/apps/{app_id_or_name}/logs/tail | 



## TailInstanceLogs

> StreamResultOfRevisionLogEntry TailInstanceLogs(ctx, appIdOrName, serviceIdOrName, revisionId, instanceId, stream).Offset(offset).Execute()



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
    appIdOrName := "appIdOrName_example" // string | 
    serviceIdOrName := "serviceIdOrName_example" // string | 
    revisionId := "revisionId_example" // string | 
    instanceId := "instanceId_example" // string | 
    stream := "stream_example" // string | 
    offset := "offset_example" // string | -1 is from the end, 0 is from the start anything else is from this offset in the file. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailInstanceLogs(context.Background(), appIdOrName, serviceIdOrName, revisionId, instanceId, stream).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailInstanceLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailInstanceLogs`: StreamResultOfRevisionLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailInstanceLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 
**serviceIdOrName** | **string** |  | 
**revisionId** | **string** |  | 
**instanceId** | **string** |  | 
**stream** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailInstanceLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **offset** | **string** | -1 is from the end, 0 is from the start anything else is from this offset in the file. | 

### Return type

[**StreamResultOfRevisionLogEntry**](StreamResultOfRevisionLogEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TailLogs

> StreamResultOfLogEntry TailLogs(ctx).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Start(start).Limit(limit).Execute()



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
    start := "start_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailLogs(context.Background()).Type_(type_).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).InstanceId(instanceId).Start(start).Limit(limit).Execute()
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


## TailRevisionBuildLogs

> StreamResultOfLogEntry TailRevisionBuildLogs(ctx, appIdOrName, serviceIdOrName, revisionId).Start(start).Limit(limit).Execute()



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
    appIdOrName := "appIdOrName_example" // string | 
    serviceIdOrName := "serviceIdOrName_example" // string | 
    revisionId := "revisionId_example" // string | 
    start := "start_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailRevisionBuildLogs(context.Background(), appIdOrName, serviceIdOrName, revisionId).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailRevisionBuildLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailRevisionBuildLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailRevisionBuildLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 
**serviceIdOrName** | **string** |  | 
**revisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailRevisionBuildLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## TailRuntimeLogs

> StreamResultOfLogEntry TailRuntimeLogs(ctx, appIdOrName).ServiceId(serviceId).RevisionId(revisionId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()



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
    appIdOrName := "appIdOrName_example" // string | 
    serviceId := "serviceId_example" // string |  (optional)
    revisionId := "revisionId_example" // string |  (optional)
    instanceId := "instanceId_example" // string |  (optional)
    stream := "stream_example" // string |  (optional)
    start := "start_example" // string | Temporary field. (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailRuntimeLogs(context.Background(), appIdOrName).ServiceId(serviceId).RevisionId(revisionId).InstanceId(instanceId).Stream(stream).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailRuntimeLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailRuntimeLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailRuntimeLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailRuntimeLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceId** | **string** |  | 
 **revisionId** | **string** |  | 
 **instanceId** | **string** |  | 
 **stream** | **string** |  | 
 **start** | **string** | Temporary field. | 
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

