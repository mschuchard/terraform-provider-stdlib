# Remove the "foo" and "baz" keys from a map:
provider::stdlib::keys_delete(
  {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  },
  ["foo", "baz"]
)
# result => {"hello" = "world"}
