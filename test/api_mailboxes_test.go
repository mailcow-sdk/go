package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_MailboxesAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "MailboxesAPI", func(client *openapiclient.APIClient) interface{} {
		return client.MailboxesAPI
	})
}