# \CatalogInstancesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogInstance**](CatalogInstancesApi.md#GetCatalogInstance) | **Get** /v1/catalog/instances/{id} | Get Instance
[**ListCatalogInstances**](CatalogInstancesApi.md#ListCatalogInstances) | **Get** /v1/catalog/instances | List Instance



## GetCatalogInstance

> GetCatalogInstanceReply GetCatalogInstance(ctx, id).Execute()

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
    resp, r, err := api_client.CatalogInstancesApi.GetCatalogInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogInstancesApi.GetCatalogInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogInstance`: GetCatalogInstanceReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogInstancesApi.GetCatalogInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The name of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCatalogInstanceReply**](GetCatalogInstanceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogInstances

> ListCatalogInstancesReply ListCatalogInstances(ctx).Limit(limit).Offset(offset).Id(id).Execute()

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
    resp, r, err := api_client.CatalogInstancesApi.ListCatalogInstances(context.Background()).Limit(limit).Offset(offset).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogInstancesApi.ListCatalogInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogInstances`: ListCatalogInstancesReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogInstancesApi.ListCatalogInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **id** | **string** | (Optional) A filter for instances. | 

### Return type

[**ListCatalogInstancesReply**](ListCatalogInstancesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

