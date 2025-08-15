# \ProfileApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptOrganizationInvitation**](ProfileApi.md#AcceptOrganizationInvitation) | **Post** /v1/account/organization_invitations/{id}/accept | Accept Organization Invitation
[**ClearIdenfyVerificationResult**](ProfileApi.md#ClearIdenfyVerificationResult) | **Post** /v1/account/idenfy | ClearIdenfyVerificationResult marks the current result for idenfy as superseded
[**DeclineOrganizationInvitation**](ProfileApi.md#DeclineOrganizationInvitation) | **Post** /v1/account/organization_invitations/{id}/decline | Decline Organization Invitation
[**GetCurrentOrganization**](ProfileApi.md#GetCurrentOrganization) | **Get** /v1/account/organization | Get Current Organization
[**GetCurrentUser**](ProfileApi.md#GetCurrentUser) | **Get** /v1/account/profile | Get Current User
[**GetIdenfyToken**](ProfileApi.md#GetIdenfyToken) | **Get** /v1/account/idenfy | Begin a session with iDenfy, emit an authToken
[**GetOAuthOptions**](ProfileApi.md#GetOAuthOptions) | **Get** /v1/account/oauth | Get OAuth Providers
[**GetUserOrganizationInvitation**](ProfileApi.md#GetUserOrganizationInvitation) | **Get** /v1/account/organization_invitations/{id} | Get User Organization Invitation
[**GetUserSettings**](ProfileApi.md#GetUserSettings) | **Get** /v1/account/settings | 
[**ListUserOrganizationInvitations**](ProfileApi.md#ListUserOrganizationInvitations) | **Get** /v1/account/organization_invitations | List User Organization Invitations
[**ListUserOrganizations**](ProfileApi.md#ListUserOrganizations) | **Get** /v1/account/organizations | List User Organizations
[**OAuthCallback**](ProfileApi.md#OAuthCallback) | **Post** /v1/account/oauth | Authenticate using OAuth
[**ResendEmailValidation**](ProfileApi.md#ResendEmailValidation) | **Post** /v1/account/resend_validation | Resend Email Verification
[**ResetPassword**](ProfileApi.md#ResetPassword) | **Post** /v1/account/reset_password | Reset Password
[**Signup**](ProfileApi.md#Signup) | **Post** /v1/account/signup | Signup
[**UpdatePassword**](ProfileApi.md#UpdatePassword) | **Post** /v1/account/update_password | Update Password
[**UpdateUser**](ProfileApi.md#UpdateUser) | **Put** /v1/account/profile | Update User
[**UpdateUser2**](ProfileApi.md#UpdateUser2) | **Patch** /v1/account/profile | Update User
[**UpdateUserSettings**](ProfileApi.md#UpdateUserSettings) | **Patch** /v1/account/settings | 
[**Validate**](ProfileApi.md#Validate) | **Post** /v1/account/validate/{id} | Validate



## AcceptOrganizationInvitation

> AcceptOrganizationInvitationReply AcceptOrganizationInvitation(ctx, id).Body(body).Execute()

Accept Organization Invitation

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
    id := "id_example" // string | The id of the organization invitation to accept
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.AcceptOrganizationInvitation(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.AcceptOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptOrganizationInvitation`: AcceptOrganizationInvitationReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.AcceptOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the organization invitation to accept | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**AcceptOrganizationInvitationReply**](AcceptOrganizationInvitationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearIdenfyVerificationResult

> map[string]interface{} ClearIdenfyVerificationResult(ctx).Body(body).Execute()

ClearIdenfyVerificationResult marks the current result for idenfy as superseded

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
    body := *openapiclient.NewClearIdenfyVerificationResultRequest() // ClearIdenfyVerificationResultRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.ClearIdenfyVerificationResult(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ClearIdenfyVerificationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearIdenfyVerificationResult`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ClearIdenfyVerificationResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearIdenfyVerificationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClearIdenfyVerificationResultRequest**](ClearIdenfyVerificationResultRequest.md) |  | 

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


## DeclineOrganizationInvitation

> DeclineOrganizationInvitationReply DeclineOrganizationInvitation(ctx, id).Body(body).Execute()

Decline Organization Invitation

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
    id := "id_example" // string | The id of the organization invitation to decline
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.DeclineOrganizationInvitation(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.DeclineOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclineOrganizationInvitation`: DeclineOrganizationInvitationReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.DeclineOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the organization invitation to decline | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**DeclineOrganizationInvitationReply**](DeclineOrganizationInvitationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentOrganization

> GetOrganizationReply GetCurrentOrganization(ctx).Execute()

Get Current Organization

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
    resp, r, err := apiClient.ProfileApi.GetCurrentOrganization(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetCurrentOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentOrganization`: GetOrganizationReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetCurrentOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentOrganizationRequest struct via the builder pattern


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


## GetCurrentUser

> UserReply GetCurrentUser(ctx).Execute()

Get Current User

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
    resp, r, err := apiClient.ProfileApi.GetCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: UserReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**UserReply**](UserReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdenfyToken

> GetIdenfyTokenReply GetIdenfyToken(ctx).Execute()

Begin a session with iDenfy, emit an authToken

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
    resp, r, err := apiClient.ProfileApi.GetIdenfyToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetIdenfyToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdenfyToken`: GetIdenfyTokenReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetIdenfyToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdenfyTokenRequest struct via the builder pattern


### Return type

[**GetIdenfyTokenReply**](GetIdenfyTokenReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthOptions

> GetOAuthOptionsReply GetOAuthOptions(ctx).Action(action).Metadata(metadata).Execute()

Get OAuth Providers

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
    action := "action_example" // string | Which authentication flow is being initiated (optional) (default to "signin")
    metadata := "metadata_example" // string | A small (limited to 400 characters) string of arbitrary metadata which will be encoded in the state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.GetOAuthOptions(context.Background()).Action(action).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetOAuthOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthOptions`: GetOAuthOptionsReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetOAuthOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Which authentication flow is being initiated | [default to &quot;signin&quot;]
 **metadata** | **string** | A small (limited to 400 characters) string of arbitrary metadata which will be encoded in the state | 

### Return type

[**GetOAuthOptionsReply**](GetOAuthOptionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserOrganizationInvitation

> GetUserOrganizationInvitationReply GetUserOrganizationInvitation(ctx, id).Execute()

Get User Organization Invitation

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
    id := "id_example" // string | The id of the organization invitation to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.GetUserOrganizationInvitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetUserOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserOrganizationInvitation`: GetUserOrganizationInvitationReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetUserOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the organization invitation to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserOrganizationInvitationReply**](GetUserOrganizationInvitationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettings

> GetUserSettingsReply GetUserSettings(ctx).Execute()



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
    resp, r, err := apiClient.ProfileApi.GetUserSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.GetUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettings`: GetUserSettingsReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.GetUserSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsRequest struct via the builder pattern


### Return type

[**GetUserSettingsReply**](GetUserSettingsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOrganizationInvitations

> ListUserOrganizationInvitationsReply ListUserOrganizationInvitations(ctx).Limit(limit).Offset(offset).Statuses(statuses).Execute()

List User Organization Invitations

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
    statuses := []string{"Statuses_example"} // []string | (Optional) Filter on organization invitation statuses (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.ListUserOrganizationInvitations(context.Background()).Limit(limit).Offset(offset).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ListUserOrganizationInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserOrganizationInvitations`: ListUserOrganizationInvitationsReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ListUserOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserOrganizationInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **statuses** | **[]string** | (Optional) Filter on organization invitation statuses | 

### Return type

[**ListUserOrganizationInvitationsReply**](ListUserOrganizationInvitationsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOrganizations

> ListUserOrganizationsReply ListUserOrganizations(ctx).Limit(limit).Offset(offset).Order(order).Search(search).Execute()

List User Organizations



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
    limit := "limit_example" // string | (Optional) Define pagination limit (optional)
    offset := "offset_example" // string | (Optional) Define pagination offset (optional)
    order := "order_example" // string | (Optional) Sorts the list in the ascending or the descending order (optional)
    search := "search_example" // string | (Optional) Fuzzy case-insensitive search based on organization name or organization id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.ListUserOrganizations(context.Background()).Limit(limit).Offset(offset).Order(order).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ListUserOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserOrganizations`: ListUserOrganizationsReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ListUserOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) Define pagination limit | 
 **offset** | **string** | (Optional) Define pagination offset | 
 **order** | **string** | (Optional) Sorts the list in the ascending or the descending order | 
 **search** | **string** | (Optional) Fuzzy case-insensitive search based on organization name or organization id | 

### Return type

[**ListUserOrganizationsReply**](ListUserOrganizationsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OAuthCallback

> OAuthCallbackReply OAuthCallback(ctx).Body(body).SeonFp(seonFp).Execute()

Authenticate using OAuth

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
    body := *openapiclient.NewOAuthCallbackRequest() // OAuthCallbackRequest | 
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.OAuthCallback(context.Background()).Body(body).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.OAuthCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OAuthCallback`: OAuthCallbackReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.OAuthCallback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOAuthCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OAuthCallbackRequest**](OAuthCallbackRequest.md) |  | 
 **seonFp** | **string** | Seon Fingerprint | 

### Return type

[**OAuthCallbackReply**](OAuthCallbackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailValidation

> map[string]interface{} ResendEmailValidation(ctx).Body(body).Execute()

Resend Email Verification

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.ResendEmailValidation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ResendEmailValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendEmailValidation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ResendEmailValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## ResetPassword

> map[string]interface{} ResetPassword(ctx).Body(body).Execute()

Reset Password

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
    body := *openapiclient.NewResetPasswordRequest() // ResetPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.ResetPassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

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


## Signup

> LoginReply Signup(ctx).Body(body).SeonFp(seonFp).Execute()

Signup

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
    body := *openapiclient.NewCreateAccountRequest("Email_example", "Password_example") // CreateAccountRequest | Create new account
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.Signup(context.Background()).Body(body).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Signup`: LoginReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAccountRequest**](CreateAccountRequest.md) | Create new account | 
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


## UpdatePassword

> LoginReply UpdatePassword(ctx).Body(body).SeonFp(seonFp).Execute()

Update Password

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
    body := *openapiclient.NewUpdatePasswordRequest() // UpdatePasswordRequest | 
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdatePassword(context.Background()).Body(body).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassword`: LoginReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) |  | 
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


## UpdateUser

> UserReply UpdateUser(ctx).User(user).UpdateMask(updateMask).Execute()

Update User

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
    user := *openapiclient.NewUpdateUserRequestUserUpdateBody() // UpdateUserRequestUserUpdateBody | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdateUser(context.Background()).User(user).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**UpdateUserRequestUserUpdateBody**](UpdateUserRequestUserUpdateBody.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UserReply**](UserReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser2

> UserReply UpdateUser2(ctx).User(user).UpdateMask(updateMask).Execute()

Update User

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
    user := *openapiclient.NewUpdateUserRequestUserUpdateBody() // UpdateUserRequestUserUpdateBody | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdateUser2(context.Background()).User(user).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser2`: UserReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateUser2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**UpdateUserRequestUserUpdateBody**](UpdateUserRequestUserUpdateBody.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UserReply**](UserReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettings

> UpdateUserSettingsReply UpdateUserSettings(ctx).Body(body).Execute()



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
    body := *openapiclient.NewUpdateUserSettingsRequest() // UpdateUserSettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdateUserSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserSettings`: UpdateUserSettingsReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateUserSettingsRequest**](UpdateUserSettingsRequest.md) |  | 

### Return type

[**UpdateUserSettingsReply**](UpdateUserSettingsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> LoginReply Validate(ctx, id).SeonFp(seonFp).Execute()

Validate

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
    seonFp := "seonFp_example" // string | Seon Fingerprint (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.Validate(context.Background(), id).SeonFp(seonFp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: LoginReply
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.Validate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

