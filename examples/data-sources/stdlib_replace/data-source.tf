# Return the list with beginning value replaced.
data "stdlib_replace" "begin" {
  list_param     = ["foo", "two", "three"]
  replace_values = ["zero", "one"]
  index          = 0
}
# => ["zero", "one", "two", "three"]

# Return the list with middle values replaced.
data "stdlib_replace" "replace" {
  list_param     = ["zero", "foo", "bar", "four", "five"]
  replace_values = ["one", "two", "three"]
  index          = 1
}
# => ["zero", "one", "two", "three", "four", "five"]

# Return the list with terminating values replaced.
data "stdlib_replace" "append" {
  list_param     = ["zero", "foo", "bar"]
  replace_values = ["one", "two", "three"]
  index          = length(["zero", "foo", "bar"]) - (length(["one", "two", "three"]) - 1)
}
# => ["zero", "one", "two", "three"]
