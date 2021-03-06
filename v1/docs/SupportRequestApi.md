# \SupportRequestApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestSuppert**](SupportRequestApi.md#RequestSuppert) | **Post** /support-requests | Create a support ticket



## RequestSuppert

> RequestSuppert(ctx).SupportRequest(supportRequest).Execute()

Create a support ticket



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
    supportRequest := *openapiclient.NewSupportRequestInput("Subject_example", "Message_example") // SupportRequestInput | Support Request to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportRequestApi.RequestSuppert(context.Background()).SupportRequest(supportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestApi.RequestSuppert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSuppertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportRequest** | [**SupportRequestInput**](SupportRequestInput.md) | Support Request to create | 

### Return type

 (empty response body)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

