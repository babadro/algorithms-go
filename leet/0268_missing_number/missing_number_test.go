package _268_missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{0}, 0},
		{[]int{0, 1}, 0},
		{[]int{0, 2}, 1},
		{[]int{10, 9, 8, 3, 2, 4, 5, 6, 1}, 7},
	}

	for i, c := range cases {
		if fact := missingNumber(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
