package _347_top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{[]int{1, 1, 1, 2, 2, 3}, 2}, want: []int{1, 2}},
		{name: "2", args: args{[]int{1}, 1}, want: []int{1}},
		{name: "3", args: args{[]int{3, 3, 3, 3, 4, 5, 5, 2, 4}, 2}, want: []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent3(tt.args.nums, tt.args.k)
			//sort.Ints(tt.want)
			//sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
