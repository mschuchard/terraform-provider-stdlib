---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_cut Data Source - stdlib"
subcategory: ""
description: |-
  Returns the strings before and after the first instance of the separator in the input string. Also returns whether or not the separator was found in the input string. If the separator is not found in the input string, then found will be false, before will be equal to param, and after will be an empty string.
---

# stdlib_cut (Data Source)

Returns the strings before and after the first instance of the separator in the input string. Also returns whether or not the separator was found in the input string. If the separator is not found in the input string, then `found` will be false, `before` will be equal to `param`, and `after` will be an empty string.

## Example Usage

```terraform
# Return the separated strings:
data "stdlib_cut" "foobarbaz" {
  param     = "foobarbaz"
  separator = "bar"
}
# before => "foo", after => "baz", found = true

# Return the separated strings with absent separator:
data "stdlib_cut" "pizza" {
  param     = "foobarbaz"
  separator = "pizza"
}
# before => "foobarbaz", after => "", found = false
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `param` (String) Input string parameter for cutting around a separator.
- `separator` (String) The separator for cutting the input string.

### Read-Only

- `after` (String) Function result storing the input string after the separator.
- `before` (String) Function result storing the input string before the separator.
- `found` (Boolean) Function result storing whether the input string contained the separator.
- `id` (String) Aliased to string input parameter(s) for efficiency and proper plan diff detection.
