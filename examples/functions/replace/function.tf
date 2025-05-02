# Return the list with beginning values replaced:
provider::stdlib::replace(["foo", "bar", "two", "three"], ["zero", "one"], 0)
# result => ["zero", "one", "two", "three"]

# Return the list with middle values replaced:
provider::stdlib::replace(["zero", "foo", "bar", "baz", "four", "five"], ["one", "two", "three"], 1)
# result => ["zero", "one", "two", "three", "four", "five"]

# Return the list with middle values replaced and zeroed:
provider::stdlib::replace(["zero", "foo", "bar", "four", "five"], ["one"], 1, 2)
# result => ["zero", "one", "four", "five"]

# Return the list with terminating values replaced:
provider::stdlib::replace(["zero", "foo", "bar", "baz"], ["one", "two", "three"], length(["zero", "foo", "bar", "baz"]) - length(["one", "two", "three"]))
# result => ["zero", "one", "two", "three"]