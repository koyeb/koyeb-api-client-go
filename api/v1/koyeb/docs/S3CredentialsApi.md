# \S3CredentialsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateS3Credential**](S3CredentialsApi.md#CreateS3Credential) | **Post** /v1/s3_credentials | 
[**DeleteS3Credential**](S3CredentialsApi.md#DeleteS3Credential) | **Delete** /v1/s3_credentials/{id} | 
[**GetS3Credential**](S3CredentialsApi.md#GetS3Credential) | **Get** /v1/s3_credentials/{id} | 
[**ListS3Credentials**](S3CredentialsApi.md#ListS3Credentials) | **Get** /v1/s3_credentials | 
[**UpdateS3Credential**](S3CredentialsApi.md#UpdateS3Credential) | **Put** /v1/s3_credentials/{id} | 
[**UpdateS3Credential2**](S3CredentialsApi.md#UpdateS3Credential2) | **Patch** /v1/s3_credentials/{id} | 



## CreateS3Credential

> S3CredentialReply CreateS3Credential(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateS3Credential() // CreateS3Credential | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.CreateS3Credential(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.CreateS3Credential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateS3Credential`: S3CredentialReply
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.CreateS3Credential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateS3CredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateS3Credential**](CreateS3Credential.md) |  | 

### Return type

[**S3CredentialReply**](S3CredentialReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteS3Credential

> interface{} DeleteS3Credential(ctx, id).Execute()



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.DeleteS3Credential(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.DeleteS3Credential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteS3Credential`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.DeleteS3Credential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3CredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3Credential

> S3CredentialReply GetS3Credential(ctx, id).Execute()



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.GetS3Credential(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.GetS3Credential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetS3Credential`: S3CredentialReply
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.GetS3Credential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3CredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3CredentialReply**](S3CredentialReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListS3Credentials

> ListS3CredentialsReply ListS3Credentials(ctx).Name(name).Status(status).Limit(limit).Offset(offset).Execute()



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
    name := "name_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.ListS3Credentials(context.Background()).Name(name).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.ListS3Credentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListS3Credentials`: ListS3CredentialsReply
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.ListS3Credentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListS3CredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **status** | **string** |  | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListS3CredentialsReply**](ListS3CredentialsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateS3Credential

> S3CredentialReply UpdateS3Credential(ctx, id).Body(body).UpdateMask(updateMask).Execute()



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
    body := *openapiclient.NewS3Credential() // S3Credential | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.UpdateS3Credential(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.UpdateS3Credential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateS3Credential`: S3CredentialReply
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.UpdateS3Credential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateS3CredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3Credential**](S3Credential.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**S3CredentialReply**](S3CredentialReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateS3Credential2

> S3CredentialReply UpdateS3Credential2(ctx, id).Body(body).UpdateMask(updateMask).Execute()



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
    body := *openapiclient.NewS3Credential() // S3Credential | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.S3CredentialsApi.UpdateS3Credential2(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3CredentialsApi.UpdateS3Credential2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateS3Credential2`: S3CredentialReply
    fmt.Fprintf(os.Stdout, "Response from `S3CredentialsApi.UpdateS3Credential2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateS3Credential2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3Credential**](S3Credential.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**S3CredentialReply**](S3CredentialReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

