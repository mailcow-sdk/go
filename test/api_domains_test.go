package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_DomainsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "DomainsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.DomainsAPI
	})
}