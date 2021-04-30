package _1837_sum_of_digits_in_base_k

import (
	"fmt"
	"testing"
)

func Test_sumBase(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{2050, 10, 7},
		{10, 10, 1},
		{0, 10, 0},
		{10, 2, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_%d", tt.n, tt.k), func(t *testing.T) {
			if got := sumBase(tt.n, tt.k); got != tt.want {
				t.Errorf("sumBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
