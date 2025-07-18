---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "insert function - stdlib"
subcategory: ""
description: |-
  Insert elements into a list
---

# function: insert

Return the list where values are inserted into a list at a specific index. The elements at the index in the original list are shifted up to make room. This function errors if the specified index is out of range for the list (length + 1).

## Example Usage

```terraform
# Return the list with value prepended:
provider::stdlib::insert(["one", "two", "three"], ["zero"], 0)
# result => ["zero", "one", "two", "three"]

# Return the list with values inserted in middle.
provider::stdlib::insert(["zero", "one", "four", "five"], ["two", "three"], 2)
# result => ["zero", "one", "two", "three", "four", "five"]

# Return the list with value appended (similar to concat).
provider::stdlib::insert(["zero", "one", "two"], ["three"], length(["zero", "one", "two"]))
# result => ["zero", "one", "two", "three"]
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
insert(list list of string, insert_values list of string, index number) list of string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `list` (List of String) Input list parameter into which the values will be inserted.
1. `insert_values` (List of String) Input list of values which will be inserted into the list.
1. `index` (Number) Index in the list at which to insert the values.
