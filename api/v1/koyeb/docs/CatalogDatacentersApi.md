# \CatalogDatacentersApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatacenters**](CatalogDatacentersApi.md#ListDatacenters) | **Get** /v1/catalog/datacenters | List datacenters



## ListDatacenters

> ListDatacentersReply ListDatacenters(ctx).Execute()

List datacenters

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/koyeb/koyeb-api-client-go/koyeb"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogDatacentersApi.ListDatacenters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDatacentersApi.ListDatacenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatacenters`: ListDatacentersReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogDatacentersApi.ListDatacenters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDatacentersRequest struct via the builder pattern


### Return type

[**ListDatacentersReply**](ListDatacentersReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

