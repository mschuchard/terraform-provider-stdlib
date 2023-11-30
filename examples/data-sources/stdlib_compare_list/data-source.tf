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
