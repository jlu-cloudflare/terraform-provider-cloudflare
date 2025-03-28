// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_acl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSZoneTransfersACLDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_acl.DNSZoneTransfersACLDataSourceModel)(nil)
	schema := dns_zone_transfers_acl.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
