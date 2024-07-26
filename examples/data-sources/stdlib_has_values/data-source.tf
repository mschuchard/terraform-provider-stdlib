# Check existence of either "foo" or "bar" values in map:
data "stdlib_has_values" "foo_bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "bar"]
}
# => true

# Check existence of either "foo" or "pizza" keys in map:
data "stdlib_has_values" "foo_pizza" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "pizza"]
}
# => false

# Check existence of "foo" and "bar" values in map:
data "stdlib_has_values" "foo_bar_all" {
  map    = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  values = ["foo", "bar"]
  all    = true
}
# => false

# Check existence of "hello", "bar", and "bat" values in map:
data "stdlib_has_values" "three_values_all" {
  map    = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
  values = ["world", "bar", "bat"]
  all    = true
}
# => true
