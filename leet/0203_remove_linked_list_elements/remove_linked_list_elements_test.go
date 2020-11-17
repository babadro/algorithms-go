package _203_remove_linked_list_elements

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		val   int
		want  []int
	}{
		{"1", []int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{"2", []int{}, 6, nil},
		{"3", []int{6}, 6, nil},
		{"4", []int{6, 1, 1, 1}, 6, []int{1, 1, 1}},
		{"5", []int{6, 1, 1, 1, 6, 6, 6}, 6, []int{1, 1, 1}},
		{"5", []int{6, 6, 6}, 6, nil},
		{"5", []int{6, 6, 6, 1, 2, 6, 3, 4, 6, 6}, 6, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := single.ArrToLinkedList(tt.input)
			if got := removeElements2(head, tt.val); !reflect.DeepEqual(single.LinkedListToArr(got), tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
