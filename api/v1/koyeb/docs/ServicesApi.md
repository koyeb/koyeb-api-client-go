# \ServicesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/services | Create Service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/services/{id} | Delete Service
[**DeprecatedCreateService**](ServicesApi.md#DeprecatedCreateService) | **Post** /v1/apps/{app_id_or_name}/services | Create Service
[**DeprecatedDeleteService**](ServicesApi.md#DeprecatedDeleteService) | **Delete** /v1/apps/{app_id_or_name}/services/{id_or_name} | Delete Service
[**DeprecatedGetRevision**](ServicesApi.md#DeprecatedGetRevision) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name}/revisions/{id} | Get Revision
[**DeprecatedGetService**](ServicesApi.md#DeprecatedGetService) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name} | Get Service
[**DeprecatedListRevisions**](ServicesApi.md#DeprecatedListRevisions) | **Get** /v1/apps/{app_id_or_name}/services/{id_or_name}/revisions | List Revisions
[**DeprecatedListServices**](ServicesApi.md#DeprecatedListServices) | **Get** /v1/apps/{app_id_or_name}/services | List Service
[**DeprecatedReDeploy**](ServicesApi.md#DeprecatedReDeploy) | **Post** /v1/apps/{app_id_or_name}/services/{id_or_name}/redeploy | ReDeploy Service
[**DeprecatedUpdateService**](ServicesApi.md#DeprecatedUpdateService) | **Put** /v1/apps/{app_id_or_name}/services/{id_or_name} | Update Service
[**DeprecatedUpdateService2**](ServicesApi.md#DeprecatedUpdateService2) | **Patch** /v1/apps/{app_id_or_name}/services/{id_or_name} | Update Service
[**GetService**](ServicesApi.md#GetService) | **Get** /v1/services/{id} | Get Service
[**ListServices**](ServicesApi.md#ListServices) | **Get** /v1/services | List Service
[**ReDeploy**](ServicesApi.md#ReDeploy) | **Post** /v1/services/{id}/redeploy | ReDeploy Service
[**UpdateService**](ServicesApi.md#UpdateService) | **Put** /v1/services/{id} | Update Service
[**UpdateService2**](ServicesApi.md#UpdateService2) | **Patch** /v1/services/{id} | Update Service



## CreateService

> CreateServiceReply CreateService(ctx).Body(body).DryRun(dryRun).Execute()

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
    body := *openapiclient.NewCreateService() // CreateService | 
    dryRun := true // bool | If set only run validation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.CreateService(context.Background()).Body(body).DryRun(dryRun).Execute()
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

> interface{} DeleteService(ctx, id).Execute()

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
    id := "id_example" // string | The id of the entity to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeleteService(context.Background(), id).Execute()
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
**id** | **string** | The id of the entity to delete | 

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


## DeprecatedCreateService

> CreateServiceReply DeprecatedCreateService(ctx, appIdOrName).Body(body).DryRun(dryRun).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedCreateService(context.Background(), appIdOrName).Body(body).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedCreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedCreateService`: CreateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedCreateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedCreateServiceRequest struct via the builder pattern


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


## DeprecatedDeleteService

> interface{} DeprecatedDeleteService(ctx, appIdOrName, idOrName).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedDeleteService(context.Background(), appIdOrName, idOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedDeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedDeleteService`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedDeleteService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedDeleteServiceRequest struct via the builder pattern


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


## DeprecatedGetRevision

> DeprecatedGetServiceRevisionReply DeprecatedGetRevision(ctx, appIdOrName, idOrName, id).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedGetRevision(context.Background(), appIdOrName, idOrName, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedGetRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedGetRevision`: DeprecatedGetServiceRevisionReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedGetRevision`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeprecatedGetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeprecatedGetServiceRevisionReply**](DeprecatedGetServiceRevisionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedGetService

> GetServiceReply DeprecatedGetService(ctx, appIdOrName, idOrName).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedGetService(context.Background(), appIdOrName, idOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedGetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedGetService`: GetServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedGetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedGetServiceRequest struct via the builder pattern


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


## DeprecatedListRevisions

> DeprecatedListServiceRevisionsReply DeprecatedListRevisions(ctx, appIdOrName, idOrName).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedListRevisions(context.Background(), appIdOrName, idOrName).Limit(limit).Offset(offset).Statuses(statuses).DeploymentGroups(deploymentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedListRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedListRevisions`: DeprecatedListServiceRevisionsReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedListRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedListRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **statuses** | **[]string** | (Optional) Statuses to filter on. | 
 **deploymentGroups** | **[]string** | (Optional) Only fetch revisions in this deployment group. | 

### Return type

[**DeprecatedListServiceRevisionsReply**](DeprecatedListServiceRevisionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedListServices

> ListServicesReply DeprecatedListServices(ctx, appIdOrName).Limit(limit).Offset(offset).Name(name).Execute()

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
    name := "name_example" // string | (Optional) A filter for name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeprecatedListServices(context.Background(), appIdOrName).Limit(limit).Offset(offset).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedListServices`: ListServicesReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedListServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **name** | **string** | (Optional) A filter for name. | 

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


## DeprecatedReDeploy

> DeprecatedRedeployReply DeprecatedReDeploy(ctx, appIdOrName, idOrName).Body(body).Execute()

ReDeploy Service

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
    body := *openapiclient.NewDeprecatedRedeployRequestInfo() // DeprecatedRedeployRequestInfo | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeprecatedReDeploy(context.Background(), appIdOrName, idOrName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedReDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedReDeploy`: DeprecatedRedeployReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedReDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** |  | 
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedReDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeprecatedRedeployRequestInfo**](DeprecatedRedeployRequestInfo.md) |  | 

### Return type

[**DeprecatedRedeployReply**](DeprecatedRedeployReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecatedUpdateService

> UpdateServiceReply DeprecatedUpdateService(ctx, appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedUpdateService(context.Background(), appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedUpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedUpdateService`: UpdateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedUpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedUpdateServiceRequest struct via the builder pattern


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


## DeprecatedUpdateService2

> UpdateServiceReply DeprecatedUpdateService2(ctx, appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    resp, r, err := api_client.ServicesApi.DeprecatedUpdateService2(context.Background(), appIdOrName, idOrName).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeprecatedUpdateService2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedUpdateService2`: UpdateServiceReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeprecatedUpdateService2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | The id or the name of the App | 
**idOrName** | **string** | The id or the name of the entity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedUpdateService2Request struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the Service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetService(context.Background(), id).Execute()
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
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | (Optional) The id of the app. (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    name := "name_example" // string | (Optional) A filter for name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ListServices(context.Background()).AppId(appId).Limit(limit).Offset(offset).Name(name).Execute()
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
 **appId** | **string** | (Optional) The id of the app. | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **name** | **string** | (Optional) A filter for name. | 

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


## ReDeploy

> RedeployReply ReDeploy(ctx, id).Body(body).Execute()

ReDeploy Service

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
    body := *openapiclient.NewRedeployRequestInfo() // RedeployRequestInfo | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ReDeploy(context.Background(), id).Body(body).Execute()
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

 **body** | [**RedeployRequestInfo**](RedeployRequestInfo.md) |  | 

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


## UpdateService

> UpdateServiceReply UpdateService(ctx, id).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    id := "id_example" // string | The id of the entity to update
    body := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set, run validation and check that the service exists. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateService(context.Background(), id).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
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

 **body** | [**UpdateService**](UpdateService.md) |  | 
 **updateMask** | **string** |  | 
 **dryRun** | **bool** | If set, run validation and check that the service exists. | 

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

> UpdateServiceReply UpdateService2(ctx, id).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    id := "id_example" // string | The id of the entity to update
    body := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set, run validation and check that the service exists. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateService2(context.Background(), id).Body(body).UpdateMask(updateMask).DryRun(dryRun).Execute()
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

 **body** | [**UpdateService**](UpdateService.md) |  | 
 **updateMask** | **string** |  | 
 **dryRun** | **bool** | If set, run validation and check that the service exists. | 

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

