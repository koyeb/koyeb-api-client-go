# \StackApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBuild**](StackApi.md#CancelBuild) | **Post** /v1/stacks/{stack_id}/revisions/{sha}/build/cancel | Cancel the execution of a build
[**CreateStack**](StackApi.md#CreateStack) | **Post** /v1/stacks | Create a stack
[**CreateStackEvent**](StackApi.md#CreateStackEvent) | **Post** /v1/stacks/{stack_id}/events | Send event to stack
[**CreateStackRevision**](StackApi.md#CreateStackRevision) | **Post** /v1/stacks/{stack_id}/revisions | Create stack revisions
[**DeleteStack**](StackApi.md#DeleteStack) | **Delete** /v1/stacks/{id} | Delete an existing stack
[**GetStack**](StackApi.md#GetStack) | **Get** /v1/stacks/{id} | Fetch a stack
[**GetStackActivities**](StackApi.md#GetStackActivities) | **Get** /v1/stacks/{id}/activities | View stack activities
[**GetStackRevision**](StackApi.md#GetStackRevision) | **Get** /v1/stacks/{stack_id}/revisions/{sha} | Fetch a stack revision
[**GetStackRevisionActivities**](StackApi.md#GetStackRevisionActivities) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/activities | Get a stack revision activities
[**ListStackRevisions**](StackApi.md#ListStackRevisions) | **Get** /v1/stacks/{stack_id}/revisions | List stack revisions
[**ListStacks**](StackApi.md#ListStacks) | **Get** /v1/stacks | List stacks
[**RetryBuild**](StackApi.md#RetryBuild) | **Post** /v1/stacks/{stack_id}/revisions/{sha}/build/retrigger | Relaunch the execution of a build
[**UpdateStack**](StackApi.md#UpdateStack) | **Put** /v1/stacks/{id} | Update an existing stack
[**UpdateStack2**](StackApi.md#UpdateStack2) | **Patch** /v1/stacks/{id} | Update an existing stack



## CancelBuild

> interface{} CancelBuild(ctx, stackId, sha).Execute()

Cancel the execution of a build

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
    stackId := "stackId_example" // string | 
    sha := "sha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.CancelBuild(context.Background(), stackId, sha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.CancelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelBuild`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StackApi.CancelBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBuildRequest struct via the builder pattern


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


## CreateStack

> CreateStackReply CreateStack(ctx).Body(body).Execute()

Create a stack

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
    body := *openapiclient.NewStackUpsert() // StackUpsert | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.CreateStack(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.CreateStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStack`: CreateStackReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.CreateStack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StackUpsert**](StackUpsert.md) |  | 

### Return type

[**CreateStackReply**](CreateStackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackEvent

> CreateStackEventReply CreateStackEvent(ctx, stackId).Body(body).Execute()

Send event to stack

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
    stackId := "stackId_example" // string | 
    body := interface{}(Object) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.CreateStackEvent(context.Background(), stackId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.CreateStackEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStackEvent`: CreateStackEventReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.CreateStackEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**CreateStackEventReply**](CreateStackEventReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackRevision

> GetStackRevisionReply CreateStackRevision(ctx, stackId).Body(body).Execute()

Create stack revisions

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
    stackId := "stackId_example" // string | 
    body := *openapiclient.NewCreateStackRevisionRequest() // CreateStackRevisionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.CreateStackRevision(context.Background(), stackId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.CreateStackRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStackRevision`: GetStackRevisionReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.CreateStackRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateStackRevisionRequest**](CreateStackRevisionRequest.md) |  | 

### Return type

[**GetStackRevisionReply**](GetStackRevisionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStack

> interface{} DeleteStack(ctx, id).Execute()

Delete an existing stack

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
    resp, r, err := api_client.StackApi.DeleteStack(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.DeleteStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStack`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StackApi.DeleteStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackRequest struct via the builder pattern


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


## GetStack

> GetStackReply GetStack(ctx, id).Execute()

Fetch a stack

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
    resp, r, err := api_client.StackApi.GetStack(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.GetStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStack`: GetStackReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.GetStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStackReply**](GetStackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackActivities

> ActivityList GetStackActivities(ctx, id).Limit(limit).Offset(offset).Execute()

View stack activities

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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.GetStackActivities(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.GetStackActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackActivities`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `StackApi.GetStackActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackRevision

> GetStackRevisionReply GetStackRevision(ctx, stackId, sha).Execute()

Fetch a stack revision

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
    stackId := "stackId_example" // string | 
    sha := "sha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.GetStackRevision(context.Background(), stackId, sha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.GetStackRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackRevision`: GetStackRevisionReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.GetStackRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStackRevisionReply**](GetStackRevisionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackRevisionActivities

> ActivityList GetStackRevisionActivities(ctx, stackId, sha).Limit(limit).Offset(offset).Execute()

Get a stack revision activities

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
    stackId := "stackId_example" // string | 
    sha := "sha_example" // string | 
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.GetStackRevisionActivities(context.Background(), stackId, sha).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.GetStackRevisionActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStackRevisionActivities`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `StackApi.GetStackRevisionActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRevisionActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ActivityList**](ActivityList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStackRevisions

> ListStackRevisionsReply ListStackRevisions(ctx, stackId).Limit(limit).Offset(offset).Execute()

List stack revisions

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
    stackId := "stackId_example" // string | 
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.ListStackRevisions(context.Background(), stackId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.ListStackRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStackRevisions`: ListStackRevisionsReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.ListStackRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStackRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListStackRevisionsReply**](ListStackRevisionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStacks

> ListStacksReply ListStacks(ctx).Name(name).RepositoryType(repositoryType).RepositoryName(repositoryName).RepositoryBranch(repositoryBranch).Limit(limit).Offset(offset).Execute()

List stacks

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
    repositoryType := "repositoryType_example" // string | Where the repo lives. (optional) (default to "GITHUB")
    repositoryName := "repositoryName_example" // string | The url to find the repo (.e.g: koyeb/gateway). (optional)
    repositoryBranch := "repositoryBranch_example" // string | The branch to track changes on. (optional)
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.ListStacks(context.Background()).Name(name).RepositoryType(repositoryType).RepositoryName(repositoryName).RepositoryBranch(repositoryBranch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.ListStacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStacks`: ListStacksReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.ListStacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **repositoryType** | **string** | Where the repo lives. | [default to &quot;GITHUB&quot;]
 **repositoryName** | **string** | The url to find the repo (.e.g: koyeb/gateway). | 
 **repositoryBranch** | **string** | The branch to track changes on. | 
 **limit** | **string** |  | 
 **offset** | **string** |  | 

### Return type

[**ListStacksReply**](ListStacksReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryBuild

> interface{} RetryBuild(ctx, stackId, sha).Execute()

Relaunch the execution of a build

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
    stackId := "stackId_example" // string | 
    sha := "sha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.RetryBuild(context.Background(), stackId, sha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.RetryBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetryBuild`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StackApi.RetryBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryBuildRequest struct via the builder pattern


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


## UpdateStack

> UpdateStackReply UpdateStack(ctx, id).Body(body).UpdateMask(updateMask).Execute()

Update an existing stack

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
    body := *openapiclient.NewStackUpsert() // StackUpsert | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.UpdateStack(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.UpdateStack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStack`: UpdateStackReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.UpdateStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StackUpsert**](StackUpsert.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateStackReply**](UpdateStackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStack2

> UpdateStackReply UpdateStack2(ctx, id).Body(body).UpdateMask(updateMask).Execute()

Update an existing stack

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
    body := *openapiclient.NewStackUpsert() // StackUpsert | 
    updateMask := "updateMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StackApi.UpdateStack2(context.Background(), id).Body(body).UpdateMask(updateMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StackApi.UpdateStack2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStack2`: UpdateStackReply
    fmt.Fprintf(os.Stdout, "Response from `StackApi.UpdateStack2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStack2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StackUpsert**](StackUpsert.md) |  | 
 **updateMask** | **string** |  | 

### Return type

[**UpdateStackReply**](UpdateStackReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

