package _967_number_of_strings_that_appear_as_substring

import "testing"

func Test_numOfStrings(t *testing.T) {

	tests := []struct {
		patterns []string
		word     string
		want     int
	}{
		{[]string{"a", "abc", "bc", "d"}, "abc", 3},
		{[]string{"", "", ""}, "a", 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numOfStrings(tt.patterns, tt.word); got != tt.want {
				t.Errorf("numOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
