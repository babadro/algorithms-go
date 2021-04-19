package _168_excel_sheet_column_title

import "testing"

func Test_convertToString(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"1", 123, "123"},
		{"1", 909, "909"},
		{"1", 1200, "1200"},
		{"1", 1209, "1209"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToString(tt.n); got != tt.want {
				t.Errorf("convertToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"A", 1, "A"},
		{"Z", 26, "Z"},
		{"AA", 27, "AA"},
		{"701", 701, "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
