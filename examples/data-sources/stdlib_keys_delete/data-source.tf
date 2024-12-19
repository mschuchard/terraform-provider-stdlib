# Remove the "foo" and "baz" keys from a map:
data "stdlib_keys_delete" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "baz"]
}
# result => {"hello" = "world"}
