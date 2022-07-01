package _844_backspace_string_compare

import "testing"

func Test_backspaceCompare(t *testing.T) {

	tests := []struct {
		name string
		S    string
		T    string
		want bool
	}{
		{"1", "ab#c", "ad#c", true},
		{"2", "ab##", "c#d#", true},
		{"3", "a##c", "#a#c", true},
		{"4", "a#c", "b", false},
		{"5", "", "", true},
		{"6", "sadf", "sadf", true},
		{"7", "sadf1", "sadf2", false},
		{"8", "sadf1", "sadf2", false},
		{"9", "y#fo##f", "y#f#o##f", true},
		{"", "#", "#", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare3(tt.S, tt.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
