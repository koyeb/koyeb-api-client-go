# \ServiceViewsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryServicesView**](ServiceViewsApi.md#QueryServicesView) | **Get** /v1/views/services | Query the Services view



## QueryServicesView

> QueryServicesViewReply QueryServicesView(ctx).Search(search).FiltersTypes(filtersTypes).FiltersStatuses(filtersStatuses).FiltersRegions(filtersRegions).FiltersAppIds(filtersAppIds).FiltersActiveDeploymentStates(filtersActiveDeploymentStates).FiltersLatestDeploymentStates(filtersLatestDeploymentStates).FiltersNeedsAttention(filtersNeedsAttention).SortField(sortField).SortDirection(sortDirection).GroupBy(groupBy).Limit(limit).Offset(offset).Facets(facets).Execute()

Query the Services view



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
    search := "search_example" // string | (Optional) Case-insensitive search on the Service name and the App name (optional)
    filtersTypes := []string{"FiltersTypes_example"} // []string | (Optional) Filter on Service types (optional)
    filtersStatuses := []string{"FiltersStatuses_example"} // []string | (Optional) Filter on Service statuses (optional)
    filtersRegions := []string{"Inner_example"} // []string | (Optional) Filter on regions of the latest Deployment (optional)
    filtersAppIds := []string{"Inner_example"} // []string | (Optional) Filter on App ids (optional)
    filtersActiveDeploymentStates := []string{"FiltersActiveDeploymentStates_example"} // []string | (Optional) Filter on active Deployment states   - SLEEPING: The active Deployment is scaled to zero and wakes up on traffic (optional)
    filtersLatestDeploymentStates := []string{"FiltersLatestDeploymentStates_example"} // []string | (Optional) Filter on latest Deployment states (optional)
    filtersNeedsAttention := true // bool | (Optional) Filter on Services that need attention (optional)
    sortField := "sortField_example" // string |  (optional) (default to "CREATED_AT")
    sortDirection := "sortDirection_example" // string |  (optional) (default to "ASC")
    groupBy := "groupBy_example" // string | (Optional) Group rows by App. Groups are ordered by App name and keep a stable order across pages. Pagination always counts Services (optional) (default to "NO_GROUP")
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    facets := []string{"Facets_example"} // []string | (Optional) Facets to compute on the full result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceViewsApi.QueryServicesView(context.Background()).Search(search).FiltersTypes(filtersTypes).FiltersStatuses(filtersStatuses).FiltersRegions(filtersRegions).FiltersAppIds(filtersAppIds).FiltersActiveDeploymentStates(filtersActiveDeploymentStates).FiltersLatestDeploymentStates(filtersLatestDeploymentStates).FiltersNeedsAttention(filtersNeedsAttention).SortField(sortField).SortDirection(sortDirection).GroupBy(groupBy).Limit(limit).Offset(offset).Facets(facets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceViewsApi.QueryServicesView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryServicesView`: QueryServicesViewReply
    fmt.Fprintf(os.Stdout, "Response from `ServiceViewsApi.QueryServicesView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryServicesViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | (Optional) Case-insensitive search on the Service name and the App name | 
 **filtersTypes** | **[]string** | (Optional) Filter on Service types | 
 **filtersStatuses** | **[]string** | (Optional) Filter on Service statuses | 
 **filtersRegions** | **[]string** | (Optional) Filter on regions of the latest Deployment | 
 **filtersAppIds** | **[]string** | (Optional) Filter on App ids | 
 **filtersActiveDeploymentStates** | **[]string** | (Optional) Filter on active Deployment states   - SLEEPING: The active Deployment is scaled to zero and wakes up on traffic | 
 **filtersLatestDeploymentStates** | **[]string** | (Optional) Filter on latest Deployment states | 
 **filtersNeedsAttention** | **bool** | (Optional) Filter on Services that need attention | 
 **sortField** | **string** |  | [default to &quot;CREATED_AT&quot;]
 **sortDirection** | **string** |  | [default to &quot;ASC&quot;]
 **groupBy** | **string** | (Optional) Group rows by App. Groups are ordered by App name and keep a stable order across pages. Pagination always counts Services | [default to &quot;NO_GROUP&quot;]
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **facets** | **[]string** | (Optional) Facets to compute on the full result | 

### Return type

[**QueryServicesViewReply**](QueryServicesViewReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

