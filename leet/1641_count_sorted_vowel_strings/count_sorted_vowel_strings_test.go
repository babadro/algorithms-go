package _1641_count_sorted_vowel_strings

import (
	"strconv"
	"testing"
)

func Test_countVowelStrings(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 5}, {2, 15}, {33, 66045},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := countVowelStrings2(tt.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countVowelStrings2(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = countVowelStrings2(33)
	}
	_ = res
}

func Benchmark_countVowelStrings3(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = countVowelStrings3(33)
	}
	_ = res
}
