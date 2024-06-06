# \IntercomApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntercomProfile**](IntercomApi.md#GetIntercomProfile) | **Get** /v1/intercom/profile | Get intercom profile



## GetIntercomProfile

> GetIntercomProfileReply GetIntercomProfile(ctx).Execute()

Get intercom profile

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
    resp, r, err := apiClient.IntercomApi.GetIntercomProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntercomApi.GetIntercomProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntercomProfile`: GetIntercomProfileReply
    fmt.Fprintf(os.Stdout, "Response from `IntercomApi.GetIntercomProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntercomProfileRequest struct via the builder pattern


### Return type

[**GetIntercomProfileReply**](GetIntercomProfileReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

