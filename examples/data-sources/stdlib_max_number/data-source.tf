# Return the maximum number from the element(s) of a list:
data "stdlib_max_number" "fibonacci" {
  param = [0, 1, 1, 2, 3, 5, 8, 13]
}
# result => 13
