package _797_all_paths_from_source_to_target

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{"1", [][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{"2", [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
		{"3", [][]int{{1}, {}}, [][]int{{0, 1}}},
		{"4", [][]int{{1, 2, 3}, {2}, {3}, {}}, [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}}},
		{"5", [][]int{{1, 3}, {2}, {3}, {}}, [][]int{{0, 1, 2, 3}, {0, 3}}},
		{"6", [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}, [][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget2(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
