---
page_title: "cloudflare_url_normalization_settings Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_url_normalization_settings (Resource)



## Example Usage

```terraform
resource "cloudflare_url_normalization_settings" "example_url_normalization_settings" {
  zone_id = "9f1839b6152d298aca64c4e906b6d074"
  scope = "incoming"
  type = "cloudflare"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scope` (String) The scope of the URL normalization.
Available values: "incoming", "both".
- `type` (String) The type of URL normalization performed by Cloudflare.
Available values: "cloudflare", "rfc3986".
- `zone_id` (String) The unique ID of the zone.

### Read-Only

- `id` (String) The unique ID of the zone.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_url_normalization_settings.example '<zone_id>'
```
