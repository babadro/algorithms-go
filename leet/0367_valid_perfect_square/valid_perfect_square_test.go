package _367_valid_perfect_square

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"1", 1, true},
		{"2", 2, false},
		{"3", 4, true},
		{"5", 5, false},
		{"25", 25, true},
		{"2147395600", 2147395600, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare2(tt.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPerfectSquare2(2147395600)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPerfectSquare(2147395600)
	}
}
