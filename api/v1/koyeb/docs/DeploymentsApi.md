# \DeploymentsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /v1/deployments/{id} | Get Deployment
[**ListDeployments**](DeploymentsApi.md#ListDeployments) | **Get** /v1/deployments | List Deployments



## GetDeployment

> GetDeploymentReply GetDeployment(ctx, id).Execute()

Get Deployment

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
    id := "id_example" // string | The id of the deployment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployment`: GetDeploymentReply
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeploymentReply**](GetDeploymentReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployments

> ListDeploymentsReply ListDeployments(ctx).Id(id).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()

List Deployments

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
    id := "id_example" // string |  (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Statuses to filter on. (optional)
    deploymentGroups := []string{"Inner_example"} // []string | (Optional) Only fetch deployments in this deployment group. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListDeployments(context.Background()).Id(id).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeployments`: ListDeploymentsReply
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **statuses** | **[]string** | (Optional) Statuses to filter on. | 
 **deploymentGroups** | **[]string** | (Optional) Only fetch deployments in this deployment group. | 

### Return type

[**ListDeploymentsReply**](ListDeploymentsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
