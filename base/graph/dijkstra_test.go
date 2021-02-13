package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name         string
		edges        [][3]int
		vertexCount  int
		start        int
		wantDistance []int
	}{
		{
			name: "1",
			edges: [][3]int{
				{0, 1, 4},
				{1, 2, 2},
				{0, 2, 4},
				{2, 3, 1},
				{2, 5, 3},
				{2, 4, 6},
				{5, 4, 2},
				{3, 4, 3},
			},
			vertexCount:  6,
			start:        0,
			wantDistance: []int{0, 4, 4, 5, 8, 7},
		},
		{
			name:        "2",
			vertexCount: 9,
			edges: [][3]int{
				{0, 1, 4},
				{0, 7, 8},
				{1, 7, 11},
				{1, 2, 8},
				{7, 8, 7},
				{7, 6, 1},
				{2, 8, 2},
				{8, 6, 6},
				{2, 3, 7},
				{2, 5, 4},
				{6, 5, 2},
				{3, 5, 14},
				{3, 4, 9},
				{5, 4, 10},
			},
			start:        0,
			wantDistance: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adjMatrix := EdgesToAdjMatrix(tt.vertexCount, tt.edges)

			if gotDistance := Dijkstra(adjMatrix, tt.vertexCount, tt.start); !reflect.DeepEqual(gotDistance, tt.wantDistance) {
				t.Errorf("Dijkstra() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
