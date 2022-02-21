package _2157_groups_of_strings

import (
	"reflect"
	"testing"
)

func Test_groupStrings(t *testing.T) {
	tests := []struct {
		words []string
		want  []int
	}{
		{[]string{"a", "b", "ab", "cde"}, []int{2, 3}},
		{[]string{"a", "b", "ab", "cde", "cde"}, []int{2, 3}},
		{[]string{"a", "b", "ab", "cde", "cde", "cde", "cde"}, []int{2, 4}},
		{[]string{"a", "b", "ab", "cde", "edc", "ecd", "dec"}, []int{2, 4}},
		{[]string{"a", "ab", "abc"}, []int{1, 3}},

		{[]string{"b", "q"}, []int{1, 2}},
		{[]string{"qamp", "am", "khdrn"}, []int{3, 1}},
		{tleInput1, []int{10624, 4144}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := groupStrings2(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordsToNums(t *testing.T) {
	tests := []struct {
		words []string
		want  []int32
	}{
		{[]string{"a", "b", "ab", "cde"}, []int32{0b1, 0b10, 0b11, 0b11100}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := wordsToNums(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordsToNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
