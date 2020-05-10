package _008_string_to_integer

import "testing"

func TestNumTimesAllBlue(t *testing.T) {
	cases := []struct {
		str      string
		expected int
	}{
		{"42", 42},
		{" ", 0},
		{"    ", 0},
		{"    -", 0},
		{"", 0},
		{"   -", 0},
		{"   -3", -3},
		{"   -3asdf", -3},
		{"   -3 a sdf", -3},
		{"a   -3 a sdf", 0},
		{"4193 with words", 4193},
		{"+4+193 with words", 4},
		{"words and 987", 0},
		{"-2", -2},
		{"-91283472332", -2147483648},
		{"-2147483647", -2147483647},
	}

	for i, c := range cases {
		if fact := myAtoi(c.str); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
