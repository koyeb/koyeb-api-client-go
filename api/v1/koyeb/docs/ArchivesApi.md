# \ArchivesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchive**](ArchivesApi.md#CreateArchive) | **Post** /v1/archives | Create Archive
[**DeleteArchive**](ArchivesApi.md#DeleteArchive) | **Delete** /v1/archives/{id} | Delete Archive
[**GetArchive**](ArchivesApi.md#GetArchive) | **Get** /v1/archives/{id} | Get Archive
[**ListArchives**](ArchivesApi.md#ListArchives) | **Get** /v1/archives | List Archives



## CreateArchive

> CreateArchiveReply CreateArchive(ctx).Archive(archive).Execute()

Create Archive



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
    archive := *openapiclient.NewCreateArchive() // CreateArchive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchivesApi.CreateArchive(context.Background()).Archive(archive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.CreateArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArchive`: CreateArchiveReply
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.CreateArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archive** | [**CreateArchive**](CreateArchive.md) |  | 

### Return type

[**CreateArchiveReply**](CreateArchiveReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArchive

> map[string]interface{} DeleteArchive(ctx, id).Execute()

Delete Archive



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
    id := "id_example" // string | The id of the archive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchivesApi.DeleteArchive(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.DeleteArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteArchive`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.DeleteArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveRequest struct via the builder pattern


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


## GetArchive

> GetArchiveReply GetArchive(ctx, id).Execute()

Get Archive



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
    id := "id_example" // string | The id of the archive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchivesApi.GetArchive(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchive`: GetArchiveReply
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.GetArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetArchiveReply**](GetArchiveReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchives

> ListArchivesReply ListArchives(ctx).Limit(limit).Offset(offset).Ids(ids).Statuses(statuses).Execute()

List Archives



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
    offset := "offset_example" // string | (Optional) The offset in the list of items to return (optional)
    ids := []string{"Inner_example"} // []string | (Optional) Filter on archive ids (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter by archive status   - ARCHIVE_STATUS_INVALID: zero value, invalid  - ARCHIVE_STATUS_HEALTHY: the archive can be deployed and downloaded  - ARCHIVE_STATUS_DELETING: the archive is waiting for its file to be removed  - ARCHIVE_STATUS_DELETED: the archive and its file are deleted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchivesApi.ListArchives(context.Background()).Limit(limit).Offset(offset).Ids(ids).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.ListArchives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchives`: ListArchivesReply
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.ListArchives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of items to return | 
 **ids** | **[]string** | (Optional) Filter on archive ids | 
 **statuses** | **[]string** | (Optional) Filter by archive status   - ARCHIVE_STATUS_INVALID: zero value, invalid  - ARCHIVE_STATUS_HEALTHY: the archive can be deployed and downloaded  - ARCHIVE_STATUS_DELETING: the archive is waiting for its file to be removed  - ARCHIVE_STATUS_DELETED: the archive and its file are deleted | 

### Return type

[**ListArchivesReply**](ListArchivesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

