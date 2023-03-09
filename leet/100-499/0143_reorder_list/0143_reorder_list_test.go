package _143_reorder_list

import (
	"fmt"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
	"github.com/stretchr/testify/require"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		arr, want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			head := singly.ArrToLinkedList(tt.arr)
			reorderList(head)

			gotArr := singly.LinkedListToArr(head)

			require.Equal(t, tt.want, gotArr)
		})
	}
}
