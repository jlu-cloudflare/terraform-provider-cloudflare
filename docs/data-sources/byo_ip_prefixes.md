---
page_title: "cloudflare_byo_ip_prefixes Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_byo_ip_prefixes (Data Source)



## Example Usage

```terraform
data "cloudflare_byo_ip_prefixes" "example_byo_ip_prefixes" {
  account_id = "258def64c72dae45f3e4c8516e2111f2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier of a Cloudflare account.

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `account_id` (String) Identifier of a Cloudflare account.
- `advertised` (Boolean) Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.
- `advertised_modified_at` (String) Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.
- `approved` (String) Approval state of the prefix (P = pending, V = active).
- `asn` (Number) Autonomous System Number (ASN) the prefix will be advertised under.
- `cidr` (String) IP Prefix in Classless Inter-Domain Routing format.
- `created_at` (String)
- `description` (String) Description of the prefix.
- `id` (String) Identifier of an IP Prefix.
- `loa_document_id` (String) Identifier for the uploaded LOA document.
- `modified_at` (String)
- `on_demand_enabled` (Boolean) Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.
- `on_demand_locked` (Boolean) Whether advertisement status of the prefix is locked, meaning it cannot be changed.


