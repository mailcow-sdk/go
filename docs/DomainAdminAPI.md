# \DomainAdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainAdminUser**](DomainAdminAPI.md#CreateDomainAdminUser) | **Post** /api/v1/add/domain-admin | Create Domain Admin user
[**DeleteDomainAdmin**](DomainAdminAPI.md#DeleteDomainAdmin) | **Post** /api/v1/delete/domain-admin | Delete Domain Admin
[**EditDomainAdminACL**](DomainAdminAPI.md#EditDomainAdminACL) | **Post** /api/v1/edit/da-acl | Edit Domain Admin ACL
[**EditDomainAdminUser**](DomainAdminAPI.md#EditDomainAdminUser) | **Post** /api/v1/edit/domain-admin | Edit Domain Admin user
[**GetDomainAdmins**](DomainAdminAPI.md#GetDomainAdmins) | **Get** /api/v1/get/domain-admin/all | Get Domain Admins



## CreateDomainAdminUser

> CreateAlias200Response CreateDomainAdminUser(ctx).CreateDomainAdminUserRequest(createDomainAdminUserRequest).Execute()

Create Domain Admin user



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
	createDomainAdminUserRequest := *openapiclient.NewCreateDomainAdminUserRequest() // CreateDomainAdminUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAdminAPI.CreateDomainAdminUser(context.Background()).CreateDomainAdminUserRequest(createDomainAdminUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminAPI.CreateDomainAdminUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomainAdminUser`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainAdminAPI.CreateDomainAdminUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainAdminUserRequest** | [**CreateDomainAdminUserRequest**](CreateDomainAdminUserRequest.md) |  | 

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


## DeleteDomainAdmin

> CreateAlias200Response DeleteDomainAdmin(ctx).DeleteDomainAdminRequest(deleteDomainAdminRequest).Execute()

Delete Domain Admin



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
	deleteDomainAdminRequest := *openapiclient.NewDeleteDomainAdminRequest() // DeleteDomainAdminRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAdminAPI.DeleteDomainAdmin(context.Background()).DeleteDomainAdminRequest(deleteDomainAdminRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminAPI.DeleteDomainAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomainAdmin`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainAdminAPI.DeleteDomainAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDomainAdminRequest** | [**DeleteDomainAdminRequest**](DeleteDomainAdminRequest.md) |  | 

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


## EditDomainAdminACL

> CreateAlias200Response EditDomainAdminACL(ctx).EditDomainAdminACLRequest(editDomainAdminACLRequest).Execute()

Edit Domain Admin ACL



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
	editDomainAdminACLRequest := *openapiclient.NewEditDomainAdminACLRequest() // EditDomainAdminACLRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAdminAPI.EditDomainAdminACL(context.Background()).EditDomainAdminACLRequest(editDomainAdminACLRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminAPI.EditDomainAdminACL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDomainAdminACL`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainAdminAPI.EditDomainAdminACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDomainAdminACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDomainAdminACLRequest** | [**EditDomainAdminACLRequest**](EditDomainAdminACLRequest.md) |  | 

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


## EditDomainAdminUser

> CreateAlias200Response EditDomainAdminUser(ctx).EditDomainAdminUserRequest(editDomainAdminUserRequest).Execute()

Edit Domain Admin user



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
	editDomainAdminUserRequest := *openapiclient.NewEditDomainAdminUserRequest() // EditDomainAdminUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAdminAPI.EditDomainAdminUser(context.Background()).EditDomainAdminUserRequest(editDomainAdminUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminAPI.EditDomainAdminUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDomainAdminUser`: CreateAlias200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainAdminAPI.EditDomainAdminUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDomainAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDomainAdminUserRequest** | [**EditDomainAdminUserRequest**](EditDomainAdminUserRequest.md) |  | 

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


## GetDomainAdmins

> GetDomainAdmins(ctx).Execute()

Get Domain Admins



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
	r, err := apiClient.DomainAdminAPI.GetDomainAdmins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAdminAPI.GetDomainAdmins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainAdminsRequest struct via the builder pattern


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

