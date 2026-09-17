# \ServicePoolsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServicePool**](ServicePoolsApi.md#CreateServicePool) | **Post** /v1/service_pools | Create a ServicePool
[**GetServicePool**](ServicePoolsApi.md#GetServicePool) | **Get** /v1/service_pools/{id} | Get a ServicePool
[**ListServicePools**](ServicePoolsApi.md#ListServicePools) | **Get** /v1/service_pools | List ServicePools



## CreateServicePool

> CreateServicePoolReply CreateServicePool(ctx).ServicePool(servicePool).Execute()

Create a ServicePool

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
    servicePool := *openapiclient.NewCreateServicePool() // CreateServicePool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicePoolsApi.CreateServicePool(context.Background()).ServicePool(servicePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePoolsApi.CreateServicePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServicePool`: CreateServicePoolReply
    fmt.Fprintf(os.Stdout, "Response from `ServicePoolsApi.CreateServicePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServicePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicePool** | [**CreateServicePool**](CreateServicePool.md) |  | 

### Return type

[**CreateServicePoolReply**](CreateServicePoolReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePool

> GetServicePoolReply GetServicePool(ctx, id).Execute()

Get a ServicePool

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicePoolsApi.GetServicePool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePoolsApi.GetServicePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicePool`: GetServicePoolReply
    fmt.Fprintf(os.Stdout, "Response from `ServicePoolsApi.GetServicePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServicePoolReply**](GetServicePoolReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicePools

> ListServicePoolsReply ListServicePools(ctx).Name(name).Limit(limit).Offset(offset).Execute()

List ServicePools

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
    name := "name_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicePoolsApi.ListServicePools(context.Background()).Name(name).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePoolsApi.ListServicePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServicePools`: ListServicePoolsReply
    fmt.Fprintf(os.Stdout, "Response from `ServicePoolsApi.ListServicePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListServicePoolsReply**](ListServicePoolsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

