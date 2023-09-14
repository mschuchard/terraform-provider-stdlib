# Check existence of either "bar" or "foo" keys in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["bar", "foo"]
# => true
data "stdlib_has_keys" "bar_foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["bar", "foo"]
}

# Check existence of either "bar" or "pizza" keys in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["bar", "pizza"]
# => false
data "stdlib_has_keys" "bar_pizza" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["bar", "pizza"]
}
