package _055_jump_game

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 0, 2}, false},
		{[]int{1}, true},
		{[]int{0}, true},
		{[]int{1, 0}, true},
		{[]int{3, 0, 0, 0}, true},
		{[]int{0, 0}, false},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 5, 0, 0}, true},
	}

	for i, c := range cases {
		if fact := canJump(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
	}
}
