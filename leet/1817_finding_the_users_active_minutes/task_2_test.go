package _818_min_absolute_sum_difference

import (
	"reflect"
	"testing"
)

func Test_findingUsersActiveMinutes(t *testing.T) {
	tests := []struct {
		name string
		logs [][]int
		k    int
		want []int
	}{
		{
			name: "1",
			logs: [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}},
			k:    5,
			want: []int{0, 2, 0, 0, 0},
		},
		{
			name: "2",
			logs: [][]int{{1, 1}, {2, 2}, {2, 3}},
			k:    4,
			want: []int{1, 1, 0, 0},
		},
		{
			name: "3",
			logs: [][]int{{1, 1}},
			k:    2,
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findingUsersActiveMinutes(tt.logs, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findingUsersActiveMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
