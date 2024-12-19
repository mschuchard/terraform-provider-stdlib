# Return the maximum string (last by lexical ordering) from the element(s) of a list:
data "stdlib_max_string" "count" {
  param = ["zero", "one", "two", "three", "four", "five", "six", "seven"]
}
# result => zero

data "stdlib_max_string" "alphabet" {
  param = ["alpha", "beta", "gamma", "delta", "epsilon"]
}
# result => gamma
