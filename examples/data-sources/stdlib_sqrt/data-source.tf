# Return the square root of 4
data "stdlib_sqrt" "four" {
  param = 4
}
# result => 2

# Return the square root of 0
data "stdlib_sqrt" "zero" {
  param = 0
}
# result => 0

# Return the square root of 2
data "stdlib_sqrt" "two" {
  param = 2
}
# result => 1.4142135623730951
