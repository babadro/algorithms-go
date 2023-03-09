package singly

import (
	"fmt"
	"testing"

	"github.com/babadro/algorithms-go/slices"
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

func TestArrToLinkedListRandom(t *testing.T) {
	arrays := [][][2]int{
		{{}},
		{{0, -1}},
		{{0, 1}, {1, -1}},
		{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
	}

	for i, arr := range arrays {
		head := ArrToLinkedListRandom(arr)
		node := head
		fmt.Printf("\nList %d: ", i+1)
		for node != nil {
			randomVal := -1
			if node.Random != nil {
				randomVal = node.Random.Val
			}
			fmt.Printf("%d %d, ", node.Val, randomVal)
			node = node.Next
		}
	}
}

/*
func TestLinkedListRandomToArr(t *testing.T) {
	arrays := [][][2]int{
		{{}},
		{{0, -1}},
		{{0, 1}, {1, -1}},
		{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
	}

	for i, input := range arrays {
		head := ArrToLinkedListRandom(input)
		output := LinkedListRandomToArr(head)
		if !sliceOfArraysAreEqual(input, output) {
			t.Errorf("case#%d want %v, got %v", i+1, input, output)
		}
	}
}

func sliceOfArraysAreEqual(arr1, arr2 [][2]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i][0] != arr2[i][0] {
			return false
		}
		if arr1[i][1] != arr2[i][1] {
			return false
		}
	}
	return true
}
*/
