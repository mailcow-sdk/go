# \IdentityProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditExternalIdentityProviderSettings**](IdentityProviderAPI.md#EditExternalIdentityProviderSettings) | **Post** /api/v1/edit/identity-provider | Edit external Identity Provider



## EditExternalIdentityProviderSettings

> EditExternalIdentityProviderSettings(ctx).EditExternalIdentityProviderSettingsRequest(editExternalIdentityProviderSettingsRequest).Execute()

Edit external Identity Provider



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
	editExternalIdentityProviderSettingsRequest := *openapiclient.NewEditExternalIdentityProviderSettingsRequest() // EditExternalIdentityProviderSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderAPI.EditExternalIdentityProviderSettings(context.Background()).EditExternalIdentityProviderSettingsRequest(editExternalIdentityProviderSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.EditExternalIdentityProviderSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditExternalIdentityProviderSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editExternalIdentityProviderSettingsRequest** | [**EditExternalIdentityProviderSettingsRequest**](EditExternalIdentityProviderSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

