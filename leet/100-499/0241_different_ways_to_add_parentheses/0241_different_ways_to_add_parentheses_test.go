package _241_different_ways_to_add_parentheses

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	tests := []struct {
		expression string
		want       []int
	}{
		{"2-1-1", []int{0, 2}},
		{"2*3-4*5", []int{-34, -14, -10, -10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			got := diffWaysToCompute(tt.expression)
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
