### 2.3.0 (Next)
- Add `truncate` function.

### 2.2.2
- Update `exp`, `mod`, and `sqrt` functions' input parameter type from `float64` to `number`.
- Update `sqrt` return type to `number`.
- Verify `round` return does not under/overflow.
- Permit `number_of_characters` to be equal to `string` parameter length for `last_char`.
- Permit empty string input value for `cut`.
- `sort_list_number` parameter element type updated from `float32`to `float64`.
- Permit `number_of_elements` to equal length of `list` in `last_element`.

### 2.2.1
- Dynamic input emptiness error message now logs accurate parameter position.
- Fix return type for `coalesce_map` function.
- Optimize `has_value` function logic.
- Modify missing key for `key(s)_delete` function from error to warning.
- Protect against overflows in `combinations` and `factorial` functions.
- Fix input parameter reflection error check in `sqrt` function.
- Improve error messages for invalid input parameters in `combinations` function.
- Validate divisor is not 0 in `mod` function.
- Update `round` function input parameter type from `float64` to `number`.

### 2.2.0
- Rename former `sort_list` to `sort_list_string`.
- Update `repeat` function to also accept string input.
- Add `combinations`, `count`, and `sort_list_number` functions.

### 2.1.1
- Improve unit testing.
- Improve `replace` descriptions and parameter validation.
- Improve `empty` function behavior.
- Optimize function error returns.
- Add experimental `compact_map` function.

### 2.1.0
- Add `coalesce_map`, `compact_map`, `elements_delete`, `factorial`, and `repeat` functions.

### 2.0.3
- Implement `multiple` type custom functions.
- Improve previously implemented custom and data functions.

### 2.0.2
- Implement `map` type custom functions.

### 2.0.1
- Implement `slice` type custom functions.

### 2.0.0
- Implement `string` and `number` type custom functions.

### 1.6.0
- Add `product` and `cut` functions.
- Optimize data source `id` where possible (may cause superficial plan changes to existing states).

### 1.5.1
- Refine attribute error messages.
- Coerce result of `round` function to integer.
- Do not coerce `id` of `min` and `max` functions to integer.

### 1.5.0
- Do not coerce `number` type ID to `string` type ID (may cause superficial plan changes to existing states).
- Add `exp`, `mod`, `round`, and `sqrt` functions.

### 1.4.1
- Add `sorted` parameter to `list_index` function.
- Add `end_index` parameter to `replace` function.
- Fix `end_index` auto-deduction in `replace` function.
- Add config validation to `insert` and `replace` functions.

### 1.4.0
- Add `insert`, `replace`, and `sort_list` functions.

### 1.3.0
- Add `list_index`, `max_number`, `min_number`, `max_string`, and `min_string` functions.

### 1.2.0
- Add `empty`, `compare_list`, and `last_element` functions.

### 1.1.1
- Add `all` parameter to `has_keys` and `has_values`.

### 1.1.0
- Add `num_chars` parameter to `last_char`.
- Add `equal_map`, `keys_delete`, `has_keys`, and `has_values` functions.
- Miscellaneous fixes and improvements to collection functions.

### 1.0.0
- Initial Release
