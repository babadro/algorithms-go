package single

import (
	"fmt"
	"testing"
)

func TestArrToLinkedList(t *testing.T) {
	arrays := [][]int{
		{},
		{0},
		{0, 1},
		{1, 2, 3},
	}

	for i, arr := range arrays {
		head := arrToLinkedList(arr)
		node := head
		fmt.Printf("\nList %d: ", i+1)
		for node != nil {
			fmt.Printf("%d ", node.Val)
			node = node.Next
		}
	}
}
