package _1_longest_alternating_subsequence

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_findLASLen(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 2},
		{[]int{3, 2, 1, 4}, 3},
		{[]int{1, 3, 2, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findLASLenTopDown3DArr(tt.nums); got != tt.want {
				t.Errorf("findLASLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBruteForce(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLASLenBruteForce(input)
	}
	_ = res
}

func BenchmarkTopDown(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLASLenTopDown(input)
	}
	_ = res
}

func BenchmarkTopDown3DArr(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLASLenTopDown3DArr(input)
	}
	_ = res
}

func BenchmarkTopDown2DP(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLASLenTopDown2DP(input)
	}
	_ = res
}

func BenchmarkBruteForceAlternative(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLASLenBruteForceAlternative(input)
	}
	_ = res
}

func BenchmarkTopDownAlternative(b *testing.B) {
	var res int
	input := make([]int, 40)
	for i := range input {
		input[i] = rand.Intn(10)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = findLasLenTopDownAlternative(input)
	}
	_ = res
}
