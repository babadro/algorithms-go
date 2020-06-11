package _299_bulls_and_cows

import "testing"

func TestGetHint(t *testing.T) {
	cases := []struct {
		secret, guess, expected string
	}{
		{"", "", "0A0B"},
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
		{"1111", "1111", "4A0B"},
		{"1234", "4321", "0A4B"},
		{"1211", "2121", "1A2B"},
	}
	for i, c := range cases {
		if fact := getHint(c.secret, c.guess); fact != c.expected {
			t.Errorf("case#%d want %s, got %s", i+1, c.expected, fact)
		}
	}
}
