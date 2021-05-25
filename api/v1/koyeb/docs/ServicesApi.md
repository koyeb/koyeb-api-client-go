# \ServicesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/apps/{app_id_or_name}/services | Create Service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/apps/{app_id_or_name}/services/{id_or_name} | Delete Service
[**GetRevision**](ServicesApi.md#GetRevision) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name}/revisions/{id} | Get Revision
[**GetService**](ServicesApi.md#GetService) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name} | Get Service
[**ListRevisions**](ServicesApi.md#ListRevisions) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name}/revisions | List Revisions
[**ListServices**](ServicesApi.md#ListServices) | **Get** /v1/apps/{app_id_or_name}/services | List Service
[**UpdateService**](ServicesApi.md#UpdateService) | **Put** /v1/apps/{app_id_or_name}/services/{id_or_name} | Update Service
[**UpdateService2**](ServicesApi.md#UpdateService2) | **Patch** /v1/apps/{app_id_or_name}/services/{id_or_name} | Update Service



## CreateService

> CreateServiceReply CreateService(ctx, appIdOrName).Body(body).DryRun(dryRun).Execute()

Create Service

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the App
    body := *openapiclient.NewCreateService() // CreateService | 
    dryRun := true // bool | If set only run validation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.CreateService(context.Background(), appIdOrName).Body(body).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: CreateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.CreateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateService**](CreateService.md) |  | 
 **dryRun** | **bool** | If set only run validation. | 

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

> interface{} DeleteService(ctx, appIdOrName, idOrName).Execute()

Delete Service

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the App
    idOrName := "idOrName_example" // string | The id or the name of the entity to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeleteService(context.Background(), appIdOrName, idOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteService`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeleteService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetRevision

> GetServiceRevisionReply GetRevision(ctx, appIdOrName, idOrName, id).Execute()

Get Revision

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the app
    idOrName := "idOrName_example" // string | The id or the name of the service
    id := "id_example" // string | The id of the revision to fetch or `_latest` to get the latest one

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetRevision(context.Background(), appIdOrName, idOrName, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevision`: GetServiceRevisionReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the app | 
**idOrName** | **string** | The id or the name of the service | 
**id** | **string** | The id of the revision to fetch or &#x60;_latest&#x60; to get the latest one | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetServiceRevisionReply**](GetServiceRevisionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> GetServiceReply GetService(ctx, appIdOrName, idOrName).Execute()

Get Service

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the App
    idOrName := "idOrName_example" // string | The id or the name of the Service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetService(context.Background(), appIdOrName, idOrName).Execute()
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
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the Service | 

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


## ListRevisions

> ListServiceRevisionsReply ListRevisions(ctx, appIdOrName, idOrName).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()

List Revisions

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
    appIdOrName := "appIdOrName_example" // string | 
    idOrName := "idOrName_example" // string | 
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Statuses to filter on. (optional)
    deploymentGroups := []string{"Inner_example"} // []string | (Optional) Only fetch revisions in this deployment group. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ListRevisions(context.Background(), appIdOrName, idOrName).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ListRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRevisions`: ListServiceRevisionsReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ListRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **statuses** | **[]string** | (Optional) Statuses to filter on. | 
 **deploymentGroups** | **[]string** | (Optional) Only fetch revisions in this deployment group. | 

### Return type

[**ListServiceRevisionsReply**](ListServiceRevisionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ListServicesReply ListServices(ctx, appIdOrName).Limit(limit).Offset(offset).Name(name).Execute()

List Service

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
    appIdOrName := "appIdOrName_example" // string | 
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    name := "name_example" // string | (Optional) A filter for regions. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ListServices(context.Background(), appIdOrName).Limit(limit).Offset(offset).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServices`: ListServicesReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ListServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **name** | **string** | (Optional) A filter for regions. | 

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


## UpdateService

> UpdateServiceReply UpdateService(ctx, appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

Update Service

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the App
    idOrName := "idOrName_example" // string | The id or the name of the entity to update
    body := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set only run validation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateService(context.Background(), appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
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
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateService**](UpdateService.md) |  | 
 **updateMask** | **string** |  | 
 **dryRun** | **bool** | If set only run validation. | 

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

> UpdateServiceReply UpdateService2(ctx, appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

Update Service

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
    appIdOrName := "appIdOrName_example" // string | The id or the name of the App
    idOrName := "idOrName_example" // string | The id or the name of the entity to update
    body := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set only run validation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateService2(context.Background(), appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
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
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateService2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateService**](UpdateService.md) |  | 
 **updateMask** | **string** |  | 
 **dryRun** | **bool** | If set only run validation. | 

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

