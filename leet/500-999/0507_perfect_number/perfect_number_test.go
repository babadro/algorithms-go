package _507_perfect_number

import (
	"strconv"
	"testing"
)

func Test_checkPerfectNumber(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{6, true},
		{496, true},
		{8128, true},
		{2, false},
		{28, true},
		{1, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.num), func(t *testing.T) {
			if got := checkPerfectNumber2(tt.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkPerfectNumber(b *testing.B) {
	var res bool
	for i := 0; i < b.N; i++ {
		res = checkPerfectNumber(8128)
	}
	_ = res
}

func Benchmark_checkPerfectNumber2(b *testing.B) {
	var res bool
	for i := 0; i < b.N; i++ {
		res = checkPerfectNumber2(8128)
	}
	_ = res
}
