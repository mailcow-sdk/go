package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_StatusAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "StatusAPI", func(client *openapiclient.APIClient) interface{} {
		return client.StatusAPI
	})
}