# Return the list with beginning value replaced.
data "stdlib_replace" "begin" {
  list_param     = ["foo", "bar", "two", "three"]
  replace_values = ["zero", "one"]
  index          = 0
}
# => ["zero", "one", "two", "three"]

# Return the list with middle values replaced.
data "stdlib_replace" "replace" {
  list_param     = ["zero", "foo", "bar", "baz", "four", "five"]
  replace_values = ["one", "two", "three"]
  index          = 1
}
# => ["zero", "one", "two", "three", "four", "five"]

# Return the list with middle values replaced and zeroed.
data "stdlib_replace" "zeroed" {
  list_param     = ["zero", "foo", "bar", "four", "five"]
  replace_values = ["one"]
  index          = 1
  end_index      = 2
}
# => ["zero", "one", "four", "five"]

# Return the list with terminating values replaced.
data "stdlib_replace" "append" {
  list_param     = ["zero", "foo", "bar", "baz"]
  replace_values = ["one", "two", "three"]
  index          = length(["zero", "foo", "bar", "baz"]) - (length(["one", "two", "three"]))
}
# => ["zero", "one", "two", "three"]
