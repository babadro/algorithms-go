package _1882_process_tasks_using_servers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_assignTasks(t *testing.T) {
	tests := []struct {
		servers []int
		tasks   []int
		want    []int
	}{
		{[]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}, []int{2, 2, 0, 2, 1, 2}},
		{[]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}, []int{1, 4, 1, 4, 1, 3, 2}},
		{[]int{10, 63, 95, 16, 85, 57, 83, 95, 6, 29, 71}, []int{70, 31, 83, 15, 32, 67, 98, 65, 56, 48, 38, 90, 5}, []int{8, 0, 3, 9, 5, 1, 10, 6, 4, 2, 7, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.servers, tt.tasks), func(t *testing.T) {
			if got := assignTasks(tt.servers, tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
