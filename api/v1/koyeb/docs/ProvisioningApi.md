# \ProvisioningApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStageAttempt**](ProvisioningApi.md#CreateStageAttempt) | **Post** /v1/provisioning/{deployment_id}/status/{stage}/{attempt} | Create an attempt for a stage
[**DeclareStageProgress**](ProvisioningApi.md#DeclareStageProgress) | **Patch** /v1/provisioning/{deployment_id}/status/{stage}/{attempt} | Declare stage progress
[**DeclareStepProgress**](ProvisioningApi.md#DeclareStepProgress) | **Patch** /v1/provisioning/{deployment_id}/status/{stage}/{attempt}/{step} | Declare step progress



## CreateStageAttempt

> map[string]interface{} CreateStageAttempt(ctx, deploymentId, stage, attempt).Body(body).Execute()

Create an attempt for a stage

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
    deploymentId := "deploymentId_example" // string | 
    stage := "stage_example" // string | 
    attempt := "attempt_example" // string | 
    body := *openapiclient.NewCreateStageAttemptRequest() // CreateStageAttemptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.CreateStageAttempt(context.Background(), deploymentId, stage, attempt).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.CreateStageAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStageAttempt`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.CreateStageAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**stage** | **string** |  | 
**attempt** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStageAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**CreateStageAttemptRequest**](CreateStageAttemptRequest.md) |  | 

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


## DeclareStageProgress

> map[string]interface{} DeclareStageProgress(ctx, deploymentId, stage, attempt).Body(body).Execute()

Declare stage progress

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
    deploymentId := "deploymentId_example" // string | 
    stage := "stage_example" // string | 
    attempt := "attempt_example" // string | 
    body := *openapiclient.NewDeclareStageProgressRequest() // DeclareStageProgressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.DeclareStageProgress(context.Background(), deploymentId, stage, attempt).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.DeclareStageProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareStageProgress`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.DeclareStageProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**stage** | **string** |  | 
**attempt** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclareStageProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**DeclareStageProgressRequest**](DeclareStageProgressRequest.md) |  | 

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


## DeclareStepProgress

> map[string]interface{} DeclareStepProgress(ctx, deploymentId, stage, attempt, step).Body(body).Execute()

Declare step progress

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
    deploymentId := "deploymentId_example" // string | 
    stage := "stage_example" // string | 
    attempt := "attempt_example" // string | 
    step := "step_example" // string | 
    body := *openapiclient.NewDeclareStepProgressRequest() // DeclareStepProgressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.DeclareStepProgress(context.Background(), deploymentId, stage, attempt, step).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.DeclareStepProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareStepProgress`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.DeclareStepProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 
**stage** | **string** |  | 
**attempt** | **string** |  | 
**step** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclareStepProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**DeclareStepProgressRequest**](DeclareStepProgressRequest.md) |  | 

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

