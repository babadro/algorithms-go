package _1738_find_kth_largest_xor_coordinate_value

import "testing"

func Test_kthLargestValue(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		k      int
		want   int
	}{
		{
			"1",
			[][]int{{5, 2}, {1, 6}},
			1,
			7,
		},
		{
			"2",
			[][]int{{5, 2}, {1, 6}},
			2,
			5,
		},
		{
			"3",
			[][]int{{5, 2}, {1, 6}},
			3,
			4,
		},
		{
			"4",
			[][]int{{5, 2}, {1, 6}},
			4,
			0,
		},
		{
			"5",
			[][]int{{2}},
			1,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.matrix, tt.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
