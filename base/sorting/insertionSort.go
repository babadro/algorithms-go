package sorting

// it has a O(N) best case behavior (when already sorted).
func InsertionSort(arr []int) {
	var ele, j int
	for i := 1; i < len(arr); i++ {
		ele = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > ele {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = ele
	}
}
