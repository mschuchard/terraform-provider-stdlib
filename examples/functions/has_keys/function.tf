# Check existence of either "bar" or "foo" keys in map:
provider::stdlib::has_keys(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["bar", "foo"]
)
# result => true

# Check existence of either "bar" or "pizza" keys in map:
provider::stdlib::has_keys(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["bar", "pizza"]
)
# result => false

# Check existence of "bar" and "foo" keys in map:
provider::stdlib::has_keys(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["bar", "foo"],
  true
)
# result => false

# Check existence of "hello", "foo", and "baz" keys in map:
provider::stdlib::has_keys(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["hello", "foo", "baz"],
  true
)
# result => true