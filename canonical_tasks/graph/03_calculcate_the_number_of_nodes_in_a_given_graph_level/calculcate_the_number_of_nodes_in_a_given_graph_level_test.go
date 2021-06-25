package _3_calculcate_the_number_of_nodes_in_a_given_graph_level

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		g     [][]int
		level int
		want  int
	}{
		{
			g: [][]int{
				0: {1, 2},
				1: {3},
				2: {3, 4},
				3: {5},
				4: nil,
				5: nil,
			},
			level: 4,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.g, tt.level), func(t *testing.T) {
			if got := CalculateBFS(tt.g, tt.level); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
