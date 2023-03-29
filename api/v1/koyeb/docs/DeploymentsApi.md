# \DeploymentsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeployment**](DeploymentsApi.md#CancelDeployment) | **Post** /v1/deployments/{id}/cancel | Cancel Deployment Deployment cancellation is allowed for the following status:  - pending  - provisioning  - scheduled
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /v1/deployments/{id} | Get Deployment
[**ListDeployments**](DeploymentsApi.md#ListDeployments) | **Get** /v1/deployments | List Deployments



## CancelDeployment

> map[string]interface{} CancelDeployment(ctx, id).Execute()

Cancel Deployment Deployment cancellation is allowed for the following status:  - pending  - provisioning  - scheduled

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
    id := "id_example" // string | The id of the deployment to cancel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.CancelDeployment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CancelDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelDeployment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CancelDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the deployment to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/koyeb/koyeb-api-client-go/koyeb"
)

func main() {
    id := "id_example" // string | The id of the deployment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.GetDeployment(context.Background(), id).Execute()
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

> ListDeploymentsReply ListDeployments(ctx).AppId(appId).ServiceId(serviceId).Limit(limit).Offset(offset).Statuses(statuses).Execute()

List Deployments

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
    appId := "appId_example" // string | (Optional) Filter on application id (optional)
    serviceId := "serviceId_example" // string | (Optional) Filter on service id (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter on statuses (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ListDeployments(context.Background()).AppId(appId).ServiceId(serviceId).Limit(limit).Offset(offset).Statuses(statuses).Execute()
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
 **appId** | **string** | (Optional) Filter on application id | 
 **serviceId** | **string** | (Optional) Filter on service id | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **statuses** | **[]string** | (Optional) Filter on statuses | 

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

