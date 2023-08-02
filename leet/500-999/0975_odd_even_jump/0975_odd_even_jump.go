package _975_odd_even_jump

import (
	"math"
	"sort"
)

// bnsrg #hard #ntu
func oddEvenJumps(arr []int) int {
	n := len(arr)
	l := []int{n, n - 1, n + 1}
	arr = append(arr, math.MinInt32, math.MaxInt32)
	evenOK := make([]uint8, n+2)
	oddOK := make([]uint8, n+2)
	evenOK[n-1] = 1
	oddOK[n-1] = 1

	res := 1
	for i := n - 2; i >= 0; i-- {
		j := sort.Search(len(l), func(k int) bool {
			return arr[l[k]] >= arr[i]
		})
		// A later element is equal to the current
		// Same target for even and odd jumps
		if arr[l[j]] == arr[i] {
			evenOK[i] = oddOK[l[j]]
			oddOK[i] = evenOK[l[j]]
			l[j] = i
		} else {
			// Element is larger than the current
			evenOK[i] = oddOK[l[j-1]]
			oddOK[i] = evenOK[l[j]]

			// Make space and insert into sorted list
			l = append(l, 0)
			copy(l[j+1:], l[j:])
			l[j] = i
		}
		res += int(oddOK[i])
	}
	return res
}
