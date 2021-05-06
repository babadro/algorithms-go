package _1849_splitting_a_string_into_descending_consecutive_values

import "testing"

func Test_splitString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"1234", false},
		{"050043", true},
		{"9080701", false},
		{"10009998", true},
		{"1", false},
		{"0", false},
		{"00", false},
		{"11", false},
		{"21", true},
		{"10", true},
		{"010", true},
		{"00000100990000098", true},
		{"000001009900000980", false},
		{"987654321", true},
		{"0987654321", true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := splitString(tt.s); got != tt.want {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
