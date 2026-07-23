# \InstanceSnapshotsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstanceSnapshot**](InstanceSnapshotsApi.md#CreateInstanceSnapshot) | **Post** /v1/instance_snapshots | 
[**DeleteInstanceSnapshot**](InstanceSnapshotsApi.md#DeleteInstanceSnapshot) | **Delete** /v1/instance_snapshots/{id} | 
[**GetInstanceSnapshot**](InstanceSnapshotsApi.md#GetInstanceSnapshot) | **Get** /v1/instance_snapshots/{id} | 
[**ListInstanceSnapshotEvents**](InstanceSnapshotsApi.md#ListInstanceSnapshotEvents) | **Get** /v1/instance_snapshot_events | 
[**ListInstanceSnapshots**](InstanceSnapshotsApi.md#ListInstanceSnapshots) | **Get** /v1/instance_snapshots | 



## CreateInstanceSnapshot

> CreateInstanceSnapshotReply CreateInstanceSnapshot(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateInstanceSnapshotRequest() // CreateInstanceSnapshotRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceSnapshotsApi.CreateInstanceSnapshot(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotsApi.CreateInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstanceSnapshot`: CreateInstanceSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotsApi.CreateInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateInstanceSnapshotRequest**](CreateInstanceSnapshotRequest.md) |  | 

### Return type

[**CreateInstanceSnapshotReply**](CreateInstanceSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceSnapshot

> DeleteInstanceSnapshotReply DeleteInstanceSnapshot(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceSnapshotsApi.DeleteInstanceSnapshot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotsApi.DeleteInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstanceSnapshot`: DeleteInstanceSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotsApi.DeleteInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInstanceSnapshotReply**](DeleteInstanceSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceSnapshot

> GetInstanceSnapshotReply GetInstanceSnapshot(ctx, id).Execute()



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceSnapshotsApi.GetInstanceSnapshot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotsApi.GetInstanceSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceSnapshot`: GetInstanceSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotsApi.GetInstanceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceSnapshotReply**](GetInstanceSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceSnapshotEvents

> ListInstanceSnapshotEventsReply ListInstanceSnapshotEvents(ctx).InstanceSnapshotId(instanceSnapshotId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()



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
    instanceSnapshotId := "instanceSnapshotId_example" // string | (Optional) Filter on instance snapshot id (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on instance snapshot event types (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceSnapshotsApi.ListInstanceSnapshotEvents(context.Background()).InstanceSnapshotId(instanceSnapshotId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotsApi.ListInstanceSnapshotEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceSnapshotEvents`: ListInstanceSnapshotEventsReply
    fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotsApi.ListInstanceSnapshotEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceSnapshotEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceSnapshotId** | **string** | (Optional) Filter on instance snapshot id | 
 **types** | **[]string** | (Optional) Filter on instance snapshot event types | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**ListInstanceSnapshotEventsReply**](ListInstanceSnapshotEventsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceSnapshots

> ListInstanceSnapshotsReply ListInstanceSnapshots(ctx).Limit(limit).Offset(offset).Name(name).Statuses(statuses).Type_(type_).Ids(ids).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    statuses := []string{"Statuses_example"} // []string |  (optional)
    type_ := "type__example" // string |  (optional) (default to "INSTANCE_SNAPSHOT_TYPE_INVALID")
    ids := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceSnapshotsApi.ListInstanceSnapshots(context.Background()).Limit(limit).Offset(offset).Name(name).Statuses(statuses).Type_(type_).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceSnapshotsApi.ListInstanceSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstanceSnapshots`: ListInstanceSnapshotsReply
    fmt.Fprintf(os.Stdout, "Response from `InstanceSnapshotsApi.ListInstanceSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | 
 **offset** | **string** |  | 
 **name** | **string** |  | 
 **statuses** | **[]string** |  | 
 **type_** | **string** |  | [default to &quot;INSTANCE_SNAPSHOT_TYPE_INVALID&quot;]
 **ids** | **[]string** |  | 

### Return type

[**ListInstanceSnapshotsReply**](ListInstanceSnapshotsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

