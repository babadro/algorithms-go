package graph

import (
	"reflect"
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	tests := []struct {
		name         string
		graphBuilder func() *Graph
		start        int
		want         []int
	}{
		{
			"1",
			func() *Graph {
				g := New(4)

				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 0)
				g.AddEdge(0, 2)
				g.AddEdge(2, 3)
				g.AddEdge(3, 3)

				return g
			},
			2,
			[]int{2, 0, 3, 1},
		},
		{
			"2",
			func() *Graph {
				g := New(6)

				g.AddEdge(5, 2)
				g.AddEdge(5, 0)
				g.AddEdge(4, 0)
				g.AddEdge(4, 1)
				g.AddEdge(2, 3)
				g.AddEdge(3, 1)

				return g
			},
			2,
			[]int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graphBuilder()
			var got []int
			f := func(v int) {
				got = append(got, v)
			}
			g.BFS(tt.start, f)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_TopologicalSort(t *testing.T) {
	tests := []struct {
		name         string
		graphBuilder func() *Graph
		want         []int
	}{
		{
			"1",
			func() *Graph {
				g := New(6)

				g.AddEdge(5, 2)
				g.AddEdge(5, 0)
				g.AddEdge(4, 0)
				g.AddEdge(4, 1)
				g.AddEdge(2, 3)
				g.AddEdge(3, 1)

				return g
			},
			[]int{5, 4, 2, 3, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graphBuilder()
			var got []int
			f := func(v int) {
				got = append(got, v)
			}
			g.TopologicalSort(f)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
