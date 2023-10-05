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

# Check existence of "bar" and "foo" keys in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["bar", "foo"], true
# => false
data "stdlib_has_keys" "bar_foo_all" {
  map  = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  keys = ["bar", "foo"]
  all  = true
}

# Check existence of "hello", "foo", and "baz" keys in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["hello", "foo", "baz"], true
# => true
data "stdlib_has_keys" "three_keys_all" {
  map  = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  keys = ["hello", "foo", "baz"]
  all  = true
}
