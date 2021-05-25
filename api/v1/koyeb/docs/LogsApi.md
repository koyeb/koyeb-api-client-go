# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TailConnectorLogs**](LogsApi.md#TailConnectorLogs) | **Get** /v1/connectors/{idOrName}/logs/tail | 
[**TailInstanceLogs**](LogsApi.md#TailInstanceLogs) | **Get** /v1/apps/{app_id_or_name}/services/{service_id_or_name}/revisions/{revision_id}/instances/{instance_id}/logs/{stream}/tail | 
[**TailStackEvents**](LogsApi.md#TailStackEvents) | **Get** /v1/stacks/{stack_id}/events/tail | 
[**TailStackRevisionBuildLogs**](LogsApi.md#TailStackRevisionBuildLogs) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/builds/tail | 
[**TailStackRevisionLogs**](LogsApi.md#TailStackRevisionLogs) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/logs/tail | 
[**TailStackRevisionLogsForFunction**](LogsApi.md#TailStackRevisionLogsForFunction) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/functions/{fn_name}/logs/tail | 



## TailConnectorLogs

> StreamResultOfLogEntry TailConnectorLogs(ctx, idOrName).Start(start).Execute()



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
    idOrName := "idOrName_example" // string | The id or name of the connector to tail
    start := "start_example" // string | A timestamp to indicate when to pull the logs from. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailConnectorLogs(context.Background(), idOrName).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailConnectorLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailConnectorLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailConnectorLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** | The id or name of the connector to tail | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailConnectorLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** | A timestamp to indicate when to pull the logs from. | 

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


## TailInstanceLogs

> StreamResultOfLogEntry TailInstanceLogs(ctx, appIdOrName, serviceIdOrName, revisionId, instanceId, stream).Offset(offset).Execute()



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
    // response from `TailInstanceLogs`: StreamResultOfLogEntry
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

[**StreamResultOfLogEntry**](StreamResultOfLogEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TailStackEvents

> StreamResultOfEvent TailStackEvents(ctx, stackId).Start(start).Execute()



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
    stackId := "stackId_example" // string | 
    start := "start_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailStackEvents(context.Background(), stackId).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailStackEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailStackEvents`: StreamResultOfEvent
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailStackEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailStackEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** |  | 

### Return type

[**StreamResultOfEvent**](StreamResultOfEvent.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TailStackRevisionBuildLogs

> StreamResultOfLogEntry TailStackRevisionBuildLogs(ctx, stackId, sha).Start(start).Execute()



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
    stackId := "stackId_example" // string | The name of the stack
    sha := "sha_example" // string | The sha or _latest to indicate the latest revision
    start := "start_example" // string | A timestamp to indicate when to pull the logs from. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailStackRevisionBuildLogs(context.Background(), stackId, sha).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailStackRevisionBuildLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailStackRevisionBuildLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailStackRevisionBuildLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | The name of the stack | 
**sha** | **string** | The sha or _latest to indicate the latest revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailStackRevisionBuildLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **string** | A timestamp to indicate when to pull the logs from. | 

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


## TailStackRevisionLogs

> StreamResultOfLogEntry TailStackRevisionLogs(ctx, stackId, sha).Start(start).Execute()



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
    stackId := "stackId_example" // string | The name of the stack
    sha := "sha_example" // string | The sha or _latest to indicate the latest revision
    start := "start_example" // string | A timestamp to indicate when to pull the logs from. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailStackRevisionLogs(context.Background(), stackId, sha).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailStackRevisionLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailStackRevisionLogs`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailStackRevisionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | The name of the stack | 
**sha** | **string** | The sha or _latest to indicate the latest revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailStackRevisionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **string** | A timestamp to indicate when to pull the logs from. | 

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


## TailStackRevisionLogsForFunction

> StreamResultOfLogEntry TailStackRevisionLogsForFunction(ctx, stackId, sha, fnName).Start(start).EventId(eventId).Execute()



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
    stackId := "stackId_example" // string | The name of the stack
    sha := "sha_example" // string | The sha or _latest to indicate the latest revision
    fnName := "fnName_example" // string | The name of the function
    start := "start_example" // string | A timestamp. (optional)
    eventId := "eventId_example" // string | An optional event_id to filter on to only see runs from this id. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.TailStackRevisionLogsForFunction(context.Background(), stackId, sha, fnName).Start(start).EventId(eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.TailStackRevisionLogsForFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailStackRevisionLogsForFunction`: StreamResultOfLogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.TailStackRevisionLogsForFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** | The name of the stack | 
**sha** | **string** | The sha or _latest to indicate the latest revision | 
**fnName** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailStackRevisionLogsForFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **string** | A timestamp. | 
 **eventId** | **string** | An optional event_id to filter on to only see runs from this id. | 

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

