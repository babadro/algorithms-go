package _2213_longest_subsring_of_one_repeating_character

import (
	"reflect"
	"testing"
)

func Test_longestRepeating(t *testing.T) {
	tests := []struct {
		s               string
		queryCharacters string
		queryIndices    []int
		want            []int
	}{
		{"babacc", "bcb", []int{1, 1, 3}, []int{3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestRepeating2(tt.s, tt.queryCharacters, tt.queryIndices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestRepeating2() = %v, want %v", got, tt.want)
			}
		})
	}
}
