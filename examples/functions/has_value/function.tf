# Check existence of "foo" value in map:
provider::stdlib::has_key({"hello" = "world", "foo" = "bar"}, "foo")
# result => false

# Check existence of "bar" value in map:
provider::stdlib::has_key({"hello" = "world", "foo" = "bar"}, "bar")
# result => true