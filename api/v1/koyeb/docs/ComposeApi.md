# \ComposeApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Compose**](ComposeApi.md#Compose) | **Post** /v1/compose | Create resources from compose.



## Compose

> ComposeReply Compose(ctx).Compose(compose).Execute()

Create resources from compose.

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
    compose := *openapiclient.NewCreateCompose() // CreateCompose | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComposeApi.Compose(context.Background()).Compose(compose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComposeApi.Compose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Compose`: ComposeReply
    fmt.Fprintf(os.Stdout, "Response from `ComposeApi.Compose`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compose** | [**CreateCompose**](CreateCompose.md) |  | 

### Return type

[**ComposeReply**](ComposeReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

