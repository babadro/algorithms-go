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
		{10, 3, []int{8, 9, 8, -1, 7, 1, 2, 0, 3, 0}, []int{224, 943, 160, 909, 0, 0, 0, 643, 867, 722}, 3665},
		{11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}, 2560},
	}

	for i, c := range cases {
		if fact := numOfMinutes(c.n, c.headId, c.manager, c.informTime); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}

/*
11
4
[5,9,6,10,-1,8,9,1,9,3,4]
[0,213,0,253,686,170,975,0,261,309,337]
*/
