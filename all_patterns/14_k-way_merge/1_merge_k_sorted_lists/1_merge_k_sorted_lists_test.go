package _1_merge_k_sorted_lists

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func Test_mergeSortedLists(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, []int{1, 2, 3, 3, 4, 6, 6, 7, 8}},
		{[][]int{{5, 8, 9}, {1, 7}}, []int{1, 5, 7, 8, 9}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			lists := make([]*single.ListNode, len(tt.input))
			for i := range tt.input {
				lists[i] = single.ArrToLinkedList(tt.input[i])
			}

			got := single.LinkedListToArr(mergeSortedLists(lists))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
