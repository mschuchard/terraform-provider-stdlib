# Return the last element of a list:
provider::stdlib::last_element(["h", "e", "l", "l", "o"])
# result => ["o"]

# Return the last three elements of a list (reverse slice):
provider::stdlib::last_element(["h", "e", "l", "l", "o"], 3)
# result => ["l", "l", "o"]