package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_OutgoingTLSPolicyMapOverridesAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "OutgoingTLSPolicyMapOverridesAPI", func(client *openapiclient.APIClient) interface{} {
		return client.OutgoingTLSPolicyMapOverridesAPI
	})
}