package _900_RLE_Iterator

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	it := Constructor([]int{3, 8, 0, 9, 2, 5})

	t.Log(it.Next(2)) // expected 8
	t.Log(it.Next(1)) // expected 8
	t.Log(it.Next(1)) // expected 5
	t.Log(it.Next(2)) // expected -1
}
