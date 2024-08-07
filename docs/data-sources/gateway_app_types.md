---
page_title: "cloudflare_gateway_app_types Data Source - Cloudflare"
subcategory: ""
description: |-
  Use this data source to retrieve all Gateway application types for an account.
---

# cloudflare_gateway_app_types (Data Source)

Use this data source to retrieve all Gateway application types for an account.

## Example Usage

```terraform
data "cloudflare_gateway_app_types" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account ID to fetch Gateway App Types from.

### Read-Only

- `app_types` (Attributes List) A list of Gateway App Types. (see [below for nested schema](#nestedatt--app_types))

<a id="nestedatt--app_types"></a>
### Nested Schema for `app_types`

Read-Only:

- `application_type_id` (Number) The identifier for the application type of this app.
- `description` (String) A short summary of the app type.
- `id` (Number) The identifier for this app type. There is only one app type per ID.
- `name` (String) The name of the app type.


