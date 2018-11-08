package array

// Equal ArrayEqual
func Equal(array1 []string, array2 []string) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}
