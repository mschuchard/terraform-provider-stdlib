# Return the remainder of 4 / 2
data "stdlib_mod" "zero" {
  dividend = 4
  divisor  = 2
}
# => 0

# Return the remainder of 5 / 3
data "stdlib_mod" "integer" {
  dividend = 5
  divisor  = 3
}
# => 2

# Return the remainder of 10 / 3.5
data "stdlib_mod" "decimal" {
  dividend = 10
  divisor  = 4.75
}
# => 0.5
