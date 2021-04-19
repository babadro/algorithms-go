package _1409_queries_on_a_permutation_with_key

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_processQueries(t *testing.T) {
	tests := []struct {
		queries []int
		m       int
		want    []int
	}{
		{[]int{3, 1, 2, 1}, 5, []int{2, 1, 2, 1}},
		{[]int{4, 1, 2, 2}, 4, []int{3, 1, 2, 0}},
		{[]int{7, 5, 5, 8, 3}, 8, []int{6, 5, 0, 7, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.queries), func(t *testing.T) {
			if got := processQueries(tt.queries, tt.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
