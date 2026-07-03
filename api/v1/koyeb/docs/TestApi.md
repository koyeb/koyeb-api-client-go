# \TestApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestAnon**](TestApi.md#TestAnon) | **Get** /v1/test/anon | Test anonymous call
[**TestAuth**](TestApi.md#TestAuth) | **Post** /v1/test/auth | Test authenticated call



## TestAnon

> map[string]interface{} TestAnon(ctx).Execute()

Test anonymous call

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestApi.TestAnon(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestApi.TestAnon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestAnon`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TestApi.TestAnon`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestAnonRequest struct via the builder pattern


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


## TestAuth

> KtestTestAuthReply TestAuth(ctx).Body(body).Execute()

Test authenticated call

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
    body := *openapiclient.NewKtestTestAuthRequest() // KtestTestAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestApi.TestAuth(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestApi.TestAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestAuth`: KtestTestAuthReply
    fmt.Fprintf(os.Stdout, "Response from `TestApi.TestAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KtestTestAuthRequest**](KtestTestAuthRequest.md) |  | 

### Return type

[**KtestTestAuthReply**](KtestTestAuthReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

