package _35

import "testing"

func Test_minAbsoluteSumDiff(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "1",
			nums1: []int{1, 7, 5},
			nums2: []int{2, 3, 5},
			want:  3,
		},
		{
			name:  "2",
			nums1: []int{2, 4, 6, 8, 10},
			nums2: []int{2, 4, 6, 8, 10},
			want:  0,
		},
		{
			name:  "3",
			nums1: []int{1, 10, 4, 4, 2, 7},
			nums2: []int{9, 3, 5, 1, 7, 4},
			want:  20,
		},
		{
			name:  "4",
			nums1: []int{2},
			nums2: []int{5},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsoluteSumDiff(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
