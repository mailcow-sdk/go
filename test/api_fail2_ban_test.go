package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_Fail2BanAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "Fail2BanAPI", func(client *openapiclient.APIClient) interface{} {
		return client.Fail2BanAPI
	})
}