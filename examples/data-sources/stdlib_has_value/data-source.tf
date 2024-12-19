# Check existence of "foo" value in map:
data "stdlib_has_value" "foo" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  value = "foo"
}
# result => false

# Check existence of "bar" value in map:
data "stdlib_has_value" "bar" {
  map = {
    "hello" = "world",
    "foo"   = "bar"
  }
  value = "bar"
}
# result => true
