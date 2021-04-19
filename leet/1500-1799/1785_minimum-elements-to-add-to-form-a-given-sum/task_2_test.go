package _785_minimum_elements_to_add_to_form_a_given_sum

import "testing"

func Test_minElements(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		limit int
		goal  int
		want  int
	}{
		{
			nums: []int{1, -1, 1}, limit: 3, goal: -4, want: 2,
		},
		{nums: []int{1, -10, 9, 1}, limit: 100, goal: 0, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElements(tt.nums, tt.limit, tt.goal); got != tt.want {
				t.Errorf("minElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
