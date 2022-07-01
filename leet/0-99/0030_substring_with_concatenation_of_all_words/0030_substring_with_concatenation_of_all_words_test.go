package _030_substring_with_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, nil},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findSubstring(tt.s, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
