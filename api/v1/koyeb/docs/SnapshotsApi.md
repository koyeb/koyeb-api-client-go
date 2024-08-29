# \SnapshotsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotsApi.md#CreateSnapshot) | **Post** /v1/snapshots | Create a Snapshot
[**DeleteSnapshot**](SnapshotsApi.md#DeleteSnapshot) | **Delete** /v1/snapshots/{id} | Delete a Snapshot
[**GetSnapshot**](SnapshotsApi.md#GetSnapshot) | **Get** /v1/snapshots/{id} | Get a Snapshot
[**ListSnapshots**](SnapshotsApi.md#ListSnapshots) | **Get** /v1/snapshots | List all Snapshots
[**UpdateSnapshot**](SnapshotsApi.md#UpdateSnapshot) | **Post** /v1/snapshots/{id} | Update a Snapshot



## CreateSnapshot

> CreateSnapshotReply CreateSnapshot(ctx).Body(body).Execute()

Create a Snapshot

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
    body := *openapiclient.NewCreateSnapshotRequest() // CreateSnapshotRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.CreateSnapshot(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshot`: CreateSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**CreateSnapshotReply**](CreateSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshotReply DeleteSnapshot(ctx, id).Execute()

Delete a Snapshot

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
    resp, r, err := apiClient.SnapshotsApi.DeleteSnapshot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshot`: DeleteSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSnapshotReply**](DeleteSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> GetSnapshotReply GetSnapshot(ctx, id).Execute()

Get a Snapshot

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
    resp, r, err := apiClient.SnapshotsApi.GetSnapshot(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshot`: GetSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSnapshotReply**](GetSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshots

> ListSnapshotsReply ListSnapshots(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Statuses(statuses).Region(region).Execute()

List all Snapshots

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
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    organizationId := "organizationId_example" // string | (Optional) Filter by organization_id (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter by status   - SNAPSHOT_STATUS_INVALID: zero value, invalid  - SNAPSHOT_STATUS_CREATING: the snapshot is being created  - SNAPSHOT_STATUS_AVAILABLE: the snapshot is complete and available  - SNAPSHOT_STATUS_MIGRATING: the snapshot is being migrated  - SNAPSHOT_STATUS_DELETING: the snapshot is being deleted  - SNAPSHOT_STATUS_DELETED: the snapshot is deleted (optional)
    region := "region_example" // string | (Optional) A filter for the region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.ListSnapshots(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Statuses(statuses).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.ListSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSnapshots`: ListSnapshotsReply
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **organizationId** | **string** | (Optional) Filter by organization_id | 
 **statuses** | **[]string** | (Optional) Filter by status   - SNAPSHOT_STATUS_INVALID: zero value, invalid  - SNAPSHOT_STATUS_CREATING: the snapshot is being created  - SNAPSHOT_STATUS_AVAILABLE: the snapshot is complete and available  - SNAPSHOT_STATUS_MIGRATING: the snapshot is being migrated  - SNAPSHOT_STATUS_DELETING: the snapshot is being deleted  - SNAPSHOT_STATUS_DELETED: the snapshot is deleted | 
 **region** | **string** | (Optional) A filter for the region | 

### Return type

[**ListSnapshotsReply**](ListSnapshotsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> UpdateSnapshotReply UpdateSnapshot(ctx, id).Body(body).Execute()

Update a Snapshot

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
    id := "id_example" // string | The id of the snapshot
    body := *openapiclient.NewUpdateSnapshotRequest() // UpdateSnapshotRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.UpdateSnapshot(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.UpdateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnapshot`: UpdateSnapshotReply
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the snapshot | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateSnapshotRequest**](UpdateSnapshotRequest.md) |  | 

### Return type

[**UpdateSnapshotReply**](UpdateSnapshotReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

