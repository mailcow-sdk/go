# \LogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetACMELogs**](LogsAPI.md#GetACMELogs) | **Get** /api/v1/get/logs/acme/{count} | Get ACME logs
[**GetApiLogs**](LogsAPI.md#GetApiLogs) | **Get** /api/v1/get/logs/api/{count} | Get Api logs
[**GetAutodiscoverLogs**](LogsAPI.md#GetAutodiscoverLogs) | **Get** /api/v1/get/logs/autodiscover/{count} | Get Autodiscover logs
[**GetDovecotLogs**](LogsAPI.md#GetDovecotLogs) | **Get** /api/v1/get/logs/dovecot/{count} | Get Dovecot logs
[**GetNetfilterLogs**](LogsAPI.md#GetNetfilterLogs) | **Get** /api/v1/get/logs/netfilter/{count} | Get Netfilter logs
[**GetPostfixLogs**](LogsAPI.md#GetPostfixLogs) | **Get** /api/v1/get/logs/postfix/{count} | Get Postfix logs
[**GetRatelimitLogs**](LogsAPI.md#GetRatelimitLogs) | **Get** /api/v1/get/logs/ratelimited/{count} | Get Ratelimit logs
[**GetRspamdLogs**](LogsAPI.md#GetRspamdLogs) | **Get** /api/v1/get/logs/rspamd-history/{count} | Get Rspamd logs
[**GetSOGoLogs**](LogsAPI.md#GetSOGoLogs) | **Get** /api/v1/get/logs/sogo/{count} | Get SOGo logs
[**GetWatchdogLogs**](LogsAPI.md#GetWatchdogLogs) | **Get** /api/v1/get/logs/watchdog/{count} | Get Watchdog logs



## GetACMELogs

> GetACMELogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get ACME logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetACMELogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetACMELogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACMELogsRequest struct via the builder pattern


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


## GetApiLogs

> GetApiLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Api logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetApiLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetApiLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiLogsRequest struct via the builder pattern


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


## GetAutodiscoverLogs

> GetAutodiscoverLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Autodiscover logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetAutodiscoverLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetAutodiscoverLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutodiscoverLogsRequest struct via the builder pattern


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


## GetDovecotLogs

> GetDovecotLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Dovecot logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetDovecotLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetDovecotLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDovecotLogsRequest struct via the builder pattern


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


## GetNetfilterLogs

> GetNetfilterLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Netfilter logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetNetfilterLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetNetfilterLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetfilterLogsRequest struct via the builder pattern


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


## GetPostfixLogs

> GetPostfixLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Postfix logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetPostfixLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetPostfixLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostfixLogsRequest struct via the builder pattern


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


## GetRatelimitLogs

> GetRatelimitLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Ratelimit logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetRatelimitLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetRatelimitLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRatelimitLogsRequest struct via the builder pattern


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


## GetRspamdLogs

> GetRspamdLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Rspamd logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetRspamdLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetRspamdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRspamdLogsRequest struct via the builder pattern


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


## GetSOGoLogs

> GetSOGoLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get SOGo logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetSOGoLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetSOGoLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSOGoLogsRequest struct via the builder pattern


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


## GetWatchdogLogs

> GetWatchdogLogs(ctx, count).XAPIKey(xAPIKey).Execute()

Get Watchdog logs



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
	count := float32(8.14) // float32 | Number of logs to return
	xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetWatchdogLogs(context.Background(), count).XAPIKey(xAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetWatchdogLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32** | Number of logs to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchdogLogsRequest struct via the builder pattern


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

