package _925_long_pressed_name

import "testing"

func Test_isLongPressedName(t *testing.T) {
	tests := []struct {
		name  string
		typed string
		want  bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
		{"a", "a", true},
		{"aa", "a", false},
		{"b", "a", false},
		{"baaa", "ba", false},
		{"da", "daa", true},
		{"dadada", "ddaaddaaddaa", true},
		{"alex", "ablex", false},
	}
	for _, tt := range tests {
		t.Run(tt.typed, func(t *testing.T) {
			if got := isLongPressedName2(tt.name, tt.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
