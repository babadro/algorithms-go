package graph

import "testing"

func TestUndirectedGraphHasEulerianPathOrCycle(t *testing.T) {
	tests := []struct {
		edges        [][2]int
		n            int
		wantHasPath  bool
		wantHasCycle bool
	}{
		{
			[][2]int{{0, 1}, {0, 5}, {0, 4}, {1, 5}, {1, 3}, {1, 2}, {2, 4}, {2, 3}, {3, 4}, {4, 5}},
			6, false, false,
		},
		{
			[][2]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {3, 4}},
			5, true, true,
		},
		{
			[][2]int{{0, 1}, {0, 4}, {0, 5}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 5}},
			6, true, false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			g := New(tt.n)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
				g.AddEdge(edge[1], edge[0])
			}
			gotHasPath, gotHasCycle := UndirectedGraphHasEulerianPathOrCycle(*g)
			if gotHasPath != tt.wantHasPath {
				t.Errorf("UndirectedGraphHasEulerianPathOrCycle() gotHasPath = %v, want %v", gotHasPath, tt.wantHasPath)
			}
			if gotHasCycle != tt.wantHasCycle {
				t.Errorf("UndirectedGraphHasEulerianPathOrCycle() gotHasCycle = %v, want %v", gotHasCycle, tt.wantHasCycle)
			}
		})
	}
}
