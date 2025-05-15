# Check existence of "foo" value in map:
provider::stdlib::has_value({"hello" = "world", "foo" = "bar"}, "foo")
# result => false

# Check existence of "bar" value in map:
provider::stdlib::has_value({"hello" = "world", "foo" = "bar"}, "bar")
# result => true
