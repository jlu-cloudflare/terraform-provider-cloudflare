---
page_title: "cloudflare_turnstile_widget Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_turnstile_widget (Resource)



## Example Usage

```terraform
resource "cloudflare_turnstile_widget" "example_turnstile_widget" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  domains = ["203.0.113.1", "cloudflare.com", "blog.example.com"]
  mode = "invisible"
  name = "blog.cloudflare.com login form"
  bot_fight_mode = false
  clearance_level = "interactive"
  ephemeral_id = false
  offlabel = false
  region = "world"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier
- `domains` (List of String)
- `mode` (String) Widget Mode
Available values: "non-interactive", "invisible", "managed".
- `name` (String) Human readable widget name. Not unique. Cloudflare suggests that you
set this to a meaningful string to make it easier to identify your
widget, and where it is used.

### Optional

- `bot_fight_mode` (Boolean) If bot_fight_mode is set to `true`, Cloudflare issues computationally
expensive challenges in response to malicious bots (ENT only).
- `clearance_level` (String) If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,
this setting can determine the clearance level to be set
Available values: "no_clearance", "jschallenge", "managed", "interactive".
- `ephemeral_id` (Boolean) Return the Ephemeral ID in /siteverify (ENT only).
- `offlabel` (Boolean) Do not show any Cloudflare branding on the widget (ENT only).
- `region` (String) Region where this widget can be used.
Available values: "world".

### Read-Only

- `created_on` (String) When the widget was created.
- `id` (String) Widget item identifier tag.
- `modified_on` (String) When the widget was modified.
- `secret` (String) Secret key for this widget.
- `sitekey` (String) Widget item identifier tag.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_turnstile_widget.example '<account_id>/<sitekey>'
```
