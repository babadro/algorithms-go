package _34

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{"a123b34c34", 2},
		{"uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17", 9},
		{"leet1234code234", 2},
		{"a1b01c001", 1},
		{"0a0", 1},
		{"a", 0},
		{"0", 1},
		{"1", 1},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := numDifferentIntegers(tt.word); got != tt.want {
				t.Errorf("numDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
