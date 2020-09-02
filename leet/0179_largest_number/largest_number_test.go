package _179_largest_number

import "testing"

func Test_largestNumber(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want string
	}{
		{name: "example 1", nums: []int{10, 2}, want: "210"},
		{name: "example 2", nums: []int{3, 30, 34, 5, 9}, want: "9534330"},
		{name: "example 2", nums: []int{3, 30, 3, 5, 9}, want: "953330"},
		{name: "example 2", nums: []int{3, 34, 3, 5, 9}, want: "953433"},
		{name: "example 2", nums: []int{3, 30, 3}, want: "3330"},
		{name: "example 2", nums: []int{}, want: ""},
		{name: "example 2", nums: []int{0, 0}, want: "0"},
		{name: "example 2", nums: []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}, want: "9609938824824769735703560743981399"}, // todo 1 - doesn't match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
