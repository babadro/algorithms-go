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
		{
			"3",
			func() *Graph {
				arr := [][]int{
					0: {1, 2, 3},
					1: {4, 5, 6},
					2: nil,
					3: nil,
					4: nil,
					5: nil,
					6: {7},
					7: nil,
				}

				return &Graph{adj: arr}
			},
			0,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
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
		{
			"2",
			func() *Graph {
				g := New(6)

				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 3)
				g.AddEdge(3, 4)
				g.AddEdge(4, 5)

				return g
			},
			[]int{0, 1, 2, 3, 4, 5},
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

func TestGraph_AllPathsFromSourceToTarget(t *testing.T) {
	tests := []struct {
		name   string
		graph  [][]int
		source int
		target int
		want   [][]int
	}{
		{"1", [][]int{{1, 2}, {3}, {3}, {}}, 0, 3, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{"2", [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, 0, 4, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
		{"3", [][]int{{1}, {}}, 0, 1, [][]int{{0, 1}}},
		{"4", [][]int{{1, 2, 3}, {2}, {3}, {}}, 0, 3, [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}}},
		{"5", [][]int{{1, 3}, {2}, {3}, {}}, 0, 3, [][]int{{0, 1, 2, 3}, {0, 3}}},
		{"6", [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}, 0, 7, [][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Constructor(tt.graph)

			var res [][]int
			f := func(path []int) {
				item := make([]int, len(path))
				copy(item, path)
				res = append(res, item)
			}

			g.AllPathsFromSourceToTarget(tt.source, tt.target, f)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", res, tt.want)
			}
		})
	}
}

func Benchmark_AllPaths(b *testing.B) {
	var res [][]int

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		res = nil
		f := func(path []int) {
			item := make([]int, len(path))
			copy(item, path)
			res = append(res, item)
		}
		arr := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
		g := Constructor(arr)
		b.StartTimer()

		g.AllPathsFromSourceToTarget(0, g.V()-1, f)
	}

	_ = res
}
