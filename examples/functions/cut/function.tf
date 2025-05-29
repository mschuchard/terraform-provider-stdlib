# Return the separated strings:
provider::stdlib::cut("foobarbaz", "bar")
# result => ("foo", "baz", true)

# Return the separated strings with absent separator:
provider::stdlib::cut("foobarbaz", "pizza")
# result => ("foobarbaz", "", false)