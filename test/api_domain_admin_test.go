package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_DomainAdminAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "DomainAdminAPI", func(client *openapiclient.APIClient) interface{} {
		return client.DomainAdminAPI
	})
}