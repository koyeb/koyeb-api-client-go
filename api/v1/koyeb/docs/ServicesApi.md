# \ServicesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/services | Create Service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/services/{id} | Delete Service Service deletion is allowed for all status.
[**GetService**](ServicesApi.md#GetService) | **Get** /v1/services/{id} | Get Service
[**ListServices**](ServicesApi.md#ListServices) | **Get** /v1/services | List Service
[**PauseService**](ServicesApi.md#PauseService) | **Post** /v1/services/{id}/pause | Pause Service Service pause action is allowed for the following status:  - starting  - healthy  - degraded  - unhealthy  - resuming
[**ReDeploy**](ServicesApi.md#ReDeploy) | **Post** /v1/services/{id}/redeploy | ReDeploy Service
[**ResumeService**](ServicesApi.md#ResumeService) | **Post** /v1/services/{id}/resume | Resume Service Service resume action is allowed for the following status:  - paused
[**UpdateService**](ServicesApi.md#UpdateService) | **Put** /v1/services/{id} | Update Service
[**UpdateService2**](ServicesApi.md#UpdateService2) | **Patch** /v1/services/{id} | Update Service



## CreateService

> CreateServiceReply CreateService(ctx).Service(service).DryRun(dryRun).Execute()

Create Service

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
    service := *openapiclient.NewCreateService() // CreateService | 
    dryRun := true // bool | If set only run validation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.CreateService(context.Background()).Service(service).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: CreateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.CreateService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**CreateService**](CreateService.md) |  | 
 **dryRun** | **bool** | If set only run validation | 

### Return type

[**CreateServiceReply**](CreateServiceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> map[string]interface{} DeleteService(ctx, id).Execute()

Delete Service Service deletion is allowed for all status.

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
    id := "id_example" // string | The id of the entity to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.DeleteService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeleteService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the entity to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetService

> GetServiceReply GetService(ctx, id).Execute()

Get Service

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
    id := "id_example" // string | The id of the Service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: GetServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceReply**](GetServiceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ListServicesReply ListServices(ctx).AppId(appId).Limit(limit).Offset(offset).Name(name).Execute()

List Service

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
    appId := "appId_example" // string | (Optional) The id of the app (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    name := "name_example" // string | (Optional) A filter for name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ListServices(context.Background()).AppId(appId).Limit(limit).Offset(offset).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServices`: ListServicesReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | (Optional) The id of the app | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **name** | **string** | (Optional) A filter for name | 

### Return type

[**ListServicesReply**](ListServicesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseService

> map[string]interface{} PauseService(ctx, id).Execute()

Pause Service Service pause action is allowed for the following status:  - starting  - healthy  - degraded  - unhealthy  - resuming

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
    id := "id_example" // string | The id of the service to pause.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.PauseService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.PauseService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.PauseService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the service to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseServiceRequest struct via the builder pattern


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


## ReDeploy

> RedeployReply ReDeploy(ctx, id).Info(info).Execute()

ReDeploy Service

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
    id := "id_example" // string | 
    info := *openapiclient.NewRedeployRequestInfo() // RedeployRequestInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ReDeploy(context.Background(), id).Info(info).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ReDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReDeploy`: RedeployReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ReDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **info** | [**RedeployRequestInfo**](RedeployRequestInfo.md) |  | 

### Return type

[**RedeployReply**](RedeployReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeService

> map[string]interface{} ResumeService(ctx, id).Execute()

Resume Service Service resume action is allowed for the following status:  - paused

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
    id := "id_example" // string | The id of the service to pause.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ResumeService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ResumeService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeService`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ResumeService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the service to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeServiceRequest struct via the builder pattern


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


## UpdateService

> UpdateServiceReply UpdateService(ctx, id).Service(service).DryRun(dryRun).Execute()

Update Service

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
    id := "id_example" // string | The id of the entity to update
    service := *openapiclient.NewUpdateService() // UpdateService | 
    dryRun := true // bool | If set, run validation and check that the service exists (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateService(context.Background(), id).Service(service).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: UpdateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | [**UpdateService**](UpdateService.md) |  | 
 **dryRun** | **bool** | If set, run validation and check that the service exists | 

### Return type

[**UpdateServiceReply**](UpdateServiceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService2

> UpdateServiceReply UpdateService2(ctx, id).Service(service).DryRun(dryRun).Execute()

Update Service

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
    id := "id_example" // string | The id of the entity to update
    service := *openapiclient.NewUpdateService() // UpdateService | 
    dryRun := true // bool | If set, run validation and check that the service exists (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateService2(context.Background(), id).Service(service).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateService2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService2`: UpdateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateService2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateService2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | [**UpdateService**](UpdateService.md) |  | 
 **dryRun** | **bool** | If set, run validation and check that the service exists | 

### Return type

[**UpdateServiceReply**](UpdateServiceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

