# Reports whether the two maps contain the same key/value pairs:
provider::stdlib::equal_map({ "hello" = "world" }, { "hello" = "world" })
# result => true

# Reports whether the two maps contain the same key/value pairs:
provider::stdlib::equal_map({ "hello" = "world" }, { "foo" = "bar" })
# result => false