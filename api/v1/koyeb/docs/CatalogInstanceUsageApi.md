# \CatalogInstanceUsageApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsage**](CatalogInstanceUsageApi.md#ListUsage) | **Get** /v1/catalog/usage | 



## ListUsage

> ListUsageReply ListUsage(ctx).Region(region).Execute()



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
    region := "region_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogInstanceUsageApi.ListUsage(context.Background()).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogInstanceUsageApi.ListUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsage`: ListUsageReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogInstanceUsageApi.ListUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 

### Return type

[**ListUsageReply**](ListUsageReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

