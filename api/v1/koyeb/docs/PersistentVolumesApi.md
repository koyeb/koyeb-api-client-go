# \PersistentVolumesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersistentVolume**](PersistentVolumesApi.md#CreatePersistentVolume) | **Post** /v1/volumes | Create a PersistentVolume
[**DeletePersistentVolume**](PersistentVolumesApi.md#DeletePersistentVolume) | **Delete** /v1/volumes/{id} | Delete a PersistentVolume
[**GetPersistentVolume**](PersistentVolumesApi.md#GetPersistentVolume) | **Get** /v1/volumes/{id} | Get a PersistentVolume
[**ListPersistentVolumeEvents**](PersistentVolumesApi.md#ListPersistentVolumeEvents) | **Get** /v1/volume_events | List Persistent Volume events
[**ListPersistentVolumes**](PersistentVolumesApi.md#ListPersistentVolumes) | **Get** /v1/volumes | List all PersistentVolumes
[**UpdatePersistentVolume**](PersistentVolumesApi.md#UpdatePersistentVolume) | **Post** /v1/volumes/{id} | Update a PersistentVolume



## CreatePersistentVolume

> CreatePersistentVolumeReply CreatePersistentVolume(ctx).Body(body).Execute()

Create a PersistentVolume

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
    body := *openapiclient.NewCreatePersistentVolumeRequest() // CreatePersistentVolumeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.CreatePersistentVolume(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.CreatePersistentVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersistentVolume`: CreatePersistentVolumeReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.CreatePersistentVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersistentVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreatePersistentVolumeRequest**](CreatePersistentVolumeRequest.md) |  | 

### Return type

[**CreatePersistentVolumeReply**](CreatePersistentVolumeReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersistentVolume

> DeletePersistentVolumeReply DeletePersistentVolume(ctx, id).Execute()

Delete a PersistentVolume

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.DeletePersistentVolume(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.DeletePersistentVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersistentVolume`: DeletePersistentVolumeReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.DeletePersistentVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersistentVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePersistentVolumeReply**](DeletePersistentVolumeReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersistentVolume

> GetPersistentVolumeReply GetPersistentVolume(ctx, id).Execute()

Get a PersistentVolume

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.GetPersistentVolume(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.GetPersistentVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersistentVolume`: GetPersistentVolumeReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.GetPersistentVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersistentVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPersistentVolumeReply**](GetPersistentVolumeReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersistentVolumeEvents

> ListPersistentVolumeEventsReply ListPersistentVolumeEvents(ctx).PersistentVolumeId(persistentVolumeId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()

List Persistent Volume events

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
    persistentVolumeId := "persistentVolumeId_example" // string | (Optional) Filter on persistent volume id (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on persistent volume event types (optional)
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.ListPersistentVolumeEvents(context.Background()).PersistentVolumeId(persistentVolumeId).Types(types).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.ListPersistentVolumeEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersistentVolumeEvents`: ListPersistentVolumeEventsReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.ListPersistentVolumeEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersistentVolumeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **persistentVolumeId** | **string** | (Optional) Filter on persistent volume id | 
 **types** | **[]string** | (Optional) Filter on persistent volume event types | 
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 

### Return type

[**ListPersistentVolumeEventsReply**](ListPersistentVolumeEventsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersistentVolumes

> ListPersistentVolumesReply ListPersistentVolumes(ctx).Limit(limit).Offset(offset).ServiceId(serviceId).Region(region).Name(name).Execute()

List all PersistentVolumes

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
    limit := "limit_example" // string | (Optional) The number of items to return (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return (optional)
    serviceId := "serviceId_example" // string | (Optional) A filter for the service id (optional)
    region := "region_example" // string | (Optional) A filter for the region (optional)
    name := "name_example" // string | (Optional) A filter for the name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.ListPersistentVolumes(context.Background()).Limit(limit).Offset(offset).ServiceId(serviceId).Region(region).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.ListPersistentVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersistentVolumes`: ListPersistentVolumesReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.ListPersistentVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersistentVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **serviceId** | **string** | (Optional) A filter for the service id | 
 **region** | **string** | (Optional) A filter for the region | 
 **name** | **string** | (Optional) A filter for the name | 

### Return type

[**ListPersistentVolumesReply**](ListPersistentVolumesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersistentVolume

> UpdatePersistentVolumeReply UpdatePersistentVolume(ctx, id).Body(body).Execute()

Update a PersistentVolume

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
    body := *openapiclient.NewUpdatePersistentVolumeRequest() // UpdatePersistentVolumeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersistentVolumesApi.UpdatePersistentVolume(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentVolumesApi.UpdatePersistentVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersistentVolume`: UpdatePersistentVolumeReply
    fmt.Fprintf(os.Stdout, "Response from `PersistentVolumesApi.UpdatePersistentVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersistentVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdatePersistentVolumeRequest**](UpdatePersistentVolumeRequest.md) |  | 

### Return type

[**UpdatePersistentVolumeReply**](UpdatePersistentVolumeReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

