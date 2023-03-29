# \RegionalDeploymentsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegionalDeployment**](RegionalDeploymentsApi.md#GetRegionalDeployment) | **Get** /v1/regional_deployments/{id} | Experimental: Get regional deployment Use at your own risk
[**ListRegionalDeployments**](RegionalDeploymentsApi.md#ListRegionalDeployments) | **Get** /v1/regional_deployments | Experimental: List regional deployments Use at your own risk



## GetRegionalDeployment

> GetRegionalDeploymentReply GetRegionalDeployment(ctx, id).Execute()

Experimental: Get regional deployment Use at your own risk

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
    id := "id_example" // string | The id of the regional deployment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionalDeploymentsApi.GetRegionalDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionalDeploymentsApi.GetRegionalDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionalDeployment`: GetRegionalDeploymentReply
    fmt.Fprintf(os.Stdout, "Response from `RegionalDeploymentsApi.GetRegionalDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the regional deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionalDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegionalDeploymentReply**](GetRegionalDeploymentReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegionalDeployments

> ListRegionalDeploymentsReply ListRegionalDeployments(ctx).DeploymentId(deploymentId).Limit(limit).Offset(offset).Execute()

Experimental: List regional deployments Use at your own risk

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
    deploymentId := "deploymentId_example" // string | (Optional) Filter on deployment id (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionalDeploymentsApi.ListRegionalDeployments(context.Background()).DeploymentId(deploymentId).Limit(limit).Offset(offset).Execute()
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
 **deploymentId** | **string** | (Optional) Filter on deployment id | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 

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

