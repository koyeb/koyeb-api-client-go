# \DockerHelperApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyDockerImage**](DockerHelperApi.md#VerifyDockerImage) | **Get** /v1/docker-helper/verify | Verify if a docker image is reachable



## VerifyDockerImage

> VerifyDockerImageReply VerifyDockerImage(ctx).Image(image).SecretId(secretId).Execute()

Verify if a docker image is reachable

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
    image := "image_example" // string | The full image uri to be verified (optional)
    secretId := "secretId_example" // string | (Optional) the id of the secret to use to authenticate to the registry (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DockerHelperApi.VerifyDockerImage(context.Background()).Image(image).SecretId(secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DockerHelperApi.VerifyDockerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyDockerImage`: VerifyDockerImageReply
    fmt.Fprintf(os.Stdout, "Response from `DockerHelperApi.VerifyDockerImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyDockerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string** | The full image uri to be verified | 
 **secretId** | **string** | (Optional) the id of the secret to use to authenticate to the registry | 

### Return type

[**VerifyDockerImageReply**](VerifyDockerImageReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

