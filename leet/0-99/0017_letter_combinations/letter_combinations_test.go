package _017_letter_combinations

import (
	"reflect"
	"sort"
	"testing"
)

func Test_letterCombinations2(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.digits, func(t *testing.T) {
			got := letterCombinations2(tt.digits)
			sort.Strings(tt.want)
			sort.Strings(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations2() = %v, want %v", got, tt.want)
			}
		})
	}
}
