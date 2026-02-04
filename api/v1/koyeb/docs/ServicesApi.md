# \ServicesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Autocomplete**](ServicesApi.md#Autocomplete) | **Post** /v1/services-autocomplete | Autocomplete definition
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/services | Create Service
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/services/{id} | Delete Service
[**DeleteServiceScaling**](ServicesApi.md#DeleteServiceScaling) | **Delete** /v1/services/{id}/scale | Delete Service Scaling
[**GetService**](ServicesApi.md#GetService) | **Get** /v1/services/{id} | Get Service
[**GetServiceScaling**](ServicesApi.md#GetServiceScaling) | **Get** /v1/services/{id}/scale | Get Service Scaling
[**ListServiceEvents**](ServicesApi.md#ListServiceEvents) | **Get** /v1/service_events | List Service events
[**ListServices**](ServicesApi.md#ListServices) | **Get** /v1/services | List Services
[**PauseService**](ServicesApi.md#PauseService) | **Post** /v1/services/{id}/pause | Pause Service
[**ReDeploy**](ServicesApi.md#ReDeploy) | **Post** /v1/services/{id}/redeploy | ReDeploy Service
[**ResumeService**](ServicesApi.md#ResumeService) | **Post** /v1/services/{id}/resume | Resume Service
[**UpdateService**](ServicesApi.md#UpdateService) | **Put** /v1/services/{id} | Update Service
[**UpdateService2**](ServicesApi.md#UpdateService2) | **Patch** /v1/services/{id} | Update Service
[**UpdateServiceScaling**](ServicesApi.md#UpdateServiceScaling) | **Put** /v1/services/{id}/scale | Update Service Scaling



## Autocomplete

> AutocompleteReply Autocomplete(ctx).Body(body).Execute()

Autocomplete definition



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
    body := *openapiclient.NewAutocompleteRequest() // AutocompleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.Autocomplete(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.Autocomplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Autocomplete`: AutocompleteReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.Autocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AutocompleteRequest**](AutocompleteRequest.md) |  | 

### Return type

[**AutocompleteReply**](AutocompleteReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "./openapi"
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


## DeleteServiceScaling

> map[string]interface{} DeleteServiceScaling(ctx, id).Execute()

Delete Service Scaling



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
    id := "id_example" // string | The id of the service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.DeleteServiceScaling(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteServiceScaling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceScaling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.DeleteServiceScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceScalingRequest struct via the builder pattern


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
    openapiclient "./openapi"
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


## GetServiceScaling

> GetServiceScalingReply GetServiceScaling(ctx, id).Execute()

Get Service Scaling



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
    id := "id_example" // string | The id of the service

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.GetServiceScaling(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceScaling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceScaling`: GetServiceScalingReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetServiceScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceScalingReply**](GetServiceScalingReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceEvents

> ListServiceEventsReply ListServiceEvents(ctx).ServiceId(serviceId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()

List Service events

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
    serviceId := "serviceId_example" // string | (Optional) Filter on service id (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on service event types (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ListServiceEvents(context.Background()).ServiceId(serviceId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ListServiceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceEvents`: ListServiceEventsReply
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ListServiceEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | (Optional) Filter on service id | 
 **types** | **[]string** | (Optional) Filter on service event types | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**ListServiceEventsReply**](ListServiceEventsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ListServicesReply ListServices(ctx).AppId(appId).Limit(limit).Offset(offset).Name(name).Types(types).Statuses(statuses).Regions(regions).Execute()

List Services

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
    appId := "appId_example" // string | (Optional) The id of the app (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    name := "name_example" // string | (Optional) A filter for name (optional)
    types := []string{"Types_example"} // []string | (Optional) Filter on service types (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter on service statuses (optional)
    regions := []string{"Inner_example"} // []string | (Optional) Filter on regions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ListServices(context.Background()).AppId(appId).Limit(limit).Offset(offset).Name(name).Types(types).Statuses(statuses).Regions(regions).Execute()
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
 **types** | **[]string** | (Optional) Filter on service types | 
 **statuses** | **[]string** | (Optional) Filter on service statuses | 
 **regions** | **[]string** | (Optional) Filter on regions | 

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

Pause Service



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
    openapiclient "./openapi"
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

> map[string]interface{} ResumeService(ctx, id).SkipBuild(skipBuild).UseCache(useCache).Execute()

Resume Service



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
    id := "id_example" // string | The id of the service to pause.
    skipBuild := true // bool | If set to true, the build stage will be skipped and the image coming from the last successful build step will be used instead. The call fails if no previous successful builds happened. (optional)
    useCache := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.ResumeService(context.Background(), id).SkipBuild(skipBuild).UseCache(useCache).Execute()
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

 **skipBuild** | **bool** | If set to true, the build stage will be skipped and the image coming from the last successful build step will be used instead. The call fails if no previous successful builds happened. | 
 **useCache** | **bool** |  | 

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

> UpdateServiceReply UpdateService(ctx, id).Service(service).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    service := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set, run validation and check that the service exists (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateService(context.Background(), id).Service(service).UpdateMask(updateMask).DryRun(dryRun).Execute()
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
 **updateMask** | **string** |  | 
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

> UpdateServiceReply UpdateService2(ctx, id).Service(service).UpdateMask(updateMask).DryRun(dryRun).Execute()

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
    service := *openapiclient.NewUpdateService() // UpdateService | 
    updateMask := "updateMask_example" // string |  (optional)
    dryRun := true // bool | If set, run validation and check that the service exists (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateService2(context.Background(), id).Service(service).UpdateMask(updateMask).DryRun(dryRun).Execute()
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
 **updateMask** | **string** |  | 
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


## UpdateServiceScaling

> map[string]interface{} UpdateServiceScaling(ctx, id).Body(body).Execute()

Update Service Scaling



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
    id := "id_example" // string | The id of the service to scale
    body := *openapiclient.NewUpdateServiceScalingRequest() // UpdateServiceScalingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServicesApi.UpdateServiceScaling(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateServiceScaling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceScaling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateServiceScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the service to scale | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateServiceScalingRequest**](UpdateServiceScalingRequest.md) |  | 

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

