package _384_shuffle_an_array

import (
	"testing"
)

func TestSolution(t *testing.T) {
	obj := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	for i := 0; i < 30; i++ {
		t.Log(obj.Shuffle())
	}

	t.Log(obj.Reset())

}
