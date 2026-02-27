package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_ResourcesAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "ResourcesAPI", func(client *openapiclient.APIClient) interface{} {
		return client.ResourcesAPI
	})
}