# \UsagesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationUsage**](UsagesApi.md#GetOrganizationUsage) | **Get** /v1/usages | Get organization usage
[**GetOrganizationUsageDetails**](UsagesApi.md#GetOrganizationUsageDetails) | **Get** /v1/usages/details | Get organization usage details



## GetOrganizationUsage

> GetOrganizationUsageReply GetOrganizationUsage(ctx).StartingTime(startingTime).EndingTime(endingTime).Execute()

Get organization usage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/koyeb/koyeb-api-client-go/koyeb"
)

func main() {
    startingTime := time.Now() // time.Time | The starting time of the period to get data from (optional)
    endingTime := time.Now() // time.Time | The ending time of the period to get date from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsagesApi.GetOrganizationUsage(context.Background()).StartingTime(startingTime).EndingTime(endingTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesApi.GetOrganizationUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationUsage`: GetOrganizationUsageReply
    fmt.Fprintf(os.Stdout, "Response from `UsagesApi.GetOrganizationUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startingTime** | **time.Time** | The starting time of the period to get data from | 
 **endingTime** | **time.Time** | The ending time of the period to get date from | 

### Return type

[**GetOrganizationUsageReply**](GetOrganizationUsageReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationUsageDetails

> GetOrganizationUsageDetailsReply GetOrganizationUsageDetails(ctx).StartingTime(startingTime).EndingTime(endingTime).Limit(limit).Offset(offset).Order(order).Execute()

Get organization usage details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/koyeb/koyeb-api-client-go/koyeb"
)

func main() {
    startingTime := time.Now() // time.Time | The starting time of the period to get data from (optional)
    endingTime := time.Now() // time.Time | The ending time of the period to get date from (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsagesApi.GetOrganizationUsageDetails(context.Background()).StartingTime(startingTime).EndingTime(endingTime).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsagesApi.GetOrganizationUsageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationUsageDetails`: GetOrganizationUsageDetailsReply
    fmt.Fprintf(os.Stdout, "Response from `UsagesApi.GetOrganizationUsageDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUsageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startingTime** | **time.Time** | The starting time of the period to get data from | 
 **endingTime** | **time.Time** | The ending time of the period to get date from | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**GetOrganizationUsageDetailsReply**](GetOrganizationUsageDetailsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

