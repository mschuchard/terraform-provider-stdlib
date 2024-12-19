# Return the base e exponential of 0
data "stdlib_exp" "zero" {
  param = 0
}
# result => 1

# Return the base e exponential of 1.0986122
data "stdlib_exp" "decimal" {
  param = 1.0986122
}
# result => 2.9999997339956828
