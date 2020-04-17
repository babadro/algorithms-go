package _171_excel_sheet_column_number

import "testing"

func TestTitleToNumber(t *testing.T) {
	cases := []struct {
		title    string
		expected int
	}{
		{"A", 1},
		{"B", 2},
		{"Z", 26},
		{"AA", 27},
		{"AB", 28},
		{"ZY", 701},
		{"AAA", 703},
	}

	for i, c := range cases {
		if fact := titleToNumber(c.title); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
