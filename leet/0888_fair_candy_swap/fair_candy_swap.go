package _0888_fair_candy_swap

// tptl

// Brutforce 18% 59%
func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	for _, v := range A {
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	for _, x := range A {
		for _, y := range B {
			if x-y == (sumA-sumB)/2 {
				return []int{x, y}
			}
		}
	}

	return nil
}

// Best solution
func fairCandySwap2(A []int, B []int) []int {
	sumA, sumB := 0, 0
	m := make(map[int]bool)
	for _, v := range A {
		m[v] = true
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	delta := (sumA - sumB) / 2 // a-b = (sumA - sumB)/2

	for _, b := range B {
		if _, ok := m[b+delta]; ok {
			return []int{b + delta, b}
		}
	}

	return nil
}
