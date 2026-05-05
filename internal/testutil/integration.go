// Package testutil provides shared test helpers — currently just the
// integration-test gating function. Lives under internal/ so it can't be
// imported from outside the module by mistake.
package testutil

import (
	"os"
	"testing"
)

// integrationEnv is the magic env var that a CI runner sets when the
// integration-tier infrastructure (MongoDB, DBC RPC node, IP2Location DB
// file, websocket peer) is available. The default empty value makes
// `go test ./...` skip every test wrapped with RequireIntegration so a
// developer can run unit tests on a laptop without any external services.
const integrationEnv = "AIL2_RUN_INTEGRATION_TESTS"

// RequireIntegration skips the calling test unless AIL2_RUN_INTEGRATION_TESTS
// is set to a non-empty value. Use it at the top of any test that needs
// external infrastructure (DB, RPC, network peer). The skip message is
// loud enough that it's obvious in CI logs why a test ran/didn't.
//
//	func TestMongoConnect(t *testing.T) {
//	    testutil.RequireIntegration(t)
//	    ...
//	}
func RequireIntegration(t *testing.T) {
	t.Helper()
	if os.Getenv(integrationEnv) == "" {
		t.Skipf("integration test: set %s=1 to enable (needs Mongo / RPC / external infra)", integrationEnv)
	}
}
