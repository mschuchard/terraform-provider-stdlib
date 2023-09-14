# Check existence of "foo" value in map:
# {"hello" = "world", "foo" = "bar"}, "foo"
# => false
data "stdlib_has_value" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  value = "foo"
}

# Check existence of "bar" value in map:
# {"hello" = "world", "foo" = "bar"}, "bar"
# => true
data "stdlib_has_value" "bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  value = "bar"
}
