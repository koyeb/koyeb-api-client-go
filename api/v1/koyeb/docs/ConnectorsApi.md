# \ConnectorsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorsApi.md#CreateConnector) | **Post** /v1/connectors | Create a new connector
[**DeleteConnector**](ConnectorsApi.md#DeleteConnector) | **Delete** /v1/connectors/{idOrName} | Delete a connector
[**GetConnector**](ConnectorsApi.md#GetConnector) | **Get** /v1/connectors/{idOrName} | Fetch a connector
[**ListConnectors**](ConnectorsApi.md#ListConnectors) | **Get** /v1/connectors | List connectors
[**UpdateConnector**](ConnectorsApi.md#UpdateConnector) | **Put** /v1/connectors/{idOrName} | Update a connector



## CreateConnector

> CreateConnectorReply CreateConnector(ctx).Body(body).Execute()

Create a new connector

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
    body := *openapiclient.NewConnectorUpsert() // ConnectorUpsert | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsApi.CreateConnector(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.CreateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnector`: CreateConnectorReply
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.CreateConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConnectorUpsert**](ConnectorUpsert.md) |  | 

### Return type

[**CreateConnectorReply**](CreateConnectorReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnector

> interface{} DeleteConnector(ctx, idOrName).Execute()

Delete a connector

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
    idOrName := "idOrName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsApi.DeleteConnector(context.Background(), idOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.DeleteConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnector`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.DeleteConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRequest struct via the builder pattern


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


## GetConnector

> GetConnectorReply GetConnector(ctx, idOrName).Execute()

Fetch a connector

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
    idOrName := "idOrName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsApi.GetConnector(context.Background(), idOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.GetConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnector`: GetConnectorReply
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConnectorReply**](GetConnectorReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> ListConnectorsReply ListConnectors(ctx).Limit(limit).Offset(offset).NameFilter(nameFilter).Type_(type_).Execute()

List connectors

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
    nameFilter := "nameFilter_example" // string | (Optional) Filter by name. (optional)
    type_ := "type__example" // string | (Optional) A way to retrieve connectors only of a specific type. (optional) (default to "unknown")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsApi.ListConnectors(context.Background()).Limit(limit).Offset(offset).NameFilter(nameFilter).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.ListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectors`: ListConnectorsReply
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 
 **nameFilter** | **string** | (Optional) Filter by name. | 
 **type_** | **string** | (Optional) A way to retrieve connectors only of a specific type. | [default to &quot;unknown&quot;]

### Return type

[**ListConnectorsReply**](ListConnectorsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnector

> UpdateConnectorReply UpdateConnector(ctx, idOrName).Body(body).Execute()

Update a connector

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
    idOrName := "idOrName_example" // string | 
    body := *openapiclient.NewConnectorUpsert() // ConnectorUpsert | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorsApi.UpdateConnector(context.Background(), idOrName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.UpdateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnector`: UpdateConnectorReply
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.UpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectorUpsert**](ConnectorUpsert.md) |  | 

### Return type

[**UpdateConnectorReply**](UpdateConnectorReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

