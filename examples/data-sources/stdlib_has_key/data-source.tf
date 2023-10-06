# Check existence of "foo" key in map:
data "stdlib_has_key" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "foo"
}
# => true

# Check existence of "bar" key in map:
data "stdlib_has_key" "bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  key = "bar"
}
# => false
