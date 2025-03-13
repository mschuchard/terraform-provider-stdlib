# Return the element's index within a list:
provider::stdlib::compare_list(["zero", "one", "two"], "one")
# result: 1

# Return the element's index within a sorted list:
provider::stdlib::compare_list(["a", "b", "c", "d"], "c", true)
# result: 2

# Return the element's first occurrence index within a list:
provider::stdlib::compare_list(["zero", "one", "two", "three", "two", "one", "zero"], "two")
# result: 2

# Return the element's nonexistence within a list:
provider::stdlib::compare_list(["hundred", "thousand", "million", "billion"], "infinity")
# result: -1