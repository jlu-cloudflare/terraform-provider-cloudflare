// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config/migration/v500"
)

func init() {
	// Provide target schema to migration package (avoids circular import)
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*HyperdriveConfigResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths via one slot:
//
//   - Slot 0: AMBIGUOUS -- v4 framework state OR early v5 state (both schema_version=0)
//     v4: caching is types.Object, no created_on/modified_on/service_id/mtls fields
//     v5: caching is struct pointer, has created_on/modified_on/service_id/mtls
//     PriorSchema is nil; runtime detection via raw JSON disambiguates the two formats.
func (r *HyperdriveConfigResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Handle state at schema_version=0, which is AMBIGUOUS:
		// - v4 framework provider: no created_on field
		// - early v5 provider: has created_on field
		// PriorSchema=nil so we can inspect req.RawState.JSON to detect format.
		0: {
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}
