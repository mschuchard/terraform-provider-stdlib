# Return the rounding of 1.2:
provider::stdlib::round(1.2)
# result => 1

# Return the rounding of 1.8:
provider::stdlib::round(1.8)
# result => 2

# Return the rounding of 1.5:
provider::stdlib::round(1.5)
# result => 2

# Return the even tiebreak rounding of 2.5:
provider::stdlib::round(2.5, true)
# result => 2

# Return the normal tiebreak rounding of 2.5:
provider::stdlib::round(2.5, false)
# result => 3