---
page_title: "cloudflare_hostname_tls_setting Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_hostname_tls_setting (Resource)



## Example Usage

```terraform
resource "cloudflare_hostname_tls_setting" "example_hostname_tls_setting" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  setting_id = "ciphers"
  hostname = "app.example.com"
  value = ["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hostname` (String) The hostname for which the tls settings are set.
- `setting_id` (String) The TLS Setting name.
Available values: "ciphers", "min_tls_version", "http2".
- `value` (Dynamic) The tls setting value.
- `zone_id` (String) Identifier

### Read-Only

- `created_at` (String) This is the time the tls setting was originally created for this hostname.
- `id` (String) The TLS Setting name.
Available values: "ciphers", "min_tls_version", "http2".
- `status` (String) Deployment status for the given tls setting.
- `updated_at` (String) This is the time the tls setting was updated.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_hostname_tls_setting.example '<zone_id>/<setting_id>'
```
