package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_AliasesAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "AliasesAPI", func(client *openapiclient.APIClient) interface{} {
		return client.AliasesAPI
	})
}