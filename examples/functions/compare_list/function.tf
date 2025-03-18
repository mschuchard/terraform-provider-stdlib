# Returns a comparison between two lists (final element of first list is less than final element of second list):
provider::stdlib::compare_list(["foo", "bar", "b"], ["foo", "bar", "baz"])
# result => -1

# Returns a comparison between two lists (lists are equal):
provider::stdlib::compare_list(["pizza", "cake"], ["pizza", "cake"])
# result => 0

# Returns a comparison between two lists (second element of first list is greater than second element of second list):
provider::stdlib::compare_list(["super", "hyper", "turbo"], ["pizza", "cake", "punch"])
# result => 1

# Returns a comparison between two lists (lists are equal until first list has more elements):
provider::stdlib::compare_list(["pizza", "cake", "punch"], ["pizza", "cake"])
# result => 1