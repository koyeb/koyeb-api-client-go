# \AppsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsApi.md#CreateApp) | **Post** /v1/apps | Create App
[**DeleteApp**](AppsApi.md#DeleteApp) | **Delete** /v1/apps/{id} | Delete App
[**GetApp**](AppsApi.md#GetApp) | **Get** /v1/apps/{id} | Get App
[**ListApps**](AppsApi.md#ListApps) | **Get** /v1/apps | List App
[**PauseApp**](AppsApi.md#PauseApp) | **Post** /v1/apps/{id}/pause | Pause App
[**ResumeApp**](AppsApi.md#ResumeApp) | **Post** /v1/apps/{id}/resume | Resume App
[**UpdateApp**](AppsApi.md#UpdateApp) | **Put** /v1/apps/{id} | Update App
[**UpdateApp2**](AppsApi.md#UpdateApp2) | **Patch** /v1/apps/{id} | Update App



## CreateApp

> CreateAppReply CreateApp(ctx).Body(body).Execute()

Create App

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
    body := *openapiclient.NewCreateApp() // CreateApp | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.CreateApp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: CreateAppReply
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateApp**](CreateApp.md) |  | 

### Return type

[**CreateAppReply**](CreateAppReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> interface{} DeleteApp(ctx, id).Execute()

Delete App

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
    id := "id_example" // string | The id of the App to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.DeleteApp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApp`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.DeleteApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the App to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## GetApp

> GetAppReply GetApp(ctx, id).Execute()

Get App

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
    id := "id_example" // string | The id of the App

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.GetApp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: GetAppReply
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the App | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppReply**](GetAppReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> ListAppsReply ListApps(ctx).Limit(limit).Offset(offset).Name(name).Execute()

List App

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
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)
    name := "name_example" // string | (Optional) A filter for regions. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.ListApps(context.Background()).Limit(limit).Offset(offset).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: ListAppsReply
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **name** | **string** | (Optional) A filter for regions. | 

### Return type

[**ListAppsReply**](ListAppsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseApp

> interface{} PauseApp(ctx, id).Execute()

Pause App

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
    id := "id_example" // string | The id of the app to pause.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.PauseApp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.PauseApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseApp`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.PauseApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the app to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseAppRequest struct via the builder pattern


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


## ResumeApp

> interface{} ResumeApp(ctx, id).Execute()

Resume App

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
    id := "id_example" // string | The id of the app to resume.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.ResumeApp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ResumeApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeApp`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ResumeApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the app to resume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeAppRequest struct via the builder pattern


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


## UpdateApp

> UpdateAppReply UpdateApp(ctx, id).Execute()

Update App

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
    id := "id_example" // string | The id of the app to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.UpdateApp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UpdateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApp`: UpdateAppReply
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateAppReply**](UpdateAppReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp2

> UpdateAppReply UpdateApp2(ctx, id).Body(body).UpdateMask(updateMask).Execute()

Update App

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
    id := "id_example" // string | The id of the app to update.
    body := *openapiclient.NewUpdateApp() // UpdateApp | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.UpdateApp2(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UpdateApp2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApp2`: UpdateAppReply
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.UpdateApp2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the app to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApp2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateApp**](UpdateApp.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateAppReply**](UpdateAppReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

