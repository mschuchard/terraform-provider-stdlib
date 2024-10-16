# Return the square root of 4
data "stdlib_sqrt" "four" {
  param = 4
}
# => 2

# Return the square root of 0
data "stdlib_sqrt" "zero" {
  param = 0
}
# => 0

# Return the square root of 2
data "stdlib_sqrt" "two" {
  param = 2
}
# => 1.4142135623730951
