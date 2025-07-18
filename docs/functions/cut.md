---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cut function - stdlib"
subcategory: ""
description: |-
  Cut a string in two
---

# function: cut

Returns the strings before and after the first instance of the separator in the input string. Also returns whether or not the separator was found in the input string. The return is a tuple: `before`, `after`, `found`. If the separator is not found in the input string, then `found` will be false, `before` will be equal to the `string` parameter, and `after` will be an empty string.

## Example Usage

```terraform
# Return the separated strings:
provider::stdlib::cut("foobarbaz", "bar")
# result => ("foo", "baz", true)

# Return the separated strings with absent separator:
provider::stdlib::cut("foobarbaz", "pizza")
# result => ("foobarbaz", "", false)
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
cut(string string, separator string) dynamic
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `string` (String) Input string parameter for cutting around a separator.
1. `separator` (String) The separator for cutting the input string.
