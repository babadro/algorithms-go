package _0888_fair_candy_swap

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		B    []int
		want []int
	}{
		{"1", []int{1, 2, 5}, []int{2, 4}, []int{5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap2(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
