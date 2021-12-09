package _2097_valid_arrangement_of_pairs

import (
	"testing"
)

func Test_validArrangement(t *testing.T) {

	tests := []struct {
		pairs [][]int
		want  [][]int
	}{
		{[][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}, nil},
		//{[][]int{{0, 1}, {1, 2}, {2, 0}}, nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := validArrangement(tt.pairs)
			t.Log(got)
		})
	}
}

func Test_EulerianPath(t *testing.T) {
	tests := []struct {
		pairs [][]int
		start int
	}{
		{[][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}, 11},
		{[][]int{{0, 1}, {1, 2}, {2, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := eulerianPath(edjesToAdj(tt.pairs), tt.start)
			t.Log(got)
		})
	}
}

func edjesToAdj(edges [][]int) map[int][]int {
	m := make(map[int][]int)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
	}

	return m
}
