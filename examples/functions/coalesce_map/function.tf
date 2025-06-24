# Return first non-empty map.
provider::stdlib::coalesce_map(
  {},
  { "hello" = "world" },
  { "foo"   = "bar" },
  {}
}
# result => { "hello" = "world" }