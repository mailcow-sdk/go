package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_RatelimitsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "RatelimitsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.RatelimitsAPI
	})
}