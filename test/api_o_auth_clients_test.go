package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_OAuthClientsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "OAuthClientsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.OAuthClientsAPI
	})
}