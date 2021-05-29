# \LogsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TailInstanceLogs**](LogsApi.md#TailInstanceLogs) | **Get** /v1/apps/{app_id_or_name}/services/{service_id_or_name}/revisions/{revision_id}/instances/{instance_id}/logs/{stream}/tail | 



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

