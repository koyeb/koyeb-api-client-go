# \CouponsApi

All URIs are relative to *https://app.koyeb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedeemCoupon**](CouponsApi.md#RedeemCoupon) | **Post** /v1/coupons | This API allows to redeem a coupon. Pass the code you received in the body.



## RedeemCoupon

> map[string]interface{} RedeemCoupon(ctx).Body(body).Execute()

This API allows to redeem a coupon. Pass the code you received in the body.

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
    body := *openapiclient.NewRedeemCouponRequest() // RedeemCouponRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsApi.RedeemCoupon(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsApi.RedeemCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeemCoupon`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CouponsApi.RedeemCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedeemCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RedeemCouponRequest**](RedeemCouponRequest.md) |  | 

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

