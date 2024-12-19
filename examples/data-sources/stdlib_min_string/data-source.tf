# Return the minimum string (first by lexical ordering) from the element(s) of a list:
data "stdlib_min_string" "count" {
  param = ["zero", "one", "two", "three", "four", "five", "six", "seven"]
}
# result => five

data "stdlib_min_string" "alphabet" {
  param = ["alpha", "beta", "gamma", "delta", "epsilon"]
}
# result => alpha
