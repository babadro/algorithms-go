package _1_find_the_median_of_a_number_stream

import (
	"testing"
)

func TestStreamMedian_InsertNum(t *testing.T) {
	s := StreamMedian{}

	s.InsertNum(1)
	s.InsertNum(2)
	s.InsertNum(3)

	t.Log(s.FindMedian())

	s.InsertNum(4)

	t.Log(s.FindMedian())
}
