package lexicographic_rank_of_string

import (
	"testing"
)

//func Test_factorial(t *testing.T) {
//	tests := []struct {
//		n    int
//		want int
//	}{
//		{0, 1},
//		{1, 1},
//		{2, 2},
//		{3, 6},
//		{4, 24},
//	}
//	for _, tt := range tests {
//		t.Run(fmt.Sprintf("%v", tt.n), func(t *testing.T) {
//			if got := modFactorial(tt.n, mod); got != tt.want {
//				t.Errorf("factorial() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_findRank(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"acb", 2},
		{"abc", 1},
		{"cba", 6},
		{"cab", 5},
		{"DTNGJPURFHYEW", 342501},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findRank(tt.s); got != tt.want {
				t.Errorf("findRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
