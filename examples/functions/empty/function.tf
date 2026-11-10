# Returns whether the string is empty.
provider::stdlib::empty("")
# result => true

# Returns whether the set is empty.
provider::stdlib::empty(toset(["no"]))
# result => false

# Returns whether the list is empty.
provider::stdlib::empty(tolist([]))
# result => true

# Returns whether the map is empty.
provider::stdlib::empty(tomap({ "foo" = "bar" }))
# result => false

# Returns whether the tuple is empty.
provider::stdlib::empty([])
# result => true

# Returns whether the map is empty.
provider::stdlib::empty({ "foo" = "bar" })
# result => false