# Return the list with value prepended:
provider::stdlib::insert(["one", "two", "three"], ["zero"], 0)
# result => ["zero", "one", "two", "three"]

# Return the list with values inserted in middle.
provider::stdlib::insert(["zero", "one", "four", "five"], ["two", "three"], 2)
# result => ["zero", "one", "two", "three", "four", "five"]

# Return the list with value appended (similar to concat).
provider::stdlib::insert(["zero", "one", "two"], ["three"], length(["zero", "one", "two"]))
# result => ["zero", "one", "two", "three"]