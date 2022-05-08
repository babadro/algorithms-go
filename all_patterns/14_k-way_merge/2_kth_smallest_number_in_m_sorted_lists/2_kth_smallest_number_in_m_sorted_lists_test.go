package __kth_smallest_number_in_m_sorted_lists

import "testing"

func Test_findKthSmallest(t *testing.T) {
	tests := []struct {
		lists [][]int
		k     int
		want  int
	}{
		{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, 5, 4},
		{[][]int{{5, 8, 9}, {1, 7}}, 3, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findKthSmallest(tt.lists, tt.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
