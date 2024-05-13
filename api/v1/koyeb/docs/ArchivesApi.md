# \ArchivesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchive**](ArchivesApi.md#CreateArchive) | **Post** /v1/archives | Create a signed URL to upload an archive.



## CreateArchive

> CreateArchiveReply CreateArchive(ctx).Archive(archive).Execute()

Create a signed URL to upload an archive.

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

