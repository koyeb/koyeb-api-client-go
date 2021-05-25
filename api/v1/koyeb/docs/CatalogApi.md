# \CatalogApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogFunction**](CatalogApi.md#GetCatalogFunction) | **Get** /v1/catalog/functions/{name} | Fetch an item of the function catalog
[**GetCatalogStack**](CatalogApi.md#GetCatalogStack) | **Get** /v1/catalog/stacks/{name} | Fetch an item of the stack catalog
[**GetCatalogStore**](CatalogApi.md#GetCatalogStore) | **Get** /v1/catalog/stores/{name} | Fetch an item of the store catalog
[**ListCatalogFunctions**](CatalogApi.md#ListCatalogFunctions) | **Get** /v1/catalog/functions | Show catalog of functions
[**ListCatalogStacks**](CatalogApi.md#ListCatalogStacks) | **Get** /v1/catalog/stacks | Show catalog of stacks
[**ListCatalogStores**](CatalogApi.md#ListCatalogStores) | **Get** /v1/catalog/stores | Show catalog of stores
[**ValidateYaml**](CatalogApi.md#ValidateYaml) | **Post** /v1/validate_yaml | Validate a yaml stack



## GetCatalogFunction

> GetCatalogFunctionsReply GetCatalogFunction(ctx, name).Version(version).Execute()

Fetch an item of the function catalog

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
    name := "name_example" // string | The name of the function
    version := "version_example" // string | An optional function version, if it is not set return the latest. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogFunction(context.Background(), name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogFunction`: GetCatalogFunctionsReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the function | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | An optional function version, if it is not set return the latest. | 

### Return type

[**GetCatalogFunctionsReply**](GetCatalogFunctionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogStack

> GetCatalogStackReply GetCatalogStack(ctx, name).Execute()

Fetch an item of the stack catalog

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogStack(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogStack`: GetCatalogStackReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCatalogStackReply**](GetCatalogStackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogStore

> CatalogStoreReply GetCatalogStore(ctx, name).Execute()

Fetch an item of the store catalog

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogStore(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogStore`: CatalogStoreReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogStoreReply**](CatalogStoreReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogFunctions

> ListCatalogFunctionsReply ListCatalogFunctions(ctx).Name(name).Limit(limit).Offset(offset).Execute()

Show catalog of functions

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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ListCatalogFunctions(context.Background()).Name(name).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ListCatalogFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogFunctions`: ListCatalogFunctionsReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ListCatalogFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListCatalogFunctionsReply**](ListCatalogFunctionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogStacks

> ListCatalogStacksReply ListCatalogStacks(ctx).Name(name).Limit(limit).Offset(offset).Execute()

Show catalog of stacks

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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ListCatalogStacks(context.Background()).Name(name).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ListCatalogStacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogStacks`: ListCatalogStacksReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ListCatalogStacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListCatalogStacksReply**](ListCatalogStacksReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogStores

> ListCatalogStoresReply ListCatalogStores(ctx).Name(name).Limit(limit).Offset(offset).Execute()

Show catalog of stores

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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ListCatalogStores(context.Background()).Name(name).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ListCatalogStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCatalogStores`: ListCatalogStoresReply
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ListCatalogStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListCatalogStoresReply**](ListCatalogStoresReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateYaml

> interface{} ValidateYaml(ctx).Body(body).Execute()

Validate a yaml stack

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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ValidateYaml(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ValidateYaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateYaml`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ValidateYaml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateYamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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

