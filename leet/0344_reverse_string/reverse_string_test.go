package _344_reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	for i := range s {
		t.Log(string(s[i]))
	}
}
