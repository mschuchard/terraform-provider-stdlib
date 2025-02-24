---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_empty Data Source - stdlib"
subcategory: ""
description: |-
  Return whether the input parameter of one of four possible different types is empty or not.
---

# stdlib_empty (Data Source)

Return whether the input parameter of one of four possible different types is empty or not.

## Example Usage

```terraform
# Returns whether the list is empty.
data "stdlib_empty" "list" {
  list_param = []
}
# result => true

# Returns whether the map is empty.
data "stdlib_empty" "map" {
  map_param = { "foo" = "bar" }
}
# result => false

# Returns whether the set is empty.
data "stdlib_empty" "set" {
  set_param = ["no"]
}
# result => false

# Returns whether the string is empty.
data "stdlib_empty" "string" {
  string_param = ""
}
# result => true
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `list_param` (List of String) List type parameter to check for emptiness. Must be single-level.
- `map_param` (Map of String) Map type parameter to check for emptiness. Must be single-level.
- `set_param` (Set of String) Set type parameter to check for emptiness. Must be single-level.
- `string_param` (String) String type parameter to check for emptiness.

### Read-Only

- `id` (String) Aliased to string input parameter(s) for efficiency and proper plan diff detection.
- `result` (Boolean) Function result storing whether input parameter is empty or not.
