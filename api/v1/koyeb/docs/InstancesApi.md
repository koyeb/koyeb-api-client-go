# \InstancesApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecCommand**](InstancesApi.md#ExecCommand) | **Get** /v1/instances/exec | Exec Command
[**ListServiceInstances**](InstancesApi.md#ListServiceInstances) | **Get** /v1/apps/{app_id_or_name}/services/{service_id_or_name}/instances | List Instances for a service



## ExecCommand

> StreamResultOfExecCommandReply ExecCommand(ctx).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).Execute()

Exec Command

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
    id := "id_example" // string |  (optional)
    bodyCommand := []string{"Inner_example"} // []string |  (optional)
    bodyTtySizeHeight := int32(56) // int32 |  (optional)
    bodyTtySizeWidth := int32(56) // int32 |  (optional)
    bodyStdinData := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ExecCommand(context.Background()).Id(id).BodyCommand(bodyCommand).BodyTtySizeHeight(bodyTtySizeHeight).BodyTtySizeWidth(bodyTtySizeWidth).BodyStdinData(bodyStdinData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ExecCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecCommand`: StreamResultOfExecCommandReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ExecCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **bodyCommand** | **[]string** |  | 
 **bodyTtySizeHeight** | **int32** |  | 
 **bodyTtySizeWidth** | **int32** |  | 
 **bodyStdinData** | **string** |  | 

### Return type

[**StreamResultOfExecCommandReply**](StreamResultOfExecCommandReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceInstances

> ListServiceInstancesReply ListServiceInstances(ctx, appIdOrName, serviceIdOrName).Limit(limit).Offset(offset).Execute()

List Instances for a service

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
    appIdOrName := "appIdOrName_example" // string | Name or id of the application
    serviceIdOrName := "serviceIdOrName_example" // string | Name or id of the service
    limit := "limit_example" // string | (Optional) The number of items to return. (optional)
    offset := "offset_example" // string | (Optional) The offset in the list of item to return. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ListServiceInstances(context.Background(), appIdOrName, serviceIdOrName).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ListServiceInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceInstances`: ListServiceInstancesReply
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ListServiceInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appIdOrName** | **string** | Name or id of the application | 
**serviceIdOrName** | **string** | Name or id of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **string** | (Optional) The number of items to return. | 
 **offset** | **string** | (Optional) The offset in the list of item to return. | 

### Return type

[**ListServiceInstancesReply**](ListServiceInstancesReply.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

