package _1556_thousand_separator

import (
	"strconv"
	"testing"
)

func Test_thousandSeparator(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{12, "12"},
		{123, "123"},
		{1234, "1.234"},
		{12345, "12.345"},
		{123456, "123.456"},
		{1234567, "1.234.567"},
		{12345678, "12.345.678"},
		{123456789, "123.456.789"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := thousandSeparator(tt.n); got != tt.want {
				t.Errorf("thousandSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
