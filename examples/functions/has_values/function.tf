# Check existence of either "foo" or "bar" values in map:
provider::stdlib::has_values(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["foo", "bar"]
)
# result => true

# Check existence of either "foo" or "pizza" keys in map:
provider::stdlib::has_values(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["foo", "pizza"]
)
# result => false

# Check existence of "foo" and "bar" values in map:
provider::stdlib::has_values(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["foo", "bar"],
  true
)
# result => false

# Check existence of "hello", "bar", and "bat" values in map:
provider::stdlib::has_values(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["world", "bar", "bat"],
  true
)
# result => true