package _769_minimum_number_of_operations_to_move_all_balls_to_each_box

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		boxes string
		want  []int
	}{
		{
			"1", "110", []int{1, 1, 3},
		},
		{"1", "001011", []int{11, 8, 5, 4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.boxes, func(t *testing.T) {
			if got := minOperations(tt.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
