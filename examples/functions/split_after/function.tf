# Return the strings split after the separator:
provider::stdlib::split_after("foo-bar-baz", "-")
# result => ["foo-", "bar-", "baz"]

# Return the strings split after the absent separator:
provider::stdlib::split_after("foo-bar-baz", "pizza")
# result => ["foo-bar-baz"]