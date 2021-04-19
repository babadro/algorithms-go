package _619_mean_of_array_after_removing_some_elements

import (
	"math"
	"sort"
)

// 100% 100%
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	start, end := int((float64(n)/100)*5), n-int((float64(n)/100)*5)

	arr = arr[start:end]

	sum := float64(0)
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}

	res := sum / float64(len(arr))

	return math.Ceil(res*100000) / 100000
}

// after refac:
func trimMean2(arr []int) float64 {
	sort.Ints(arr)
	sum, count, n := float64(0), 0, len(arr)

	for i := n / 20; i < n-(n/20); i++ {
		sum += float64(arr[i])
		count++
	}

	res := sum / float64(count)

	return math.Ceil(res*100000) / 100000
}
