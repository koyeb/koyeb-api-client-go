# \RepositoriesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBranches**](RepositoriesApi.md#ListBranches) | **Get** /v1/git/branches | 
[**ListRepositories**](RepositoriesApi.md#ListRepositories) | **Get** /v1/git/repositories | 
[**ResyncOrganization**](RepositoriesApi.md#ResyncOrganization) | **Post** /v1/git/sync/organization/{organization_id} | 



## ListBranches

> KgitproxyListBranchesReply ListBranches(ctx).RepositoryId(repositoryId).Name(name).Limit(limit).Offset(offset).Execute()



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
    repositoryId := "repositoryId_example" // string | (Optional) Filter on one repository. (optional)
    name := "name_example" // string | (Optional) Filter on branch name using a fuzzy search. Repository filter is required to enable this filter. (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.ListBranches(context.Background()).RepositoryId(repositoryId).Name(name).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ListBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBranches`: KgitproxyListBranchesReply
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ListBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryId** | **string** | (Optional) Filter on one repository. | 
 **name** | **string** | (Optional) Filter on branch name using a fuzzy search. Repository filter is required to enable this filter. | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 

### Return type

[**KgitproxyListBranchesReply**](KgitproxyListBranchesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepositories

> KgitproxyListRepositoriesReply ListRepositories(ctx).Name(name).NameSearchOp(nameSearchOp).Limit(limit).Offset(offset).Execute()



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
    name := "name_example" // string | (Optional) Filter on repository name using a fuzzy search. (optional)
    nameSearchOp := "nameSearchOp_example" // string | (Optional) Define search operation for repository name. Accept either \"fuzzy\" or \"equality\", use \"fuzzy\" by default. (optional)
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.ListRepositories(context.Background()).Name(name).NameSearchOp(nameSearchOp).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ListRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRepositories`: KgitproxyListRepositoriesReply
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ListRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (Optional) Filter on repository name using a fuzzy search. | 
 **nameSearchOp** | **string** | (Optional) Define search operation for repository name. Accept either \&quot;fuzzy\&quot; or \&quot;equality\&quot;, use \&quot;fuzzy\&quot; by default. | 
 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 

### Return type

[**KgitproxyListRepositoriesReply**](KgitproxyListRepositoriesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResyncOrganization

> map[string]interface{} ResyncOrganization(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.ResyncOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.ResyncOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResyncOrganization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.ResyncOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResyncOrganizationRequest struct via the builder pattern


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

