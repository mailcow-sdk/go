package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_QuarantineAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "QuarantineAPI", func(client *openapiclient.APIClient) interface{} {
		return client.QuarantineAPI
	})
}