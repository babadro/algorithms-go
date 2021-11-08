package _1830_minimum_number_of_operations_to_make_string_sorted

import "fmt"

// todo 1
func makeStringSorted(s string) int {
	b := []byte(s)
	res, n := 0, len(b)

Loop:
	for {
		for i := n - 1; i > 0; i-- {
			if b[i] < b[i-1] {
				res++
				j := i
				for ; j < n-1 && b[j+1] < b[i-1]; j++ {
				}

				b[i-1], b[j] = b[j], b[i-1]
				for idx := i; idx < (n+idx)/2; idx++ {
					b[idx], b[n-1-idx] = b[n-1-idx], b[idx]
				}

				fmt.Println(string(b))
				continue Loop
			}
		}

		return res
	}
}
