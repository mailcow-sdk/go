# \AddressRewritingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBCCMap**](AddressRewritingAPI.md#CreateBCCMap) | **Post** /api/v1/add/bcc | Create BCC Map
[**CreateRecipientMap**](AddressRewritingAPI.md#CreateRecipientMap) | **Post** /api/v1/add/recipient_map | Create Recipient Map
[**DeleteBCCMap**](AddressRewritingAPI.md#DeleteBCCMap) | **Post** /api/v1/delete/bcc | Delete BCC Map
[**DeleteRecipientMap**](AddressRewritingAPI.md#DeleteRecipientMap) | **Post** /api/v1/delete/recipient_map | Delete Recipient Map
[**GetBCCMap**](AddressRewritingAPI.md#GetBCCMap) | **Get** /api/v1/get/bcc/{id} | Get BCC Map
[**GetRecipientMap**](AddressRewritingAPI.md#GetRecipientMap) | **Get** /api/v1/get/recipient_map/{id} | Get Recipient Map



## CreateBCCMap

> CreateAlias200Response CreateBCCMap(ctx).CreateBCCMapRequest(createBCCMapRequest).Execute()

Create BCC Map



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
	createBCCMapRequest := *openapiclient.NewCreateBCCMapRequest() // CreateBCCMapRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressRewritingAPI.CreateBCCMap(context.Background()).CreateBCCMapRequest(createBCCMapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.CreateBCCMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBCCMap`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressRewritingAPI.CreateBCCMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBCCMapRequest** | [**CreateBCCMapRequest**](CreateBCCMapRequest.md) |  | 

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


## CreateRecipientMap

> CreateAlias200Response CreateRecipientMap(ctx).CreateRecipientMapRequest(createRecipientMapRequest).Execute()

Create Recipient Map



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
	createRecipientMapRequest := *openapiclient.NewCreateRecipientMapRequest() // CreateRecipientMapRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressRewritingAPI.CreateRecipientMap(context.Background()).CreateRecipientMapRequest(createRecipientMapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.CreateRecipientMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecipientMap`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressRewritingAPI.CreateRecipientMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecipientMapRequest** | [**CreateRecipientMapRequest**](CreateRecipientMapRequest.md) |  | 

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


## DeleteBCCMap

> CreateAlias200Response DeleteBCCMap(ctx).DeleteBCCMapRequest(deleteBCCMapRequest).Execute()

Delete BCC Map



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
	deleteBCCMapRequest := *openapiclient.NewDeleteBCCMapRequest() // DeleteBCCMapRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressRewritingAPI.DeleteBCCMap(context.Background()).DeleteBCCMapRequest(deleteBCCMapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.DeleteBCCMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBCCMap`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressRewritingAPI.DeleteBCCMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBCCMapRequest** | [**DeleteBCCMapRequest**](DeleteBCCMapRequest.md) |  | 

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


## DeleteRecipientMap

> CreateAlias200Response DeleteRecipientMap(ctx).DeleteRecipientMapRequest(deleteRecipientMapRequest).Execute()

Delete Recipient Map



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
	deleteRecipientMapRequest := *openapiclient.NewDeleteRecipientMapRequest() // DeleteRecipientMapRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressRewritingAPI.DeleteRecipientMap(context.Background()).DeleteRecipientMapRequest(deleteRecipientMapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.DeleteRecipientMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecipientMap`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressRewritingAPI.DeleteRecipientMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRecipientMapRequest** | [**DeleteRecipientMapRequest**](DeleteRecipientMapRequest.md) |  | 

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


## GetBCCMap

> GetBCCMap(ctx, id).XAPIKey(xAPIKey).Execute()

Get BCC Map



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
	id := "all" // string | id of entry you want to get
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AddressRewritingAPI.GetBCCMap(context.Background(), id).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.GetBCCMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBCCMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

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


## GetRecipientMap

> GetRecipientMap(ctx, id).XAPIKey(xAPIKey).Execute()

Get Recipient Map



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
	id := "all" // string | id of entry you want to get
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AddressRewritingAPI.GetRecipientMap(context.Background(), id).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressRewritingAPI.GetRecipientMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipientMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

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

