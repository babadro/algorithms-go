package _441_arranging_coins

import (
	"fmt"
	"math"
	"testing"
)

func Test_arrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 2},
		{8, 3},
		{10, 4},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			if got := arrangeCoins3(tt.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test2(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(i, arrangeCoins(i))
	}

	t.Log("\n")
	t.Log(arrangeCoins(math.MaxInt32))
}

func Benchmark_arrangeCoinsBruteForce(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = arrangeCoins(math.MaxInt32)
	}
	_ = res
}

func Benchmark_arrangeCoinsBinarySearch(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = arrangeCoins2(math.MaxInt32)
	}
	_ = res
}

func Benchmark_arrangeCoinsMath(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = arrangeCoins3(math.MaxInt32)
	}
	_ = res
}
