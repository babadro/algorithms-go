package _31

import (
	"testing"
)

func Test_countRestrictedPaths(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"1", 5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}, 3},
		{"2", 7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}}, 1},
		{"312 size input", 312, bigArrOfEdges, 683659},       // I'm not sure about 683659
		{"422 size input", 422, bigArrOfEdges_422, 13944179}, // todo 1 fails
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths2(tt.n, tt.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
