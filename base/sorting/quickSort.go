package sorting

// tptl. todo 1 iterative
func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func QuickSortRecursive(arr []int, low, high int) {
	if low < high {
		pivotIDx := partition(arr, low, high)

		QuickSortRecursive(arr, low, pivotIDx-1)
		QuickSortRecursive(arr, pivotIDx+1, high)
	}
}
