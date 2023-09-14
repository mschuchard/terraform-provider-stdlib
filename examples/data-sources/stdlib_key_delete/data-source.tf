# Remove the "foo" key from a map:
# {"hello" = "world", "foo" = "bar"}, "foo"
# => {"hello" = "world"}
data "stdlib_key_delete" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "foo"
}
