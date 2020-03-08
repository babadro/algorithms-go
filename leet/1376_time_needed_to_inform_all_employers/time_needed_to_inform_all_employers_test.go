package _1376_time_needed_to_inform_all_employers

import "testing"

func TestNumOfMinutes(t *testing.T) {
	cases := []struct {
		n          int
		headId     int
		manager    []int
		informTime []int
		expected   int
	}{
		{1, 0, []int{-1}, []int{0}, 0},
		{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}, 1},
		{7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}, 21},
		{15, 0, []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, 3},
		{4, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}, 1076},
	}

	for i, c := range cases {
		if fact := numOfMinutes(c.n, c.headId, c.manager, c.informTime); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
