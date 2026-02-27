package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_FordwardingHostsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "FordwardingHostsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.FordwardingHostsAPI
	})
}