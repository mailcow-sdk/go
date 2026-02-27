package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_RoutingAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "RoutingAPI", func(client *openapiclient.APIClient) interface{} {
		return client.RoutingAPI
	})
}