package single

import (
	"fmt"
	"github.com/babadro/algorithms-go/slices"
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
		head := ArrToLinkedList(arr)
		node := head
		fmt.Printf("\nList %d: ", i+1)
		for node != nil {
			fmt.Printf("%d ", node.Val)
			node = node.Next
		}
	}
}

func TestLinkedListToArr(t *testing.T) {
	arrays := [][]int{
		{},
		{0},
		{0, 1},
		{1, 2, 3},
	}
	for i, input := range arrays {
		head := ArrToLinkedList(input)
		output := LinkedListToArr(head)
		if !slices.IntSlicesAreEqual(input, output) {
			t.Errorf("case#%d want %v, got %v", i+1, input, output)
		}
	}
}
