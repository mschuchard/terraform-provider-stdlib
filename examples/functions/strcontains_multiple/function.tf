# Determine if empty list of substrings are contained within a string:
provider::stdlib::strcontains_multiple("foobarbazbot", [])
# result => false

# Determine if any substrings are contained within a string:
provider::stdlib::strcontains_multiple("foobarbazbot", ["foo", "pizza"])
# result => true

# Determine if any substrings are contained within a string:
provider::stdlib::strcontains_multiple("foobarbazbot", ["pizza", "cake"])
# result => false

# Determine if all substrings are contained within a string:
provider::stdlib::strcontains_multiple("foobarbazbot", ["foo", "pizza"], true)
# result => false

# Determine if all substrings are contained within a string:
provider::stdlib::strcontains_multiple("foobarbazbot", ["foo", "baz"], true)
# result => true