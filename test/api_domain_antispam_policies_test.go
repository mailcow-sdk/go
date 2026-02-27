package mailcowsdk

import (
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mailcowsdk_DomainAntispamPoliciesAPIService(t *testing.T) {
	runServiceGeneratedTests(t, "DomainAntispamPoliciesAPI", func(client *openapiclient.APIClient) interface{} {
		return client.DomainAntispamPoliciesAPI
	})
}