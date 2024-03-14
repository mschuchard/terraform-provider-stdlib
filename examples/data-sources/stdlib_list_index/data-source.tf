# Return the element's index within a list:
data "stdlib_list_index" "one" {
  list_param = ["zero", "one", "two"]
  elem_param = "one"
}
# => 1

# Return the element's first occurrence index within a list:
data "stdlib_list_index" "two" {
  list_param = ["zero", "one", "two", "three", "two", "one", "zero"]
  elem_param = "two"
}
# => 2


# Return the element's nonexistence within a list:
data "stdlib_list_index" "infinity" {
  list_param = ["hundred", "thousand", "million", "billion"]
  elem_param = "infinity"
}
# => -1
