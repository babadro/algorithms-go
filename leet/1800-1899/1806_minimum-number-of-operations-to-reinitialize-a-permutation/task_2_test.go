package _1806_minimum_number_of_operations_to_reinitialize_a_permutation

import (
	"fmt"
	"testing"
)

func Test_reinitializePermutation(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 1},
		{4, 2},
		{6, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := reinitializePermutation(tt.n); got != tt.want {
				t.Errorf("reinitializePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test2(t *testing.T) {
	for n := 2; n <= 1000; n += 2 {
		t.Log(reinitializePermutation(n))
	}
}
