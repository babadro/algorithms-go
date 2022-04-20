package _2_remove_duplicates

// leetCode 26
func removeDuplicates(sortedArr []int) int {
	nextNonDuplicate := 1
	for i := 0; i < len(sortedArr); i++ {
		if sortedArr[nextNonDuplicate-1] != sortedArr[i] {
			sortedArr[nextNonDuplicate] = sortedArr[i]
			nextNonDuplicate++
		}
	}

	return nextNonDuplicate
}

// similar question - remove key from unsorted arr in place. Don't change order of result array.
func remove(arr []int, key int) int {
	idx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != key {
			arr[idx] = arr[i]
			idx++
		}
	}

	return idx
}
