package _013_Roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		romanNum string
		expected int
	}{
		//{"", 0},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"C", 100},
		{"CC", 200},
		{"CCC", 300},
		{"CD", 400},
		{"D", 500},
		{"DC", 600},
		{"DCC", 700},
		{"DCCC", 800},
		{"CM", 900},
		{"M", 1000},
	}

	for i, c := range cases {
		if fact := romanToInt(c.romanNum); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
