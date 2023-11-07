# Returns whether the list is empty.
data "stdlib_empty" "list" {
  list_param = []
}
# => true

# Returns whether the map is empty.
data "stdlib_empty" "map" {
  map_param = { "foo" = "bar" }
}
# => false

# Returns whether the set is empty.
data "stdlib_empty" "set" {
  set_param = ["no"]
}
# => false

# Returns whether the string is empty.
data "stdlib_empty" "string" {
  string_param = ""
}
# => true
