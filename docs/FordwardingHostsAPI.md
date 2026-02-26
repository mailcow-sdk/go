# \FordwardingHostsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddForwardHost**](FordwardingHostsAPI.md#AddForwardHost) | **Post** /api/v1/add/fwdhost | Add Forward Host
[**DeleteForwardHost**](FordwardingHostsAPI.md#DeleteForwardHost) | **Post** /api/v1/delete/fwdhost | Delete Forward Host
[**GetForwardingHosts**](FordwardingHostsAPI.md#GetForwardingHosts) | **Get** /api/v1/get/fwdhost/all | Get Forwarding Hosts



## AddForwardHost

> CreateAlias200Response AddForwardHost(ctx).AddForwardHostRequest(addForwardHostRequest).Execute()

Add Forward Host



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	addForwardHostRequest := *openapiclient.NewAddForwardHostRequest() // AddForwardHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FordwardingHostsAPI.AddForwardHost(context.Background()).AddForwardHostRequest(addForwardHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsAPI.AddForwardHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddForwardHost`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `FordwardingHostsAPI.AddForwardHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddForwardHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addForwardHostRequest** | [**AddForwardHostRequest**](AddForwardHostRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForwardHost

> CreateAlias200Response DeleteForwardHost(ctx).DeleteForwardHostRequest(deleteForwardHostRequest).Execute()

Delete Forward Host



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deleteForwardHostRequest := *openapiclient.NewDeleteForwardHostRequest() // DeleteForwardHostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FordwardingHostsAPI.DeleteForwardHost(context.Background()).DeleteForwardHostRequest(deleteForwardHostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsAPI.DeleteForwardHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteForwardHost`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `FordwardingHostsAPI.DeleteForwardHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForwardHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteForwardHostRequest** | [**DeleteForwardHostRequest**](DeleteForwardHostRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForwardingHosts

> GetForwardingHosts(ctx).Execute()

Get Forwarding Hosts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FordwardingHostsAPI.GetForwardingHosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FordwardingHostsAPI.GetForwardingHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetForwardingHostsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

