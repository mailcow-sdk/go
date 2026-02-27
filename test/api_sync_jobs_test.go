package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_SyncJobsAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "SyncJobsAPI", func(client *openapiclient.APIClient) interface{} {
		return client.SyncJobsAPI
	})
}