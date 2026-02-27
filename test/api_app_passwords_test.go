package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_AppPasswordsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "AppPasswordsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.AppPasswordsAPI
	})
}