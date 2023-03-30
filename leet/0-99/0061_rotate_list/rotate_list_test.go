package _061_rotate_list

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4, 5}, k: 2, want: []int{4, 5, 1, 2, 3},
		},
		{input: []int{0, 1, 2}, k: 4, want: []int{2, 0, 1}},
		{input: []int{1}, k: 0, want: []int{1}},
		{input: []int{1, 2, 3}, k: 3, want: []int{1, 2, 3}},
		{input: []int{1, 2}, k: 0, want: []int{1, 2}},
		{input: []int{}, k: 0, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := singly.ArrToLinkedList(tt.input)
			got := singly.LinkedListToArr(rotateRight2(head, tt.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
