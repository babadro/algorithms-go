package _1792_maximum_average_pass_ratio

import "testing"

func Test_maxAverageRatio(t *testing.T) {
	tests := []struct {
		name          string
		classes       [][]int
		extraStudents int
		want          float64
	}{
		{
			name:          "1",
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			want:          0.78333,
		},
		{
			name:          "2",
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			want:          0.53485,
		},
		{
			name:          "3",
			classes:       [][]int{{1, 2}},
			extraStudents: 1,
			want:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAverageRatio(tt.classes, tt.extraStudents); got != tt.want {
				t.Errorf("maxAverageRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
