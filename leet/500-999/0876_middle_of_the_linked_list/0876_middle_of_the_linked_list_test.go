package _876_middle_of_the_linked_list

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
		{[]int{1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			head := single.ArrToLinkedList(tt.arr)
			got := middleNode(head)

			want := single.ArrToLinkedList(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("middleNode() = %v, want %v", got, want)
			}
		})
	}
}
