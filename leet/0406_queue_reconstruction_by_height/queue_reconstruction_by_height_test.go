package _406_queue_reconstruction_by_height

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {

	tests := []struct {
		people [][]int
		want   [][]int
	}{
		{
			[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
		{
			[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
			[][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}},
		},
		{
			[][]int{{2, 4}, {3, 4}, {9, 0}, {0, 6}, {7, 1}, {6, 0}, {7, 3}, {2, 5}, {1, 1}, {8, 0}},
			[][]int{{6, 0}, {1, 1}, {8, 0}, {7, 1}, {9, 0}, {2, 4}, {2, 5}, {3, 4}, {7, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.people), func(t *testing.T) {
			if got := reconstructQueue(tt.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueueOld() = %v, want %v", got, tt.want)
			}
		})
	}
}
