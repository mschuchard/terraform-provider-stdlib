---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_flatten_map Data Source - stdlib"
subcategory: ""
description: |-
  Return the flattened map of an input list of maps parameter.
---

# stdlib_flatten_map (Data Source)

Return the flattened map of an input list of maps parameter.

## Example Usage

```terraform
# Flatten a list(map) into map: [{"hello" = "world"}, {"foo" = "bar"}] => {"hello" = "world", "foo = "bar}
data "stdlib_flatten_map" "foo" {
  param = [
    { "hello" = "world" },
    { "foo" = "bar" }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `param` (List of Map of String) Input list of maps to flatten.

### Read-Only

- `id` (String) Aliased to string input parameter for efficiency.
- `result` (Map of String) Function result storing the flattened map.


