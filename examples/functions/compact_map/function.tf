# Return map with null and empty values removed.
provider::stdlib::compact_map({
  "hello" = "world",
  "foo"   = "",
  "bar"   = null
})
# result => { "hello" = "world" }