package _844_backspace_string_compare

import "testing"

func Test_backspaceCompare4(t *testing.T) {
	tests := []struct {
		S    string
		T    string
		want bool
	}{
		{"abc", "abc", true},
		{"bc", "abc", false},
		{"cba", "abc", false},
		{"", "a", false},
		{"", "", true},
		// backspace usecases
		{"ab#c", "ab#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
		{"######", "", true},
		{"bxj##tw", "bxo#j##tw", true},
	}
	for _, tt := range tests {
		t.Run(tt.S+"_"+tt.T, func(t *testing.T) {
			if got := backspaceCompare4(tt.S, tt.T); got != tt.want {
				t.Errorf("backspaceCompare4() = %v, want %v", got, tt.want)
			}
		})
	}
}
