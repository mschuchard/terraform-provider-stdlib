# Flatten a list(map) into map:
data "stdlib_flatten_map" "foo" {
  param = [
    { "hello" = "world" },
    { "foo" = "bar" }
  ]
}
# => {"hello" = "world", "foo = "bar}
