package graph

import (
	"encoding/json"
	"github.com/babadro/algorithms-go/base/graph/bigInput"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name         string
		edges        [][3]int
		_1_indexed   bool
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
			_1_indexed:   false,
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
			_1_indexed:   false,
			start:        0,
			wantDistance: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
		},
		{name: "312 size input", vertexCount: 312, edges: bigInput.Edges312, _1_indexed: true, wantDistance: bigInput.Res312},
		{name: "422 size input", vertexCount: 422, edges: bigInput.Edges422, _1_indexed: true, wantDistance: bigInput.Res422},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adjMatrix := EdgesToAdjMatrix(tt.vertexCount, tt.edges, tt._1_indexed)

			if gotDistance := Dijkstra(adjMatrix, tt.vertexCount, tt.start); !reflect.DeepEqual(gotDistance, tt.wantDistance) {
				gotDistanceBytes, err := json.Marshal(gotDistance)
				if err != nil {
					t.Error(err)
				}

				if err = ioutil.WriteFile(tt.name, gotDistanceBytes, 0644); err != nil {
					t.Error(err)
				}

				t.Errorf("Dijkstra() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
