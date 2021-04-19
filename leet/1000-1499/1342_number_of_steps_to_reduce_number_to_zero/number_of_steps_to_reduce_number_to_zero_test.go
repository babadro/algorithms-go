package _1342_number_of_steps_to_reduce_number_to_zero

import "testing"

func Test_numberOfSteps(t *testing.T) {

	tests := []struct {
		name string
		num  int
		want int
	}{
		{"1", 14, 6},
		{"2", 8, 4},
		{"3", 123, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
