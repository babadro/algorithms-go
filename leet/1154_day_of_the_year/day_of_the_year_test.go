package _1154_day_of_the_year

import "testing"

func Test_dayOfYear(t *testing.T) {
	tests := []struct {
		date string
		want int
	}{
		{"2019-01-09", 9},
		{"2019-02-10", 41},
		{"2003-03-01", 60},
		{"2020-03-01", 61},
		{"1900-03-25", 84},
	}
	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			if got := dayOfYear(tt.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
