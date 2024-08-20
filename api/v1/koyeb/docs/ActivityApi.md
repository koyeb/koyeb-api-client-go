# \ActivityApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountActivities**](ActivityApi.md#GetAccountActivities) | **Get** /v1/account/activities | 
[**ListActivities**](ActivityApi.md#ListActivities) | **Get** /v1/activities | 
[**ListNotifications**](ActivityApi.md#ListNotifications) | **Get** /v1/notifications | 



## GetAccountActivities

> ActivityList GetAccountActivities(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetAccountActivities(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetAccountActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountActivities`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetAccountActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountActivitiesRequest struct via the builder pattern


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


## ListActivities

> ActivityList ListActivities(ctx).Limit(limit).Offset(offset).Types(types).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)
    types := []string{"Inner_example"} // []string | (Optional) Filter on object type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListActivities(context.Background()).Limit(limit).Offset(offset).Types(types).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActivities`: ActivityList
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | 
 **offset** | **string** |  | 
 **types** | **[]string** | (Optional) Filter on object type | 

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


## ListNotifications

> NotificationList ListNotifications(ctx).Limit(limit).Offset(offset).MarkRead(markRead).MarkSeen(markSeen).Unread(unread).Unseen(unseen).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)
    markRead := "markRead_example" // string |  (optional)
    markSeen := "markSeen_example" // string |  (optional)
    unread := "unread_example" // string |  (optional)
    unseen := "unseen_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListNotifications(context.Background()).Limit(limit).Offset(offset).MarkRead(markRead).MarkSeen(markSeen).Unread(unread).Unseen(unseen).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotifications`: NotificationList
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | 
 **offset** | **string** |  | 
 **markRead** | **string** |  | 
 **markSeen** | **string** |  | 
 **unread** | **string** |  | 
 **unseen** | **string** |  | 

### Return type

[**NotificationList**](NotificationList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

