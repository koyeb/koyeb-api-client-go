# \ArchivesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchive**](ArchivesApi.md#CreateArchive) | **Post** /v1/archives | Create Archive
[**GetArchive**](ArchivesApi.md#GetArchive) | **Get** /v1/archives/{id} | Get Archive



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

