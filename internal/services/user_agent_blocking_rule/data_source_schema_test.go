// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/user_agent_blocking_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestUserAgentBlockingRuleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user_agent_blocking_rule.UserAgentBlockingRuleDataSourceModel)(nil)
	schema := user_agent_blocking_rule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
