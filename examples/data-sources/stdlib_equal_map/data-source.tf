# Reports whether the two maps contain the same key/value pairs.
data "stdlib_equal_map" "foo" {
  map_one = { "hello" = "world" }
  map_two = { "hello" = "world" }
}
# result => true

# Reports whether the two maps contain the same key/value pairs.
data "stdlib_equal_map" "bar" {
  map_one = { "hello" = "world" }
  map_two = { "foo" = "bar" }
}
# result => false
