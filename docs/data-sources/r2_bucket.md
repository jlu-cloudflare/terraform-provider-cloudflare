---
page_title: "cloudflare_r2_bucket Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_r2_bucket (Data Source)



## Example Usage

```terraform
data "cloudflare_r2_bucket" "example_r2_bucket" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account ID
- `bucket_name` (String) Name of the bucket

### Read-Only

- `creation_date` (String) Creation timestamp
- `location` (String) Location of the bucket
Available values: "apac", "eeur", "enam", "weur", "wnam", "oc".
- `name` (String) Name of the bucket
- `storage_class` (String) Storage class for newly uploaded objects, unless specified otherwise.
Available values: "Standard", "InfrequentAccess".


