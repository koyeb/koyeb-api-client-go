# \BillingApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HasUnpaidInvoices**](BillingApi.md#HasUnpaidInvoices) | **Get** /v1/billing/has_unpaid_invoices | Experimental: Has unpaid invoices
[**Manage**](BillingApi.md#Manage) | **Get** /v1/billing/manage | 
[**NextInvoice**](BillingApi.md#NextInvoice) | **Get** /v1/billing/next_invoice | Experimental: Fetch next invoice



## HasUnpaidInvoices

> HasUnpaidInvoicesReply HasUnpaidInvoices(ctx).Execute()

Experimental: Has unpaid invoices



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.HasUnpaidInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.HasUnpaidInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasUnpaidInvoices`: HasUnpaidInvoicesReply
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.HasUnpaidInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHasUnpaidInvoicesRequest struct via the builder pattern


### Return type

[**HasUnpaidInvoicesReply**](HasUnpaidInvoicesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Manage

> ManageReply Manage(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.Manage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.Manage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Manage`: ManageReply
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.Manage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageRequest struct via the builder pattern


### Return type

[**ManageReply**](ManageReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NextInvoice

> NextInvoiceReply NextInvoice(ctx).Execute()

Experimental: Fetch next invoice



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.NextInvoice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.NextInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NextInvoice`: NextInvoiceReply
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.NextInvoice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNextInvoiceRequest struct via the builder pattern


### Return type

[**NextInvoiceReply**](NextInvoiceReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

