# Reports whether the two maps contain the same key/value pairs.
# {"hello" = "world"}, {"hello" = "world"}
# => true
data "stdlib_equal_map" "foo" {
  map_one = { "hello" = "world" }
  map_two = { "hello" = "world" }
}

# Reports whether the two maps contain the same key/value pairs.
# {"hello" = "world"}, {"foo" = "bar"}
# => false
data "stdlib_equal_map" "bar" {
  map_one = { "hello" = "world" }
  map_two = { "foo" = "bar" }
}
