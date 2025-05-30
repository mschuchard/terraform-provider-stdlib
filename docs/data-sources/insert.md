---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_insert Data Source - stdlib"
subcategory: ""
description: |-
  Return the list where values are inserted into a list at a specific index. The elements at the index in the original list are shifted up to make room. This function errors if the specified index is out of range for the list (length + 1).
---

# stdlib_insert (Data Source)

Return the list where values are inserted into a list at a specific index. The elements at the index in the original list are shifted up to make room. This function errors if the specified index is out of range for the list (length + 1).

## Example Usage

```terraform
# Return the list with value prepended.
data "stdlib_insert" "prepend" {
  list_param    = ["one", "two", "three"]
  insert_values = ["zero"]
  index         = 0
}
# result => ["zero", "one", "two", "three"]

# Return the list with values inserted in middle.
data "stdlib_insert" "insert" {
  list_param    = ["zero", "one", "four", "five"]
  insert_values = ["two", "three"]
  index         = 2
}
# result => ["zero", "one", "two", "three", "four", "five"]

# Return the list with value appended (similar to concat).
data "stdlib_insert" "append" {
  list_param    = ["zero", "one", "two"]
  insert_values = ["three"]
  index         = length(["zero", "one", "two"])
}
# result => ["zero", "one", "two", "three"]
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `index` (Number) Index in the list at which to insert the values.
- `insert_values` (List of String) Input list of values which will be inserted into the list.
- `list_param` (List of String) Input list parameter into which the values will be inserted.

### Read-Only

- `id` (String) Aliased to string input parameter(s) for efficiency and proper plan diff detection.
- `result` (List of String) The resulting list with the inserted values.
