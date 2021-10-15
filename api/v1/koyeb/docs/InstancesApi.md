# \InstancesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecCommand**](InstancesApi.md#ExecCommand) | **Post** /v1/instances/exec | Exec Command
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /v1/instances/{id} | Get Instance
[**ListInstances**](InstancesApi.md#ListInstances) | **Get** /v1/instances | List Instance



## ExecCommand

> StreamResultOfExecCommandReply ExecCommand(ctx).Body(body).Id(id).Execute()

Exec Command

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
    body := *openapiclient.NewExecCommandRequestBody() // ExecCommandRequestBody |  (streaming inputs)
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ExecCommand(context.Background()).Body(body).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExecCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecCommand`: StreamResultOfExecCommandReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExecCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExecCommandRequestBody**](ExecCommandRequestBody.md) |  (streaming inputs) | 
 **id** | **string** |  | 

### Return type

[**StreamResultOfExecCommandReply**](StreamResultOfExecCommandReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> GetInstanceReply GetInstance(ctx, id).Execute()

Get Instance

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
    id := "id_example" // string | The name of the instance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: GetInstanceReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The name of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceReply**](GetInstanceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> ListInstancesReply ListInstances(ctx).Limit(limit).Offset(offset).Id(id).Execute()

List Instance

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
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    id := "id_example" // string | (Optional) A filter for instances. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListInstances(context.Background()).Limit(limit).Offset(offset).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstances`: ListInstancesReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **id** | **string** | (Optional) A filter for instances. | 

### Return type

[**ListInstancesReply**](ListInstancesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

