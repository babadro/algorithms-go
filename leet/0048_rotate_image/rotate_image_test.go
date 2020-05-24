package _048_rotate_image

import "testing"

// TODO 1
func TestRotate(t *testing.T) {
	_ = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	input2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	input3 := [][]int{
		{1},
	}
	_ = input2
	rotate(input3)
	t.Log(input3)
}

func BenchmarkMyRotate(b *testing.B) {
	input2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	for i := 0; i < b.N; i++ {
		rotate(input2)
	}
}

func BenchmarkRotate2(b *testing.B) {
	input2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	for i := 0; i < b.N; i++ {
		rotate2(input2)
	}
}
