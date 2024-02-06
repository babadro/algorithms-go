package _2013_detect_squares

import "testing"

func TestDetectSquares_Count(t *testing.T) {
	ds := Constructor()

	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})

	t.Log(ds.Count([]int{11, 10})) // 1
	t.Log(ds.Count([]int{14, 8}))  // 0

	ds.Add([]int{11, 2})

	t.Log(ds.Count([]int{11, 10})) // 2
}
