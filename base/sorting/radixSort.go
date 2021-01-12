package sorting

// todo 1 make unit test and understand
// it's not a comparison algorithm and takes time O(n lg u) (which is MUCH faster than Quicksort). Van Embde Boas if you want to get really performant with these concepts with time O(n lg lg u).

func maxElement(arr []int, n int) int {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func countSortHelper(arr []int, n, exp int) {
	output := make([]int, n)
	count := make([]int, 10)

	i := 0
	for ; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i = 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i = n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i = 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func Radix(arr []int, n int) {
	m := maxElement(arr, n)

	for exp := 1; m/exp > 0; exp *= 10 {
		countSortHelper(arr, n, exp)
	}
}
