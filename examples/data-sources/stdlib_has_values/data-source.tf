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

# Check existence of "foo" and "bar" values in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["foo", "bar"], true
# => false
data "stdlib_has_values" "foo_bar_all" {
  map    = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  values = ["foo", "bar"]
  all    = true
}

# Check existence of "hello", "bar", and "bat" values in map:
# {"hello" = "world", "foo" = "bar", "baz" = "bat"}, ["world", "bar". "bat"], true
# => true
data "stdlib_has_values" "three_values_all" {
  map    = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  values = ["world", "bar", "bat"]
  all    = true
}
