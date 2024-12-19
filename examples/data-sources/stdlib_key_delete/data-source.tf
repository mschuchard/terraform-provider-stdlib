# Remove the "foo" key from a map:
data "stdlib_key_delete" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "foo"
}
# result => {"hello" = "world"}
