# Return map with null and empty values of type String removed.
provider::stdlib::compact_map({
  "hello" = "world",
  "foo"   = "",
  "bar"   = null
})
# result => { "hello" = "world" }

# Return map with null and empty values of type Set removed.
provider::stdlib::compact_map({
  "hello" = toset(["world"]),
  "foo"   = toset([]),
  "bar"   = null
})
# result => { "hello" = toset(["world"]) }

# Return map with null and empty values of type List removed.
provider::stdlib::compact_map({
  "hello" = ["world"],
  "foo"   = [],
  "bar"   = null
})
# result => { "hello" = ["world"] }

# Return map with null and empty values of type Map removed.
provider::stdlib::compact_map({
  "hello" = {"world" = "!"},
  "foo"   = {},
  "bar"   = null
})
# result => { "hello" = {"world" = "!"} }