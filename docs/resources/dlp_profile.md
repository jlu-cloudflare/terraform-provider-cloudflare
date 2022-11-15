---
page_title: "cloudflare_dlp_profile Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare DLP Profile resource. Data Loss Prevention profiles
  are a set of entries that can be matched in HTTP bodies or files.
  They are referenced in Zero Trust Gateway rules.
---

# cloudflare_dlp_profile (Resource)

Provides a Cloudflare DLP Profile resource. Data Loss Prevention profiles
are a set of entries that can be matched in HTTP bodies or files.
They are referenced in Zero Trust Gateway rules.

## Example Usage

```terraform
# Predefined profile
resource "cloudflare_dlp_profile" "example_predefined" {
  account_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Example Predefined Profile"
  type        = "predefined"

  entry {
	name = "Mastercard Card Number"
	enabled = true
  }

  entry {
	name = "Union Pay Card Number"
	enabled = false
  }
}

# Custom profile
resource "cloudflare_dlp_profile" "example_custom" {
  account_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "Example Custom Profile"
  description = "A profile with example entries"
  type        = "custom"

  entry {
	name = "Matches visa credit cards"
	enabled = true
	pattern {
		regex = "4\d{3}([-\. ])?\d{4}([-\. ])?\d{4}([-\. ])?\d{4}"
		validation = "luhn"
	}
  }

  entry {
	name = "Matches diners club card"
	enabled = true
	pattern {
		regex = "(?:0[0-5]|[68][0-9])[0-9]{11}"
		validation = "luhn"
	}
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) The account identifier to target for the resource.
- `entry` (Block Set, Min: 1) List of entries to apply to the profile. (see [below for nested schema](#nestedblock--entry))
- `name` (String) Name of the profile.
- `type` (String) The type of the profile. Available values: `custom`, `predefined`.

### Optional

- `description` (String) Brief summary of the profile and its intended use.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--entry"></a>
### Nested Schema for `entry`

Required:

- `name` (String) Name of the entry to deploy.

Optional:

- `enabled` (Boolean) Whether the entry is active. Defaults to `false`.
- `id` (String) Unique entry identifier.
- `pattern` (Block List, Max: 1) (see [below for nested schema](#nestedblock--entry--pattern))

<a id="nestedblock--entry--pattern"></a>
### Nested Schema for `entry.pattern`

Required:

- `regex` (String) The regex that defines the pattern.

Optional:

- `validation` (String) The validation algorithm to apply with this pattern.

## Import

Import is supported using the following syntax:
```shell
$ terraform import cloudflare_dlp_profile.example <account_id>/<dlp_profile_id>
```