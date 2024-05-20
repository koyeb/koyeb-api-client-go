# \InstancesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecCommand**](InstancesApi.md#ExecCommand) | **Get** /v1/streams/instances/exec | Exec Command
[**GetInstance**](InstancesApi.md#GetInstance) | **Get** /v1/instances/{id} | Get Instance
[**ListInstanceEvents**](InstancesApi.md#ListInstanceEvents) | **Get** /v1/instance_events | List Instance events
[**ListInstances**](InstancesApi.md#ListInstances) | **Get** /v1/instances | List Instances



## ExecCommand

> StreamResultOfExecCommandReply ExecCommand(ctx).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).BodyStdinClose(bodyStdinClose).BodyDisableTty(bodyDisableTty).IdType(idType).Execute()

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
    id := "id_example" // string | ID of the resource to exec on. (optional)
    bodyCommand := []string{"Inner_example"} // []string | Command to exec. Mandatory in the first frame sent (optional)
    bodyTtySizeHeight := int32(56) // int32 |  (optional)
    bodyTtySizeWidth := int32(56) // int32 |  (optional)
    bodyStdinData := string(BYTE_ARRAY_DATA_HERE) // string | Data is base64 encoded (optional)
    bodyStdinClose := true // bool | Indicate last data frame (optional)
    bodyDisableTty := true // bool | Disable TTY. It's enough to specify it in the first frame (optional)
    idType := "idType_example" // string | When specified, it is used to determine if the kind of resource the id refers to. If missing, defaults to the instance id. (optional) (default to "INVALID")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.ExecCommand(context.Background()).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).BodyStdinClose(bodyStdinClose).BodyDisableTty(bodyDisableTty).IdType(idType).Execute()
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
 **id** | **string** | ID of the resource to exec on. | 
 **bodyCommand** | **[]string** | Command to exec. Mandatory in the first frame sent | 
 **bodyTtySizeHeight** | **int32** |  | 
 **bodyTtySizeWidth** | **int32** |  | 
 **bodyStdinData** | **string** | Data is base64 encoded | 
 **bodyStdinClose** | **bool** | Indicate last data frame | 
 **bodyDisableTty** | **bool** | Disable TTY. It&#39;s enough to specify it in the first frame | 
 **idType** | **string** | When specified, it is used to determine if the kind of resource the id refers to. If missing, defaults to the instance id. | [default to &quot;INVALID&quot;]

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
    id := "id_example" // string | The id of the instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetInstance(context.Background(), id).Execute()
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
**id** | **string** | The id of the instance | 

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


## ListInstanceEvents

> ListInstanceEventsReply ListInstanceEvents(ctx).InstanceIds(instanceIds).Types(types).Limit(limit).Offset(offset).Order(order).Execute()

List Instance events

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
    instanceIds := []string{"Inner_example"} // []string | (Optional) Filter on list of instance id (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on instance event types (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.ListInstanceEvents(context.Background()).InstanceIds(instanceIds).Types(types).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListInstanceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceEvents`: ListInstanceEventsReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListInstanceEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceIds** | **[]string** | (Optional) Filter on list of instance id | 
 **types** | **[]string** | (Optional) Filter on instance event types | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**ListInstanceEventsReply**](ListInstanceEventsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> ListInstancesReply ListInstances(ctx).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).AllocationId(allocationId).ReplicaIndex(replicaIndex).Statuses(statuses).Limit(limit).Offset(offset).Order(order).StartingTime(startingTime).EndingTime(endingTime).Execute()

List Instances

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | (Optional) Filter on application id (optional)
    serviceId := "serviceId_example" // string | (Optional) Filter on service id (optional)
    deploymentId := "deploymentId_example" // string | (Optional) Filter on deployment id (optional)
    regionalDeploymentId := "regionalDeploymentId_example" // string | (Optional) Filter on regional deployment id (optional)
    allocationId := "allocationId_example" // string | (Optional) Filter on allocation id (optional)
    replicaIndex := "replicaIndex_example" // string | (Optional) Filter on replica index (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter on instance statuses (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)
    startingTime := time.Now() // time.Time | (Optional) The starting time of the period of running instance (optional)
    endingTime := time.Now() // time.Time | (Optional) The ending time of the period of running instance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.ListInstances(context.Background()).AppId(appId).ServiceId(serviceId).DeploymentId(deploymentId).RegionalDeploymentId(regionalDeploymentId).AllocationId(allocationId).ReplicaIndex(replicaIndex).Statuses(statuses).Limit(limit).Offset(offset).Order(order).StartingTime(startingTime).EndingTime(endingTime).Execute()
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
 **appId** | **string** | (Optional) Filter on application id | 
 **serviceId** | **string** | (Optional) Filter on service id | 
 **deploymentId** | **string** | (Optional) Filter on deployment id | 
 **regionalDeploymentId** | **string** | (Optional) Filter on regional deployment id | 
 **allocationId** | **string** | (Optional) Filter on allocation id | 
 **replicaIndex** | **string** | (Optional) Filter on replica index | 
 **statuses** | **[]string** | (Optional) Filter on instance statuses | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 
 **startingTime** | **time.Time** | (Optional) The starting time of the period of running instance | 
 **endingTime** | **time.Time** | (Optional) The ending time of the period of running instance | 

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

