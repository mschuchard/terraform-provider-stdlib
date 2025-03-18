# Return the element's index within a list:
provider::stdlib::list_index(["zero", "one", "two"], "one")
# result => 1

# Return the element's index within a sorted list:
provider::stdlib::list_index(["a", "b", "c", "d"], "c", true)
# result => 2

# Return the element's first occurrence index within a list:
provider::stdlib::list_index(["zero", "one", "two", "three", "two", "one", "zero"], "two")
# result => 2

# Return the element's nonexistence within a list:
provider::stdlib::list_index(["hundred", "thousand", "million", "billion"], "infinity")
# result => -1