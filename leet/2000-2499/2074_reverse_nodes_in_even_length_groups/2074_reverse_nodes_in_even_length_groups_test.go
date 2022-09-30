package _2074_reverse_nodes_in_even_length_groups

import (
	"fmt"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"github.com/stretchr/testify/require"
)

func Test_reverseEvenLengthGroups(t *testing.T) {
	tests := []struct {
		input, want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		//{[]int{1, 2, 3}, []int{1, 3, 2}},
		//{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
		//{[]int{1, 2, 3, 4, 5}, []int{1, 3, 2, 5, 4}},
		//{[]int{1, 2, 3, 4, 5, 6}, []int{1, 3, 2, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 2, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 3, 2, 4, 5, 6, 8, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 3, 2, 4, 5, 6, 9, 8, 7}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			head := single.ArrToLinkedList(tt.input)
			got := reverseEvenLengthGroups(head)
			gotArr := single.LinkedListToArr(got)

			require.Equal(t, tt.want, gotArr)
		})
	}
}
