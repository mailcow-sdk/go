package mailcowsdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

type captureTransport struct {
	status   int
	body     string
	headers  http.Header
	requests []*http.Request
}

func (t *captureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.requests = append(t.requests, req)
	h := make(http.Header)
	for k, v := range t.headers {
		h[k] = append([]string{}, v...)
	}
	return &http.Response{
		StatusCode: t.status,
		Status:     fmt.Sprintf("%d generated", t.status),
		Header:     h,
		Body:       io.NopCloser(strings.NewReader(t.body)),
		Request:    req,
	}, nil
}

func newMockClient(status int, body string, headers http.Header) (*openapiclient.APIClient, *captureTransport) {
	cfg := openapiclient.NewConfiguration()
	cfg.AddDefaultHeader("X-API-Key", "test-api-key")
	transport := &captureTransport{status: status, body: body, headers: headers}
	cfg.HTTPClient = &http.Client{Transport: transport}
	return openapiclient.NewAPIClient(cfg), transport
}

func sampleArg(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf("value/with space").Convert(t)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(1).Convert(t)
	case reflect.Bool:
		return reflect.ValueOf(true).Convert(t)
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(1.0).Convert(t)
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map:
		return reflect.Zero(t)
	default:
		return reflect.Zero(t)
	}
}

func runServiceGeneratedTests(t *testing.T, serviceName string, serviceFactory func(*openapiclient.APIClient) interface{}) {
	t.Run(serviceName+"_request_auth_mapping", func(t *testing.T) {
		client, transport := newMockClient(200, "null", http.Header{"Content-Type": []string{"application/json"}})
		service := reflect.ValueOf(serviceFactory(client))
		serviceType := service.Type()

		ctx := context.WithValue(context.Background(), openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
			"ApiKeyAuth": {Key: "test-api-key"},
		})

		for i := 0; i < serviceType.NumMethod(); i++ {
			method := serviceType.Method(i)
			if strings.HasSuffix(method.Name, "Execute") {
				continue
			}
			call := service.Method(i)
			args := make([]reflect.Value, 0, call.Type().NumIn())
			for j := 0; j < call.Type().NumIn(); j++ {
				argType := call.Type().In(j)
				if argType.String() == "context.Context" {
					args = append(args, reflect.ValueOf(ctx))
				} else {
					args = append(args, sampleArg(argType))
				}
			}

			builder := call.Call(args)[0]
			exec := builder.MethodByName("Execute")
			if !exec.IsValid() {
				t.Fatalf("builder for %s has no Execute()", method.Name)
			}
			out := exec.Call(nil)
			errVal := out[len(out)-1]
			if !errVal.IsNil() {
				t.Fatalf("%s failed: %v", method.Name, errVal.Interface())
			}

			lastReq := transport.requests[len(transport.requests)-1]
			if !strings.HasPrefix(lastReq.URL.Path, "/api/v1/") {
				t.Fatalf("%s built unexpected path: %s", method.Name, lastReq.URL.Path)
			}
			if lastReq.Method == "" {
				t.Fatalf("%s built empty method", method.Name)
			}
			if got := lastReq.Header.Get("X-API-Key"); got != "test-api-key" {
				t.Fatalf("%s missing auth header, got: %q", method.Name, got)
			}
		}
	})

	t.Run(serviceName+"_error_mapping", func(t *testing.T) {
		client, _ := newMockClient(500, `{"type":"danger","msg":"server error"}`, http.Header{
			"Content-Type": []string{"application/json"},
			"x-request-id": []string{"req-500"},
		})
		service := reflect.ValueOf(serviceFactory(client))
		serviceType := service.Type()

		ctx := context.WithValue(context.Background(), openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
			"ApiKeyAuth": {Key: "test-api-key"},
		})

		for i := 0; i < serviceType.NumMethod(); i++ {
			method := serviceType.Method(i)
			if strings.HasSuffix(method.Name, "Execute") {
				continue
			}
			call := service.Method(i)
			args := make([]reflect.Value, 0, call.Type().NumIn())
			for j := 0; j < call.Type().NumIn(); j++ {
				argType := call.Type().In(j)
				if argType.String() == "context.Context" {
					args = append(args, reflect.ValueOf(ctx))
				} else {
					args = append(args, sampleArg(argType))
				}
			}
			builder := call.Call(args)[0]
			exec := builder.MethodByName("Execute")
			if !exec.IsValid() {
				continue
			}
			out := exec.Call(nil)
			errVal := out[len(out)-1]
			if errVal.IsNil() {
				t.Fatalf("%s expected error for 500 response", method.Name)
			}
			if !strings.Contains(errVal.Interface().(error).Error(), "500") {
				t.Fatalf("%s expected 500 in error, got: %v", method.Name, errVal.Interface())
			}
			break
		}
	})
}
