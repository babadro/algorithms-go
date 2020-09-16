package _328_odd_even_linked_list

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{"2", []int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
		{"3", []int{1}, []int{1}},
		{"4", []int{1, 2}, []int{1, 2}},
		{"5", []int{1, 2, 3}, []int{1, 3, 2}},
		{"6", []int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := single.ArrToLinkedList(tt.arr)
			gotLinkedList := oddEvenList(head)
			resultArr := single.LinkedListToArr(gotLinkedList)
			if !reflect.DeepEqual(resultArr, tt.want) {
				t.Errorf("resultArr = %v, want %v", resultArr, tt.want)
			}
		})
	}
}
