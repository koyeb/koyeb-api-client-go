# \OrganizationApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationApi.md#CreateOrganization) | **Post** /v1/organizations | Create organization
[**DeactivateOrganization**](OrganizationApi.md#DeactivateOrganization) | **Post** /v1/organizations/{id}/deactivate | Deactivate an organization
[**DeleteOrganization**](OrganizationApi.md#DeleteOrganization) | **Delete** /v1/organizations/{id} | Delete an organization
[**GetGithubInstallation**](OrganizationApi.md#GetGithubInstallation) | **Get** /v1/github/installation | Fetch github installation configuration
[**GetOrganization**](OrganizationApi.md#GetOrganization) | **Get** /v1/organizations/{id} | Get organization
[**GithubInstallation**](OrganizationApi.md#GithubInstallation) | **Post** /v1/github/installation | Start github installation
[**ReactivateOrganization**](OrganizationApi.md#ReactivateOrganization) | **Post** /v1/organizations/{id}/reactivate | Reactivate an organization
[**SwitchOrganization**](OrganizationApi.md#SwitchOrganization) | **Post** /v1/organizations/{id}/switch | Switch organization context
[**UnscopeOrganizationToken**](OrganizationApi.md#UnscopeOrganizationToken) | **Post** /v1/unscope_organization_token | UnscopeOrganizationToken removes the organization scope from a token. This endpoint is useful when a user wants to remove an organization: by unscoping the token first, the user can then delete the organization without invalidating his token.
[**UpdateOrganization**](OrganizationApi.md#UpdateOrganization) | **Put** /v1/organizations/{id} | Update organization
[**UpdateOrganization2**](OrganizationApi.md#UpdateOrganization2) | **Patch** /v1/organizations/{id} | Update organization
[**UpdateOrganizationPlan**](OrganizationApi.md#UpdateOrganizationPlan) | **Post** /v1/organizations/{id}/plan | Update organization plan
[**UpsertSignupQualification**](OrganizationApi.md#UpsertSignupQualification) | **Post** /v1/organizations/{id}/signup_qualification | Upsert organization&#39;s signup qualification



## CreateOrganization

> CreateOrganizationReply CreateOrganization(ctx).Body(body).Execute()

Create organization

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
    body := *openapiclient.NewCreateOrganizationRequest() // CreateOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.CreateOrganization(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: CreateOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) |  | 

### Return type

[**CreateOrganizationReply**](CreateOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOrganization

> DeactivateOrganizationReply DeactivateOrganization(ctx, id).Body(body).Execute()

Deactivate an organization

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.DeactivateOrganization(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.DeactivateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateOrganization`: DeactivateOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.DeactivateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**DeactivateOrganizationReply**](DeactivateOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganizationReply DeleteOrganization(ctx, id).Execute()

Delete an organization

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
    resp, r, err := apiClient.OrganizationApi.DeleteOrganization(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganization`: DeleteOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOrganizationReply**](DeleteOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubInstallation

> GetGithubInstallationReply GetGithubInstallation(ctx).Execute()

Fetch github installation configuration

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
    resp, r, err := apiClient.OrganizationApi.GetGithubInstallation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetGithubInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubInstallation`: GetGithubInstallationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetGithubInstallation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubInstallationRequest struct via the builder pattern


### Return type

[**GetGithubInstallationReply**](GetGithubInstallationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> GetOrganizationReply GetOrganization(ctx, id).Execute()

Get organization

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
    resp, r, err := apiClient.OrganizationApi.GetOrganization(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: GetOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationReply**](GetOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GithubInstallation

> GithubInstallationReply GithubInstallation(ctx).Body(body).Execute()

Start github installation

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
    body := *openapiclient.NewGithubInstallationRequest() // GithubInstallationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.GithubInstallation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GithubInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GithubInstallation`: GithubInstallationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GithubInstallation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGithubInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GithubInstallationRequest**](GithubInstallationRequest.md) |  | 

### Return type

[**GithubInstallationReply**](GithubInstallationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateOrganization

> ReactivateOrganizationReply ReactivateOrganization(ctx, id).Body(body).Execute()

Reactivate an organization

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.ReactivateOrganization(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.ReactivateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactivateOrganization`: ReactivateOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.ReactivateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**ReactivateOrganizationReply**](ReactivateOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchOrganization

> LoginReply SwitchOrganization(ctx, id).Body(body).SeonFp(seonFp).Execute()

Switch organization context

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.SwitchOrganization(context.Background(), id).Body(body).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.SwitchOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchOrganization`: LoginReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.SwitchOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **seonFp** | **string** | Seon Fingerprint | 

### Return type

[**LoginReply**](LoginReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnscopeOrganizationToken

> LoginReply UnscopeOrganizationToken(ctx).Body(body).SeonFp(seonFp).Execute()

UnscopeOrganizationToken removes the organization scope from a token. This endpoint is useful when a user wants to remove an organization: by unscoping the token first, the user can then delete the organization without invalidating his token.

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
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.UnscopeOrganizationToken(context.Background()).Body(body).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UnscopeOrganizationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnscopeOrganizationToken`: LoginReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UnscopeOrganizationToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnscopeOrganizationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 
 **seonFp** | **string** | Seon Fingerprint | 

### Return type

[**LoginReply**](LoginReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> UpdateOrganizationReply UpdateOrganization(ctx, id).Organization(organization).UpdateMask(updateMask).Execute()

Update organization

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
    organization := *openapiclient.NewOrganization() // Organization | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.UpdateOrganization(context.Background(), id).Organization(organization).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: UpdateOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateOrganizationReply**](UpdateOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization2

> UpdateOrganizationReply UpdateOrganization2(ctx, id).Organization(organization).UpdateMask(updateMask).Execute()

Update organization

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
    organization := *openapiclient.NewOrganization() // Organization | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.UpdateOrganization2(context.Background(), id).Organization(organization).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateOrganization2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization2`: UpdateOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateOrganization2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganization2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateOrganizationReply**](UpdateOrganizationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPlan

> UpdateOrganizationPlanReply UpdateOrganizationPlan(ctx, id).Body(body).Execute()

Update organization plan

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
    body := *openapiclient.NewUpdateOrganizationPlanRequest() // UpdateOrganizationPlanRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.UpdateOrganizationPlan(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateOrganizationPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPlan`: UpdateOrganizationPlanReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateOrganizationPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateOrganizationPlanRequest**](UpdateOrganizationPlanRequest.md) |  | 

### Return type

[**UpdateOrganizationPlanReply**](UpdateOrganizationPlanReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertSignupQualification

> UpsertSignupQualificationReply UpsertSignupQualification(ctx, id).Body(body).Execute()

Upsert organization's signup qualification

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
    body := *openapiclient.NewUpsertSignupQualificationRequest() // UpsertSignupQualificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.UpsertSignupQualification(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpsertSignupQualification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertSignupQualification`: UpsertSignupQualificationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpsertSignupQualification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertSignupQualificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpsertSignupQualificationRequest**](UpsertSignupQualificationRequest.md) |  | 

### Return type

[**UpsertSignupQualificationReply**](UpsertSignupQualificationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

