# \DeploymentsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeployment**](DeploymentsApi.md#CancelDeployment) | **Post** /v1/deployments/{id}/cancel | Cancel Deployment
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /v1/deployments/{id} | Get Deployment
[**GetDeploymentScaling**](DeploymentsApi.md#GetDeploymentScaling) | **Get** /v1/deployment/{id}/scaling | Get Deployment Scaling
[**ListDeploymentEvents**](DeploymentsApi.md#ListDeploymentEvents) | **Get** /v1/deployment_events | List Deployment events
[**ListDeployments**](DeploymentsApi.md#ListDeployments) | **Get** /v1/deployments | List Deployments



## CancelDeployment

> map[string]interface{} CancelDeployment(ctx, id).Execute()

Cancel Deployment



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
    openapiclient "./openapi"
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


## GetDeploymentScaling

> GetDeploymentScalingReply GetDeploymentScaling(ctx, id).Region(region).ReplicaIndex(replicaIndex).Execute()

Get Deployment Scaling

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
    id := "id_example" // string | 
    region := "region_example" // string | (Optional) Filter on region (optional)
    replicaIndex := int64(789) // int64 | (Optional) Filter on replica_index (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.GetDeploymentScaling(context.Background(), id).Region(region).ReplicaIndex(replicaIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentScaling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentScaling`: GetDeploymentScalingReply
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | (Optional) Filter on region | 
 **replicaIndex** | **int64** | (Optional) Filter on replica_index | 

### Return type

[**GetDeploymentScalingReply**](GetDeploymentScalingReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeploymentEvents

> ListDeploymentEventsReply ListDeploymentEvents(ctx).DeploymentId(deploymentId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()

List Deployment events

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
    deploymentId := "deploymentId_example" // string | (Optional) Filter on deployment id (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on deployment event types (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.ListDeploymentEvents(context.Background()).DeploymentId(deploymentId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDeploymentEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeploymentEvents`: ListDeploymentEventsReply
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDeploymentEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | (Optional) Filter on deployment id | 
 **types** | **[]string** | (Optional) Filter on deployment event types | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**ListDeploymentEventsReply**](ListDeploymentEventsReply.md)

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
    openapiclient "./openapi"
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

