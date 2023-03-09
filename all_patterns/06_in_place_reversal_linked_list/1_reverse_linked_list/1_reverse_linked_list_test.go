package __reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func Test_reverseLinkedList(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{2, 4, 6, 8, 10}, []int{10, 8, 6, 4, 2}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			head := singly.ArrToLinkedList(tt.input)
			got := singly.LinkedListToArr(reverseLinkedList(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
