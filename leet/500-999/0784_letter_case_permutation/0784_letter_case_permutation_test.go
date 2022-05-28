package _784_letter_case_permutation

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"ad52", []string{"ad52", "Ad52", "aD52", "AD52"}},
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"3z4", []string{"3z4", "3Z4"}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := letterCasePermutation(tt.s)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	t.Log(fmt.Sprintf("%d", byte('A')))
}
