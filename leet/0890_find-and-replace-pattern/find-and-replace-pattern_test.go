package _890_find_and_replace_pattern

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	tests := []struct {
		words   []string
		pattern string
		want    []string
	}{
		{
			words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			pattern: "abb",
			want:    []string{"mee", "aqq"},
		},
		{
			words: []string{"a"}, pattern: "b", want: []string{"a"},
		},
		{
			words: []string{"aaa"}, pattern: "ccc", want: []string{"aaa"},
		},
		{
			words: []string{"ab"}, pattern: "cac", want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := findAndReplacePattern(tt.words, tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
