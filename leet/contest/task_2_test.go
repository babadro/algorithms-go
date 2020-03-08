package contest

import "testing"

func TestNumTimesAllBlue(t *testing.T) {
	cases := []struct {
		light    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{2, 1, 3, 5, 4}, 3},
		{[]int{3, 2, 4, 1, 5}, 2},
		{[]int{4, 1, 2, 3}, 1},
		{[]int{2, 1, 4, 3, 6, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 6},
	}

	for i, c := range cases {
		if fact := numTimesAllBlue(c.light); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
