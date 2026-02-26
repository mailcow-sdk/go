# \CrossOriginResourceSharingCORSAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditCrossOriginResourceSharingCORSSettings**](CrossOriginResourceSharingCORSAPI.md#EditCrossOriginResourceSharingCORSSettings) | **Post** /api/v1/edit/cors | Edit Cross-Origin Resource Sharing (CORS) settings



## EditCrossOriginResourceSharingCORSSettings

> EditCrossOriginResourceSharingCORSSettings(ctx).EditCrossOriginResourceSharingCORSSettingsRequest(editCrossOriginResourceSharingCORSSettingsRequest).Execute()

Edit Cross-Origin Resource Sharing (CORS) settings



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
	editCrossOriginResourceSharingCORSSettingsRequest := *openapiclient.NewEditCrossOriginResourceSharingCORSSettingsRequest() // EditCrossOriginResourceSharingCORSSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrossOriginResourceSharingCORSAPI.EditCrossOriginResourceSharingCORSSettings(context.Background()).EditCrossOriginResourceSharingCORSSettingsRequest(editCrossOriginResourceSharingCORSSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossOriginResourceSharingCORSAPI.EditCrossOriginResourceSharingCORSSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditCrossOriginResourceSharingCORSSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editCrossOriginResourceSharingCORSSettingsRequest** | [**EditCrossOriginResourceSharingCORSSettingsRequest**](EditCrossOriginResourceSharingCORSSettingsRequest.md) |  | 

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

