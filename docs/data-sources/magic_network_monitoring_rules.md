---
page_title: "cloudflare_magic_network_monitoring_rules Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_magic_network_monitoring_rules (Data Source)



## Example Usage

```terraform
data "cloudflare_magic_network_monitoring_rules" "example_magic_network_monitoring_rules" {
  account_id = "6f91088a406011ed95aed352566e8d4c"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `automatic_advertisement` (Boolean) Toggle on if you would like Cloudflare to automatically advertise the IP Prefixes within the rule via Magic Transit when the rule is triggered. Only available for users of Magic Transit.
- `bandwidth_threshold` (Number) The number of bits per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
- `duration` (String) The amount of time that the rule threshold must be exceeded to send an alert notification. The final value must be equivalent to one of the following 8 values ["1m","5m","10m","15m","20m","30m","45m","60m"].
Available values: "1m", "5m", "10m", "15m", "20m", "30m", "45m", "60m".
- `id` (String) The id of the rule. Must be unique.
- `name` (String) The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9, underscore (_), dash (-), period (.), and tilde (~). You can’t have a space in the rule name. Max 256 characters.
- `packet_threshold` (Number) The number of packets per second for the rule. When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
- `prefix_match` (String) Prefix match type to be applied for a prefix auto advertisement when using an advanced_ddos rule.
Available values: "exact", "subnet", "supernet".
- `prefixes` (List of String)
- `type` (String) MNM rule type.
Available values: "threshold", "zscore", "advanced_ddos".
- `zscore_sensitivity` (String) Level of sensitivity set for zscore rules.
Available values: "low", "medium", "high".
- `zscore_target` (String) Target of the zscore rule analysis.
Available values: "bits", "packets".


