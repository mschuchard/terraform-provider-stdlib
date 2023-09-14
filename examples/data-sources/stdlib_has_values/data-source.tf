# Check existence of either "foo" or "bar" values in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["foo", "bar"]
# => true
data "stdlib_has_keys" "foo_bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "bar"]
}

# Check existence of either "foo" or "pizza" keys in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["foo", "pizza"]
# => false
data "stdlib_has_keys" "foo_pizza" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "pizza"]
}
