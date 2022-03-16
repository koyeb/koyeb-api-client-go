# \RegionalDeploymentsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRegionalDeployments**](RegionalDeploymentsApi.md#ListRegionalDeployments) | **Get** /v1/regional_deployments | Temporary: List regional deployments



## ListRegionalDeployments

> ListRegionalDeploymentsReply ListRegionalDeployments(ctx).DeploymentId(deploymentId).Limit(limit).Offset(offset).Execute()

Temporary: List regional deployments

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
    deploymentId := "deploymentId_example" // string | (Optional) Filter on deployment id. (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionalDeploymentsApi.ListRegionalDeployments(context.Background()).DeploymentId(deploymentId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionalDeploymentsApi.ListRegionalDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegionalDeployments`: ListRegionalDeploymentsReply
    fmt.Fprintf(os.Stdout, "Response from `RegionalDeploymentsApi.ListRegionalDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegionalDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | (Optional) Filter on deployment id. | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 

### Return type

[**ListRegionalDeploymentsReply**](ListRegionalDeploymentsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

