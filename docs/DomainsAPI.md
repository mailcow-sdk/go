# \DomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainsAPI.md#CreateDomain) | **Post** /api/v1/add/domain | Create domain
[**DeleteDomain**](DomainsAPI.md#DeleteDomain) | **Post** /api/v1/delete/domain | Delete domain
[**DeleteDomainTags**](DomainsAPI.md#DeleteDomainTags) | **Post** /api/v1/delete/domain/tag/{domain} | Delete domain tags
[**GetDomains**](DomainsAPI.md#GetDomains) | **Get** /api/v1/get/domain/{id} | Get domains
[**UpdateDomain**](DomainsAPI.md#UpdateDomain) | **Post** /api/v1/edit/domain | Update domain
[**UpdateDomainWideFooter**](DomainsAPI.md#UpdateDomainWideFooter) | **Post** /api/v1/edit/domain/footer | Update domain wide footer



## CreateDomain

> []CreateAlias200Response CreateDomain(ctx).CreateDomainRequest(createDomainRequest).Execute()

Create domain



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
	createDomainRequest := *openapiclient.NewCreateDomainRequest() // CreateDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.CreateDomain(context.Background()).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: []CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainRequest** | [**CreateDomainRequest**](CreateDomainRequest.md) |  | 

### Return type

[**[]CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> []CreateAlias200Response DeleteDomain(ctx).DeleteDomainRequest(deleteDomainRequest).Execute()

Delete domain



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
	deleteDomainRequest := *openapiclient.NewDeleteDomainRequest() // DeleteDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteDomain(context.Background()).DeleteDomainRequest(deleteDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomain`: []CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDomainRequest** | [**DeleteDomainRequest**](DeleteDomainRequest.md) |  | 

### Return type

[**[]CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomainTags

> CreateAlias200Response DeleteDomainTags(ctx, domain).DeleteDomainTagsRequest(deleteDomainTagsRequest).Execute()

Delete domain tags



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
	domain := "domain.tld" // string | name of domain
	deleteDomainTagsRequest := *openapiclient.NewDeleteDomainTagsRequest() // DeleteDomainTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteDomainTags(context.Background(), domain).DeleteDomainTagsRequest(deleteDomainTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomainTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomainTags`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteDomainTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | name of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteDomainTagsRequest** | [**DeleteDomainTagsRequest**](DeleteDomainTagsRequest.md) |  | 

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


## GetDomains

> GetDomains(ctx, id).Tags(tags).XAPIKey(xAPIKey).Execute()

Get domains



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
	tags := "tag1,tag2" // string | comma seperated list of tags to filter by (optional)
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.GetDomains(context.Background(), id).Tags(tags).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **string** | comma seperated list of tags to filter by | 
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


## UpdateDomain

> []CreateAlias200Response UpdateDomain(ctx).UpdateDomainRequest(updateDomainRequest).Execute()

Update domain



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
	updateDomainRequest := *openapiclient.NewUpdateDomainRequest() // UpdateDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.UpdateDomain(context.Background()).UpdateDomainRequest(updateDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomain`: []CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDomainRequest** | [**UpdateDomainRequest**](UpdateDomainRequest.md) |  | 

### Return type

[**[]CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainWideFooter

> CreateAlias200Response UpdateDomainWideFooter(ctx).UpdateDomainWideFooterRequest(updateDomainWideFooterRequest).Execute()

Update domain wide footer



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
	updateDomainWideFooterRequest := *openapiclient.NewUpdateDomainWideFooterRequest() // UpdateDomainWideFooterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.UpdateDomainWideFooter(context.Background()).UpdateDomainWideFooterRequest(updateDomainWideFooterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainWideFooter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomainWideFooter`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainWideFooter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainWideFooterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDomainWideFooterRequest** | [**UpdateDomainWideFooterRequest**](UpdateDomainWideFooterRequest.md) |  | 

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

