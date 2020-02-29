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
		{[]int{1, 3, 10}, []int{10}, 5, false},
		{[]int{1, 11}, []int{1, 5, 9, 11}, 1, true},
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
		{[]int{1, 3, 10}, []int{10}, 10},
		{[]int{1, 11}, []int{1, 5, 9, 11}, 1},
		{[]int{1, 11}, []int{3}, 8},
	}

	for i, c := range cases {
		if factRadius := findRadius(c.houses, c.heaters); factRadius != c.expectedRadius {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expectedRadius, factRadius)
		}
	}
}
