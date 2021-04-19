package _138_copy_list_with_random_pointer

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	originArr := [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	originList := single.ArrToLinkedListRandom(originArr)
	copyList := copyRandomList(originList)

	originSlice, copySlice := single.LinkedListRandomToArr(single.ArrToLinkedListRandom(originArr)), single.LinkedListRandomToArr(copyList)
	if !sliceOfArraysAreEqual(originSlice, copySlice) {
		t.Errorf("not equal %v %v", originSlice, copySlice)
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
