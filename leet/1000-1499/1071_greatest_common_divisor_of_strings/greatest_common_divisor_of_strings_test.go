package _1071_greatest_common_divisor_of_strings

import (
	"fmt"
	"testing"
)

func Test_gcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{3, 1, 1},
		{5, 13, 1},
		{36, 42, 6},
		{42, 36, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.a, tt.b), func(t *testing.T) {
			if got := gcd(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"ABAABAABA", "ABAABA", "ABA"},
		{"LEET", "CODE", ""},
		{"AAA", "AA", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.str1+" "+tt.str2, func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
