// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*DNSZoneTransfersPeerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the peer.",
				Required:    true,
			},
			"ip": schema.StringAttribute{
				Description: "IPv4/IPv6 address of primary or secondary nameserver, depending on what zone this peer is linked to. For primary zones this IP defines the IP of the secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary zones this IP defines the IP of the primary nameserver Cloudflare will send AXFR/IXFR requests to.",
				Optional:    true,
			},
			"ixfr_enable": schema.BoolAttribute{
				Description: "Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary zones.",
				Optional:    true,
			},
			"port": schema.Float64Attribute{
				Description: "DNS port of primary or secondary nameserver, depending on what zone this peer is linked to.",
				Optional:    true,
			},
			"tsig_id": schema.StringAttribute{
				Description: "TSIG authentication will be used for zone transfer if configured.",
				Optional:    true,
			},
		},
	}
}

func (r *DNSZoneTransfersPeerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *DNSZoneTransfersPeerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
