package _438_find_all_anagrams_in_a_string

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"a", "b", []int{}},
		{"a", "a", []int{0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.s, tt.p), func(t *testing.T) {
			if got := findAnagrams2(tt.s, tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
