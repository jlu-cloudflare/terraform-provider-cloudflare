// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestNotificationPoliciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*notification_policy.NotificationPoliciesDataSourceModel)(nil)
	schema := notification_policy.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
