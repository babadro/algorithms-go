package _40

import "testing"

func Test_largestPathValue(t *testing.T) {
	tests := []struct {
		colors string
		edges  [][]int
		want   int
	}{
		{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
		{"a", [][]int{{0, 0}}, -1},
		{"hhqhuqhqff", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}}, 3},
		{"eeyyeeyeye", [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {4, 6}, {5, 7}, {6, 8}, {8, 9}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.colors, func(t *testing.T) {
			if got := largestPathValue(tt.colors, tt.edges); got != tt.want {
				t.Errorf("largestPathValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
