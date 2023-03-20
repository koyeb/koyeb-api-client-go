# \OrganizationInvitationsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInvitation**](OrganizationInvitationsApi.md#CreateOrganizationInvitation) | **Post** /v1/organization_invitations | 
[**DeleteOrganizationInvitation**](OrganizationInvitationsApi.md#DeleteOrganizationInvitation) | **Delete** /v1/organization_invitations/{id} | 
[**ListOrganizationInvitations**](OrganizationInvitationsApi.md#ListOrganizationInvitations) | **Get** /v1/organization_invitations | 
[**ResendOrganizationInvitation**](OrganizationInvitationsApi.md#ResendOrganizationInvitation) | **Post** /v1/organization_invitations/{id}/resend | 



## CreateOrganizationInvitation

> CreateOrganizationInvitationReply CreateOrganizationInvitation(ctx).Body(body).Execute()



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
    body := *openapiclient.NewCreateOrganizationInvitationRequest() // CreateOrganizationInvitationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationInvitationsApi.CreateOrganizationInvitation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsApi.CreateOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInvitation`: CreateOrganizationInvitationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsApi.CreateOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrganizationInvitationRequest**](CreateOrganizationInvitationRequest.md) |  | 

### Return type

[**CreateOrganizationInvitationReply**](CreateOrganizationInvitationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationInvitation

> map[string]interface{} DeleteOrganizationInvitation(ctx, id).Execute()



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
    id := "id_example" // string | The id of the organization invitation to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationInvitationsApi.DeleteOrganizationInvitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsApi.DeleteOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganizationInvitation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsApi.DeleteOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the organization invitation to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationInvitationRequest struct via the builder pattern


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


## ListOrganizationInvitations

> ListOrganizationInvitationsReply ListOrganizationInvitations(ctx).Limit(limit).Offset(offset).Statuses(statuses).UserId(userId).Execute()



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
    userId := "userId_example" // string | (Optional) Filter on invitee ID. Will match both invitations sent to that user_id and invitations sent to the email of that user_id. The only valid value is the requester's user_id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationInvitationsApi.ListOrganizationInvitations(context.Background()).Limit(limit).Offset(offset).Statuses(statuses).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsApi.ListOrganizationInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationInvitations`: ListOrganizationInvitationsReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsApi.ListOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **statuses** | **[]string** | (Optional) Filter on organization invitation statuses | 
 **userId** | **string** | (Optional) Filter on invitee ID. Will match both invitations sent to that user_id and invitations sent to the email of that user_id. The only valid value is the requester&#39;s user_id | 

### Return type

[**ListOrganizationInvitationsReply**](ListOrganizationInvitationsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendOrganizationInvitation

> ResendOrganizationInvitationReply ResendOrganizationInvitation(ctx, id).Body(body).Execute()



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
    id := "id_example" // string | The id of the organization invitation to resend
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationInvitationsApi.ResendOrganizationInvitation(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsApi.ResendOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendOrganizationInvitation`: ResendOrganizationInvitationReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsApi.ResendOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the organization invitation to resend | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**ResendOrganizationInvitationReply**](ResendOrganizationInvitationReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

