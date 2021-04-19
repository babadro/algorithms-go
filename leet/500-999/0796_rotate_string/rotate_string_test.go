package _796_rotate_string

import "testing"

func TestRotateString(t *testing.T) {
	cases := []struct {
		A, B     string
		expected bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"b", "b", true},
		{"b", "ba", false},
		{"ba", "ab", true},
		{"ba", "ab", true},
		{"aaa", "aaa", true},
		{"aaa", "aa", false},
	}
	for i, c := range cases {
		if fact := rotateString(c.A, c.B); fact != c.expected {
			t.Errorf("case#%d want %t, got %t", i+1, c.expected, fact)
		}
	}
}
