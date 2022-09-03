package _1

import "fmt"

func bioHazard(n int32, allergic []int32, poisonous []int32) int64 {
	poisDict := make(map[[2]int32]bool)
	for i := range allergic {
		a, b := sort(allergic[i], poisonous[i])

		poisDict[[2]int32{a, b}] = true
	}

	var res int64
	l := int32(1)
	for r := int32(1); r <= n; r++ {
		if r == 3 {
			fmt.Sprintf("")
		}

		for i := r; i >= l; i-- {
			if !poisDict[[2]int32{i, r}] {
				res++
				if i == r {
					fmt.Printf("(%d)", i)
				} else {
					fmt.Printf("(%d,%d)", i, r)
				}

				continue

			}

			l = i + 1

			break
		}
	}

	return res
}

func sort(a, b int32) (int32, int32) {
	if a > b {
		return b, a
	}

	return a, b
}
