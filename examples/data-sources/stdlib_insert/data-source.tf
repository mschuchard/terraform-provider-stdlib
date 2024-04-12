# Return the list with value prepended.
data "stdlib_insert" "prepend" {
  list_param    = ["one", "two", "three"]
  insert_values = ["zero"]
  index         = 0
}
# => ["zero", "one", "two", "three"]

# Return the list with values inserted in middle.
data "stdlib_insert" "insert" {
  list_param    = ["zero", "one", "four", "five"]
  insert_values = ["two", "three"]
  index         = 2
}
# => ["zero", "one", "two", "three", "four", "five"]


# Return the list with value appended (similar to concat).
data "stdlib_insert" "append" {
  list_param    = ["zero", "one", "two"]
  insert_values = ["three"]
  index         = length(["zero", "one", "two"])
}
# => ["zero", "one", "two", "three"]
