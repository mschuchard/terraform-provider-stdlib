# Delete the elements of a list from 0 to 1 inclusive:
provider::stdlib::elements_delete(["zero", "one", "two", "three"], 0, 1)
# result => ["two", "three"]

# Delete the elements of a list from 2 to 3 inclusive:
provider::stdlib::elements_delete(["zero", "one", "two", "three"], 2, 3)
# result => ["zero", "one"]