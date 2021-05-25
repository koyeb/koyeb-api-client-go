# \FunctionsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelFunctionExecution**](FunctionsApi.md#CancelFunctionExecution) | **Post** /v1/stacks/{stack_id}/executions/{run_id}/cancel | Cancel the execution of a function
[**FetchFunctionExecutions**](FunctionsApi.md#FetchFunctionExecutions) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/executions | Get the executions of function runs
[**GetFunction**](FunctionsApi.md#GetFunction) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/functions/{function} | Get function
[**InvokeFunction**](FunctionsApi.md#InvokeFunction) | **Post** /v1/stacks/{stack_id}/revisions/{sha}/functions/{function}/invoke | Send an event to a specific function
[**ListFunctions**](FunctionsApi.md#ListFunctions) | **Get** /v1/stacks/{stack_id}/revisions/{sha}/functions | List functions for a revision
[**RetryFunctionExecution**](FunctionsApi.md#RetryFunctionExecution) | **Post** /v1/stacks/{stack_id}/executions/{run_id}/retrigger | Relaunch the execution of a function



## CancelFunctionExecution

> interface{} CancelFunctionExecution(ctx, stackId, runId).Execute()

Cancel the execution of a function

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
    runId := "runId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.CancelFunctionExecution(context.Background(), stackId, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.CancelFunctionExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelFunctionExecution`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.CancelFunctionExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFunctionExecutionRequest struct via the builder pattern


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


## FetchFunctionExecutions

> FetchFunctionExecutionsReply FetchFunctionExecutions(ctx, stackId, sha, function).Offset(offset).Limit(limit).Execute()

Get the executions of function runs

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
    sha := "sha_example" // string | The sha of the revision (either short of long form, `_latest` returns the latest deployed revision)
    function := "function_example" // string | The name of the function or `:all` to get the history for all the functions
    offset := "offset_example" // string |  (optional)
    limit := "limit_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.FetchFunctionExecutions(context.Background(), stackId, sha, function).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.FetchFunctionExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFunctionExecutions`: FetchFunctionExecutionsReply
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.FetchFunctionExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** | The sha of the revision (either short of long form, &#x60;_latest&#x60; returns the latest deployed revision) | 
**function** | **string** | The name of the function or &#x60;:all&#x60; to get the history for all the functions | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchFunctionExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **offset** | **string** |  | 
 **limit** | **string** |  | 

### Return type

[**FetchFunctionExecutionsReply**](FetchFunctionExecutionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> GetFunctionReply GetFunction(ctx, stackId, sha, function).Execute()

Get function

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
    sha := "sha_example" // string | The sha of the revision (either short of long form, `_latest` returns the latest deployed revision)
    function := "function_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.GetFunction(context.Background(), stackId, sha, function).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.GetFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunction`: GetFunctionReply
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.GetFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** | The sha of the revision (either short of long form, &#x60;_latest&#x60; returns the latest deployed revision) | 
**function** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFunctionReply**](GetFunctionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeFunction

> InvokeFunctionReply InvokeFunction(ctx, stackId, sha, function).Body(body).Execute()

Send an event to a specific function

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
    function := "function_example" // string | 
    body := interface{}(Object) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.InvokeFunction(context.Background(), stackId, sha, function).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.InvokeFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvokeFunction`: InvokeFunctionReply
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.InvokeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** |  | 
**function** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **interface{}** |  | 

### Return type

[**InvokeFunctionReply**](InvokeFunctionReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctions

> ListFunctionsReply ListFunctions(ctx, stackId, sha).Execute()

List functions for a revision

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
    sha := "sha_example" // string | The sha of the revision (either short of long form, `_latest` returns the latest deployed revision)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.ListFunctions(context.Background(), stackId, sha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.ListFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFunctions`: ListFunctionsReply
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.ListFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**sha** | **string** | The sha of the revision (either short of long form, &#x60;_latest&#x60; returns the latest deployed revision) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListFunctionsReply**](ListFunctionsReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryFunctionExecution

> interface{} RetryFunctionExecution(ctx, stackId, runId).Execute()

Relaunch the execution of a function

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
    runId := "runId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FunctionsApi.RetryFunctionExecution(context.Background(), stackId, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FunctionsApi.RetryFunctionExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetryFunctionExecution`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `FunctionsApi.RetryFunctionExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stackId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryFunctionExecutionRequest struct via the builder pattern


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

