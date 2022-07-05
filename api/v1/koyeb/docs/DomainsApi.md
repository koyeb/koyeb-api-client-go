# \DomainsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainsApi.md#CreateDomain) | **Post** /v1/domains | 
[**DeleteDomain**](DomainsApi.md#DeleteDomain) | **Delete** /v1/domains/{id} | 
[**GetDomain**](DomainsApi.md#GetDomain) | **Get** /v1/domains/{id} | 
[**ListDomains**](DomainsApi.md#ListDomains) | **Get** /v1/domains | 
[**RefreshDomainStatus**](DomainsApi.md#RefreshDomainStatus) | **Post** /v1/domains/{id}/refresh | 
[**UpdateDomain**](DomainsApi.md#UpdateDomain) | **Patch** /v1/domains/{id} | 



## CreateDomain

> CreateDomainReply CreateDomain(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateDomain() // CreateDomain | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.CreateDomain(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: CreateDomainReply
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDomain**](CreateDomain.md) |  | 

### Return type

[**CreateDomainReply**](CreateDomainReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> interface{} DeleteDomain(ctx, id).Execute()



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
    resp, r, err := api_client.DomainsApi.DeleteDomain(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomain`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


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


## GetDomain

> GetDomainReply GetDomain(ctx, id).Execute()



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
    resp, r, err := api_client.DomainsApi.GetDomain(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: GetDomainReply
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDomainReply**](GetDomainReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> ListDomainsReply ListDomains(ctx).Limit(limit).Offset(offset).Types(types).Statuses(statuses).AppIds(appIds).Name(name).Execute()



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
    types := []string{"Types_example"} // []string | (Optional) A filter for types. (optional)
    statuses := []string{"Statuses_example"} // []string | (Optional) A filter for statuses. (optional)
    appIds := []string{"Inner_example"} // []string | (Optional) A filter for apps. (optional)
    name := "name_example" // string | (Optional) A filter for name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.ListDomains(context.Background()).Limit(limit).Offset(offset).Types(types).Statuses(statuses).AppIds(appIds).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.ListDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomains`: ListDomainsReply
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.ListDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **types** | **[]string** | (Optional) A filter for types. | 
 **statuses** | **[]string** | (Optional) A filter for statuses. | 
 **appIds** | **[]string** | (Optional) A filter for apps. | 
 **name** | **string** | (Optional) A filter for name. | 

### Return type

[**ListDomainsReply**](ListDomainsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshDomainStatus

> interface{} RefreshDomainStatus(ctx, id).Execute()



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
    resp, r, err := api_client.DomainsApi.RefreshDomainStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.RefreshDomainStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshDomainStatus`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.RefreshDomainStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshDomainStatusRequest struct via the builder pattern


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


## UpdateDomain

> UpdateDomainReply UpdateDomain(ctx, id).Body(body).UpdateMask(updateMask).Execute()



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
    body := *openapiclient.NewUpdateDomain() // UpdateDomain | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.UpdateDomain(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UpdateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomain`: UpdateDomainReply
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateDomain**](UpdateDomain.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateDomainReply**](UpdateDomainReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

