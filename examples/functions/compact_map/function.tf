# Return map with null and empty values of type String removed.
provider::stdlib::compact_map({
  "hello" = "world",
  "foo"   = "",
  "bar"   = null
})
# result => { "hello" = "world" }