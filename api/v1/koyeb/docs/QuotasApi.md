# \QuotasApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationQuotasUsage**](QuotasApi.md#GetOrganizationQuotasUsage) | **Get** /v1/quotas/organizations/{organization_id}/usage | Return the organization&#39;s current quota usage alongside the plan&#39;s limits. Response is cached in Redis for 60s. Accept text/plain (or ?format&#x3D;prometheus) to receive the response rendered as Prometheus text exposition format (\&quot;koyeb_quota_&lt;x&gt;\&quot; for usage, \&quot;koyeb_quota_&lt;x&gt;_limit\&quot; for the plan limit) suitable for scraping into an external Prometheus.
[**ReviewOrganizationCapacity**](QuotasApi.md#ReviewOrganizationCapacity) | **Post** /v1/quotas/capacity | DEPRECATED: Review Organization Capacity



## GetOrganizationQuotasUsage

> GetOrganizationQuotasUsageReply GetOrganizationQuotasUsage(ctx, organizationId).Execute()

Return the organization's current quota usage alongside the plan's limits. Response is cached in Redis for 60s. Accept text/plain (or ?format=prometheus) to receive the response rendered as Prometheus text exposition format (\"koyeb_quota_<x>\" for usage, \"koyeb_quota_<x>_limit\" for the plan limit) suitable for scraping into an external Prometheus.

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
    organizationId := "organizationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotasApi.GetOrganizationQuotasUsage(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotasApi.GetOrganizationQuotasUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationQuotasUsage`: GetOrganizationQuotasUsageReply
    fmt.Fprintf(os.Stdout, "Response from `QuotasApi.GetOrganizationQuotasUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationQuotasUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationQuotasUsageReply**](GetOrganizationQuotasUsageReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewOrganizationCapacity

> ReviewOrganizationCapacityReply ReviewOrganizationCapacity(ctx).Body(body).Execute()

DEPRECATED: Review Organization Capacity

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

