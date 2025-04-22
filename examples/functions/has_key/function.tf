# Check existence of "foo" key in map:
provider::stdlib::has_key({"hello" = "world", "foo" = "bar"}, "foo")
# result => true

# Check existence of "bar" key in map:
provider::stdlib::has_key({"hello" = "world", "foo" = "bar"}, "bar")
# result => false
