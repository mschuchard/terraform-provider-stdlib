# Check existence of "foo" key in map:
# {"hello" = "world", "foo" = "bar"}, "foo"
# => true
data "stdlib_has_key" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "foo"
}

# Check existence of "bar" key in map:
# {"hello" = "world", "foo" = "bar"}, "bar"
# => false
data "stdlib_has_key" "bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "bar"
}
