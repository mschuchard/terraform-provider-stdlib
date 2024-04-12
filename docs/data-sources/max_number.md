---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_max_number Data Source - stdlib"
subcategory: ""
description: |-
  Return the maximum number from the elements of an input list parameter.
---

# stdlib_max_number (Data Source)

Return the maximum number from the elements of an input list parameter.

## Example Usage

```terraform
# Return the maximum number from the element(s) of a list:
data "stdlib_max_number" "fibonacci" {
  param = [0, 1, 1, 2, 3, 5, 8, 13]
}
# => 13
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `param` (List of Number) Input list parameter for determining the maximum number.

### Read-Only

- `id` (String) Aliased to string input parameter(s) for efficiency and proper plan diff detection.
- `result` (Number) Function result storing the maximum number from the element(s) of the input list.