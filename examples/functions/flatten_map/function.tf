# Flatten a list(map) into map:
provider::stdlib::flatten_map([
  { "hello" = "world" },
  { "foo" = "bar" }
])
# result => {"hello" = "world", "foo = "bar}