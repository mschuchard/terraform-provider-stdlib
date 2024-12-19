# Check existence of either "bar" or "foo" keys in map:
data "stdlib_has_keys" "bar_foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["bar", "foo"]
}
# result => true

# Check existence of either "bar" or "pizza" keys in map:
data "stdlib_has_keys" "bar_pizza" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["bar", "pizza"]
}
# result => false

# Check existence of "bar" and "foo" keys in map:
data "stdlib_has_keys" "bar_foo_all" {
  map  = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  keys = ["bar", "foo"]
  all  = true
}
# result => false

# Check existence of "hello", "foo", and "baz" keys in map:
data "stdlib_has_keys" "three_keys_all" {
  map  = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  keys = ["hello", "foo", "baz"]
  all  = true
}
# result => true
