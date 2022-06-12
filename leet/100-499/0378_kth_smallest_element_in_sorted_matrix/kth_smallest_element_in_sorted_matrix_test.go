package _378_kth_smallest_element_in_sorted_matrix

import "testing"

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		k      int
		want   int
	}{
		{
			name: "1",
			matrix: [][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			k:    8,
			want: 13,
		},
		{
			name: "2",
			matrix: [][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			k:    4,
			want: 10,
		},
		{
			name: "3",
			matrix: [][]int{
				{1, 5, 9, 16},
				{10, 11, 13, 18},
				{12, 13, 15, 20},
				{14, 15, 17, 21},
			},
			k:    10,
			want: 15,
		},
		{
			name: "4",
			matrix: [][]int{
				{1, 5, 9, 16},
				{10, 11, 13, 18},
				{12, 13, 15, 20},
				{14, 15, 17, 21},
			},
			k:    15,
			want: 20,
		},
		{
			name: "5",
			matrix: [][]int{
				{1},
			},
			k:    1,
			want: 1,
		},
		{
			name: "6",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			k:    3,
			want: 3,
		},
		{
			name: "7",
			matrix: [][]int{
				{1, 3, 5},
				{6, 7, 12},
				{11, 14, 14},
			},
			k:    3,
			want: 5,
		},
		{
			name: "8",
			matrix: [][]int{
				{2, 6, 8},
				{3, 7, 10},
				{5, 8, 11},
			},
			k:    5,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.matrix, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
