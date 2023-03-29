# \InviteApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvite**](InviteApi.md#CreateInvite) | **Post** /v1/account/invite | 



## CreateInvite

> map[string]interface{} CreateInvite(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/koyeb/koyeb-api-client-go/koyeb"
)

func main() {
    body := *openapiclient.NewInviteUserRequest() // InviteUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InviteApi.CreateInvite(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InviteApi.CreateInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvite`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InviteApi.CreateInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InviteUserRequest**](InviteUserRequest.md) |  | 

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

