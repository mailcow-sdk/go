package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_AddressRewritingAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "AddressRewritingAPI", func(client *openapiclient.APIClient) interface{} {
		return client.AddressRewritingAPI
	})
}