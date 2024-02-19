---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stdlib_compare_list Data Source - stdlib"
subcategory: ""
description: |-
  Returns a comparison between two lists. The elements are compared sequentially, starting at index 0, until one element is not equal to the other. The result of comparing the first non-matching elements is returned. If both lists are equal until one of them ends, then the shorter list is considered less than the longer one. The result is 0 if listone == listtwo, -1 if listone < listtwo, and +1 if listone > listtwo. The input lists must be single-level
---

# stdlib_compare_list (Data Source)

Returns a comparison between two lists. The elements are compared sequentially, starting at index 0, until one element is not equal to the other. The result of comparing the first non-matching elements is returned. If both lists are equal until one of them ends, then the shorter list is considered less than the longer one. The result is 0 if list_one == list_two, -1 if list_one < list_two, and +1 if list_one > list_two. The input lists must be single-level

## Example Usage

```terraform
# Returns a comparison between two lists.
data "stdlib_compare_list" "lesser" {
  list_one = ["foo", "bar", "b"]
  list_two = ["foo", "bar", "baz"]
}
# => -1

# Returns a comparison between two lists.
data "stdlib_compare_list" "equals" {
  list_one = ["pizza", "cake"]
  list_two = ["pizza", "cake"]
}
# => 0

# Returns a comparison between two lists.
data "stdlib_compare_list" "greater" {
  list_one = ["super", "hyper", "turbo"]
  list_two = ["pizza", "cake"]
}
# => 1
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `list_one` (List of String) First input list parameter to compare with the second.
- `list_two` (List of String) Second input list parameter to compare with the first.

### Read-Only

- `id` (String) Aliased to string input parameter(s) for efficiency and proper plan diff detection.
- `result` (Number) Function result storing whether the two maps are equal.