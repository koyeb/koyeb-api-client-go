# \InstancesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecCommand**](InstancesApi.md#ExecCommand) | **Get** /v1/streams/instances/exec | Exec Command
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /v1/instances/{id} | Get Instance
[**ListInstances**](InstancesApi.md#ListInstances) | **Get** /v1/instances | List Instances



## ExecCommand

> StreamResultOfExecCommandReply ExecCommand(ctx).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).Execute()

Exec Command

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
    bodyCommand := []string{"Inner_example"} // []string |  (optional)
    bodyTtySizeHeight := int32(56) // int32 |  (optional)
    bodyTtySizeWidth := int32(56) // int32 |  (optional)
    bodyStdinData := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ExecCommand(context.Background()).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExecCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecCommand`: StreamResultOfExecCommandReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExecCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **bodyCommand** | **[]string** |  | 
 **bodyTtySizeHeight** | **int32** |  | 
 **bodyTtySizeWidth** | **int32** |  | 
 **bodyStdinData** | **string** |  | 

### Return type

[**StreamResultOfExecCommandReply**](StreamResultOfExecCommandReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> GetInstanceReply GetInstance(ctx, id).Execute()

Get Instance

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: GetInstanceReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceReply**](GetInstanceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> ListInstancesReply ListInstances(ctx).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).AllocationId(allocationId).Statuses(statuses).Limit(limit).Offset(offset).Order(order).Execute()

List Instances

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
    appId := "appId_example" // string | (Optional) Filter on application id. (optional)
    serviceId := "serviceId_example" // string | (Optional) Filter on service id. (optional)
    deploymentId := "deploymentId_example" // string | (Optional) Filter on deployment id. (optional)
    allocationId := "allocationId_example" // string | (Optional) Filter on allocation id. (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter on instance statuses. (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListInstances(context.Background()).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).AllocationId(allocationId).Statuses(statuses).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstances`: ListInstancesReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | (Optional) Filter on application id. | 
 **serviceId** | **string** | (Optional) Filter on service id. | 
 **deploymentId** | **string** | (Optional) Filter on deployment id. | 
 **allocationId** | **string** | (Optional) Filter on allocation id. | 
 **statuses** | **[]string** | (Optional) Filter on instance statuses. | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order. | 

### Return type

[**ListInstancesReply**](ListInstancesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

