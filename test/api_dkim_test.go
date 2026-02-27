package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_DKIMAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "DKIMAPI", func(client *openapiclient.APIClient) interface{} {
		return client.DKIMAPI
	})
}