package _0220_contains_duplicate_iii

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		t    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{[]int{1, 3, 6, 2}, 1, 2, true},
		{[]int{8, 7, 15, 1, 6, 1, 9, 15}, 1, 3, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate2(tt.nums, tt.k, tt.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
