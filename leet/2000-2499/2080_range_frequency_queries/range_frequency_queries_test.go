package _2080_range_frequency_queries

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	q := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})

	t.Log(q.Query(1, 2, 4))
	t.Log(q.Query(0, 11, 33))

	t.Log(q.Query(0, 1, 35))
	t.Log(q.Query(0, 0, 12))
	t.Log(q.Query(0, 0, 1))
}
