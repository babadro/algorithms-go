package _7_employee_free_time

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findEmployeeFreeTime(t *testing.T) {
	tests := []struct {
		schedule [][][]int
		want     [][]int
	}{
		{[][][]int{{{1, 2}, {5, 6}}}, [][]int{{2, 5}}},
		{[][][]int{{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, [][]int{{2, 3}, {4, 5}, {6, 7}}},
		{[][][]int{{{1, 2}, {5, 6}}, {{1, 3}}, {{4, 10}}}, [][]int{{3, 4}}},
		{[][][]int{{{1, 3}, {6, 7}}, {{2, 4}}, {{2, 5}, {9, 12}}}, [][]int{{5, 6}, {7, 9}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.schedule), func(t *testing.T) {
			if got := findEmployeeFreeTime(tt.schedule); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEmployeeFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
