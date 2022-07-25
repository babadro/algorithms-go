package _092_reverse_linked_list_ii

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		inputArr    []int
		left, right int
		want        []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{1}, 1, 1, []int{1}},
		{[]int{1, 2}, 1, 2, []int{2, 1}},
		{[]int{1, 2}, 2, 2, []int{1, 2}},
		{[]int{1, 2}, 1, 1, []int{1, 2}},
		{[]int{1, 2, 3, 4}, 3, 4, []int{1, 2, 4, 3}},
		{[]int{1, 2, 3, 4}, 1, 2, []int{2, 1, 3, 4}},
		{[]int{5}, 1, 1, []int{5}},
		{[]int{2, 3, 4, 5, 6}, 2, 4, []int{2, 5, 4, 3, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.inputArr), func(t *testing.T) {
			head := single.ArrToLinkedList(tt.inputArr)
			got := single.LinkedListToArr(reverseBetween(head, tt.left, tt.right))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
