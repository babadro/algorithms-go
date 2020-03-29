package _912_sort_an_array

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

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

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIdx := partition(arr, low, high)

		quickSort(arr, low, pivotIdx-1)
		quickSort(arr, pivotIdx+1, high)
	}
}
