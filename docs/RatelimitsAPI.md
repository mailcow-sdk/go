# \RatelimitsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditDomainRatelimits**](RatelimitsAPI.md#EditDomainRatelimits) | **Post** /api/v1/edit/rl-domain/ | Edit domain ratelimits
[**EditMailboxRatelimits**](RatelimitsAPI.md#EditMailboxRatelimits) | **Post** /api/v1/edit/rl-mbox/ | Edit mailbox ratelimits
[**GetDomainRatelimits**](RatelimitsAPI.md#GetDomainRatelimits) | **Get** /api/v1/get/rl-domain/{domain} | Get domain ratelimits
[**GetMailboxRatelimits**](RatelimitsAPI.md#GetMailboxRatelimits) | **Get** /api/v1/get/rl-mbox/{mailbox} | Get mailbox ratelimits



## EditDomainRatelimits

> CreateAlias200Response EditDomainRatelimits(ctx).EditDomainRatelimitsRequest(editDomainRatelimitsRequest).Execute()

Edit domain ratelimits



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
	editDomainRatelimitsRequest := *openapiclient.NewEditDomainRatelimitsRequest() // EditDomainRatelimitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatelimitsAPI.EditDomainRatelimits(context.Background()).EditDomainRatelimitsRequest(editDomainRatelimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatelimitsAPI.EditDomainRatelimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDomainRatelimits`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `RatelimitsAPI.EditDomainRatelimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDomainRatelimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDomainRatelimitsRequest** | [**EditDomainRatelimitsRequest**](EditDomainRatelimitsRequest.md) |  | 

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


## EditMailboxRatelimits

> CreateAlias200Response EditMailboxRatelimits(ctx).EditMailboxRatelimitsRequest(editMailboxRatelimitsRequest).Execute()

Edit mailbox ratelimits



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
	editMailboxRatelimitsRequest := *openapiclient.NewEditMailboxRatelimitsRequest() // EditMailboxRatelimitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatelimitsAPI.EditMailboxRatelimits(context.Background()).EditMailboxRatelimitsRequest(editMailboxRatelimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatelimitsAPI.EditMailboxRatelimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditMailboxRatelimits`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `RatelimitsAPI.EditMailboxRatelimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditMailboxRatelimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editMailboxRatelimitsRequest** | [**EditMailboxRatelimitsRequest**](EditMailboxRatelimitsRequest.md) |  | 

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


## GetDomainRatelimits

> GetDomainRatelimits(ctx, domain).XAPIKey(xAPIKey).Execute()

Get domain ratelimits



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
	domain := "domain_example" // string | name of domain or all
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RatelimitsAPI.GetDomainRatelimits(context.Background(), domain).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatelimitsAPI.GetDomainRatelimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | name of domain or all | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRatelimitsRequest struct via the builder pattern


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


## GetMailboxRatelimits

> GetMailboxRatelimits(ctx, mailbox).XAPIKey(xAPIKey).Execute()

Get mailbox ratelimits



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
	mailbox := "mailbox_example" // string | name of mailbox or all
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RatelimitsAPI.GetMailboxRatelimits(context.Background(), mailbox).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatelimitsAPI.GetMailboxRatelimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** | name of mailbox or all | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxRatelimitsRequest struct via the builder pattern


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

