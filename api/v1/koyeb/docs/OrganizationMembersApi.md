# \OrganizationMembersApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOrganizationMembers**](OrganizationMembersApi.md#ListOrganizationMembers) | **Get** /v1/organization_members | List organization members
[**RemoveOrganizationMember**](OrganizationMembersApi.md#RemoveOrganizationMember) | **Delete** /v1/organization_members/{id} | Remove an organization member



## ListOrganizationMembers

> ListOrganizationMembersReply ListOrganizationMembers(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).UserId(userId).OrganizationStatuses(organizationStatuses).Execute()

List organization members

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
    organizationId := "organizationId_example" // string | (Optional) Filter for an organization (optional)
    userId := "userId_example" // string | (Optional) Filter for an user (optional)
    organizationStatuses := []string{"OrganizationStatuses_example"} // []string | (Optional) Filter for organization statuses (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMembersApi.ListOrganizationMembers(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).UserId(userId).OrganizationStatuses(organizationStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMembersApi.ListOrganizationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationMembers`: ListOrganizationMembersReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMembersApi.ListOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | (Optional) The number of items to return | 
 **offset** | **string** | (Optional) The offset in the list of item to return | 
 **organizationId** | **string** | (Optional) Filter for an organization | 
 **userId** | **string** | (Optional) Filter for an user | 
 **organizationStatuses** | **[]string** | (Optional) Filter for organization statuses | 

### Return type

[**ListOrganizationMembersReply**](ListOrganizationMembersReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationMember

> RemoveOrganizationMemberReply RemoveOrganizationMember(ctx, id).Execute()

Remove an organization member

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
    resp, r, err := apiClient.OrganizationMembersApi.RemoveOrganizationMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMembersApi.RemoveOrganizationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationMember`: RemoveOrganizationMemberReply
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMembersApi.RemoveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoveOrganizationMemberReply**](RemoveOrganizationMemberReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

