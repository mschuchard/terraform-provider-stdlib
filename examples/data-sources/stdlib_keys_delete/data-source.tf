# Remove the "foo" and "baz" keys from a map:
# {"hello" = "world", "foo" = "bar", "baz" => "bat"}, ["foo", "baz"]
# => {"hello" = "world"}
data "stdlib_keys_delete" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar",
    "baz"   = "bat"
  }
  keys = ["foo", "baz"]
}
