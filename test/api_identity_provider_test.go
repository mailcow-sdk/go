package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_IdentityProviderAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "IdentityProviderAPI", func(client *openapiclient.APIClient) interface{} {
		return client.IdentityProviderAPI
	})
}