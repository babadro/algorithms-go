package _475_heaters

import "testing"

func TestCheckRadius(t *testing.T) {
	cases := []struct {
		houses   []int
		heaters  []int
		radius   int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{2}, 1, true},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1, true},
		{[]int{1, 3, 10}, []int{10}, 9, true},
		{[]int{1, 11}, []int{1, 5, 9, 10}, 0, false},
	}

	for i, c := range cases {
		if fact := checkRadius(c.houses, c.heaters, c.radius); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
	}
}

func TestFindRadius(t *testing.T) {
	cases := []struct {
		houses         []int
		heaters        []int
		expectedRadius int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1, 3, 10}, []int{10}, 9},
		{[]int{1, 11}, []int{1, 5, 9, 11}, 0},
		{[]int{1, 11}, []int{3}, 8},
		{[]int{1, 5}, []int{10}, 9},
		{[]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}, []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}, 161834419},
	}

	for i, c := range cases {
		if factRadius := findRadius(c.houses, c.heaters); factRadius != c.expectedRadius {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expectedRadius, factRadius)
		}
	}
}

/*
[282475249,622650073,984943658,144108930,470211272,101027544,457850878,458777923]
[823564440,115438165,784484492,74243042,114807987,137522503,441282327,16531729,823378840,143542612]
output: 176302675
expected: 161834419

*/
