package _025_reverse_nodes_in_k_group

import (
	"fmt"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"github.com/stretchr/testify/require"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{[]int{1, 2, 3}, 3, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3}, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			head := single.ArrToLinkedList(tt.input)

			got := single.LinkedListToArr(reverseKGroup(head, tt.k))

			require.Equal(t, tt.want, got)
		})
	}
}
