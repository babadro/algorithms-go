package slices

func IntSlicesAreEqual(slice1, slice2 []int) bool {
	length := len(slice1)
	if length != len(slice2) {
		return false
	}
	for i := 0; i < length; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func SliceOfIntSlicesAreEqual(slice1, slice2 [][]int) bool {
	length := len(slice1)
	if length != len(slice2) {
		return false
	}
	for i := 0; i < length; i++ {
		if !IntSlicesAreEqual(slice1[i], slice2[i]) {
			return false
		}
	}
	return true
}
