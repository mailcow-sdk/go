package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_QueueManagerAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "QueueManagerAPI", func(client *openapiclient.APIClient) interface{} {
		return client.QueueManagerAPI
	})
}