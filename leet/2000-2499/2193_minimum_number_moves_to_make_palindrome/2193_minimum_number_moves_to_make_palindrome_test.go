package _2193_minimum_number_moves_to_make_palindrome

import "testing"

func Test_minMovesToMakePalindrome(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"aabb", 2},
		{"letelt", 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minMovesToMakePalindrome2(tt.s); got != tt.want {
				t.Errorf("minMovesToMakePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
