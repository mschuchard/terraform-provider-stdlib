# Return the last index of a substring instance within a string:
provider::stdlib::last_index("terra terraform", "terra")
# result => 6

# Return the last index of a substring instance absent from a string:
provider::stdlib::last_index("terra terraform", "vault")
# result => -1