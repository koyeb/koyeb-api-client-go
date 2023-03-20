# \QuotasApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReviewOrganizationCapacity**](QuotasApi.md#ReviewOrganizationCapacity) | **Post** /v1/quotas/capacity | Review Organization Capacity



## ReviewOrganizationCapacity

> ReviewOrganizationCapacityReply ReviewOrganizationCapacity(ctx).Body(body).Execute()

Review Organization Capacity

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
    body := *openapiclient.NewReviewOrganizationCapacityRequest() // ReviewOrganizationCapacityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotasApi.ReviewOrganizationCapacity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotasApi.ReviewOrganizationCapacity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReviewOrganizationCapacity`: ReviewOrganizationCapacityReply
    fmt.Fprintf(os.Stdout, "Response from `QuotasApi.ReviewOrganizationCapacity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewOrganizationCapacityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReviewOrganizationCapacityRequest**](ReviewOrganizationCapacityRequest.md) |  | 

### Return type

[**ReviewOrganizationCapacityReply**](ReviewOrganizationCapacityReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

