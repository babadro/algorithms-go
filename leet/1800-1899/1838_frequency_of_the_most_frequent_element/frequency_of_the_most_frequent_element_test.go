package _1838_frequency_of_the_most_frequent_element

import (
	"fmt"
	"testing"
)

var longArr = []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966}

func Test_maxFrequency(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 4}, 5, 3},
		{[]int{1, 4, 8, 13}, 5, 2},
		{[]int{3, 9, 6}, 2, 1},
		{[]int{2, 3, 2, 100}, 3, 3},
		{[]int{2, 3, 2, 100}, 98 + 98 + 97, 4},
		{[]int{1}, 1, 1},
		{longArr, 3056, 73},
		{[]int{1, 2, 10, 11, 11, 11}, 1, 4},
		{[]int{1, 2, 10, 11, 12, 13}, 6, 4},
		{[]int{1, 1, 1, 3, 4}, 0, 3},
		{[]int{10_000}, 10_000, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := maxFrequency2(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething(t *testing.T) {
	res := 0
	for _, num := range longArr {
		if 9983 >= num {
			res++
		}
	}

	t.Log(res)
}
