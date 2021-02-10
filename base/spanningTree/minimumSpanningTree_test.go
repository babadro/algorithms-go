package spanTree

import (
	"reflect"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	tests := []struct {
		name    string
		V       int
		edges   [][3]int
		wantMst []Edge
	}{
		{
			name: "1",
			V:    9,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMst := KruskalMST(tt.V, tt.edges); !reflect.DeepEqual(gotMst, tt.wantMst) {
				t.Errorf("KruskalMST() = %v, want %v", gotMst, tt.wantMst)
			}
		})
	}
}
