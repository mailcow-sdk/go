package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_SingleSignOnAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "SingleSignOnAPI", func(client *openapiclient.APIClient) interface{} {
		return client.SingleSignOnAPI
	})
}