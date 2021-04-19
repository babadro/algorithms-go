package _434_number_of_segments_in_a_string

import "testing"

func Test_countSegments(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", "Hello, my name is John", 5},
		{"2", "Hello, my name is   John", 5},
		{"3", "Hello, my name is   John ", 5},
		{"4", " Hello, my name is   John ", 5},
		{"5", "hello", 1},
		{"6", "", 0},
		{"7", "  ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments2(tt.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
